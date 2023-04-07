import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';

class SheBeiMingXiItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for sheBeiId widget.
  TextEditingController? sheBeiIdController;
  String? Function(BuildContext, String?)? sheBeiIdControllerValidator;
  // State field(s) for zuLinDanJuId widget.
  TextEditingController? zuLinDanJuIdController;
  String? Function(BuildContext, String?)? zuLinDanJuIdControllerValidator;
  // State field(s) for zuLinYongTu widget.
  TextEditingController? zuLinYongTuController;
  String? Function(BuildContext, String?)? zuLinYongTuControllerValidator;
  TextEditingController? danJiaController;
  String? Function(BuildContext, String?)? danJiaControllerValidator;
  // State field(s) for zuLinShuLiang widget.
  TextEditingController? zuLinShuLiangController;
  String? Function(BuildContext, String?)? zuLinShuLiangControllerValidator;
  DateTime? datePicked1;
  DateTime? datePicked2;
  DateTime? datePicked3;
  DateTime? datePicked4;
  // State field(s) for zuLinTianShu widget.
  TextEditingController? zuLinTianShuController;
  String? Function(BuildContext, String?)? zuLinTianShuControllerValidator;
  // State field(s) for Switch widget.
  bool? switchValue;
  // State field(s) for bei_zhu widget.
  TextEditingController? beiZhuController;
  String? Function(BuildContext, String?)? beiZhuControllerValidator;
  // Stores action output result for [Backend Call - API (deleteSheBeiMingXiBiao)] action in Button widget.
  ApiCallResponse? deleteRes;
  // Stores action output result for [Backend Call - API (updateSheBeiMingXiBiao)] action in Button widget.
  ApiCallResponse? updateRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    danJiaController?.dispose();
    groupController?.dispose();
    sheBeiIdController?.dispose();
    zuLinDanJuIdController?.dispose();
    zuLinYongTuController?.dispose();
    zuLinShuLiangController?.dispose();
    zuLinTianShuController?.dispose();
    beiZhuController?.dispose();
  }

  /// Additional helper methods are added here.
}
