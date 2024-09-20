import 'package:file_picker/file_picker.dart';
import 'package:flutter/material.dart';
import 'package:phimtor_app/constants/enviroment_vars.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/locale_provider.dart';
import 'package:phimtor_app/services/preferences/preferences_service.dart';
import 'package:phimtor_app/services/updater/updater_service.dart';
import 'package:phimtor_app/utilities/dialogs/updater_dialog.dart';
import 'package:provider/provider.dart';

class SettingsView extends StatefulWidget {
  const SettingsView({super.key});

  @override
  State<SettingsView> createState() => _SettingsViewState();
}

class _SettingsViewState extends State<SettingsView> {
  late String _dataDirPath;
  bool _isDataDirPathChanged = false;
  late bool _deleteAfterClose;

  @override
  void initState() {
    super.initState();
    _dataDirPath = PreferencesService.getInstance().dataDirPath;
    _deleteAfterClose = PreferencesService.getInstance().deleteAfterClose;
  }

  void updateDataDirPath() async {
    String? selectedDirectory = await FilePicker.platform.getDirectoryPath();
    if (selectedDirectory != null) {
      await PreferencesService.getInstance().setDataDirPath(selectedDirectory);
      setState(() {
        _dataDirPath = selectedDirectory;
      });
      _isDataDirPathChanged = true;
    }
  }

  void updateDeleteAfterClose(bool value) async {
    await PreferencesService.getInstance().setDeleteAfterClose(value);
    setState(() {
      _deleteAfterClose = value;
    });
  }

  @override
  Widget build(BuildContext context) {
    final localeProvider = Provider.of<LocaleProvider>(context);

    return Scaffold(
      appBar: AppBar(
        title: Text(context.loc.setting_title),
      ),
      body: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          buildVersionSection(context),
          const SizedBox(height: 8),
          ListTile(
            title: Text(context.loc.setting_lang),
            subtitle: Text(getLanguageName(localeProvider.locale)),
            onTap: () {
              showDialog(
                context: context,
                builder: (context) {
                  return AlertDialog(
                    title: Text(context.loc.setting_lang),
                    content: Column(
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        ListTile(
                          title: Text(context.loc.setting_lang_en),
                          onTap: () {
                            localeProvider.setLocale(const Locale("en"));
                            Navigator.of(context).pop();
                          },
                        ),
                        ListTile(
                          title: Text(context.loc.setting_lang_vi),
                          onTap: () {
                            localeProvider.setLocale(const Locale("vi"));
                            Navigator.of(context).pop();
                          },
                        ),
                      ],
                    ),
                  );
                },
              );
            },
          ),
          ListTile(
            title: Text(context.loc.setting_data_dir),
            subtitle: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(_dataDirPath),
                if (_isDataDirPathChanged)
                  Text(
                    context.loc.setting_data_dir_changed_note,
                    style: Theme.of(context)
                        .textTheme
                        .labelMedium
                        ?.merge(const TextStyle(fontStyle: FontStyle.italic)),
                  ),
              ],
            ),
            onTap: updateDataDirPath,
          ),
          SwitchListTile(
            title: Text(context.loc.setting_delete_after_close),
            subtitle: Text(context.loc.setting_delete_after_close_note),
            value: _deleteAfterClose,
            onChanged: updateDeleteAfterClose,
          ),
        ],
      ),
    );
  }

  Widget buildVersionSection(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.fromLTRB(16, 0, 16, 8),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            "Version: ${Constants.appVersion}",
            style: Theme.of(context).textTheme.labelSmall!.merge(
                  const TextStyle(fontStyle: FontStyle.italic),
                ),
          ),
          if (UpdaterService().hasNewVersion) ...[
            const SizedBox(height: 8),
            ElevatedButton.icon(
              onPressed: () async {
                await showUpdaterDialog(context, UpdaterService().newVersion!);
              },
              icon: const Icon(Icons.update),
              label: Text(
                context.loc.has_new_version,
                style: Theme.of(context).textTheme.labelMedium,
              ),
            ),
          ],
        ],
      ),
    );
  }
}

String getLanguageName(Locale locale) {
  switch (locale.languageCode) {
    case "en":
      return "English";
    case "vi":
      return "Tiếng Việt";
    default:
      return "Unknown";
  }
}
