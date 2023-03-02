// ignore_for_file: unused_import

import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_drop_down.dart';
import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/flutter_flow_widgets.dart';
import 'package:styled_divider/styled_divider.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:provider/provider.dart';

class ZhengShuJieChuJiLuItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for zhengShuId widget.
  TextEditingController? zhengShuIdController;
  String? Function(BuildContext, String?)? zhengShuIdControllerValidator;
  // State field(s) for zhengShuMingCheng widget.
  TextEditingController? zhengShuMingChengController;
  String? Function(BuildContext, String?)? zhengShuMingChengControllerValidator;
  // State field(s) for zhengShuBianHao widget.
  TextEditingController? zhengShuBianHaoController;
  String? Function(BuildContext, String?)? zhengShuBianHaoControllerValidator;
  // State field(s) for DropDown widget.
  int? dropDownValue1;
  // State field(s) for DropDown widget.
  int? dropDownValue2;
  DateTime? datePicked1;
  // State field(s) for jieChuShiYou widget.
  TextEditingController? jieChuShiYouController;
  String? Function(BuildContext, String?)? jieChuShiYouControllerValidator;
  // State field(s) for jieChuBeiZhu widget.
  TextEditingController? jieChuBeiZhuController;
  String? Function(BuildContext, String?)? jieChuBeiZhuControllerValidator;
  DateTime? datePicked2;
  // State field(s) for guiHaiBeiZhu widget.
  TextEditingController? guiHaiBeiZhuController;
  String? Function(BuildContext, String?)? guiHaiBeiZhuControllerValidator;
  DateTime? datePicked3;
  // Stores action output result for [Backend Call - API (deleteZhengShuJieChuJiLuBia)] action in Button widget.
  ApiCallResponse? deleteRes;
  // Stores action output result for [Backend Call - API (updateZhengShuJieChuJiLuBiao)] action in Button widget.
  ApiCallResponse? updateRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    groupController?.dispose();
    zhengShuIdController?.dispose();
    zhengShuMingChengController?.dispose();
    zhengShuBianHaoController?.dispose();
    jieChuShiYouController?.dispose();
    jieChuBeiZhuController?.dispose();
    guiHaiBeiZhuController?.dispose();
  }

  /// Additional helper methods are added here.
}
