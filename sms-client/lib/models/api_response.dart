class ApiResponse<T> {
  final bool success;
  final String? message;
  final T? data;
  final int? code;

  ApiResponse({
    required this.success,
    this.message,
    this.data,
    this.code,
  });

  factory ApiResponse.fromJson(
    Map<String, dynamic> json,
    T Function(dynamic)? fromJsonT,
  ) {
    final code = json['code'] as int?;
    final isSuccess = code != null && code >= 200 && code < 300;
    
    return ApiResponse<T>(
      success: isSuccess,
      message: json['msg'] as String?,
      data: json['data'] != null && fromJsonT != null
          ? fromJsonT(json['data'])
          : json['data'] as T?,
      code: code,
    );
  }

  factory ApiResponse.success({T? data, String? message}) {
    return ApiResponse(
      success: true,
      data: data,
      message: message,
    );
  }

  factory ApiResponse.error({required String message, int? code}) {
    return ApiResponse(
      success: false,
      message: message,
      code: code,
    );
  }
}
