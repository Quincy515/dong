// ignore_for_file: unused_import

import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/flutter_flow_widgets.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:provider/provider.dart';

class DictionaryDetailItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for label widget.
  TextEditingController? labelController;
  String? Function(BuildContext, String?)? labelControllerValidator;
  // State field(s) for value widget.
  TextEditingController? valueController;
  String? Function(BuildContext, String?)? valueControllerValidator;
  // State field(s) for sort widget.
  TextEditingController? sortController;
  String? Function(BuildContext, String?)? sortControllerValidator;
  // State field(s) for Switch widget.
  bool? switchValue;
  // Stores action output result for [Backend Call - API (deleteSysDictionaryDetail)] action in Button widget.
  ApiCallResponse? deleteRes;
  // Stores action output result for [Backend Call - API (updateSysDictionaryDetail)] action in Button widget.
  ApiCallResponse? updateRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    groupController?.dispose();
    labelController?.dispose();
    valueController?.dispose();
    sortController?.dispose();
  }

  /// Additional helper methods are added here.
}
