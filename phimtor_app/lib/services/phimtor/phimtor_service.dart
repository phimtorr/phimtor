import 'package:phimtor_app/constants/enviroment_vars.dart';
import 'package:phimtor_app/services/auth/auth_service.dart';
import 'package:phimtor_openapi_client/api.dart' as phimtor_api;

class PhimtorService {
  final phimtor_api.DefaultApi _defaultApi;

  static final PhimtorService _instance = PhimtorService._internal();

  factory PhimtorService() {
    return _instance;
  }

  PhimtorService._internal()
      : _defaultApi = phimtor_api.DefaultApi(
          phimtor_api.ApiClient(
            basePath: '${Constants.apiUrl}/api/v1',
            authentication: _FirebaseAuthentication(),
          ),
        );

  phimtor_api.DefaultApi get defaultApi {
    return _defaultApi;
  }
}

class _FirebaseAuthentication implements phimtor_api.Authentication {
  @override
  Future<void> applyToParams(List<phimtor_api.QueryParam> queryParams,
      Map<String, String> headerParams) async {
    final token = await AuthService().authToken;
    if (token != null) {
      headerParams['Authorization'] = 'Bearer $token';
    }
  }
}
