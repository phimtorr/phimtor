import 'package:file_picker/file_picker.dart';
import 'package:flutter/material.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/services/preferences/preferences_service.dart';

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
    return Scaffold(
      appBar: AppBar(
        title: Text(context.loc.setting_title),
      ),
      body: Column(
        children: [
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
            value: _deleteAfterClose,
            onChanged: updateDeleteAfterClose,
          ),
        ],
      ),
    );
  }
}
