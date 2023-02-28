// ignore_for_file: unused_import

import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_icon_button.dart';
import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/flutter_flow_widgets.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:provider/provider.dart';

class LoginPageModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this page.

  // State field(s) for username widget.
  TextEditingController? usernameController;
  String? Function(BuildContext, String?)? usernameControllerValidator;
  // State field(s) for password widget.
  TextEditingController? passwordController;
  late bool passwordVisibility;
  String? Function(BuildContext, String?)? passwordControllerValidator;
  // State field(s) for captcha widget.
  TextEditingController? captchaController;
  String? Function(BuildContext, String?)? captchaControllerValidator;
  // Stores action output result for [Backend Call - API (login)] action in ButtonLogin widget.
  ApiCallResponse? loginRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {
    passwordVisibility = false;
  }

  void dispose() {
    usernameController?.dispose();
    passwordController?.dispose();
    captchaController?.dispose();
  }

  /// Additional helper methods are added here.
}
