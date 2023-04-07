import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';

class CaiGouJiHuaMingXiItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for caiLiaoId widget.
  TextEditingController? caiLiaoIdController;
  String? Function(BuildContext, String?)? caiLiaoIdControllerValidator;
  // State field(s) for caiGouJiHuaId widget.
  TextEditingController? caiGouJiHuaIdController;
  String? Function(BuildContext, String?)? caiGouJiHuaIdControllerValidator;
  // State field(s) for gongYingShangId widget.
  TextEditingController? gongYingShangIdController;
  String? Function(BuildContext, String?)? gongYingShangIdControllerValidator;
  // State field(s) for danJia widget.
  TextEditingController? danJiaController;
  String? Function(BuildContext, String?)? danJiaControllerValidator;
  // State field(s) for ruKuLiang widget.
  TextEditingController? ruKuLiangController;
  String? Function(BuildContext, String?)? ruKuLiangControllerValidator;
  // State field(s) for caiGouLiang widget.
  TextEditingController? caiGouLiangController;
  String? Function(BuildContext, String?)? caiGouLiangControllerValidator;
  // State field(s) for zongJia widget.
  TextEditingController? zongJiaController;
  String? Function(BuildContext, String?)? zongJiaControllerValidator;
  // State field(s) for Switch widget.
  bool? switchValue;
  // State field(s) for bei_zhu widget.
  TextEditingController? beiZhuController;
  String? Function(BuildContext, String?)? beiZhuControllerValidator;
  // Stores action output result for [Backend Call - API (deleteMeiRiKaoQin)] action in Button widget.
  ApiCallResponse? deleteRes;
  // Stores action output result for [Backend Call - API (updateMeiRiKaoQin)] action in Button widget.
  ApiCallResponse? updateRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    groupController?.dispose();
    caiLiaoIdController?.dispose();
    caiGouJiHuaIdController?.dispose();
    gongYingShangIdController?.dispose();
    danJiaController?.dispose();
    ruKuLiangController?.dispose();
    caiGouLiangController?.dispose();
    zongJiaController?.dispose();
    beiZhuController?.dispose();
  }

  /// Additional helper methods are added here.
}
