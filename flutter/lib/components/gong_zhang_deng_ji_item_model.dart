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

class GongZhangDengJiItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for gong_zhang_ming_cheng widget.
  TextEditingController? gongZhangMingChengController;
  String? Function(BuildContext, String?)?
      gongZhangMingChengControllerValidator;
  // State field(s) for shu_liang widget.
  TextEditingController? shuLiangController;
  String? Function(BuildContext, String?)? shuLiangControllerValidator;
  // State field(s) for shi_xiang widget.
  TextEditingController? shiXiangController;
  String? Function(BuildContext, String?)? shiXiangControllerValidator;
  // State field(s) for shen_qing_ren widget.
  TextEditingController? shenQingRenController;
  String? Function(BuildContext, String?)? shenQingRenControllerValidator;
  DateTime? datePicked1;
  // State field(s) for jing_shou_ren widget.
  TextEditingController? jingShouRenController;
  String? Function(BuildContext, String?)? jingShouRenControllerValidator;
  DateTime? datePicked2;
  // State field(s) for bei_zhu widget.
  TextEditingController? beiZhuController;
  String? Function(BuildContext, String?)? beiZhuControllerValidator;
  // State field(s) for xiang_qing widget.
  TextEditingController? xiangQingController;
  String? Function(BuildContext, String?)? xiangQingControllerValidator;
  // Stores action output result for [Backend Call - API (deleteGongZhangDengJi)] action in Button widget.
  ApiCallResponse? deleteRes;
  // Stores action output result for [Backend Call - API (updateGongZhangDengJi)] action in Button widget.
  ApiCallResponse? updateRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    groupController?.dispose();
    gongZhangMingChengController?.dispose();
    shuLiangController?.dispose();
    shiXiangController?.dispose();
    shenQingRenController?.dispose();
    jingShouRenController?.dispose();
    beiZhuController?.dispose();
    xiangQingController?.dispose();
  }

  /// Additional helper methods are added here.
}
