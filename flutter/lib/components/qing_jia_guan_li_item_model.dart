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

class QingJiaGuanLiItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for xing_ming widget.
  TextEditingController? xingMingController;
  String? Function(BuildContext, String?)? xingMingControllerValidator;
  // State field(s) for lei_xing widget.
  TextEditingController? leiXingController;
  String? Function(BuildContext, String?)? leiXingControllerValidator;
  // State field(s) for tian_shu widget.
  TextEditingController? tianShuController;
  String? Function(BuildContext, String?)? tianShuControllerValidator;
  // State field(s) for shi_you widget.
  TextEditingController? shiYouController;
  String? Function(BuildContext, String?)? shiYouControllerValidator;
  // State field(s) for jin_du widget.
  TextEditingController? jinDuController;
  String? Function(BuildContext, String?)? jinDuControllerValidator;
  DateTime? datePicked1;
  DateTime? datePicked2;
  // State field(s) for shen_he_ren widget.
  TextEditingController? shenHeRenController;
  String? Function(BuildContext, String?)? shenHeRenControllerValidator;
  // State field(s) for shen_he_zhuang_tai widget.
  TextEditingController? shenHeZhuangTaiController;
  String? Function(BuildContext, String?)? shenHeZhuangTaiControllerValidator;
  DateTime? datePicked3;
  // State field(s) for shen_he_xiang_qing widget.
  TextEditingController? shenHeXiangQingController;
  String? Function(BuildContext, String?)? shenHeXiangQingControllerValidator;
  // State field(s) for bei_zhu widget.
  TextEditingController? beiZhuController;
  String? Function(BuildContext, String?)? beiZhuControllerValidator;
  // Stores action output result for [Backend Call - API (deleteQingJiaGuanLi)] action in Button widget.
  ApiCallResponse? deleteRes;
  // Stores action output result for [Backend Call - API (updateQingJiaGuanLi)] action in Button widget.
  ApiCallResponse? updateRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    groupController?.dispose();
    xingMingController?.dispose();
    leiXingController?.dispose();
    tianShuController?.dispose();
    shiYouController?.dispose();
    jinDuController?.dispose();
    shenHeRenController?.dispose();
    shenHeZhuangTaiController?.dispose();
    shenHeXiangQingController?.dispose();
    beiZhuController?.dispose();
  }

  /// Additional helper methods are added here.
}
