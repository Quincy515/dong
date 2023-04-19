// ignore_for_file: unused_element

import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'flutter_flow/lat_lng.dart';

class FFAppState extends ChangeNotifier {
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

  void update(VoidCallback callback) {
    callback();
    notifyListeners();
  }

  late SharedPreferences prefs;

  String _xtoken = '';
  String get xtoken => _xtoken;
  set xtoken(String _value) {
    _xtoken = _value;
    prefs.setString('ff_xtoken', _value);
  }

  String _username = '';
  String get username => _username;
  set username(String _value) {
    _username = _value;
    prefs.setString('ff_username', _value);
  }

  String _password = '';
  String get password => _password;
  set password(String _value) {
    _password = _value;
    prefs.setString('ff_password', _value);
  }

  int _userID = 3;
  int get userID => _userID;
  set userID(int _value) {
    _userID = _value;
    prefs.setInt('ff_userID', _value);
  }

  String _nickName = '';
  String get nickName => _nickName;
  set nickName(String _value) {
    _nickName = _value;
    prefs.setString('ff_nickName', _value);
  }

  String _headerImg = '';
  String get headerImg => _headerImg;
  set headerImg(String _value) {
    _headerImg = _value;
    prefs.setString('ff_headerImg', _value);
  }

  String _authorityName = '';
  String get authorityName => _authorityName;
  set authorityName(String _value) {
    _authorityName = _value;
    prefs.setString('ff_authorityName', _value);
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
