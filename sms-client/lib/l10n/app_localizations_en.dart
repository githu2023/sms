// ignore: unused_import
import 'package:intl/intl.dart' as intl;
import 'app_localizations.dart';

// ignore_for_file: type=lint

/// The translations for English (`en`).
class AppLocalizationsEn extends AppLocalizations {
  AppLocalizationsEn([String locale = 'en']) : super(locale);

  @override
  String get appTitle => 'SMS Platform';

  @override
  String get login => 'Login';

  @override
  String get register => 'Register';

  @override
  String get username => 'Username';

  @override
  String get password => 'Password';

  @override
  String get email => 'Email';

  @override
  String get loginTitle => 'Login SMS Platform';

  @override
  String get registerTitle => 'Register New Account';

  @override
  String get noAccount => 'No account? Register';

  @override
  String get hasAccount => 'Already have an account? Login';

  @override
  String get home => 'Home';

  @override
  String get profile => 'Profile';

  @override
  String get whitelist => 'Whitelist';

  @override
  String get settings => 'Settings';

  @override
  String get balance => 'Current Balance';

  @override
  String get frozenAmount => 'Frozen Amount';

  @override
  String get availableBalance => 'Available Balance';

  @override
  String get recharge => 'Recharge';

  @override
  String get assignPhone => 'Assign Phone';

  @override
  String get getPhone => 'Get Phone';

  @override
  String get getCode => 'Get Code';

  @override
  String get history => 'Assignment History';

  @override
  String get recentAssignments => 'Recent Assignments';

  @override
  String get refreshCode => 'Refresh Code';

  @override
  String get businessType => 'Business Type';

  @override
  String get cardType => 'Card Type';

  @override
  String get phone => 'Phone';

  @override
  String get code => 'Code';

  @override
  String get cost => 'Cost';

  @override
  String get status => 'Status';

  @override
  String get time => 'Time';

  @override
  String get completed => 'Completed';

  @override
  String get expired => 'Expired';

  @override
  String get apiKey => 'API Key';

  @override
  String get lastLogin => 'Last Login';

  @override
  String get regIp => 'Registration IP';

  @override
  String get changePassword => 'Change Password';

  @override
  String get logout => 'Logout';

  @override
  String get language => 'Language';

  @override
  String get theme => 'Theme';

  @override
  String get light => 'Light';

  @override
  String get dark => 'Dark';

  @override
  String get loading => 'Loading';

  @override
  String get pleaseWait => 'Please wait...';

  @override
  String get error => 'Error';

  @override
  String get noData => 'No Data';

  @override
  String get noDataMessage => 'There is no content to display.';

  @override
  String get back => 'Back';

  @override
  String get add => 'Add';

  @override
  String get delete => 'Delete';

  @override
  String get ipAddress => 'IP Address';

  @override
  String get notes => 'Notes';

  @override
  String get usernameRequired => 'Please enter username';

  @override
  String get passwordRequired => 'Please enter password';

  @override
  String get emailRequired => 'Please enter email';

  @override
  String get noAccountRegister => 'No account? Register now';

  @override
  String get getPhoneTitle => 'Assign Phone Number';

  @override
  String get getCodeTitle => 'Get Verification Code';

  @override
  String get selectBusinessType => 'Select Business Type';

  @override
  String get selectCardType => 'Select Card Type';

  @override
  String get physicalCard => 'Physical Card';

  @override
  String get virtualCard => 'Virtual Card';

  @override
  String get assignNow => 'Assign Now';

  @override
  String get assignSuccess => 'Assigned Successfully';

  @override
  String get phoneNumber => 'Phone Number';

  @override
  String get enterPhoneNumber => 'Enter phone number';

  @override
  String get getCodeNow => 'Get Code Now';

  @override
  String get codeReceived => 'Code Received';

  @override
  String get copyPhone => 'Copy Phone';

  @override
  String get copyCode => 'Copy Code';

  @override
  String get copied => 'Copied';

  @override
  String get businessTypeRequired => 'Please select business type';

  @override
  String get cardTypeRequired => 'Please select card type';

  @override
  String get phoneRequired => 'Please enter phone number';

  @override
  String get gettingCode => 'Getting code, please wait...';

  @override
  String get codeTimeout => 'Code retrieval timeout';

  @override
  String get noBusinessTypes => 'No business types available';

  @override
  String get loadingBusinessTypes => 'Loading business types...';

  @override
  String get assignmentResult => 'Assignment Result';

  @override
  String get verificationCode => 'Verification Code';

  @override
  String get waitingForCode => 'Waiting for code...';

  @override
  String get noCodeReceived => 'No code received';

  @override
  String get retryGetCode => 'Retry';

  @override
  String get requestCount => 'Request Count';

  @override
  String get requestCountHelper => 'Maximum 10 numbers per request';

  @override
  String get totalCostLabel => 'Total Cost';

  @override
  String get remainingBalanceLabel => 'Remaining Balance';

  @override
  String get successCountLabel => 'Success';

  @override
  String get failedCountLabel => 'Failed';

  @override
  String get pendingCountLabel => 'Pending';

  @override
  String get countryCode => 'Country Code';

  @override
  String get providerLabel => 'Provider';

  @override
  String get validUntilLabel => 'Valid Until';

  @override
  String get codesSummary => 'Verification Summary';

  @override
  String get enterPhonesHint => 'Enter multiple phone numbers separated by newline or comma';

  @override
  String get enterPhonesHelper => 'Up to 10 phone numbers per request';

  @override
  String get invalidPhoneInput => 'Please enter between 1 and 10 valid phone numbers';

  @override
  String get noCodesYet => 'No verification codes yet. Keep polling if status is pending.';

  @override
  String get messageLabel => 'Message';

  @override
  String get receivedAtLabel => 'Received At';

  @override
  String get statusSuccess => 'Success';

  @override
  String get statusPending => 'Pending';

  @override
  String get statusActive => 'Active';

  @override
  String get statusFailed => 'Failed';

  @override
  String get getCodeInstructions => 'Enter up to 10 phone numbers. If the status is pending, continue polling this interface.';

  @override
  String get currentPasswordLabel => 'Current Password';

  @override
  String get newPasswordLabel => 'New Password';

  @override
  String get confirmPasswordLabel => 'Confirm Password';

  @override
  String get passwordTooShort => 'Password must be at least 6 characters';

  @override
  String get passwordUnchanged => 'New password must be different from current password';

  @override
  String get passwordMismatch => 'Passwords do not match';

  @override
  String get changePasswordSuccess => 'Password changed successfully';

  @override
  String get changePasswordFailed => 'Failed to change password';

  @override
  String get addWhitelistSuccess => 'Whitelist entry added';

  @override
  String get addWhitelistFailed => 'Failed to add whitelist entry';

  @override
  String get deleteWhitelistSuccess => 'Whitelist entry deleted';

  @override
  String get deleteWhitelistFailed => 'Failed to delete whitelist entry';

  @override
  String get deleteWhitelistConfirm => 'Delete the following IP?';

  @override
  String get notesOptional => 'Notes (optional)';

  @override
  String get ipAddressRequired => 'Please enter IP address or CIDR';

  @override
  String get invalidIpFormat => 'Invalid IP or CIDR format';

  @override
  String get viewAll => 'View All';

  @override
  String get recentRecords => 'Recent Records';

  @override
  String get pullDownToRefresh => 'Pull down to refresh, scroll up to load more';

  @override
  String get noMoreData => 'No more data';

  @override
  String get loadMore => 'Load More';

  @override
  String get copy => 'Copy';

  @override
  String totalRecords(int count) {
    return 'Total $count records';
  }

  @override
  String get refresh => 'Refresh';

  @override
  String get loadingFailed => 'Loading failed';

  @override
  String get noAvailableBusinessTypes => 'No business types available, please contact administrator';

  @override
  String get getBusinessTypesFailed => 'Failed to get business types';

  @override
  String get getFailed => 'Failed to get';

  @override
  String get registerFailed => 'Registration failed';

  @override
  String get registerSuccess => 'Registration successful! Please login';
}
