// ignore_for_file: unnecessary_question_mark, unnecessary_brace_in_string_interps, unused_import, unused_element

import 'dart:convert';
import 'dart:typed_data';

import '../../flutter_flow/flutter_flow_util.dart';

import 'api_manager.dart';

export 'api_manager.dart' show ApiCallResponse;

const _kPrivateApiFunctionName = 'ffPrivateApiCall';

/// Start base Group Code

class BaseGroup {
  static String baseUrl = 'http://119.45.234.161:8888';
  static Map<String, String> headers = {};
  static GetCaptchaCall getCaptchaCall = GetCaptchaCall();
  static LoginCall loginCall = LoginCall();
  static GetMenuCall getMenuCall = GetMenuCall();
  static AdminregisterCall adminregisterCall = AdminregisterCall();
}

class GetCaptchaCall {
  Future<ApiCallResponse> call() {
    return ApiManager.instance.makeApiCall(
      callName: 'getCaptcha',
      apiUrl: '${BaseGroup.baseUrl}/base/captcha',
      callType: ApiCallType.POST,
      headers: {
        ...BaseGroup.headers,
      },
      params: {},
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic captchaId(dynamic response) => getJsonField(
        response,
        r'''$.data.captchaId''',
      );
  dynamic picPath(dynamic response) => getJsonField(
        response,
        r'''$.data.picPath''',
      );
  dynamic captchaLength(dynamic response) => getJsonField(
        response,
        r'''$.data.captchaLength''',
      );
}

class LoginCall {
  Future<ApiCallResponse> call({
    String? username = '',
    String? password = '',
    String? captcha = '',
    String? captchaId = '',
  }) {
    final body = '''
{
  "username": "${username}",
  "password": "${password}",
  "captcha": "${captcha}",
  "captchaId": "${captchaId}"
}''';
    return ApiManager.instance.makeApiCall(
      callName: 'login',
      apiUrl: '${BaseGroup.baseUrl}/base/login',
      callType: ApiCallType.POST,
      headers: {
        ...BaseGroup.headers,
      },
      params: {},
      body: body,
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic token(dynamic response) => getJsonField(
        response,
        r'''$.data.token''',
      );
}

class GetMenuCall {
  Future<ApiCallResponse> call() {
    return ApiManager.instance.makeApiCall(
      callName: 'getMenu',
      apiUrl: '${BaseGroup.baseUrl}/menu/getMenu',
      callType: ApiCallType.POST,
      headers: {
        ...BaseGroup.headers,
        'x-token': FFAppState().xtoken,
      },
      params: {},
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic menus(dynamic response) => getJsonField(
        response,
        r'''$.data.menus''',
        true,
      );
  dynamic code(dynamic response) => getJsonField(
        response,
        r'''$.code''',
      );
  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class AdminregisterCall {
  Future<ApiCallResponse> call({
    String? userName = '',
    String? password = '',
    String? nickName = '',
    String? headerImg = '',
  }) {
    final body = '''
{
  "password": "${password}",
  "nickName": "${nickName}",
  "headerImg": "${headerImg}",
  "authorityId": 888,
  "authorityIds": [
    888,
    8881
  ],
  "enable": 1,
  "userName": "${userName}"
}''';
    return ApiManager.instance.makeApiCall(
      callName: 'adminregister',
      apiUrl: '${BaseGroup.baseUrl}/user/admin_register',
      callType: ApiCallType.POST,
      headers: {
        ...BaseGroup.headers,
        'x-token': FFAppState().xtoken,
      },
      params: {},
      body: body,
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
  dynamic code(dynamic response) => getJsonField(
        response,
        r'''$.code''',
      );
}

/// End base Group Code

/// Start user Group Code

class UserGroup {
  static String baseUrl = 'http://119.45.234.161:8888/user';
  static Map<String, String> headers = {
    'x-token': FFAppState().xtoken,
  };
  static GetUserListCall getUserListCall = GetUserListCall();
}

class GetUserListCall {
  Future<ApiCallResponse> call({
    int? page = 1,
    int? pageSize = 10,
  }) {
    final body = '''
{
  "page": ${page},
  "pageSize": ${pageSize}
}''';
    return ApiManager.instance.makeApiCall(
      callName: 'getUserList',
      apiUrl: '${UserGroup.baseUrl}/getUserList',
      callType: ApiCallType.POST,
      headers: {
        ...UserGroup.headers,
      },
      params: {},
      body: body,
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic list(dynamic response) => getJsonField(
        response,
        r'''$.data.list''',
        true,
      );
  dynamic total(dynamic response) => getJsonField(
        response,
        r'''$.data.total''',
      );
  dynamic page(dynamic response) => getJsonField(
        response,
        r'''$.data.page''',
      );
  dynamic pageSize(dynamic response) => getJsonField(
        response,
        r'''$.data.pageSize''',
      );
  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
  dynamic code(dynamic response) => getJsonField(
        response,
        r'''$.code''',
      );
}

/// End user Group Code

/// Start file Group Code

class FileGroup {
  static String baseUrl = 'http://119.45.234.161:8888';
  static Map<String, String> headers = {
    'x-token': FFAppState().xtoken,
  };
  static GetFileListCall getFileListCall = GetFileListCall();
}

class GetFileListCall {
  Future<ApiCallResponse> call({
    int? page = 1,
    int? pageSize = 10,
  }) {
    final body = '''
{
  "page": ${page},
  "pageSize": ${pageSize}
}''';
    return ApiManager.instance.makeApiCall(
      callName: 'getFileList',
      apiUrl: '${FileGroup.baseUrl}/fileUploadAndDownload/getFileList',
      callType: ApiCallType.POST,
      headers: {
        ...FileGroup.headers,
      },
      params: {},
      body: body,
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic code(dynamic response) => getJsonField(
        response,
        r'''$.code''',
      );
  dynamic list(dynamic response) => getJsonField(
        response,
        r'''$.data.list''',
        true,
      );
  dynamic total(dynamic response) => getJsonField(
        response,
        r'''$.data.total''',
      );
  dynamic page(dynamic response) => getJsonField(
        response,
        r'''$.data.page''',
      );
  dynamic pageSize(dynamic response) => getJsonField(
        response,
        r'''$.data.pageSize''',
      );
  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

/// End file Group Code

/// Start jingYingBu Group Code

class JingYingBuGroup {
  static String baseUrl = 'http://119.45.234.161:8888';
  static Map<String, String> headers = {
    'x-token': FFAppState().xtoken,
  };
  static GetXiangMuDengJiBiaoListCall getXiangMuDengJiBiaoListCall =
      GetXiangMuDengJiBiaoListCall();
  static DeleteXiangMuDengJiBiaoCall deleteXiangMuDengJiBiaoCall =
      DeleteXiangMuDengJiBiaoCall();
  static UpdateXiangMuDengJiBiaoCall updateXiangMuDengJiBiaoCall =
      UpdateXiangMuDengJiBiaoCall();
  static GetTouBiaoShuJuTongJiListCall getTouBiaoShuJuTongJiListCall =
      GetTouBiaoShuJuTongJiListCall();
  static DeleteTouBiaoShuJuTongJiCall deleteTouBiaoShuJuTongJiCall =
      DeleteTouBiaoShuJuTongJiCall();
  static UpdateTouBiaoShuJuTongJiCall updateTouBiaoShuJuTongJiCall =
      UpdateTouBiaoShuJuTongJiCall();
}

class GetXiangMuDengJiBiaoListCall {
  Future<ApiCallResponse> call({
    int? page = 1,
    int? pageSize = 10,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getXiangMuDengJiBiaoList',
      apiUrl:
          '${JingYingBuGroup.baseUrl}/xiangMuDengJiBiao/getXiangMuDengJiBiaoList',
      callType: ApiCallType.GET,
      headers: {
        ...JingYingBuGroup.headers,
      },
      params: {
        'page': page,
        'pageSize': pageSize,
      },
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic list(dynamic response) => getJsonField(
        response,
        r'''$.data.list''',
        true,
      );
  dynamic total(dynamic response) => getJsonField(
        response,
        r'''$.data.total''',
      );
  dynamic page(dynamic response) => getJsonField(
        response,
        r'''$.data.page''',
      );
  dynamic pageSize(dynamic response) => getJsonField(
        response,
        r'''$.data.pageSize''',
      );
  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class DeleteXiangMuDengJiBiaoCall {
  Future<ApiCallResponse> call({
    int? id = 0,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteXiangMuDengJiBiao',
      apiUrl:
          '${JingYingBuGroup.baseUrl}/xiangMuDengJiBiao/deleteXiangMuDengJiBiao',
      callType: ApiCallType.DELETE,
      headers: {
        ...JingYingBuGroup.headers,
      },
      params: {},
      body: body,
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class UpdateXiangMuDengJiBiaoCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateXiangMuDengJiBiao',
      apiUrl:
          '${JingYingBuGroup.baseUrl}/xiangMuDengJiBiao/updateXiangMuDengJiBiao',
      callType: ApiCallType.PUT,
      headers: {
        ...JingYingBuGroup.headers,
      },
      params: {},
      body: body,
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class GetTouBiaoShuJuTongJiListCall {
  Future<ApiCallResponse> call({
    int? page = 1,
    int? pageSize = 10,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getTouBiaoShuJuTongJiList',
      apiUrl:
          '${JingYingBuGroup.baseUrl}/touBiaoShuJuTongJi/getTouBiaoShuJuTongJiList',
      callType: ApiCallType.GET,
      headers: {
        ...JingYingBuGroup.headers,
      },
      params: {
        'page': page,
        'pageSize': pageSize,
      },
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic list(dynamic response) => getJsonField(
        response,
        r'''$.data.list''',
        true,
      );
  dynamic total(dynamic response) => getJsonField(
        response,
        r'''$.data.total''',
      );
  dynamic page(dynamic response) => getJsonField(
        response,
        r'''$.data.page''',
      );
  dynamic pageSize(dynamic response) => getJsonField(
        response,
        r'''$.data.pageSize''',
      );
  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class DeleteTouBiaoShuJuTongJiCall {
  Future<ApiCallResponse> call({
    int? id = 0,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteTouBiaoShuJuTongJi',
      apiUrl:
          '${JingYingBuGroup.baseUrl}/touBiaoShuJuTongJi/deleteTouBiaoShuJuTongJi',
      callType: ApiCallType.DELETE,
      headers: {
        ...JingYingBuGroup.headers,
      },
      params: {},
      body: body,
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class UpdateTouBiaoShuJuTongJiCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateTouBiaoShuJuTongJi',
      apiUrl:
          '${JingYingBuGroup.baseUrl}/touBiaoShuJuTongJi/updateTouBiaoShuJuTongJi',
      callType: ApiCallType.PUT,
      headers: {
        ...JingYingBuGroup.headers,
      },
      params: {},
      body: body,
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

/// End jingYingBu Group Code

/// Start banGongShi Group Code

class BanGongShiGroup {
  static String baseUrl = 'http://119.45.234.161:8888';
  static Map<String, String> headers = {
    'x-token': FFAppState().xtoken,
  };
  static GetYuanGongHuaMingCeListCall getYuanGongHuaMingCeListCall =
      GetYuanGongHuaMingCeListCall();
  static DeleteYuanGongHuaMingCeCall deleteYuanGongHuaMingCeCall =
      DeleteYuanGongHuaMingCeCall();
  static UpdateYuanGongHuaMingCeCall updateYuanGongHuaMingCeCall =
      UpdateYuanGongHuaMingCeCall();
  static GetMeiRiKaoQinListCall getMeiRiKaoQinListCall =
      GetMeiRiKaoQinListCall();
  static DeleteMeiRiKaoQinCall deleteMeiRiKaoQinCall = DeleteMeiRiKaoQinCall();
  static UpdateMeiRiKaoQinCall updateMeiRiKaoQinCall = UpdateMeiRiKaoQinCall();
  static GetKaoQinTongJiListCall getKaoQinTongJiListCall =
      GetKaoQinTongJiListCall();
  static DeleteKaoQinTongJiCall deleteKaoQinTongJiCall =
      DeleteKaoQinTongJiCall();
  static UpdateKaoQinTongJiCall updateKaoQinTongJiCall =
      UpdateKaoQinTongJiCall();
  static GetQingJiaGuanLiListCall getQingJiaGuanLiListCall =
      GetQingJiaGuanLiListCall();
  static DeleteQingJiaGuanLiCall deleteQingJiaGuanLiCall =
      DeleteQingJiaGuanLiCall();
  static UpdateQingJiaGuanLiCall updateQingJiaGuanLiCall =
      UpdateQingJiaGuanLiCall();
  static GetHuiYuanGuanLiListCall getHuiYuanGuanLiListCall =
      GetHuiYuanGuanLiListCall();
  static DeleteHuiYuanGuanLiCall deleteHuiYuanGuanLiCall =
      DeleteHuiYuanGuanLiCall();
  static UpdateHuiYuanGuanLiCall updateHuiYuanGuanLiCall =
      UpdateHuiYuanGuanLiCall();
  static GetGongSiZhangHaoMiMaListCall getGongSiZhangHaoMiMaListCall =
      GetGongSiZhangHaoMiMaListCall();
  static DeleteGongSiZhangHaoMiMaCall deleteGongSiZhangHaoMiMaCall =
      DeleteGongSiZhangHaoMiMaCall();
  static UpdateGongSiZhangHaoMiMaCall updateGongSiZhangHaoMiMaCall =
      UpdateGongSiZhangHaoMiMaCall();
  static GetGongZhangDengJiListCall getGongZhangDengJiListCall =
      GetGongZhangDengJiListCall();
  static DeleteGongZhangDengJiCall deleteGongZhangDengJiCall =
      DeleteGongZhangDengJiCall();
  static UpdateGongZhangDengJiCall updateGongZhangDengJiCall =
      UpdateGongZhangDengJiCall();
}

class GetYuanGongHuaMingCeListCall {
  Future<ApiCallResponse> call({
    int? page = 1,
    int? pageSize = 10,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getYuanGongHuaMingCeList',
      apiUrl:
          '${BanGongShiGroup.baseUrl}/yuanGongHuaMingCe/getYuanGongHuaMingCeList',
      callType: ApiCallType.GET,
      headers: {
        ...BanGongShiGroup.headers,
      },
      params: {
        'page': page,
        'pageSize': pageSize,
      },
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic list(dynamic response) => getJsonField(
        response,
        r'''$.data.list''',
        true,
      );
  dynamic total(dynamic response) => getJsonField(
        response,
        r'''$.data.total''',
      );
  dynamic page(dynamic response) => getJsonField(
        response,
        r'''$.data.page''',
      );
  dynamic pageSize(dynamic response) => getJsonField(
        response,
        r'''$.data.pageSize''',
      );
  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class DeleteYuanGongHuaMingCeCall {
  Future<ApiCallResponse> call({
    int? id = 0,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteYuanGongHuaMingCe',
      apiUrl:
          '${BanGongShiGroup.baseUrl}/yuanGongHuaMingCe/deleteYuanGongHuaMingCe',
      callType: ApiCallType.DELETE,
      headers: {
        ...BanGongShiGroup.headers,
      },
      params: {},
      body: body,
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class UpdateYuanGongHuaMingCeCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateYuanGongHuaMingCe',
      apiUrl:
          '${BanGongShiGroup.baseUrl}/yuanGongHuaMingCe/updateYuanGongHuaMingCe',
      callType: ApiCallType.PUT,
      headers: {
        ...BanGongShiGroup.headers,
      },
      params: {},
      body: body,
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class GetMeiRiKaoQinListCall {
  Future<ApiCallResponse> call({
    int? page = 1,
    int? pageSize = 10,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getMeiRiKaoQinList',
      apiUrl: '${BanGongShiGroup.baseUrl}/meiRiKaoQin/getMeiRiKaoQinList',
      callType: ApiCallType.GET,
      headers: {
        ...BanGongShiGroup.headers,
      },
      params: {
        'page': page,
        'pageSize': pageSize,
      },
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic list(dynamic response) => getJsonField(
        response,
        r'''$.data.list''',
        true,
      );
  dynamic total(dynamic response) => getJsonField(
        response,
        r'''$.data.total''',
      );
  dynamic page(dynamic response) => getJsonField(
        response,
        r'''$.data.page''',
      );
  dynamic pageSize(dynamic response) => getJsonField(
        response,
        r'''$.data.pageSize''',
      );
  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class DeleteMeiRiKaoQinCall {
  Future<ApiCallResponse> call({
    int? id = 0,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteMeiRiKaoQin',
      apiUrl: '${BanGongShiGroup.baseUrl}/meiRiKaoQin/deleteMeiRiKaoQin',
      callType: ApiCallType.DELETE,
      headers: {
        ...BanGongShiGroup.headers,
      },
      params: {},
      body: body,
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class UpdateMeiRiKaoQinCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateMeiRiKaoQin',
      apiUrl: '${BanGongShiGroup.baseUrl}/meiRiKaoQin/updateMeiRiKaoQin',
      callType: ApiCallType.PUT,
      headers: {
        ...BanGongShiGroup.headers,
      },
      params: {},
      body: body,
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class GetKaoQinTongJiListCall {
  Future<ApiCallResponse> call({
    int? page = 1,
    int? pageSize = 10,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getKaoQinTongJiList',
      apiUrl: '${BanGongShiGroup.baseUrl}/kaoQinTongJi/getKaoQinTongJiList',
      callType: ApiCallType.GET,
      headers: {
        ...BanGongShiGroup.headers,
      },
      params: {
        'page': page,
        'pageSize': pageSize,
      },
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic list(dynamic response) => getJsonField(
        response,
        r'''$.data.list''',
        true,
      );
  dynamic total(dynamic response) => getJsonField(
        response,
        r'''$.data.total''',
      );
  dynamic page(dynamic response) => getJsonField(
        response,
        r'''$.data.page''',
      );
  dynamic pageSize(dynamic response) => getJsonField(
        response,
        r'''$.data.pageSize''',
      );
  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class DeleteKaoQinTongJiCall {
  Future<ApiCallResponse> call({
    int? id = 0,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteKaoQinTongJi',
      apiUrl: '${BanGongShiGroup.baseUrl}/kaoQinTongJi/deleteKaoQinTongJi',
      callType: ApiCallType.DELETE,
      headers: {
        ...BanGongShiGroup.headers,
      },
      params: {},
      body: body,
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class UpdateKaoQinTongJiCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateKaoQinTongJi',
      apiUrl: '${BanGongShiGroup.baseUrl}/kaoQinTongJi/updateKaoQinTongJi',
      callType: ApiCallType.PUT,
      headers: {
        ...BanGongShiGroup.headers,
      },
      params: {},
      body: body,
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class GetQingJiaGuanLiListCall {
  Future<ApiCallResponse> call({
    int? page = 1,
    int? pageSize = 10,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getQingJiaGuanLiList',
      apiUrl: '${BanGongShiGroup.baseUrl}/qingJiaGuanLi/getQingJiaGuanLiList',
      callType: ApiCallType.GET,
      headers: {
        ...BanGongShiGroup.headers,
      },
      params: {
        'page': page,
        'pageSize': pageSize,
      },
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic list(dynamic response) => getJsonField(
        response,
        r'''$.data.list''',
        true,
      );
  dynamic total(dynamic response) => getJsonField(
        response,
        r'''$.data.total''',
      );
  dynamic page(dynamic response) => getJsonField(
        response,
        r'''$.data.page''',
      );
  dynamic pageSize(dynamic response) => getJsonField(
        response,
        r'''$.data.pageSize''',
      );
  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class DeleteQingJiaGuanLiCall {
  Future<ApiCallResponse> call({
    int? id = 0,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteQingJiaGuanLi',
      apiUrl: '${BanGongShiGroup.baseUrl}/qingJiaGuanLi/deleteQingJiaGuanLi',
      callType: ApiCallType.DELETE,
      headers: {
        ...BanGongShiGroup.headers,
      },
      params: {},
      body: body,
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class UpdateQingJiaGuanLiCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateQingJiaGuanLi',
      apiUrl: '${BanGongShiGroup.baseUrl}/qingJiaGuanLi/updateQingJiaGuanLi',
      callType: ApiCallType.PUT,
      headers: {
        ...BanGongShiGroup.headers,
      },
      params: {},
      body: body,
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class GetHuiYuanGuanLiListCall {
  Future<ApiCallResponse> call({
    int? page = 1,
    int? pageSize = 10,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getHuiYuanGuanLiList',
      apiUrl: '${BanGongShiGroup.baseUrl}/huiYuanGuanLi/getHuiYuanGuanLiList',
      callType: ApiCallType.GET,
      headers: {
        ...BanGongShiGroup.headers,
      },
      params: {
        'page': page,
        'pageSize': pageSize,
      },
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic list(dynamic response) => getJsonField(
        response,
        r'''$.data.list''',
        true,
      );
  dynamic total(dynamic response) => getJsonField(
        response,
        r'''$.data.total''',
      );
  dynamic page(dynamic response) => getJsonField(
        response,
        r'''$.data.page''',
      );
  dynamic pageSize(dynamic response) => getJsonField(
        response,
        r'''$.data.pageSize''',
      );
  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class DeleteHuiYuanGuanLiCall {
  Future<ApiCallResponse> call({
    int? id = 0,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteHuiYuanGuanLi',
      apiUrl: '${BanGongShiGroup.baseUrl}/huiYuanGuanLi/deleteHuiYuanGuanLi',
      callType: ApiCallType.DELETE,
      headers: {
        ...BanGongShiGroup.headers,
      },
      params: {},
      body: body,
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class UpdateHuiYuanGuanLiCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateHuiYuanGuanLi',
      apiUrl: '${BanGongShiGroup.baseUrl}/huiYuanGuanLi/updateHuiYuanGuanLi',
      callType: ApiCallType.PUT,
      headers: {
        ...BanGongShiGroup.headers,
      },
      params: {},
      body: body,
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class GetGongSiZhangHaoMiMaListCall {
  Future<ApiCallResponse> call({
    int? page = 1,
    int? pageSize = 10,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getGongSiZhangHaoMiMaList',
      apiUrl:
          '${BanGongShiGroup.baseUrl}/gongSiZhangHaoMiMa/getGongSiZhangHaoMiMaList',
      callType: ApiCallType.GET,
      headers: {
        ...BanGongShiGroup.headers,
      },
      params: {
        'page': page,
        'pageSize': pageSize,
      },
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic list(dynamic response) => getJsonField(
        response,
        r'''$.data.list''',
        true,
      );
  dynamic total(dynamic response) => getJsonField(
        response,
        r'''$.data.total''',
      );
  dynamic page(dynamic response) => getJsonField(
        response,
        r'''$.data.page''',
      );
  dynamic pageSize(dynamic response) => getJsonField(
        response,
        r'''$.data.pageSize''',
      );
  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class DeleteGongSiZhangHaoMiMaCall {
  Future<ApiCallResponse> call({
    int? id = 0,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteGongSiZhangHaoMiMa',
      apiUrl:
          '${BanGongShiGroup.baseUrl}/gongSiZhangHaoMiMa/deleteGongSiZhangHaoMiMa',
      callType: ApiCallType.DELETE,
      headers: {
        ...BanGongShiGroup.headers,
      },
      params: {},
      body: body,
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class UpdateGongSiZhangHaoMiMaCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateGongSiZhangHaoMiMa',
      apiUrl:
          '${BanGongShiGroup.baseUrl}/gongSiZhangHaoMiMa/updateGongSiZhangHaoMiMa',
      callType: ApiCallType.PUT,
      headers: {
        ...BanGongShiGroup.headers,
      },
      params: {},
      body: body,
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class GetGongZhangDengJiListCall {
  Future<ApiCallResponse> call({
    int? page = 1,
    int? pageSize = 10,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getGongZhangDengJiList',
      apiUrl:
          '${BanGongShiGroup.baseUrl}/gongZhangDengJi/getGongZhangDengJiList',
      callType: ApiCallType.GET,
      headers: {
        ...BanGongShiGroup.headers,
      },
      params: {
        'page': page,
        'pageSize': pageSize,
      },
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic list(dynamic response) => getJsonField(
        response,
        r'''$.data.list''',
        true,
      );
  dynamic total(dynamic response) => getJsonField(
        response,
        r'''$.data.total''',
      );
  dynamic page(dynamic response) => getJsonField(
        response,
        r'''$.data.page''',
      );
  dynamic pageSize(dynamic response) => getJsonField(
        response,
        r'''$.data.pageSize''',
      );
  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class DeleteGongZhangDengJiCall {
  Future<ApiCallResponse> call({
    int? id = 0,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteGongZhangDengJi',
      apiUrl:
          '${BanGongShiGroup.baseUrl}/gongZhangDengJi/deleteGongZhangDengJi',
      callType: ApiCallType.DELETE,
      headers: {
        ...BanGongShiGroup.headers,
      },
      params: {},
      body: body,
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class UpdateGongZhangDengJiCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateGongZhangDengJi',
      apiUrl:
          '${BanGongShiGroup.baseUrl}/gongZhangDengJi/updateGongZhangDengJi',
      callType: ApiCallType.PUT,
      headers: {
        ...BanGongShiGroup.headers,
      },
      params: {},
      body: body,
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

/// End banGongShi Group Code

class ApiPagingParams {
  int nextPageNumber = 0;
  int numItems = 0;
  dynamic lastResponse;

  ApiPagingParams({
    required this.nextPageNumber,
    required this.numItems,
    required this.lastResponse,
  });

  @override
  String toString() =>
      'PagingParams(nextPageNumber: $nextPageNumber, numItems: $numItems, lastResponse: $lastResponse,)';
}

String _serializeList(List? list) {
  list ??= <String>[];
  try {
    return json.encode(list);
  } catch (_) {
    return '[]';
  }
}

String _serializeJson(dynamic jsonVar) {
  jsonVar ??= {};
  try {
    return json.encode(jsonVar);
  } catch (_) {
    return '{}';
  }
}
