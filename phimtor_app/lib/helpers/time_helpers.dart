class TimeHelpers {
  static String toHumanReadableDuration(int durationInMinutes) {
    final duration = Duration(minutes: durationInMinutes);
    final hours = duration.inHours;
    final minutes = duration.inMinutes.remainder(60);
    final hoursStr = hours.toString().padLeft(2, '0');
    final minutesStr = minutes.toString().padLeft(2, '0');
    return "${hoursStr}h$minutesStr";
  }
}