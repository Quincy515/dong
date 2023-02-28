// ignore_for_file: unused_element

import 'package:shared_preferences/shared_preferences.dart';
import 'flutter_flow/lat_lng.dart';

class FFAppState {
  static final FFAppState _instance = FFAppState._internal();

  factory FFAppState() {
    return _instance;
  }

  FFAppState._internal() {
    initializePersistedState();
  }

  Future initializePersistedState() async {
    prefs = await SharedPreferences.getInstance();
    _xtoken = prefs.getString('ff_xtoken') ?? _xtoken;
  }

  late SharedPreferences prefs;

  String _xtoken = '';
  String get xtoken => _xtoken;
  set xtoken(String _value) {
    _xtoken = _value;
    prefs.setString('ff_xtoken', _value);
  }
}

LatLng? _latLngFromString(String? val) {
  if (val == null) {
    return null;
  }
  final split = val.split(',');
  final lat = double.parse(split.first);
  final lng = double.parse(split.last);
  return LatLng(lat, lng);
}
