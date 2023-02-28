// ignore_for_file: unused_import

import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/flutter_flow_widgets.dart';
import '/flutter_flow/random_data_util.dart' as random_data;
import 'package:cached_network_image/cached_network_image.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:provider/provider.dart';

class AddCandidatesModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  // State field(s) for username widget.
  TextEditingController? usernameController;
  String? Function(BuildContext, String?)? usernameControllerValidator;
  // State field(s) for pasword widget.
  TextEditingController? paswordController;
  late bool paswordVisibility;
  String? Function(BuildContext, String?)? paswordControllerValidator;
  // State field(s) for nickName widget.
  TextEditingController? nickNameController;
  String? Function(BuildContext, String?)? nickNameControllerValidator;
  // Stores action output result for [Backend Call - API (adminregister)] action in Button widget.
  ApiCallResponse? adminRegisterRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {
    paswordVisibility = false;
  }

  void dispose() {
    usernameController?.dispose();
    paswordController?.dispose();
    nickNameController?.dispose();
  }

  /// Additional helper methods are added here.
}
