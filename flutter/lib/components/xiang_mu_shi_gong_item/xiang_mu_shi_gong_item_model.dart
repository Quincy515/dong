import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';

class XiangMuShiGongItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for xiangMuId widget.
  TextEditingController? xiangMuIdController;
  String? Function(BuildContext, String?)? xiangMuIdControllerValidator;
  DateTime? datePicked1;
  DateTime? datePicked2;
  // State field(s) for gongZuoMingCheng widget.
  TextEditingController? gongZuoMingChengController;
  String? Function(BuildContext, String?)? gongZuoMingChengControllerValidator;
  // State field(s) for jiHuaTianShu widget.
  TextEditingController? jiHuaTianShuController;
  String? Function(BuildContext, String?)? jiHuaTianShuControllerValidator;
  // State field(s) for zhanZongGongZuoLiangBaiFenBi widget.
  TextEditingController? zhanZongGongZuoLiangBaiFenBiController;
  String? Function(BuildContext, String?)?
      zhanZongGongZuoLiangBaiFenBiControllerValidator;
  // State field(s) for wanChengBi widget.
  TextEditingController? wanChengBiController;
  String? Function(BuildContext, String?)? wanChengBiControllerValidator;
  // State field(s) for shunXu widget.
  TextEditingController? shunXuController;
  String? Function(BuildContext, String?)? shunXuControllerValidator;
  // State field(s) for zhuangTai widget.
  bool? zhuangTaiValue;
  // State field(s) for bei_zhu widget.
  TextEditingController? beiZhuController;
  String? Function(BuildContext, String?)? beiZhuControllerValidator;
  // Stores action output result for [Backend Call - API (deleteXiangMuShiGongBiao)] action in Button widget.
  ApiCallResponse? deleteRes;
  // Stores action output result for [Backend Call - API (updateXiangMuShiGongBiao)] action in Button widget.
  ApiCallResponse? updateRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    groupController?.dispose();
    xiangMuIdController?.dispose();
    gongZuoMingChengController?.dispose();
    jiHuaTianShuController?.dispose();
    zhanZongGongZuoLiangBaiFenBiController?.dispose();
    wanChengBiController?.dispose();
    shunXuController?.dispose();
    beiZhuController?.dispose();
  }

  /// Additional helper methods are added here.
}
