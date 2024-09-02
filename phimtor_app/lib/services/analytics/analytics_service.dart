import 'package:ambilytics/ambilytics.dart' as ambilytics;
import 'package:flutter/foundation.dart' show kIsWeb, defaultTargetPlatform;
import 'package:phimtor_app/constants/enviroment_vars.dart';
import 'package:phimtor_app/firebase_options.dart' show DefaultFirebaseOptions;

class AnalyticsService {
  // Singleton
  static final AnalyticsService _instance = AnalyticsService._();
  factory AnalyticsService() => _instance;
  AnalyticsService._();

  bool initialized = false;

  Future<void> initialize({String? userId}) async {
    await ambilytics.initAnalytics(
      measurementId: 'G-N2M2ZZ2LJG',
      apiSecret: 'LEoMKbjZReqsAnGcBckQnA',
      userId: userId,
      firebaseOptions: DefaultFirebaseOptions.currentPlatform,
      sendAppLaunch: false,
    );
    initialized = true;
    _sendAppLaunchEvent();
  }

  void sendEvent({required String name, Map<String, Object>? parameters}) {
    if (!initialized) {
      throw Exception('AnalyticsService not initialized');
    }
    ambilytics.sendEvent(name: name, parameters: parameters);
  }

  void _sendAppLaunchEvent() {
    final params = {
      'platform': kIsWeb ? 'web' : defaultTargetPlatform.name,
      'version': Constants.appVersion,
    };
    sendEvent(name: ambilytics.PredefinedEvents.appLaunch, parameters: params);
  }
}
