import 'package:flutter/foundation.dart';
import '../models/user.dart';
import '../core/api_client.dart';

class AuthProvider extends ChangeNotifier {
  final ApiClient _apiClient = ApiClient();
  
  User? _user;
  bool _isLoading = false;
  String? _error;

  User? get user => _user;
  bool get isLoading => _isLoading;
  String? get error => _error;
  bool get isAuthenticated => _user != null;

  // Login
  Future<bool> login(String username, String password) async {
    _isLoading = true;
    _error = null;
    notifyListeners();

    try {
      final response = await _apiClient.login(username, password);
      
      if (response.success && response.data != null) {
        _user = response.data;
        _isLoading = false;
        notifyListeners();
        return true;
      } else {
        _error = response.message ?? 'Login failed';
        _isLoading = false;
        notifyListeners();
        return false;
      }
    } catch (e) {
      _error = 'Network error: $e';
      _isLoading = false;
      notifyListeners();
      return false;
    }
  }

  // Register
  Future<bool> register(String username, String email, String password) async {
    _isLoading = true;
    _error = null;
    notifyListeners();

    try {
      final response = await _apiClient.register(username, email, password);
      
      if (response.success) {
        _isLoading = false;
        notifyListeners();
        return true;
      } else {
        _error = response.message ?? 'Registration failed';
        _isLoading = false;
        notifyListeners();
        return false;
      }
    } catch (e) {
      _error = 'Network error: $e';
      _isLoading = false;
      notifyListeners();
      return false;
    }
  }

  // Logout
  Future<void> logout() async {
    await _apiClient.logout();
    _user = null;
    notifyListeners();
  }

  // Load user profile
  Future<void> loadUserProfile() async {
    try {
      final response = await _apiClient.getUserProfile();
      if (response.success && response.data != null) {
        _user = response.data;
        notifyListeners();
      }
    } catch (e) {
      debugPrint('Failed to load user profile: $e');
    }
  }

  // Clear error
  void clearError() {
    _error = null;
    notifyListeners();
  }
}
