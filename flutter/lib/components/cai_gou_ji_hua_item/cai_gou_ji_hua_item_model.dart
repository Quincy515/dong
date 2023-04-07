import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/form_field_controller.dart';
import 'package:flutter/material.dart';

class CaiGouJiHuaItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for xiangMuId widget.
  TextEditingController? xiangMuIdController;
  String? Function(BuildContext, String?)? xiangMuIdControllerValidator;
  // State field(s) for caiGouJiHuaMingCheng widget.
  TextEditingController? caiGouJiHuaMingChengController;
  String? Function(BuildContext, String?)?
      caiGouJiHuaMingChengControllerValidator;
  // State field(s) for caiGouFuZeRen widget.
  int? caiGouFuZeRenValue;
  FormFieldController<int>? caiGouFuZeRenController;
  DateTime? datePicked;
  // State field(s) for leiJiCaiGouJinE widget.
  TextEditingController? leiJiCaiGouJinEController;
  String? Function(BuildContext, String?)? leiJiCaiGouJinEControllerValidator;
  // State field(s) for shiJiZhiFuJinE widget.
  TextEditingController? shiJiZhiFuJinEController;
  String? Function(BuildContext, String?)? shiJiZhiFuJinEControllerValidator;
  // State field(s) for Switch widget.
  bool? switchValue;
  // State field(s) for bei_zhu widget.
  TextEditingController? beiZhuController;
  String? Function(BuildContext, String?)? beiZhuControllerValidator;
  // Stores action output result for [Backend Call - API (deleteCaiGouJiHuaBiao)] action in Button widget.
  ApiCallResponse? deleteRes;
  // Stores action output result for [Backend Call - API (updateCaiGouJiHuaBiao)] action in Button widget.
  ApiCallResponse? updateRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    groupController?.dispose();
    xiangMuIdController?.dispose();
    caiGouJiHuaMingChengController?.dispose();
    leiJiCaiGouJinEController?.dispose();
    shiJiZhiFuJinEController?.dispose();
    beiZhuController?.dispose();
  }

  /// Additional helper methods are added here.
}
