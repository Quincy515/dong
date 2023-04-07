import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/form_field_controller.dart';
import 'package:flutter/material.dart';

class XiangMuZhuCaItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for xiangMuMingCheng widget.
  int? xiangMuMingChengValue;
  FormFieldController<int>? xiangMuMingChengController;
  // State field(s) for caiLiaoBianHao widget.
  TextEditingController? caiLiaoBianHaoController;
  String? Function(BuildContext, String?)? caiLiaoBianHaoControllerValidator;
  // State field(s) for caiLiaoMingCheng widget.
  TextEditingController? caiLiaoMingChengController;
  String? Function(BuildContext, String?)? caiLiaoMingChengControllerValidator;
  // State field(s) for danWei widget.
  TextEditingController? danWeiController;
  String? Function(BuildContext, String?)? danWeiControllerValidator;
  // State field(s) for yongLiang widget.
  TextEditingController? yongLiangController;
  String? Function(BuildContext, String?)? yongLiangControllerValidator;
  // State field(s) for danJia widget.
  TextEditingController? danJiaController;
  String? Function(BuildContext, String?)? danJiaControllerValidator;
  // State field(s) for xianHangJiaGe widget.
  TextEditingController? xianHangJiaGeController;
  String? Function(BuildContext, String?)? xianHangJiaGeControllerValidator;
  // State field(s) for jiaChaHeJi widget.
  TextEditingController? jiaChaHeJiController;
  String? Function(BuildContext, String?)? jiaChaHeJiControllerValidator;
  // State field(s) for heJi widget.
  TextEditingController? heJiController;
  String? Function(BuildContext, String?)? heJiControllerValidator;
  // State field(s) for bei_zhu widget.
  TextEditingController? beiZhuController;
  String? Function(BuildContext, String?)? beiZhuControllerValidator;
  // Stores action output result for [Backend Call - API (deleteXiangMuZhuCaiBiao)] action in Button widget.
  ApiCallResponse? deleteRes;
  // Stores action output result for [Backend Call - API (updateXiangMuZhuCaiBiao)] action in Button widget.
  ApiCallResponse? updateRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    groupController?.dispose();
    caiLiaoBianHaoController?.dispose();
    caiLiaoMingChengController?.dispose();
    danWeiController?.dispose();
    yongLiangController?.dispose();
    danJiaController?.dispose();
    xianHangJiaGeController?.dispose();
    jiaChaHeJiController?.dispose();
    heJiController?.dispose();
    beiZhuController?.dispose();
  }

  /// Additional helper methods are added here.
}
