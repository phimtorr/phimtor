abstract class Constants {
  static const String apiUrl = String.fromEnvironment(
    "API_URL",
    defaultValue: 'http://localhost:8080',
  );
  static const String appVersion = String.fromEnvironment(
    "APP_VERSION",
    defaultValue: '0.0.1',
  );
}
