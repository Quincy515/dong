// ignore_for_file: unused_import

import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/flutter_flow_widgets.dart';
import 'package:styled_divider/styled_divider.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:provider/provider.dart';

class MeiRiKaoQinItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for xing_ming widget.
  TextEditingController? xingMingController;
  String? Function(BuildContext, String?)? xingMingControllerValidator;
  // State field(s) for Checkbox widget.
  bool? checkboxValue;
  // State field(s) for xiu_jia_qing_jia_chu_cha widget.
  TextEditingController? xiuJiaQingJiaChuChaController;
  String? Function(BuildContext, String?)?
      xiuJiaQingJiaChuChaControllerValidator;
  // State field(s) for xiang_qing widget.
  TextEditingController? xiangQingController;
  String? Function(BuildContext, String?)? xiangQingControllerValidator;
  DateTime? datePicked;
  // State field(s) for bei_zhu widget.
  TextEditingController? beiZhuController;
  String? Function(BuildContext, String?)? beiZhuControllerValidator;
  // Stores action output result for [Backend Call - API (deleteMeiRiKaoQin)] action in Button widget.
  ApiCallResponse? deleteRes;
  // Stores action output result for [Backend Call - API (updateMeiRiKaoQin)] action in Button widget.
  ApiCallResponse? updateRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    groupController?.dispose();
    xingMingController?.dispose();
    xiuJiaQingJiaChuChaController?.dispose();
    xiangQingController?.dispose();
    beiZhuController?.dispose();
  }

  /// Additional helper methods are added here.
}