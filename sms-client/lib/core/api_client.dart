import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';
import '../models/api_response.dart';
import '../models/user.dart';
import '../models/business_type.dart';
import '../models/assignment.dart';

class ApiClient {
  static const String baseUrl = 'http://localhost:8080';
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
    final headers = {
      'Content-Type': 'application/json',
    };

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
      final jsonData = json.decode(utf8.decode(response.bodyBytes));

      if (response.statusCode >= 200 && response.statusCode < 300) {
        return ApiResponse.fromJson(jsonData, fromJson);
      } else {
        return ApiResponse.error(
          message: jsonData['message'] ?? jsonData['msg'] ?? 'Request failed',
          code: jsonData['code'] ?? response.statusCode,
        );
      }
    } catch (e) {
      return ApiResponse.error(
        message: 'Failed to parse response: $e',
        code: response.statusCode,
      );
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
      final response = await http.post(
        Uri.parse('$baseUrl$endpoint'),
        headers: headers,
        body: json.encode(data),
      );

      return _handleResponse(response, fromJson);
    } catch (e) {
      return ApiResponse.error(message: 'Network error: $e');
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

      final response = await http.get(uri, headers: headers);
      return _handleResponse(response, fromJson);
    } catch (e) {
      return ApiResponse.error(message: 'Network error: $e');
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
      {
        'username': username,
        'email': email,
        'password': password,
      },
      fromJson: (data) => User(
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
        final types = data['business_types'] as List;
        return types.map((e) => BusinessType.fromJson(e)).toList();
      },
    );
  }

  // Phone Assignment APIs
  Future<ApiResponse<Map<String, dynamic>>> assignPhone({
    required String businessType,
    required String cardType,
    int count = 1, // 批量获取数量，默认1个
  }) async {
    return await post(
      '/client/v1/get_phone',
      {
        'business_type': businessType,
        'card_type': cardType,
        'count': count,
      },
      fromJson: (data) => data as Map<String, dynamic>,
      needsAuth: true,
    );
  }

  // Verification Code APIs
  Future<ApiResponse<Map<String, dynamic>>> getVerificationCode({
    required List<String> phoneNumbers, // 支持批量手机号
    int timeout = 60,
  }) async {
    return await post(
      '/client/v1/get_code',
      {
        'phone_numbers': phoneNumbers,
        'timeout': timeout,
      },
      fromJson: (data) => data as Map<String, dynamic>,
      needsAuth: true,
    );
  }

  // 向后兼容：单个手机号获取验证码的便捷方法
  Future<ApiResponse<String?>> getVerificationCodeSingle({
    required String phone,
    int timeout = 60,
  }) async {
    final response = await getVerificationCode(
      phoneNumbers: [phone],
      timeout: timeout,
    );
    
    if (response.success && response.data != null) {
      final codes = response.data!['codes'] as List?;
      if (codes != null && codes.isNotEmpty) {
        final codeInfo = codes.first as Map<String, dynamic>;
        return ApiResponse.success(
          data: codeInfo['code'] as String?,
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
  Future<ApiResponse<String?>> assignPhoneSingle({
    required String businessType,
    required String cardType,
  }) async {
    final response = await assignPhone(
      businessType: businessType,
      cardType: cardType,
      count: 1,
    );
    
    if (response.success && response.data != null) {
      final phones = response.data!['phones'] as List?;
      if (phones != null && phones.isNotEmpty) {
        final phoneInfo = phones.first as Map<String, dynamic>;
        return ApiResponse.success(
          data: phoneInfo['phone_number'] as String?,
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
      queryParams: {
        'page': page.toString(),
        'limit': limit.toString(),
      },
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
    return await post(
      '/client/v1/change_password',
      {
        'old_password': oldPassword,
        'new_password': newPassword,
      },
      needsAuth: true,
    );
  }

  // Phone Status API
  Future<ApiResponse<Map<String, dynamic>>> getPhoneStatus({
    required String phone,
  }) async {
    return await get(
      '/client/v1/phone_status',
      queryParams: {'phone': phone},
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
  Future<ApiResponse<List<Map<String, dynamic>>>> getWhitelists() async {
    return await get(
      '/client/v1/whitelist',
      fromJson: (data) {
        final items = data as List;
        return items.cast<Map<String, dynamic>>();
      },
    );
  }

  Future<ApiResponse<Map<String, dynamic>>> addWhitelist({
    required String ipAddress,
    String? notes,
  }) async {
    return await post(
      '/client/v1/whitelist',
      {
        'ip_address': ipAddress,
        if (notes != null) 'notes': notes,
      },
      fromJson: (data) => data as Map<String, dynamic>,
      needsAuth: true,
    );
  }

  Future<ApiResponse<void>> deleteWhitelist({
    required String ipAddress,
  }) async {
    return await delete(
      '/client/v1/whitelist',
      data: {
        'ip_address': ipAddress,
      },
      needsAuth: true,
    );
  }
}
