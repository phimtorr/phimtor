import 'package:file_picker/file_picker.dart';
import 'package:flutter/material.dart';
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
      await PreferencesService.getInstance()
          .setDataDirPath(selectedDirectory);
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
        title: const Text('Settings'),
      ),
      body: Column(
        children: [
          ListTile(
            title: const Text('Data directory'),
            subtitle: Text(_dataDirPath +
                (_isDataDirPathChanged ? '\nNeed restart app to apply' : '')),
            onTap: updateDataDirPath,
          ),
          SwitchListTile(
            title: const Text('Delete after close'),
            value: _deleteAfterClose,
            onChanged: updateDeleteAfterClose,
          ),
        ],
      ),
    );
  }
}
