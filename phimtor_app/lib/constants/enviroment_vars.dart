abstract class Constants {
  static const String apiUrl = String.fromEnvironment(
    "API_URL",
    defaultValue: 'http://localhost:8080',
  );
}
