import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/form_field_controller.dart';
import 'package:flutter/material.dart';

class WenJianZiLiaoItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for wenJianId widget.
  TextEditingController? wenJianIdController;
  String? Function(BuildContext, String?)? wenJianIdControllerValidator;
  // State field(s) for wenJianMingCheng widget.
  TextEditingController? wenJianMingChengController;
  String? Function(BuildContext, String?)? wenJianMingChengControllerValidator;
  // State field(s) for jieDianId widget.
  TextEditingController? jieDianIdController;
  String? Function(BuildContext, String?)? jieDianIdControllerValidator;
  // State field(s) for shangChuanRen widget.
  int? shangChuanRenValue;
  FormFieldController<int>? shangChuanRenController;
  DateTime? datePicked1;
  // State field(s) for Switch widget.
  bool? switchValue1;
  // State field(s) for wenJianLeiXing widget.
  int? wenJianLeiXingValue;
  FormFieldController<int>? wenJianLeiXingController;
  // State field(s) for yueDuCiShu widget.
  TextEditingController? yueDuCiShuController;
  String? Function(BuildContext, String?)? yueDuCiShuControllerValidator;
  DateTime? datePicked2;
  // State field(s) for shenHeRen widget.
  int? shenHeRenValue;
  FormFieldController<int>? shenHeRenController;
  // State field(s) for Switch widget.
  bool? switchValue2;
  // State field(s) for bei_zhu widget.
  TextEditingController? beiZhuController;
  String? Function(BuildContext, String?)? beiZhuControllerValidator;
  // Stores action output result for [Backend Call - API (deleteWenJianZiLiaoBiao)] action in Button widget.
  ApiCallResponse? deleteRes;
  // Stores action output result for [Backend Call - API (updateWenJianZiLiaoBiao)] action in Button widget.
  ApiCallResponse? updateRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    groupController?.dispose();
    wenJianIdController?.dispose();
    wenJianMingChengController?.dispose();
    jieDianIdController?.dispose();
    yueDuCiShuController?.dispose();
    beiZhuController?.dispose();
  }

  /// Additional helper methods are added here.
}
