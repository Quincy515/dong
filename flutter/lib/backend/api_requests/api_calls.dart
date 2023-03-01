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
        'x-token':
            'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiNjk5MDhjN2EtNTAzNy00YjI4LWEzNWUtYmE3MWZjMzY5NGFhIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjo4ODgsIkJ1ZmZlclRpbWUiOjg2NDAwLCJleHAiOjE2Njk2MTY0ODgsImlzcyI6InFtUGx1cyIsIm5iZiI6MTY2OTAxMDY4OH0.XZEgZ4BMT-KXqkSki1YNNPD_xnqMP9ef7DfE2XYMLno',
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
        'x-token':
            'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiNjk5MDhjN2EtNTAzNy00YjI4LWEzNWUtYmE3MWZjMzY5NGFhIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjo4ODgsIkJ1ZmZlclRpbWUiOjg2NDAwLCJleHAiOjE2Njk2MTY0ODgsImlzcyI6InFtUGx1cyIsIm5iZiI6MTY2OTAxMDY4OH0.XZEgZ4BMT-KXqkSki1YNNPD_xnqMP9ef7DfE2XYMLno',
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
    'x-token':
        'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiNjk5MDhjN2EtNTAzNy00YjI4LWEzNWUtYmE3MWZjMzY5NGFhIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjo4ODgsIkJ1ZmZlclRpbWUiOjg2NDAwLCJleHAiOjE2Njk2MTY0ODgsImlzcyI6InFtUGx1cyIsIm5iZiI6MTY2OTAxMDY4OH0.XZEgZ4BMT-KXqkSki1YNNPD_xnqMP9ef7DfE2XYMLno',
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
    'x-token':
        'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZjk4MjEyZGItMjM0NC00YWZkLWJmODQtNzU1MzFiY2JkMGI3IiwiSUQiOjgsIlVzZXJuYW1lIjoiaW1tYW51ZWwiLCJOaWNrTmFtZSI6IuacseS6kea2myIsIkF1dGhvcml0eUlkIjo5MDAwLCJCdWZmZXJUaW1lIjo4NjQwMCwiZXhwIjoxNjY5NzIzODkxLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE2NjkxMTgwOTF9.xAAyGMc8Ym5dEqFVOwGz9Oqvzh5jR_R2ZYCNhSe5-JI',
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
    'x-token':
        'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiOWQ0ZTM2ZjMtN2I5OS00MDcxLWFjN2YtZDQxZmJhMDc0OWQ1IiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Ik1yLuWlh-a3vCIsIkF1dGhvcml0eUlkIjo4ODgsIkJ1ZmZlclRpbWUiOjg2NDAwMDAwMDAwMDAwLCJleHAiOjE2NzgxNTcxNzEsImlzcyI6InFtUGx1cyIsIm5iZiI6MTY3NzQ5OTI3Nn0.SFIMj4ux7aEG-RQViU5GK6ib_3UkVWRqeyNOxLB20nY',
  };
  static GetXiangMuDengJiBiaoListCall getXiangMuDengJiBiaoListCall =
      GetXiangMuDengJiBiaoListCall();
  static DeleteXiangMuDengJiBiaoCall deleteXiangMuDengJiBiaoCall =
      DeleteXiangMuDengJiBiaoCall();
  static UpdateXiangMuDengJiBiaoCall updateXiangMuDengJiBiaoCall =
      UpdateXiangMuDengJiBiaoCall();
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

/// End jingYingBu Group Code

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
