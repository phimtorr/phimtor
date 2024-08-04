import 'package:phimtor_app/constants/enviroment_vars.dart';
import 'package:phimtor_openapi_client/api.dart';

class PhimtorService {
  final DefaultApi _defaultApi;

  static final PhimtorService _instance = PhimtorService._internal();

  factory PhimtorService() {
    return _instance;
  }

  PhimtorService._internal()
      : _defaultApi =
            DefaultApi(ApiClient(basePath: '${Constants.apiUrl}/api/v1'));

  DefaultApi get defaultApi {
    return _defaultApi;
  }
}
