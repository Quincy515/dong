import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/form_field_controller.dart';
import 'package:flutter/material.dart';

class QingJiaGuanLiItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for xingMingId widget.
  int? xingMingIdValue;
  FormFieldController<int>? xingMingIdValueController;
  // State field(s) for leiXing widget.
  int? leiXingValue;
  FormFieldController<int>? leiXingValueController;
  // State field(s) for tian_shu widget.
  TextEditingController? tianShuController;
  String? Function(BuildContext, String?)? tianShuControllerValidator;
  // State field(s) for jinDu widget.
  TextEditingController? jinDuController;
  String? Function(BuildContext, String?)? jinDuControllerValidator;
  // State field(s) for shiYou widget.
  TextEditingController? shiYouController;
  String? Function(BuildContext, String?)? shiYouControllerValidator;
  DateTime? datePicked1;
  DateTime? datePicked2;
  // State field(s) for shenHeRen widget.
  int? shenHeRenValue;
  FormFieldController<int>? shenHeRenValueController;
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
  // Stores action output result for [Backend Call - API (updateQingJiaGuanLi)] action in Button widget.
  ApiCallResponse? disagreeRes;
  // Stores action output result for [Backend Call - API (findXiaoXiTongZhi)] action in Button widget.
  ApiCallResponse? findXiaoXiRes;
  // Stores action output result for [Backend Call - API (createXiaoXiTongZhi)] action in Button widget.
  ApiCallResponse? createXiaoXiRes;
  // Stores action output result for [Backend Call - API (updateQingJiaGuanLi)] action in Button widget.
  ApiCallResponse? agreeRes;
  // Stores action output result for [Backend Call - API (createXiaoXiTongZhi)] action in Button widget.
  ApiCallResponse? createXiaoXiSuccessRes;
  // Stores action output result for [Backend Call - API (deleteQingJiaGuanLi)] action in Button widget.
  ApiCallResponse? deleteRes;
  // Stores action output result for [Backend Call - API (updateQingJiaGuanLi)] action in Button widget.
  ApiCallResponse? updateRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    groupController?.dispose();
    tianShuController?.dispose();
    jinDuController?.dispose();
    shiYouController?.dispose();
    shenHeZhuangTaiController?.dispose();
    shenHeXiangQingController?.dispose();
    beiZhuController?.dispose();
  }

  /// Additional helper methods are added here.
}
