// ignore_for_file: unused_import

import '/flutter_flow/flutter_flow_drop_down.dart';
import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/flutter_flow_widgets.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:provider/provider.dart';

class PageSearchAndNextModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  // State field(s) for search widget.
  TextEditingController? searchController;
  String? Function(BuildContext, String?)? searchControllerValidator;
  // State field(s) for TextField widget.
  TextEditingController? textController2;
  String? Function(BuildContext, String?)? textController2Validator;
  // State field(s) for DropDown widget.
  String? dropDownValue;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    searchController?.dispose();
    textController2?.dispose();
  }

  /// Additional helper methods are added here.
}
