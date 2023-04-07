// ignore_for_file: unused_import, deprecated_member_use

import 'dart:io';

import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:flutter/foundation.dart' show kIsWeb;
import 'package:flutter/material.dart';
import 'package:from_css_color/from_css_color.dart';
import 'package:intl/intl.dart';
import 'package:json_path/json_path.dart';
import 'package:timeago/timeago.dart' as timeago;
import 'package:url_launcher/url_launcher.dart';

import '../main.dart';

import 'lat_lng.dart';

export 'lat_lng.dart';
export 'place.dart';
export 'uploaded_file.dart';
export '../app_state.dart';
export 'flutter_flow_model.dart';
export 'dart:math' show min, max;
export 'dart:typed_data' show Uint8List;
export 'dart:convert' show jsonEncode, jsonDecode;
export 'package:intl/intl.dart';
export 'package:page_transition/page_transition.dart';
export 'nav/nav.dart';

T valueOrDefault<T>(T? value, T defaultValue) =>
    (value is String && value.isEmpty) || value == null ? defaultValue : value;

String dateTimeRfc3339({DateTime? dateTime}) {
  Timestamp timestamp = Timestamp.now();

  if (dateTime != null) {
    timestamp = Timestamp.fromDate(dateTime);
  }
  print('>>>timestamp: ${timestamp.millisecondsSinceEpoch}');
  var formattedDate = DateTime.fromMillisecondsSinceEpoch(
      timestamp.millisecondsSinceEpoch,
      isUtc: true);
  print('>>>formattedDate: $formattedDate');

  return formattedDate.toIso8601String();
}

String dateTimeFormat(String format, DateTime? dateTime, {String? locale}) {
  if (dateTime == null) {
    return '';
  }
  if (format == 'relative') {
    return timeago.format(dateTime, locale: locale);
  }
  return DateFormat(format).format(dateTime);
}

Future launchURL(String url) async {
  var uri = Uri.parse(url).toString();
  try {
    await launch(uri);
  } catch (e) {
    throw 'Could not launch $uri: $e';
  }
}

Color colorFromCssString(String color, {Color? defaultColor}) {
  try {
    return fromCssColor(color);
  } catch (_) {}
  return defaultColor ?? Colors.black;
}

enum FormatType {
  decimal,
  percent,
  scientific,
  compact,
  compactLong,
  custom,
}

enum DecimalType {
  automatic,
  periodDecimal,
  commaDecimal,
}

String formatNumber(
  num? value, {
  required FormatType formatType,
  DecimalType? decimalType,
  String? currency,
  bool toLowerCase = false,
  String? format,
  String? locale,
}) {
  if (value == null) {
    return '';
  }
  var formattedValue = '';
  switch (formatType) {
    case FormatType.decimal:
      switch (decimalType!) {
        case DecimalType.automatic:
          formattedValue = NumberFormat.decimalPattern().format(value);
          break;
        case DecimalType.periodDecimal:
          formattedValue = NumberFormat.decimalPattern('en_US').format(value);
          break;
        case DecimalType.commaDecimal:
          formattedValue = NumberFormat.decimalPattern('es_PA').format(value);
          break;
      }
      break;
    case FormatType.percent:
      formattedValue = NumberFormat.percentPattern().format(value);
      break;
    case FormatType.scientific:
      formattedValue = NumberFormat.scientificPattern().format(value);
      if (toLowerCase) {
        formattedValue = formattedValue.toLowerCase();
      }
      break;
    case FormatType.compact:
      formattedValue = NumberFormat.compact().format(value);
      break;
    case FormatType.compactLong:
      formattedValue = NumberFormat.compactLong().format(value);
      break;
    case FormatType.custom:
      final hasLocale = locale != null && locale.isNotEmpty;
      formattedValue =
          NumberFormat(format, hasLocale ? locale : null).format(value);
  }

  if (formattedValue.isEmpty) {
    return value.toString();
  }

  if (currency != null) {
    final currencySymbol = currency.isNotEmpty
        ? currency
        : NumberFormat.simpleCurrency().format(0.0).substring(0, 1);
    formattedValue = '$currencySymbol$formattedValue';
  }

  return formattedValue;
}

DateTime get getCurrentTimestamp => DateTime.now();

extension DateTimeComparisonOperators on DateTime {
  bool operator <(DateTime other) => isBefore(other);
  bool operator >(DateTime other) => isAfter(other);
  bool operator <=(DateTime other) => this < other || isAtSameMomentAs(other);
  bool operator >=(DateTime other) => this > other || isAtSameMomentAs(other);
}

dynamic getJsonField(
  dynamic response,
  String jsonPath, [
  bool isForList = false,
]) {
  final field = JsonPath(jsonPath).read(response);
  if (field.isEmpty) {
    return null;
  }
  if (field.length > 1) {
    return field.map((f) => f.value).toList();
  }
  final value = field.first.value;
  return isForList && value is! Iterable ? [value] : value;
}

Rect? getWidgetBoundingBox(BuildContext context) {
  try {
    final renderBox = context.findRenderObject() as RenderBox?;
    return renderBox!.localToGlobal(Offset.zero) & renderBox.size;
  } catch (_) {
    return null;
  }
}

bool get isAndroid => !kIsWeb && Platform.isAndroid;
bool get isiOS => !kIsWeb && Platform.isIOS;
bool get isWeb => kIsWeb;

const kMobileWidthCutoff = 479.0;
bool isMobileWidth(BuildContext context) =>
    MediaQuery.of(context).size.width < kMobileWidthCutoff;
bool responsiveVisibility({
  required BuildContext context,
  bool phone = true,
  bool tablet = true,
  bool tabletLandscape = true,
  bool desktop = true,
}) {
  final width = MediaQuery.of(context).size.width;
  if (width < kMobileWidthCutoff) {
    return phone;
  } else if (width < 767) {
    return tablet;
  } else if (width < 991) {
    return tabletLandscape;
  } else {
    return desktop;
  }
}

const kTextValidatorUsernameRegex = r'^[a-zA-Z][a-zA-Z0-9_-]{2,16}$';
const kTextValidatorEmailRegex =
    r"^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,253}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,253}[a-zA-Z0-9])?)*$";
const kTextValidatorWebsiteRegex =
    r'(https?:\/\/)?(www\.)[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,10}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)|(https?:\/\/)?(www\.)?(?!ww)[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,10}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)';

extension FFTextEditingControllerExt on TextEditingController? {
  String get text => this == null ? '' : this!.text;
  set text(String newText) => this?.text = newText;
}

extension IterableExt<T> on Iterable<T> {
  List<S> mapIndexed<S>(S Function(int, T) func) => toList()
      .asMap()
      .map((index, value) => MapEntry(index, func(index, value)))
      .values
      .toList();
}

void setAppLanguage(BuildContext context, String language) =>
    MyApp.of(context).setLocale(language);

void setDarkModeSetting(BuildContext context, ThemeMode themeMode) =>
    MyApp.of(context).setThemeMode(themeMode);

void showSnackbar(
  BuildContext context,
  String message, {
  bool loading = false,
  int duration = 4,
}) {
  ScaffoldMessenger.of(context).hideCurrentSnackBar();
  ScaffoldMessenger.of(context).showSnackBar(
    SnackBar(
      content: Row(
        children: [
          if (loading)
            Padding(
              padding: EdgeInsetsDirectional.only(end: 10.0),
              child: Container(
                height: 20,
                width: 20,
                child: const CircularProgressIndicator(
                  color: Colors.white,
                ),
              ),
            ),
          Text(message),
        ],
      ),
      duration: Duration(seconds: duration),
    ),
  );
}

extension FFStringExt on String {
  String maybeHandleOverflow({int? maxChars, String replacement = ''}) =>
      maxChars != null && length > maxChars
          ? replaceRange(maxChars, null, replacement)
          : this;
}

extension ListFilterExt<T> on Iterable<T?> {
  List<T> get withoutNulls => where((s) => s != null).map((e) => e!).toList();
}

// 校验身份证合法性
bool verifyCardId(String cardId) {
  const Map city = {
    11: "北京",
    12: "天津",
    13: "河北",
    14: "山西",
    15: "内蒙古",
    21: "辽宁",
    22: "吉林",
    23: "黑龙江 ",
    31: "上海",
    32: "江苏",
    33: "浙江",
    34: "安徽",
    35: "福建",
    36: "江西",
    37: "山东",
    41: "河南",
    42: "湖北 ",
    43: "湖南",
    44: "广东",
    45: "广西",
    46: "海南",
    50: "重庆",
    51: "四川",
    52: "贵州",
    53: "云南",
    54: "西藏 ",
    61: "陕西",
    62: "甘肃",
    63: "青海",
    64: "宁夏",
    65: "新疆",
    71: "台湾",
    81: "香港",
    82: "澳门",
    91: "国外 "
  };
  String tip = '';
  bool pass = true;

  RegExp cardReg = RegExp(
      r'^\d{6}(18|19|20)?\d{2}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])\d{3}(\d|X)$');
  if (cardId.isEmpty || !cardReg.hasMatch(cardId)) {
    tip = '身份证号格式错误';
    print(tip);
    pass = false;
    return pass;
  }
  if (city[int.parse(cardId.substring(0, 2))] == null) {
    tip = '地址编码错误';
    print(tip);
    pass = false;
    return pass;
  }
  // 18位身份证需要验证最后一位校验位，15位不检测了，现在也没15位的了
  if (cardId.length == 18) {
    List numList = cardId.split('');
    //∑(ai×Wi)(mod 11)
    //加权因子
    List factor = [7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2];
    //校验位
    List parity = [1, 0, 'X', 9, 8, 7, 6, 5, 4, 3, 2];
    int sum = 0;
    int ai = 0;
    int wi = 0;
    for (var i = 0; i < 17; i++) {
      ai = int.parse(numList[i]);
      wi = factor[i];
      sum += ai * wi;
    }
    if (parity[sum % 11].toString() != numList[17]) {
      tip = "校验位错误";
      print(tip);
      pass = false;
    }
  } else {
    tip = '身份证号不是18位';
    print(tip);
    pass = false;
  }
//  print('证件格式$pass');
  return pass;
}

// 根据身份证号获取年龄
int getAgeFromCardId(String cardId) {
  bool isRight = verifyCardId(cardId);
  if (!isRight) {
    return 0;
  }
  int len = (cardId + "").length;
  String strBirthday = "";
  if (len == 18) {
    //处理18位的身份证号码从号码中得到生日和性别代码
    strBirthday = cardId.substring(6, 10) +
        "-" +
        cardId.substring(10, 12) +
        "-" +
        cardId.substring(12, 14);
  }
  if (len == 15) {
    strBirthday = "19" +
        cardId.substring(6, 8) +
        "-" +
        cardId.substring(8, 10) +
        "-" +
        cardId.substring(10, 12);
  }
  int age = getAgeFromBirthday(strBirthday);
  return age;
}

// 根据出生日期获取年龄
int getAgeFromBirthday(String strBirthday) {
  if (strBirthday.isEmpty) {
    print('生日错误');
    return 0;
  }
  DateTime birth = DateTime.parse(strBirthday);
  DateTime now = DateTime.now();

  int age = now.year - birth.year;
  //再考虑月、天的因素
  if (now.month < birth.month ||
      (now.month == birth.month && now.day < birth.day)) {
    age--;
  }
  return age;
}

// 根据身份证获取性别
String getSexFromCardId(String cardId) {
  String sex = "";
  bool isRight = verifyCardId(cardId);
  if (!isRight) {
    return sex;
  }
  if (cardId.length == 18) {
    if (int.parse(cardId.substring(16, 17)) % 2 == 1) {
      sex = "男";
    } else {
      sex = "女";
    }
  }
  if (cardId.length == 15) {
    if (int.parse(cardId.substring(14, 15)) % 2 == 1) {
      sex = "男";
    } else {
      sex = "女";
    }
  }
  return sex;
}

// 返回出生年月日
String getBirthday(String cardId) {
  bool isRight = verifyCardId(cardId);
  if (!isRight) {
    return "";
  }
  return cardId.substring(6, 17);
}
