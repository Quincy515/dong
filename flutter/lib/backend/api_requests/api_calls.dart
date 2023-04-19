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
  dynamic userID(dynamic response) => getJsonField(
        response,
        r'''$.data.user.ID''',
      );
  dynamic userName(dynamic response) => getJsonField(
        response,
        r'''$.data.user.userName''',
      );
  dynamic nickName(dynamic response) => getJsonField(
        response,
        r'''$.data.user.nickName''',
      );
  dynamic headerImg(dynamic response) => getJsonField(
        response,
        r'''$.data.user.headerImg''',
      );
  dynamic authorityName(dynamic response) => getJsonField(
        response,
        r'''$.data.user.authority.authorityName''',
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
  dynamic userID(dynamic response) => getJsonField(
        response,
        r'''$.data.list[:].ID''',
        true,
      );
  dynamic userName(dynamic response) => getJsonField(
        response,
        r'''$.data.list[:].userName''',
        true,
      );
  dynamic nickName(dynamic response) => getJsonField(
        response,
        r'''$.data.list[:].nickName''',
        true,
      );
  dynamic headerImg(dynamic response) => getJsonField(
        response,
        r'''$.data.list[:].headerImg''',
      );
  dynamic authorityId(dynamic response) => getJsonField(
        response,
        r'''$.data.list[:].authorityId''',
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
  static FindSysDictionaryCall findSysDictionaryCall = FindSysDictionaryCall();
  static GetSysDictionaryListCall getSysDictionaryListCall =
      GetSysDictionaryListCall();
  static DeleteSysDictionaryCall deleteSysDictionaryCall =
      DeleteSysDictionaryCall();
  static UpdateSysDictionaryCall updateSysDictionaryCall =
      UpdateSysDictionaryCall();
  static GetSysDictionaryDetailListCall getSysDictionaryDetailListCall =
      GetSysDictionaryDetailListCall();
  static UpdateSysDictionaryDetailCall updateSysDictionaryDetailCall =
      UpdateSysDictionaryDetailCall();
  static DeleteSysDictionaryDetailCall deleteSysDictionaryDetailCall =
      DeleteSysDictionaryDetailCall();
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

class FindSysDictionaryCall {
  Future<ApiCallResponse> call({
    String? type = '',
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'findSysDictionary',
      apiUrl: '${FileGroup.baseUrl}/sysDictionary/findSysDictionary',
      callType: ApiCallType.GET,
      headers: {
        ...FileGroup.headers,
      },
      params: {
        'type': type,
      },
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }

  dynamic label(dynamic response) => getJsonField(
        response,
        r'''$.data.resysDictionary.sysDictionaryDetails[:].label''',
        true,
      );
  dynamic value(dynamic response) => getJsonField(
        response,
        r'''$.data.resysDictionary.sysDictionaryDetails[:].value''',
        true,
      );
  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class GetSysDictionaryListCall {
  Future<ApiCallResponse> call({
    int? page = 1,
    int? pageSize = 10,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getSysDictionaryList',
      apiUrl: '${FileGroup.baseUrl}/sysDictionary/getSysDictionaryList',
      callType: ApiCallType.GET,
      headers: {
        ...FileGroup.headers,
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

class DeleteSysDictionaryCall {
  Future<ApiCallResponse> call({
    int? id,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteSysDictionary',
      apiUrl: '${FileGroup.baseUrl}/sysDictionary/deleteSysDictionary',
      callType: ApiCallType.DELETE,
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

  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class UpdateSysDictionaryCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateSysDictionary',
      apiUrl: '${FileGroup.baseUrl}/sysDictionary/updateSysDictionary',
      callType: ApiCallType.PUT,
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

  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class GetSysDictionaryDetailListCall {
  Future<ApiCallResponse> call({
    int? sysDictionaryID = 1,
    int? pageSize = 10,
    int? page = 1,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getSysDictionaryDetailList',
      apiUrl:
          '${FileGroup.baseUrl}/sysDictionaryDetail/getSysDictionaryDetailList',
      callType: ApiCallType.GET,
      headers: {
        ...FileGroup.headers,
      },
      params: {
        'page': page,
        'pageSize': pageSize,
        'sysDictionaryID': sysDictionaryID,
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

class UpdateSysDictionaryDetailCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateSysDictionaryDetail',
      apiUrl:
          '${FileGroup.baseUrl}/sysDictionaryDetail/updateSysDictionaryDetail',
      callType: ApiCallType.PUT,
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

  dynamic msg(dynamic response) => getJsonField(
        response,
        r'''$.msg''',
      );
}

class DeleteSysDictionaryDetailCall {
  Future<ApiCallResponse> call({
    int? id,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteSysDictionaryDetail',
      apiUrl:
          '${FileGroup.baseUrl}/sysDictionaryDetail/deleteSysDictionaryDetail',
      callType: ApiCallType.DELETE,
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
  static GetZhengShuBiaoListCall getZhengShuBiaoListCall =
      GetZhengShuBiaoListCall();
  static DeleteZhengShuBiaoCall deleteZhengShuBiaoCall =
      DeleteZhengShuBiaoCall();
  static UpdateZhengShuBiaoCall updateZhengShuBiaoCall =
      UpdateZhengShuBiaoCall();
  static GetZhengShuJieChuJiLuBiaoListCall getZhengShuJieChuJiLuBiaoListCall =
      GetZhengShuJieChuJiLuBiaoListCall();
  static DeleteZhengShuJieChuJiLuBiaCall deleteZhengShuJieChuJiLuBiaCall =
      DeleteZhengShuJieChuJiLuBiaCall();
  static UpdateZhengShuJieChuJiLuBiaoCall updateZhengShuJieChuJiLuBiaoCall =
      UpdateZhengShuJieChuJiLuBiaoCall();
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

class GetZhengShuBiaoListCall {
  Future<ApiCallResponse> call({
    int? page = 1,
    int? pageSize = 10,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getZhengShuBiaoList',
      apiUrl: '${BanGongShiGroup.baseUrl}/zhengShuBiao/getZhengShuBiaoList',
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

class DeleteZhengShuBiaoCall {
  Future<ApiCallResponse> call({
    int? id,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteZhengShuBiao',
      apiUrl: '${BanGongShiGroup.baseUrl}/zhengShuBiao/deleteZhengShuBiao',
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

class UpdateZhengShuBiaoCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateZhengShuBiao',
      apiUrl: '${BanGongShiGroup.baseUrl}/zhengShuBiao/updateZhengShuBiao',
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

class GetZhengShuJieChuJiLuBiaoListCall {
  Future<ApiCallResponse> call({
    int? page = 1,
    int? pageSize = 10,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getZhengShuJieChuJiLuBiaoList',
      apiUrl:
          '${BanGongShiGroup.baseUrl}/zhengShuJieChuJiLuBiao/getZhengShuJieChuJiLuBiaoList',
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

class DeleteZhengShuJieChuJiLuBiaCall {
  Future<ApiCallResponse> call({
    int? id,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteZhengShuJieChuJiLuBia',
      apiUrl:
          '${BanGongShiGroup.baseUrl}/zhengShuJieChuJiLuBiao/deleteZhengShuJieChuJiLuBiao',
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

class UpdateZhengShuJieChuJiLuBiaoCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateZhengShuJieChuJiLuBiao',
      apiUrl:
          '${BanGongShiGroup.baseUrl}/zhengShuJieChuJiLuBiao/updateZhengShuJieChuJiLuBiao',
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

/// Start xiangMuBiao Group Code
class XiangMuBiaoGroup {
  static String baseUrl = 'http://119.45.234.161:8888/xiangMuBiao';
  static Map<String, String> headers = {
    'x-token': FFAppState().xtoken,
  };
  static GetXiangMuBiaoListCall getXiangMuBiaoListCall =
      GetXiangMuBiaoListCall();
  static DeleteXiangMuBiaoCall deleteXiangMuBiaoCall = DeleteXiangMuBiaoCall();
  static UpdateXiangMuBiaoCall updateXiangMuBiaoCall = UpdateXiangMuBiaoCall();
}

class GetXiangMuBiaoListCall {
  Future<ApiCallResponse> call({
    int? page = 1,
    int? pageSize = 10,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getXiangMuBiaoList',
      apiUrl: '${XiangMuBiaoGroup.baseUrl}/getXiangMuBiaoList',
      callType: ApiCallType.GET,
      headers: {
        ...XiangMuBiaoGroup.headers,
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

class DeleteXiangMuBiaoCall {
  Future<ApiCallResponse> call({
    int? id,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteXiangMuBiao',
      apiUrl: '${XiangMuBiaoGroup.baseUrl}/deleteXiangMuBiao',
      callType: ApiCallType.DELETE,
      headers: {
        ...XiangMuBiaoGroup.headers,
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

class UpdateXiangMuBiaoCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateXiangMuBiao',
      apiUrl: '${XiangMuBiaoGroup.baseUrl}/updateXiangMuBiao',
      callType: ApiCallType.PUT,
      headers: {
        ...XiangMuBiaoGroup.headers,
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

/// End xiangMuBiao Group Code

/// Start xiangMuShiGongBiao Group Code

class XiangMuShiGongBiaoGroup {
  static String baseUrl = 'http://119.45.234.161:8888/xiangMuShiGongBiao';
  static Map<String, String> headers = {
    'x-token': FFAppState().xtoken,
  };
  static GetXiangMuShiGongBiaoListCall getXiangMuShiGongBiaoListCall =
      GetXiangMuShiGongBiaoListCall();
  static DeleteXiangMuShiGongBiaoCall deleteXiangMuShiGongBiaoCall =
      DeleteXiangMuShiGongBiaoCall();
  static UpdateXiangMuShiGongBiaoCall updateXiangMuShiGongBiaoCall =
      UpdateXiangMuShiGongBiaoCall();
}

class GetXiangMuShiGongBiaoListCall {
  Future<ApiCallResponse> call({
    int? page,
    int? pageSize,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getXiangMuShiGongBiaoList',
      apiUrl: '${XiangMuShiGongBiaoGroup.baseUrl}/getXiangMuShiGongBiaoList',
      callType: ApiCallType.GET,
      headers: {
        ...XiangMuShiGongBiaoGroup.headers,
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

class DeleteXiangMuShiGongBiaoCall {
  Future<ApiCallResponse> call({
    int? id,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteXiangMuShiGongBiao',
      apiUrl: '${XiangMuShiGongBiaoGroup.baseUrl}/deleteXiangMuShiGongBiao',
      callType: ApiCallType.DELETE,
      headers: {
        ...XiangMuShiGongBiaoGroup.headers,
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

class UpdateXiangMuShiGongBiaoCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateXiangMuShiGongBiao',
      apiUrl: '${XiangMuShiGongBiaoGroup.baseUrl}/updateXiangMuShiGongBiao',
      callType: ApiCallType.PUT,
      headers: {
        ...XiangMuShiGongBiaoGroup.headers,
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

/// End xiangMuShiGongBiao Group Code

/// Start fenBaoShangBiao Group Code

class FenBaoShangBiaoGroup {
  static String baseUrl = 'http://119.45.234.161:8888/fenBaoShangBiao';
  static Map<String, String> headers = {
    'x-token': FFAppState().xtoken,
  };
  static GetFenBaoShangBiaoListCall getFenBaoShangBiaoListCall =
      GetFenBaoShangBiaoListCall();
  static DeleteFenBaoShangBiaoCall deleteFenBaoShangBiaoCall =
      DeleteFenBaoShangBiaoCall();
  static UpdateFenBaoShangBiaoCall updateFenBaoShangBiaoCall =
      UpdateFenBaoShangBiaoCall();
}

class GetFenBaoShangBiaoListCall {
  Future<ApiCallResponse> call({
    int? page,
    int? pageSize,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getFenBaoShangBiaoList',
      apiUrl: '${FenBaoShangBiaoGroup.baseUrl}/getFenBaoShangBiaoList',
      callType: ApiCallType.GET,
      headers: {
        ...FenBaoShangBiaoGroup.headers,
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

class DeleteFenBaoShangBiaoCall {
  Future<ApiCallResponse> call({
    int? id,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteFenBaoShangBiao',
      apiUrl: '${FenBaoShangBiaoGroup.baseUrl}/deleteFenBaoShangBiao',
      callType: ApiCallType.DELETE,
      headers: {
        ...FenBaoShangBiaoGroup.headers,
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

class UpdateFenBaoShangBiaoCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateFenBaoShangBiao',
      apiUrl: '${FenBaoShangBiaoGroup.baseUrl}/updateFenBaoShangBiao',
      callType: ApiCallType.PUT,
      headers: {
        ...FenBaoShangBiaoGroup.headers,
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

/// End fenBaoShangBiao Group Code

/// Start fenBaoShangYuXiangMuGuanLianBiao Group Code

class FenBaoShangYuXiangMuGuanLianBiaoGroup {
  static String baseUrl =
      'http://119.45.234.161:8888/fenBaoShangYuXiangMuGuanLianBiao';
  static Map<String, String> headers = {
    'x-token': FFAppState().xtoken,
  };
  static GetFenBaoShangYuXiangMuGuanLianBiaoListCall
      getFenBaoShangYuXiangMuGuanLianBiaoListCall =
      GetFenBaoShangYuXiangMuGuanLianBiaoListCall();
  static DeleteFenBaoShangYuXiangMuGuanLianBiaoCall
      deleteFenBaoShangYuXiangMuGuanLianBiaoCall =
      DeleteFenBaoShangYuXiangMuGuanLianBiaoCall();
  static UpdateFenBaoShangYuXiangMuGuanLianBiaoCall
      updateFenBaoShangYuXiangMuGuanLianBiaoCall =
      UpdateFenBaoShangYuXiangMuGuanLianBiaoCall();
}

class GetFenBaoShangYuXiangMuGuanLianBiaoListCall {
  Future<ApiCallResponse> call({
    int? page,
    int? pageSize,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getFenBaoShangYuXiangMuGuanLianBiaoList',
      apiUrl:
          '${FenBaoShangYuXiangMuGuanLianBiaoGroup.baseUrl}/getFenBaoShangYuXiangMuGuanLianBiaoList',
      callType: ApiCallType.GET,
      headers: {
        ...FenBaoShangYuXiangMuGuanLianBiaoGroup.headers,
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

class DeleteFenBaoShangYuXiangMuGuanLianBiaoCall {
  Future<ApiCallResponse> call({
    int? id,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteFenBaoShangYuXiangMuGuanLianBiao',
      apiUrl:
          '${FenBaoShangYuXiangMuGuanLianBiaoGroup.baseUrl}/deleteFenBaoShangYuXiangMuGuanLianBiao',
      callType: ApiCallType.DELETE,
      headers: {
        ...FenBaoShangYuXiangMuGuanLianBiaoGroup.headers,
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

class UpdateFenBaoShangYuXiangMuGuanLianBiaoCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateFenBaoShangYuXiangMuGuanLianBiao',
      apiUrl:
          '${FenBaoShangYuXiangMuGuanLianBiaoGroup.baseUrl}/updateFenBaoShangYuXiangMuGuanLianBiao',
      callType: ApiCallType.PUT,
      headers: {
        ...FenBaoShangYuXiangMuGuanLianBiaoGroup.headers,
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

/// End fenBaoShangYuXiangMuGuanLianBiao Group Code

/// Start heTongWenJianBiao Group Code

class HeTongWenJianBiaoGroup {
  static String baseUrl = 'http://119.45.234.161:8888/heTongWenJianBiao';
  static Map<String, String> headers = {
    'x-token': FFAppState().xtoken,
  };
  static GetHeTongWenJianBiaoListCall getHeTongWenJianBiaoListCall =
      GetHeTongWenJianBiaoListCall();
  static DeleteHeTongWenJianBiaoCall deleteHeTongWenJianBiaoCall =
      DeleteHeTongWenJianBiaoCall();
  static UpdateHeTongWenJianBiaoCall updateHeTongWenJianBiaoCall =
      UpdateHeTongWenJianBiaoCall();
}

class GetHeTongWenJianBiaoListCall {
  Future<ApiCallResponse> call({
    int? page,
    int? pageSize,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getHeTongWenJianBiaoList',
      apiUrl: '${HeTongWenJianBiaoGroup.baseUrl}/getHeTongWenJianBiaoList',
      callType: ApiCallType.GET,
      headers: {
        ...HeTongWenJianBiaoGroup.headers,
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

class DeleteHeTongWenJianBiaoCall {
  Future<ApiCallResponse> call({
    int? id,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteHeTongWenJianBiao',
      apiUrl: '${HeTongWenJianBiaoGroup.baseUrl}/deleteHeTongWenJianBiao',
      callType: ApiCallType.DELETE,
      headers: {
        ...HeTongWenJianBiaoGroup.headers,
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

class UpdateHeTongWenJianBiaoCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateHeTongWenJianBiao',
      apiUrl: '${HeTongWenJianBiaoGroup.baseUrl}/updateHeTongWenJianBiao',
      callType: ApiCallType.PUT,
      headers: {
        ...HeTongWenJianBiaoGroup.headers,
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

/// End heTongWenJianBiao Group Code

/// Start gongYingShangBiao Group Code

class GongYingShangBiaoGroup {
  static String baseUrl = 'http://119.45.234.161:8888/gongYingShangBiao';
  static Map<String, String> headers = {
    'x-token': FFAppState().xtoken,
  };
  static GetGongYingShangBiaoListCall getGongYingShangBiaoListCall =
      GetGongYingShangBiaoListCall();
  static DeleteGongYingShangBiaoCall deleteGongYingShangBiaoCall =
      DeleteGongYingShangBiaoCall();
  static UpdateGongYingShangBiaoCall updateGongYingShangBiaoCall =
      UpdateGongYingShangBiaoCall();
}

class GetGongYingShangBiaoListCall {
  Future<ApiCallResponse> call({
    int? page,
    int? pageSize,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getGongYingShangBiaoList',
      apiUrl: '${GongYingShangBiaoGroup.baseUrl}/getGongYingShangBiaoList',
      callType: ApiCallType.GET,
      headers: {
        ...GongYingShangBiaoGroup.headers,
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

class DeleteGongYingShangBiaoCall {
  Future<ApiCallResponse> call({
    int? id,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteGongYingShangBiao',
      apiUrl: '${GongYingShangBiaoGroup.baseUrl}/deleteGongYingShangBiao',
      callType: ApiCallType.DELETE,
      headers: {
        ...GongYingShangBiaoGroup.headers,
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

class UpdateGongYingShangBiaoCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateGongYingShangBiao',
      apiUrl: '${GongYingShangBiaoGroup.baseUrl}/updateGongYingShangBiao',
      callType: ApiCallType.PUT,
      headers: {
        ...GongYingShangBiaoGroup.headers,
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

/// End gongYingShangBiao Group Code

/// Start xiangMuZhuCaiBiao Group Code

class XiangMuZhuCaiBiaoGroup {
  static String baseUrl = 'http://119.45.234.161:8888/xiangMuZhuCaiBiao';
  static Map<String, String> headers = {
    'x-token': FFAppState().xtoken,
  };
  static GetXiangMuZhuCaiBiaoListCall getXiangMuZhuCaiBiaoListCall =
      GetXiangMuZhuCaiBiaoListCall();
  static DeleteXiangMuZhuCaiBiaoCall deleteXiangMuZhuCaiBiaoCall =
      DeleteXiangMuZhuCaiBiaoCall();
  static UpdateXiangMuZhuCaiBiaoCall updateXiangMuZhuCaiBiaoCall =
      UpdateXiangMuZhuCaiBiaoCall();
}

class GetXiangMuZhuCaiBiaoListCall {
  Future<ApiCallResponse> call({
    int? page,
    int? pageSize,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getXiangMuZhuCaiBiaoList',
      apiUrl: '${XiangMuZhuCaiBiaoGroup.baseUrl}/getXiangMuZhuCaiBiaoList',
      callType: ApiCallType.GET,
      headers: {
        ...XiangMuZhuCaiBiaoGroup.headers,
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

class DeleteXiangMuZhuCaiBiaoCall {
  Future<ApiCallResponse> call({
    int? id,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteXiangMuZhuCaiBiao',
      apiUrl: '${XiangMuZhuCaiBiaoGroup.baseUrl}/deleteXiangMuZhuCaiBiao',
      callType: ApiCallType.DELETE,
      headers: {
        ...XiangMuZhuCaiBiaoGroup.headers,
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

class UpdateXiangMuZhuCaiBiaoCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateXiangMuZhuCaiBiao',
      apiUrl: '${XiangMuZhuCaiBiaoGroup.baseUrl}/updateXiangMuZhuCaiBiao',
      callType: ApiCallType.PUT,
      headers: {
        ...XiangMuZhuCaiBiaoGroup.headers,
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

/// End xiangMuZhuCaiBiao Group Code

/// Start caiGouJiHuaBiao Group Code

class CaiGouJiHuaBiaoGroup {
  static String baseUrl = 'http://119.45.234.161:8888/caiGouJiHuaBiao';
  static Map<String, String> headers = {
    'x-token': FFAppState().xtoken,
  };
  static GetCaiGouJiHuaBiaoListCall getCaiGouJiHuaBiaoListCall =
      GetCaiGouJiHuaBiaoListCall();
  static DeleteCaiGouJiHuaBiaoCall deleteCaiGouJiHuaBiaoCall =
      DeleteCaiGouJiHuaBiaoCall();
  static UpdateCaiGouJiHuaBiaoCall updateCaiGouJiHuaBiaoCall =
      UpdateCaiGouJiHuaBiaoCall();
}

class GetCaiGouJiHuaBiaoListCall {
  Future<ApiCallResponse> call({
    int? page,
    int? pageSize,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getCaiGouJiHuaBiaoList',
      apiUrl: '${CaiGouJiHuaBiaoGroup.baseUrl}/getCaiGouJiHuaBiaoList',
      callType: ApiCallType.GET,
      headers: {
        ...CaiGouJiHuaBiaoGroup.headers,
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

class DeleteCaiGouJiHuaBiaoCall {
  Future<ApiCallResponse> call({
    int? id,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteCaiGouJiHuaBiao',
      apiUrl: '${CaiGouJiHuaBiaoGroup.baseUrl}/deleteCaiGouJiHuaBiao',
      callType: ApiCallType.DELETE,
      headers: {
        ...CaiGouJiHuaBiaoGroup.headers,
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

class UpdateCaiGouJiHuaBiaoCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateCaiGouJiHuaBiao',
      apiUrl: '${CaiGouJiHuaBiaoGroup.baseUrl}/updateCaiGouJiHuaBiao',
      callType: ApiCallType.PUT,
      headers: {
        ...CaiGouJiHuaBiaoGroup.headers,
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

/// End caiGouJiHuaBiao Group Code

/// Start chuRuKuDanJuBiao Group Code

class ChuRuKuDanJuBiaoGroup {
  static String baseUrl = 'http://119.45.234.161:8888/chuRuKuDanJuBiao';
  static Map<String, String> headers = {
    'x-token': FFAppState().xtoken,
  };
  static GetChuRuKuDanJuBiaoListCall getChuRuKuDanJuBiaoListCall =
      GetChuRuKuDanJuBiaoListCall();
  static DeleteChuRuKuDanJuBiaoCall deleteChuRuKuDanJuBiaoCall =
      DeleteChuRuKuDanJuBiaoCall();
  static UpdateChuRuKuDanJuBiaoCall updateChuRuKuDanJuBiaoCall =
      UpdateChuRuKuDanJuBiaoCall();
}

class GetChuRuKuDanJuBiaoListCall {
  Future<ApiCallResponse> call({
    int? page,
    int? pageSize,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getChuRuKuDanJuBiaoList',
      apiUrl: '${ChuRuKuDanJuBiaoGroup.baseUrl}/getChuRuKuDanJuBiaoList',
      callType: ApiCallType.GET,
      headers: {
        ...ChuRuKuDanJuBiaoGroup.headers,
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

class DeleteChuRuKuDanJuBiaoCall {
  Future<ApiCallResponse> call({
    int? id,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteChuRuKuDanJuBiao',
      apiUrl: '${ChuRuKuDanJuBiaoGroup.baseUrl}/deleteChuRuKuDanJuBiao',
      callType: ApiCallType.DELETE,
      headers: {
        ...ChuRuKuDanJuBiaoGroup.headers,
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

class UpdateChuRuKuDanJuBiaoCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateChuRuKuDanJuBiao',
      apiUrl: '${ChuRuKuDanJuBiaoGroup.baseUrl}/updateChuRuKuDanJuBiao',
      callType: ApiCallType.PUT,
      headers: {
        ...ChuRuKuDanJuBiaoGroup.headers,
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

/// End chuRuKuDanJuBiao Group Code

/// Start chuRuKuMingXiBiao Group Code

class ChuRuKuMingXiBiaoGroup {
  static String baseUrl = 'http://119.45.234.161:8888/chuRuKuMingXiBiao';
  static Map<String, String> headers = {
    'x-token': FFAppState().xtoken,
  };
  static GetChuRuKuMingXiBiaoListCall getChuRuKuMingXiBiaoListCall =
      GetChuRuKuMingXiBiaoListCall();
  static DeleteChuRuKuMingXiBiaoCall deleteChuRuKuMingXiBiaoCall =
      DeleteChuRuKuMingXiBiaoCall();
  static UpdateChuRuKuMingXiBiaoCall updateChuRuKuMingXiBiaoCall =
      UpdateChuRuKuMingXiBiaoCall();
}

class GetChuRuKuMingXiBiaoListCall {
  Future<ApiCallResponse> call({
    int? page,
    int? pageSize,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getChuRuKuMingXiBiaoList',
      apiUrl: '${ChuRuKuMingXiBiaoGroup.baseUrl}/getChuRuKuMingXiBiaoList',
      callType: ApiCallType.GET,
      headers: {
        ...ChuRuKuMingXiBiaoGroup.headers,
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

class DeleteChuRuKuMingXiBiaoCall {
  Future<ApiCallResponse> call({
    int? id,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteChuRuKuMingXiBiao',
      apiUrl: '${ChuRuKuMingXiBiaoGroup.baseUrl}/deleteChuRuKuMingXiBiao',
      callType: ApiCallType.DELETE,
      headers: {
        ...ChuRuKuMingXiBiaoGroup.headers,
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

class UpdateChuRuKuMingXiBiaoCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateChuRuKuMingXiBiao',
      apiUrl: '${ChuRuKuMingXiBiaoGroup.baseUrl}/updateChuRuKuMingXiBiao',
      callType: ApiCallType.PUT,
      headers: {
        ...ChuRuKuMingXiBiaoGroup.headers,
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

/// End chuRuKuMingXiBiao Group Code

/// Start sheBeiZuLinBiao Group Code

class SheBeiZuLinBiaoGroup {
  static String baseUrl = 'http://119.45.234.161:8888/sheBeiZuLinBiao';
  static Map<String, String> headers = {
    'x-token': FFAppState().xtoken,
  };
  static GetSheBeiZuLinBiaoListCall getSheBeiZuLinBiaoListCall =
      GetSheBeiZuLinBiaoListCall();
  static DeleteSheBeiZuLinBiaoCall deleteSheBeiZuLinBiaoCall =
      DeleteSheBeiZuLinBiaoCall();
  static UpdateSheBeiZuLinBiaoCall updateSheBeiZuLinBiaoCall =
      UpdateSheBeiZuLinBiaoCall();
}

class GetSheBeiZuLinBiaoListCall {
  Future<ApiCallResponse> call({
    int? page,
    int? pageSize,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getSheBeiZuLinBiaoList',
      apiUrl: '${SheBeiZuLinBiaoGroup.baseUrl}/getSheBeiZuLinBiaoList',
      callType: ApiCallType.GET,
      headers: {
        ...SheBeiZuLinBiaoGroup.headers,
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

class DeleteSheBeiZuLinBiaoCall {
  Future<ApiCallResponse> call({
    int? id,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteSheBeiZuLinBiao',
      apiUrl: '${SheBeiZuLinBiaoGroup.baseUrl}/deleteSheBeiZuLinBiao',
      callType: ApiCallType.DELETE,
      headers: {
        ...SheBeiZuLinBiaoGroup.headers,
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

class UpdateSheBeiZuLinBiaoCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateSheBeiZuLinBiao',
      apiUrl: '${SheBeiZuLinBiaoGroup.baseUrl}/updateSheBeiZuLinBiao',
      callType: ApiCallType.PUT,
      headers: {
        ...SheBeiZuLinBiaoGroup.headers,
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

/// End sheBeiZuLinBiao Group Code

/// Start sheBeiMingXiBiao Group Code

class SheBeiMingXiBiaoGroup {
  static String baseUrl = 'http://119.45.234.161:8888/sheBeiMingXiBiao';
  static Map<String, String> headers = {
    'x-token': FFAppState().xtoken,
  };
  static GetSheBeiMingXiBiaoListCall getSheBeiMingXiBiaoListCall =
      GetSheBeiMingXiBiaoListCall();
  static DeleteSheBeiMingXiBiaoCall deleteSheBeiMingXiBiaoCall =
      DeleteSheBeiMingXiBiaoCall();
  static UpdateSheBeiMingXiBiaoCall updateSheBeiMingXiBiaoCall =
      UpdateSheBeiMingXiBiaoCall();
}

class GetSheBeiMingXiBiaoListCall {
  Future<ApiCallResponse> call({
    int? page,
    int? pageSize,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getSheBeiMingXiBiaoList',
      apiUrl: '${SheBeiMingXiBiaoGroup.baseUrl}/getSheBeiMingXiBiaoList',
      callType: ApiCallType.GET,
      headers: {
        ...SheBeiMingXiBiaoGroup.headers,
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

class DeleteSheBeiMingXiBiaoCall {
  Future<ApiCallResponse> call({
    int? id,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteSheBeiMingXiBiao',
      apiUrl: '${SheBeiMingXiBiaoGroup.baseUrl}/deleteSheBeiMingXiBiao',
      callType: ApiCallType.DELETE,
      headers: {
        ...SheBeiMingXiBiaoGroup.headers,
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

class UpdateSheBeiMingXiBiaoCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateSheBeiMingXiBiao',
      apiUrl: '${SheBeiMingXiBiaoGroup.baseUrl}/updateSheBeiMingXiBiao',
      callType: ApiCallType.PUT,
      headers: {
        ...SheBeiMingXiBiaoGroup.headers,
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

/// End sheBeiMingXiBiao Group Code

/// Start shiGongJinDuBiao Group Code

class ShiGongJinDuBiaoGroup {
  static String baseUrl = 'http://119.45.234.161:8888/shiGongJinDuBiao';
  static Map<String, String> headers = {
    'x-token': FFAppState().xtoken,
  };
  static GetShiGongJinDuBiaoListCall getShiGongJinDuBiaoListCall =
      GetShiGongJinDuBiaoListCall();
  static DeleteShiGongJinDuBiaoCall deleteShiGongJinDuBiaoCall =
      DeleteShiGongJinDuBiaoCall();
  static UpdateShiGongJinDuBiaovCall updateShiGongJinDuBiaovCall =
      UpdateShiGongJinDuBiaovCall();
}

class GetShiGongJinDuBiaoListCall {
  Future<ApiCallResponse> call({
    int? page,
    int? pageSize,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getShiGongJinDuBiaoList',
      apiUrl: '${ShiGongJinDuBiaoGroup.baseUrl}/getShiGongJinDuBiaoList',
      callType: ApiCallType.GET,
      headers: {
        ...ShiGongJinDuBiaoGroup.headers,
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

class DeleteShiGongJinDuBiaoCall {
  Future<ApiCallResponse> call({
    int? id,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteShiGongJinDuBiao',
      apiUrl: '${ShiGongJinDuBiaoGroup.baseUrl}/deleteShiGongJinDuBiao',
      callType: ApiCallType.DELETE,
      headers: {
        ...ShiGongJinDuBiaoGroup.headers,
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

class UpdateShiGongJinDuBiaovCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateShiGongJinDuBiaov',
      apiUrl: '${ShiGongJinDuBiaoGroup.baseUrl}/updateShiGongJinDuBiao',
      callType: ApiCallType.PUT,
      headers: {
        ...ShiGongJinDuBiaoGroup.headers,
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

/// End shiGongJinDuBiao Group Code

/// Start wenJianZiLiaoJieDianBiao Group Code

class WenJianZiLiaoJieDianBiaoGroup {
  static String baseUrl = 'http://119.45.234.161:8888/wenJianZiLiaoJieDianBiao';
  static Map<String, String> headers = {
    'x-token': FFAppState().xtoken,
  };
  static GetWenJianZiLiaoJieDianBiaoListCall
      getWenJianZiLiaoJieDianBiaoListCall =
      GetWenJianZiLiaoJieDianBiaoListCall();
  static DeleteWenJianZiLiaoJieDianBiaoCall deleteWenJianZiLiaoJieDianBiaoCall =
      DeleteWenJianZiLiaoJieDianBiaoCall();
  static UpdateWenJianZiLiaoJieDianBiaoCall updateWenJianZiLiaoJieDianBiaoCall =
      UpdateWenJianZiLiaoJieDianBiaoCall();
}

class GetWenJianZiLiaoJieDianBiaoListCall {
  Future<ApiCallResponse> call({
    int? page,
    int? pageSize,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getWenJianZiLiaoJieDianBiaoList',
      apiUrl:
          '${WenJianZiLiaoJieDianBiaoGroup.baseUrl}/getWenJianZiLiaoJieDianBiaoList',
      callType: ApiCallType.GET,
      headers: {
        ...WenJianZiLiaoJieDianBiaoGroup.headers,
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

class DeleteWenJianZiLiaoJieDianBiaoCall {
  Future<ApiCallResponse> call({
    int? id,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteWenJianZiLiaoJieDianBiao',
      apiUrl:
          '${WenJianZiLiaoJieDianBiaoGroup.baseUrl}/deleteWenJianZiLiaoJieDianBiao',
      callType: ApiCallType.DELETE,
      headers: {
        ...WenJianZiLiaoJieDianBiaoGroup.headers,
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

class UpdateWenJianZiLiaoJieDianBiaoCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateWenJianZiLiaoJieDianBiao',
      apiUrl:
          '${WenJianZiLiaoJieDianBiaoGroup.baseUrl}/updateWenJianZiLiaoJieDianBiao',
      callType: ApiCallType.PUT,
      headers: {
        ...WenJianZiLiaoJieDianBiaoGroup.headers,
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

/// End wenJianZiLiaoJieDianBiao Group Code

/// Start wenJianZiLiaoBiao Group Code

class WenJianZiLiaoBiaoGroup {
  static String baseUrl = 'http://119.45.234.161:8888/wenJianZiLiaoBiao';
  static Map<String, String> headers = {
    'x-token': FFAppState().xtoken,
  };
  static GetWenJianZiLiaoBiaoListCall getWenJianZiLiaoBiaoListCall =
      GetWenJianZiLiaoBiaoListCall();
  static DeleteWenJianZiLiaoBiaoCall deleteWenJianZiLiaoBiaoCall =
      DeleteWenJianZiLiaoBiaoCall();
  static UpdateWenJianZiLiaoBiaoCall updateWenJianZiLiaoBiaoCall =
      UpdateWenJianZiLiaoBiaoCall();
}

class GetWenJianZiLiaoBiaoListCall {
  Future<ApiCallResponse> call({
    int? page,
    int? pageSize,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getWenJianZiLiaoBiaoList',
      apiUrl: '${WenJianZiLiaoBiaoGroup.baseUrl}/getWenJianZiLiaoBiaoList',
      callType: ApiCallType.GET,
      headers: {
        ...WenJianZiLiaoBiaoGroup.headers,
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

class DeleteWenJianZiLiaoBiaoCall {
  Future<ApiCallResponse> call({
    int? id,
  }) {
    final body = '{"ID": $id}';
    return ApiManager.instance.makeApiCall(
      callName: 'deleteWenJianZiLiaoBiao',
      apiUrl: '${WenJianZiLiaoBiaoGroup.baseUrl}/deleteWenJianZiLiaoBiao',
      callType: ApiCallType.DELETE,
      headers: {
        ...WenJianZiLiaoBiaoGroup.headers,
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

class UpdateWenJianZiLiaoBiaoCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateWenJianZiLiaoBiao',
      apiUrl: '${WenJianZiLiaoBiaoGroup.baseUrl}/updateWenJianZiLiaoBiao',
      callType: ApiCallType.PUT,
      headers: {
        ...WenJianZiLiaoBiaoGroup.headers,
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

/// End wenJianZiLiaoBiao Group Code

/// Start xiaoXi Group Code

class XiaoXiGroup {
  static String baseUrl = 'http://119.45.234.161:8888';
  static Map<String, String> headers = {
    'x-token': FFAppState().xtoken,
  };
  static CreateXiaoXiTongZhiCall createXiaoXiTongZhiCall =
      CreateXiaoXiTongZhiCall();
  static GetXiaoXiTongZhiListCall getXiaoXiTongZhiListCall =
      GetXiaoXiTongZhiListCall();
  static UpdateXiaoXiTongZhiCall updateXiaoXiTongZhiCall =
      UpdateXiaoXiTongZhiCall();
}

class CreateXiaoXiTongZhiCall {
  Future<ApiCallResponse> call({
    int? senderId,
    int? receiverId,
    String? title = '',
    String? content = '',
    String? redirectUrl = '',
    String? type = '',
    String? group = '',
    int? status = 0,
  }) {
    final body = '''
{
  "content": "${content}",
  "group": "${group}",
  "receiverId": ${receiverId},
  "redirectUrl": "${redirectUrl}",
  "senderId": ${senderId},
  "status": ${status},
  "title": "${title}",
  "type": "${type}"
}''';
    return ApiManager.instance.makeApiCall(
      callName: 'createXiaoXiTongZhi',
      apiUrl: '${XiaoXiGroup.baseUrl}/xiaoXiTongZhi/createXiaoXiTongZhi',
      callType: ApiCallType.POST,
      headers: {
        ...XiaoXiGroup.headers,
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

class GetXiaoXiTongZhiListCall {
  Future<ApiCallResponse> call({
    int? page = 1,
    int? pageSize = 10,
    int? receiverId,
  }) {
    return ApiManager.instance.makeApiCall(
      callName: 'getXiaoXiTongZhiList',
      apiUrl: '${XiaoXiGroup.baseUrl}/xiaoXiTongZhi/getXiaoXiTongZhiList',
      callType: ApiCallType.GET,
      headers: {
        ...XiaoXiGroup.headers,
      },
      params: {
        'page': page,
        'pageSize': pageSize,
        'receiverId': receiverId,
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

class UpdateXiaoXiTongZhiCall {
  Future<ApiCallResponse> call({
    dynamic? itemJson,
  }) {
    final item = _serializeJson(itemJson);
    final body = '''
${item}''';
    return ApiManager.instance.makeApiCall(
      callName: 'updateXiaoXiTongZhi',
      apiUrl: '${XiaoXiGroup.baseUrl}/xiaoXiTongZhi/updateXiaoXiTongZhi',
      callType: ApiCallType.PUT,
      headers: {
        ...XiaoXiGroup.headers,
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

/// End xiaoXi Group Code

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
