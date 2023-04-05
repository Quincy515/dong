import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';

class FenBaoShangYuXiangMuGuanLianItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for heTongBianHao widget.
  TextEditingController? heTongBianHaoController;
  String? Function(BuildContext, String?)? heTongBianHaoControllerValidator;
  // State field(s) for heTongMingCheng widget.
  TextEditingController? heTongMingChengController;
  String? Function(BuildContext, String?)? heTongMingChengControllerValidator;
  // State field(s) for xiangMuId widget.
  TextEditingController? xiangMuIdController;
  String? Function(BuildContext, String?)? xiangMuIdControllerValidator;
  // State field(s) for fenBaoShangId widget.
  TextEditingController? fenBaoShangIdController;
  String? Function(BuildContext, String?)? fenBaoShangIdControllerValidator;
  // State field(s) for zhuangTai widget.
  bool? zhuangTaiValue;
  // State field(s) for weiTuoRen widget.
  TextEditingController? weiTuoRenController;
  String? Function(BuildContext, String?)? weiTuoRenControllerValidator;
  // State field(s) for weiTuoRenShouJi widget.
  TextEditingController? weiTuoRenShouJiController;
  String? Function(BuildContext, String?)? weiTuoRenShouJiControllerValidator;
  DateTime? datePicked1;
  DateTime? datePicked2;
  // State field(s) for bei_zhu widget.
  TextEditingController? beiZhuController;
  String? Function(BuildContext, String?)? beiZhuControllerValidator;
  // Stores action output result for [Backend Call - API (deleteFenBaoShangYuXiangMuGuanLianBiao)] action in Button widget.
  ApiCallResponse? deleteRes;
  // Stores action output result for [Backend Call - API (updateFenBaoShangYuXiangMuGuanLianBiao)] action in Button widget.
  ApiCallResponse? updateRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    groupController?.dispose();
    heTongBianHaoController?.dispose();
    heTongMingChengController?.dispose();
    xiangMuIdController?.dispose();
    fenBaoShangIdController?.dispose();
    weiTuoRenController?.dispose();
    weiTuoRenShouJiController?.dispose();
    beiZhuController?.dispose();
  }

  /// Additional helper methods are added here.
}
