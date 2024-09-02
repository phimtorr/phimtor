import 'package:ambilytics/ambilytics.dart' as ambilytics;
import 'package:phimtor_app/firebase_options.dart';

class AnalyticsService {
  static Future<void> initialize({String? userId}) async {
    await ambilytics.initAnalytics(
      measurementId: 'G-N2M2ZZ2LJG',
      apiSecret: 'LEoMKbjZReqsAnGcBckQnA',
      userId: userId,
      firebaseOptions: DefaultFirebaseOptions.currentPlatform,
    );
  }

  static void sendEvent({required String name, Map<String, Object>? parameters}) {
    ambilytics.sendEvent(name: name, parameters: parameters);
  }
}
