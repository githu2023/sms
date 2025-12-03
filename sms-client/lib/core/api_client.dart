import 'dart:convert';
import 'package:flutter/foundation.dart';
import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';
import '../models/api_response.dart';
import '../models/assignment.dart';
import '../models/assigned_phone.dart';
import '../models/business_type.dart';
import '../models/phone_assignment_result.dart';
import '../models/user.dart';
import '../models/verification_code_entry.dart';
import '../models/verification_code_result.dart';
import '../models/whitelist.dart';
import '../models/transaction.dart';

class ApiClient {
  // 根据环境自动选择 baseUrl
  // 开发环境使用本地地址，生产环境使用服务器地址
  static String get baseUrl {
    if (kDebugMode) {
      // 本地开发环境
      return 'http://localhost:6060';
    } else {
      // 生产环境 - 请根据实际情况修改为你的服务器地址
      return 'http://38.60.203.212:6060/v1';
    }
  }
  
  static const String tokenKey = 'auth_token';

  String? _token;

  // Get stored token
  Future<String?> getToken() async {
    if (_token != null) return _token;
    final prefs = await SharedPreferences.getInstance();
    _token = prefs.getString(tokenKey);
    return _token;
  }

  // Save token
  Future<void> saveToken(String token) async {
    _token = token;
    final prefs = await SharedPreferences.getInstance();
    await prefs.setString(tokenKey, token);
  }

  // Clear token
  Future<void> clearToken() async {
    _token = null;
    final prefs = await SharedPreferences.getInstance();
    await prefs.remove(tokenKey);
  }

  // Build headers
  Future<Map<String, String>> _buildHeaders({bool needsAuth = true}) async {
    final headers = {'Content-Type': 'application/json'};

    if (needsAuth) {
      final token = await getToken();
      if (token != null) {
        headers['Authorization'] = 'Bearer $token';
      }
    }

    return headers;
  }

  // Handle response
  ApiResponse<T> _handleResponse<T>(
    http.Response response,
    T Function(dynamic)? fromJson,
  ) {
    try {
      // 先检查响应体是否为空
      if (response.body.isEmpty) {
        if (response.statusCode >= 200 && response.statusCode < 300) {
          return ApiResponse.success(message: 'Success');
        } else {
          return ApiResponse.error(
            message: 'Empty response',
            code: response.statusCode,
          );
        }
      }

      final jsonData = json.decode(utf8.decode(response.bodyBytes));

      if (response.statusCode >= 200 && response.statusCode < 300) {
        // 检查是否是标准格式（有 code 字段）
        if (jsonData is Map<String, dynamic> && jsonData.containsKey('code')) {
          return ApiResponse.fromJson(jsonData, fromJson);
        } else {
          // 直接返回数据格式（如 transactions、whitelist 接口）
          // 将整个响应作为 data 处理
          if (fromJson != null) {
            return ApiResponse<T>(
              success: true,
              message: 'Success',
              data: fromJson(jsonData),
              code: 200,
            );
          } else {
            return ApiResponse<T>(
              success: true,
              message: 'Success',
              data: jsonData as T?,
              code: 200,
            );
          }
        }
      } else {
        return ApiResponse.error(
          message: jsonData['message'] ?? jsonData['msg'] ?? 'Request failed',
          code: jsonData['code'] ?? response.statusCode,
        );
      }
    } catch (e) {
      // 打印原始响应以便调试
      debugPrint('Response parsing error: $e');
      debugPrint('Status code: ${response.statusCode}');
      debugPrint('Response body: ${response.body}');

      // 更友好的错误提示
      String errorMsg = '服务器响应格式错误';
      if (e.toString().contains('FormatException')) {
        errorMsg = '服务器返回了无效的数据格式';
      } else if (e.toString().contains('Unexpected end of JSON')) {
        errorMsg = '服务器响应不完整，请稍后重试';
      }

      return ApiResponse.error(message: errorMsg, code: response.statusCode);
    }
  }

  // POST request
  Future<ApiResponse<T>> post<T>(
    String endpoint,
    Map<String, dynamic> data, {
    T Function(dynamic)? fromJson,
    bool needsAuth = false,
  }) async {
    try {
      final headers = await _buildHeaders(needsAuth: needsAuth);
      debugPrint('POST $baseUrl$endpoint');
      debugPrint('Request body: ${json.encode(data)}');

      final response = await http.post(
        Uri.parse('$baseUrl$endpoint'),
        headers: headers,
        body: json.encode(data),
      );

      return _handleResponse(response, fromJson);
    } catch (e) {
      debugPrint('Network error: $e');
      // 判断是否是连接错误
      if (e.toString().contains('SocketException') ||
          e.toString().contains('Failed host lookup') ||
          e.toString().contains('Connection refused')) {
        return ApiResponse.error(message: '无法连接到服务器，请检查网络连接或服务器地址');
      }
      return ApiResponse.error(message: '网络错误: $e');
    }
  }

  // GET request
  Future<ApiResponse<T>> get<T>(
    String endpoint, {
    Map<String, String>? queryParams,
    T Function(dynamic)? fromJson,
    bool needsAuth = true,
  }) async {
    try {
      final headers = await _buildHeaders(needsAuth: needsAuth);
      var uri = Uri.parse('$baseUrl$endpoint');
      if (queryParams != null) {
        uri = uri.replace(queryParameters: queryParams);
      }

      debugPrint('GET $uri');
      final response = await http.get(uri, headers: headers);
      return _handleResponse(response, fromJson);
    } catch (e) {
      debugPrint('Network error: $e');
      if (e.toString().contains('SocketException') ||
          e.toString().contains('Failed host lookup') ||
          e.toString().contains('Connection refused')) {
        return ApiResponse.error(message: '无法连接到服务器，请检查网络连接或服务器地址');
      }
      return ApiResponse.error(message: '网络错误: $e');
    }
  }

  // PUT request
  Future<ApiResponse<T>> put<T>(
    String endpoint,
    Map<String, dynamic> data, {
    T Function(dynamic)? fromJson,
    bool needsAuth = true,
  }) async {
    try {
      final headers = await _buildHeaders(needsAuth: needsAuth);
      final response = await http.put(
        Uri.parse('$baseUrl$endpoint'),
        headers: headers,
        body: json.encode(data),
      );

      return _handleResponse(response, fromJson);
    } catch (e) {
      return ApiResponse.error(message: 'Network error: $e');
    }
  }

  // DELETE request
  Future<ApiResponse<T>> delete<T>(
    String endpoint, {
    Map<String, dynamic>? data,
    T Function(dynamic)? fromJson,
    bool needsAuth = true,
  }) async {
    try {
      final headers = await _buildHeaders(needsAuth: needsAuth);
      final response = await http.delete(
        Uri.parse('$baseUrl$endpoint'),
        headers: headers,
        body: data != null ? json.encode(data) : null,
      );

      return _handleResponse(response, fromJson);
    } catch (e) {
      return ApiResponse.error(message: 'Network error: $e');
    }
  }

  // Auth APIs
  Future<ApiResponse<User>> login(String username, String password) async {
    final response = await post(
      '/client/v1/login',
      {'username': username, 'password': password},
      fromJson: (data) {
        // Save token from response
        final token = data['token'] as String?;
        if (token != null) {
          saveToken(token);
        }
        // Return a minimal user object (profile will be fetched separately)
        return User(
          id: '',
          username: username,
          email: '',
          balance: 0.0,
          createdAt: DateTime.now(),
        );
      },
    );

    return response;
  }

  Future<ApiResponse<User>> register(
    String username,
    String email,
    String password,
  ) async {
    return await post(
      '/client/v1/register',
      {'username': username, 'email': email, 'password': password},
      fromJson:
          (data) => User(
            id: data['user_id'].toString(),
            username: data['username'] as String,
            email: email,
            balance: 0.0,
            createdAt: DateTime.now(),
          ),
    );
  }

  Future<ApiResponse<void>> logout() async {
    await clearToken();
    return ApiResponse.success(message: 'Logged out successfully');
  }

  Future<ApiResponse<User>> getUserProfile() async {
    return await get(
      '/client/v1/profile',
      fromJson: (data) => User.fromJson(data),
    );
  }

  // Business Type APIs
  Future<ApiResponse<List<BusinessType>>> getBusinessTypes() async {
    return await get(
      '/client/v1/business_types',
      fromJson: (data) {
        // data 可能是数组或包含 business_types 的对象
        List types;
        if (data is List) {
          types = data;
        } else if (data is Map && data.containsKey('business_types')) {
          types = data['business_types'] as List;
        } else {
          types = [];
        }
        return types.map((e) => BusinessType.fromJson(e)).toList();
      },
    );
  }

  // Phone Assignment APIs
  Future<ApiResponse<PhoneAssignmentResult>> assignPhone({
    required String businessType,
    required String cardType,
    int count = 1, // 批量获取数量，默认1个
  }) async {
    return await post(
      '/client/v1/get_phone',
      {'business_type': businessType, 'card_type': cardType, 'count': count},
      fromJson:
          (data) =>
              PhoneAssignmentResult.fromJson(data as Map<String, dynamic>),
      needsAuth: true,
    );
  }

  // Verification Code APIs
  Future<ApiResponse<VerificationCodeResult>> getVerificationCodes({
    required List<String> phoneNumbers, // 支持批量手机号
  }) async {
    return await post(
      '/client/v1/get_code',
      {'phone_numbers': phoneNumbers},
      fromJson:
          (data) =>
              VerificationCodeResult.fromJson(data as Map<String, dynamic>),
      needsAuth: true,
    );
  }

  // 向后兼容：单个手机号获取验证码的便捷方法
  Future<ApiResponse<VerificationCodeEntry?>> getVerificationCodeSingle({
    required String phone,
  }) async {
    final response = await getVerificationCodes(phoneNumbers: [phone]);

    if (response.success && response.data != null) {
      if (response.data!.codes.isNotEmpty) {
        final entry = response.data!.codes.first;
        return ApiResponse.success(
          data: entry,
          message: response.message ?? 'Success',
        );
      }
    }

    return ApiResponse.error(
      message: response.message ?? 'Failed to get verification code',
      code: response.code,
    );
  }

  // 向后兼容：单个手机号获取的便捷方法
  Future<ApiResponse<AssignedPhone?>> assignPhoneSingle({
    required String businessType,
    required String cardType,
  }) async {
    final response = await assignPhone(
      businessType: businessType,
      cardType: cardType,
      count: 1,
    );

    if (response.success && response.data != null) {
      if (response.data!.phones.isNotEmpty) {
        final phoneInfo = response.data!.phones.first;
        return ApiResponse.success(
          data: phoneInfo,
          message: response.message ?? 'Success',
        );
      }
    }

    return ApiResponse.error(
      message: response.message ?? 'Failed to assign phone',
      code: response.code,
    );
  }

  // Assignment History APIs
  Future<ApiResponse<List<Assignment>>> getAssignments({
    int page = 1,
    int limit = 20,
  }) async {
    return await get(
      '/client/v1/assignments',
      queryParams: {'page': page.toString(), 'limit': limit.toString()},
      fromJson: (data) {
        final items = data['items'] as List;
        return items.map((e) => Assignment.fromJson(e)).toList();
      },
    );
  }

  // Balance API
  Future<ApiResponse<double>> getBalance() async {
    return await get(
      '/client/v1/balance',
      fromJson: (data) => (data['balance'] as num).toDouble(),
    );
  }

  // Change Password API
  Future<ApiResponse<void>> changePassword({
    required String oldPassword,
    required String newPassword,
  }) async {
    return await post('/client/v1/password/change', {
      'old_password': oldPassword,
      'new_password': newPassword,
    }, needsAuth: true);
  }

  // Phone Status API
  Future<ApiResponse<Map<String, dynamic>>> getPhoneStatus({
    required String phone,
  }) async {
    return await get(
      '/client/v1/phone_status',
      queryParams: {'phone_number': phone},
      fromJson: (data) => data as Map<String, dynamic>,
    );
  }

  // Cost Statistics API
  Future<ApiResponse<Map<String, dynamic>>> getCostStatistics({
    String? startDate,
    String? endDate,
  }) async {
    final queryParams = <String, String>{};
    if (startDate != null) queryParams['start_date'] = startDate;
    if (endDate != null) queryParams['end_date'] = endDate;

    return await get(
      '/client/v1/assignments/statistics',
      queryParams: queryParams.isNotEmpty ? queryParams : null,
      fromJson: (data) => data as Map<String, dynamic>,
    );
  }

  // Whitelist APIs
  Future<ApiResponse<List<Whitelist>>> getWhitelists() async {
    return await get(
      '/client/v1/whitelist',
      fromJson: (data) {
        final items = data as List<dynamic>;
        return items
            .map((item) => Whitelist.fromJson(item as Map<String, dynamic>))
            .toList();
      },
    );
  }

  Future<ApiResponse<Whitelist>> addWhitelist({
    required String ipAddress,
    String? notes,
  }) async {
    return await post(
      '/client/v1/whitelist',
      {'ip_address': ipAddress, if (notes != null) 'notes': notes},
      fromJson: (data) => Whitelist.fromJson(data as Map<String, dynamic>),
      needsAuth: true,
    );
  }

  Future<ApiResponse<void>> deleteWhitelist({required String ipAddress}) async {
    return await delete(
      '/client/v1/whitelist',
      data: {'ip_address': ipAddress},
      needsAuth: true,
    );
  }

  // Transaction APIs
  Future<ApiResponse<TransactionListResponse>> getTransactions({
    int limit = 20,
    int offset = 0,
  }) async {
    return await get(
      '/client/v1/transactions',
      queryParams: {
        'limit': limit.toString(),
        'offset': offset.toString(),
      },
      fromJson: (data) =>
          TransactionListResponse.fromJson(data as Map<String, dynamic>),
    );
  }

  Future<ApiResponse<TransactionListResponse>> getTransactionsByType({
    required int type, // 1=充值，2=消费
    int limit = 20,
    int offset = 0,
  }) async {
    return await get(
      '/client/v1/transactions/by-type',
      queryParams: {
        'type': type.toString(),
        'limit': limit.toString(),
        'offset': offset.toString(),
      },
      fromJson: (data) =>
          TransactionListResponse.fromJson(data as Map<String, dynamic>),
    );
  }

  Future<ApiResponse<TransactionListResponse>> getTransactionsByDate({
    required String startDate, // 格式：2006-01-02
    required String endDate, // 格式：2006-01-02
    int limit = 20,
    int offset = 0,
  }) async {
    return await get(
      '/client/v1/transactions/by-date',
      queryParams: {
        'start_date': startDate,
        'end_date': endDate,
        'limit': limit.toString(),
        'offset': offset.toString(),
      },
      fromJson: (data) =>
          TransactionListResponse.fromJson(data as Map<String, dynamic>),
    );
  }
}
