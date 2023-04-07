import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/form_field_controller.dart';
import 'package:flutter/material.dart';

class WenJianZiLiaoJieDianItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for xiangMuId widget.
  TextEditingController? xiangMuIdController;
  String? Function(BuildContext, String?)? xiangMuIdControllerValidator;
  // State field(s) for wenJianJiaMingCheng widget.
  TextEditingController? wenJianJiaMingChengController;
  String? Function(BuildContext, String?)?
      wenJianJiaMingChengControllerValidator;
  // State field(s) for fuJieDianId widget.
  TextEditingController? fuJieDianIdController;
  String? Function(BuildContext, String?)? fuJieDianIdControllerValidator;
  // State field(s) for Switch widget.
  bool? switchValue1;
  // State field(s) for paiXu widget.
  TextEditingController? paiXuController;
  String? Function(BuildContext, String?)? paiXuControllerValidator;
  // State field(s) for tianJiaRen widget.
  int? tianJiaRenValue;
  FormFieldController<int>? tianJiaRenController;
  DateTime? datePicked;
  // State field(s) for Switch widget.
  bool? switchValue2;
  // State field(s) for bei_zhu widget.
  TextEditingController? beiZhuController;
  String? Function(BuildContext, String?)? beiZhuControllerValidator;
  // Stores action output result for [Backend Call - API (deleteWenJianZiLiaoJieDianBiao)] action in Button widget.
  ApiCallResponse? deleteRes;
  // Stores action output result for [Backend Call - API (updateWenJianZiLiaoJieDianBiao)] action in Button widget.
  ApiCallResponse? updateRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    groupController?.dispose();
    xiangMuIdController?.dispose();
    wenJianJiaMingChengController?.dispose();
    fuJieDianIdController?.dispose();
    paiXuController?.dispose();
    beiZhuController?.dispose();
  }

  /// Additional helper methods are added here.
}
