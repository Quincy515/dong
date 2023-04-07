import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/form_field_controller.dart';
import 'package:flutter/material.dart';

class SheBeiZuLinItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for xiangMuIdxiangMuId widget.
  TextEditingController? xiangMuIdxiangMuIdController;
  String? Function(BuildContext, String?)?
      xiangMuIdxiangMuIdControllerValidator;
  // State field(s) for zuLinId widget.
  TextEditingController? zuLinIdController;
  String? Function(BuildContext, String?)? zuLinIdControllerValidator;
  // State field(s) for zuLinYongTu widget.
  TextEditingController? zuLinYongTuController;
  String? Function(BuildContext, String?)? zuLinYongTuControllerValidator;
  // State field(s) for tianJiaRen widget.
  int? tianJiaRenValue;
  FormFieldController<int>? tianJiaRenController;
  DateTime? datePicked;
  // State field(s) for Switch widget.
  bool? switchValue;
  // State field(s) for bei_zhu widget.
  TextEditingController? beiZhuController;
  String? Function(BuildContext, String?)? beiZhuControllerValidator;
  // Stores action output result for [Backend Call - API (deleteSheBeiZuLinBiao)] action in Button widget.
  ApiCallResponse? deleteRes;
  // Stores action output result for [Backend Call - API (updateSheBeiZuLinBiao)] action in Button widget.
  ApiCallResponse? updateRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    groupController?.dispose();
    xiangMuIdxiangMuIdController?.dispose();
    zuLinIdController?.dispose();
    zuLinYongTuController?.dispose();
    beiZhuController?.dispose();
  }

  /// Additional helper methods are added here.
}
