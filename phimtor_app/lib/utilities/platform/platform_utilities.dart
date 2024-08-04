import 'dart:io';

class PlatformUtilities {
  static final bool isMacOS = Platform.isMacOS;
  static final bool isWindows = Platform.isWindows;
  static final bool isLinux = Platform.isLinux;
  static final bool isAndroid = Platform.isAndroid;
  static final bool isIOS = Platform.isIOS;
  static final bool isFuchsia = Platform.isFuchsia;

  static final bool isMobile = isAndroid || isIOS;
  static final bool isDesktop = isWindows || isLinux || isMacOS;
}