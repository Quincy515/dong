import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/form_field_controller.dart';
import 'package:flutter/material.dart';

class ShiGongJinDuItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for xiangMuId widget.
  TextEditingController? xiangMuIdController;
  String? Function(BuildContext, String?)? xiangMuIdControllerValidator;
  // State field(s) for gongZuoMingChengId widget.
  TextEditingController? gongZuoMingChengIdController;
  String? Function(BuildContext, String?)?
      gongZuoMingChengIdControllerValidator;
  // State field(s) for wanChengBaiFenBi widget.
  TextEditingController? wanChengBaiFenBiController;
  String? Function(BuildContext, String?)? wanChengBaiFenBiControllerValidator;
  DateTime? datePicked1;
  DateTime? datePicked2;
  // State field(s) for tianJiaRen widget.
  int? tianJiaRenValue;
  FormFieldController<int>? tianJiaRenController;
  DateTime? datePicked3;
  // State field(s) for shiJiTianShu widget.
  TextEditingController? shiJiTianShuController;
  String? Function(BuildContext, String?)? shiJiTianShuControllerValidator;
  // State field(s) for Switch widget.
  bool? switchValue;
  // State field(s) for bei_zhu widget.
  TextEditingController? beiZhuController;
  String? Function(BuildContext, String?)? beiZhuControllerValidator;
  // Stores action output result for [Backend Call - API (deleteShiGongJinDuBiao)] action in Button widget.
  ApiCallResponse? deleteRes;
  // Stores action output result for [Backend Call - API (updateShiGongJinDuBiaov)] action in Button widget.
  ApiCallResponse? updateRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    groupController?.dispose();
    xiangMuIdController?.dispose();
    gongZuoMingChengIdController?.dispose();
    wanChengBaiFenBiController?.dispose();
    shiJiTianShuController?.dispose();
    beiZhuController?.dispose();
  }

  /// Additional helper methods are added here.
}
