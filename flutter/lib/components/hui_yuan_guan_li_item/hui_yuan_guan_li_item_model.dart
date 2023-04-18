import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/form_field_controller.dart';
import 'package:flutter/material.dart';

class HuiYuanGuanLiItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for huiYuanLeiBie widget.
  int? huiYuanLeiBieValue;
  FormFieldController<int>? huiYuanLeiBieController;
  // State field(s) for huiYuanJiBie widget.
  int? huiYuanJiBieValue;
  FormFieldController<int>? huiYuanJiBieController;
  // State field(s) for hui_fei widget.
  TextEditingController? huiFeiController;
  String? Function(BuildContext, String?)? huiFeiControllerValidator;
  // State field(s) for wang_zhi widget.
  TextEditingController? wangZhiController;
  String? Function(BuildContext, String?)? wangZhiControllerValidator;
  // State field(s) for ru_hui_lian_xi widget.
  TextEditingController? ruHuiLianXiController;
  String? Function(BuildContext, String?)? ruHuiLianXiControllerValidator;
  DateTime? datePicked1;
  DateTime? datePicked2;
  DateTime? datePicked3;
  DateTime? datePicked4;
  // State field(s) for beiZhu widget.
  TextEditingController? beiZhuController;
  String? Function(BuildContext, String?)? beiZhuControllerValidator;
  // State field(s) for guoQiTiXing widget.
  bool? guoQiTiXingValue;
  // Stores action output result for [Backend Call - API (deleteHuiYuanGuanLi)] action in Button widget.
  ApiCallResponse? deleteRes;
  // Stores action output result for [Backend Call - API (updateHuiYuanGuanLi)] action in Button widget.
  ApiCallResponse? updateRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    groupController?.dispose();
    huiFeiController?.dispose();
    wangZhiController?.dispose();
    ruHuiLianXiController?.dispose();
    beiZhuController?.dispose();
  }

  /// Additional helper methods are added here.
}
