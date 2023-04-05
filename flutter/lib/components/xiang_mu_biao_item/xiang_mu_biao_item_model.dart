import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';
import '../../flutter_flow/form_field_controller.dart';

class XiangMuBiaoItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for xiangMuBianHao widget.
  TextEditingController? xiangMuBianHaoController;
  String? Function(BuildContext, String?)? xiangMuBianHaoControllerValidator;
  DateTime? datePicked1;
  DateTime? datePicked2;
  // State field(s) for gongChengZaoJia widget.
  TextEditingController? gongChengZaoJiaController;
  String? Function(BuildContext, String?)? gongChengZaoJiaControllerValidator;
  // State field(s) for nongMinGongGongZiBaoZhangJin widget.
  TextEditingController? nongMinGongGongZiBaoZhangJinController;
  String? Function(BuildContext, String?)?
      nongMinGongGongZiBaoZhangJinControllerValidator;
  // State field(s) for baoZhangJinZhuangTai widget.
  int? baoZhangJinZhuangTaiValue;
  FormFieldController<int>? baoZhangJinZhuangTaiController;
  // State field(s) for gongZuoBaoXianJinE widget.
  TextEditingController? gongZuoBaoXianJinEController;
  String? Function(BuildContext, String?)?
      gongZuoBaoXianJinEControllerValidator;
  // State field(s) for baoZhengJinZhuangTai widget.
  int? baoZhengJinZhuangTaiValue;
  FormFieldController<int>? baoZhengJinZhuangTaiController;
  // State field(s) for luYueBaoZhengJin widget.
  TextEditingController? luYueBaoZhengJinController;
  String? Function(BuildContext, String?)? luYueBaoZhengJinControllerValidator;
  // State field(s) for fuKuanFangShi widget.
  int? fuKuanFangShiValue;
  FormFieldController<int>? fuKuanFangShiController;
  // State field(s) for yuQiLiRun widget.
  TextEditingController? yuQiLiRunController;
  String? Function(BuildContext, String?)? yuQiLiRunControllerValidator;
  DateTime? datePicked3;
  DateTime? datePicked4;
  // State field(s) for jieSuanZhuangTai widget.
  int? jieSuanZhuangTaiValue;
  FormFieldController<int>? jieSuanZhuangTaiController;
  // State field(s) for xiangMuZhuangTai widget.
  int? xiangMuZhuangTaiValue;
  FormFieldController<int>? xiangMuZhuangTaiController;
  // State field(s) for xiangMuJingLi widget.
  int? xiangMuJingLiValue;
  FormFieldController<int>? xiangMuJingLiController;
  // State field(s) for xiangMuZhuGuan widget.
  int? xiangMuZhuGuanValue;
  FormFieldController<int>? xiangMuZhuGuanController;
  // State field(s) for bei_zhu widget.
  TextEditingController? beiZhuController;
  String? Function(BuildContext, String?)? beiZhuControllerValidator;
  // Stores action output result for [Backend Call - API (deleteXiangMuBiao)] action in Button widget.
  ApiCallResponse? deleteRes;
  // Stores action output result for [Backend Call - API (updateXiangMuBiao)] action in Button widget.
  ApiCallResponse? updateRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    groupController?.dispose();
    xiangMuBianHaoController?.dispose();
    gongChengZaoJiaController?.dispose();
    nongMinGongGongZiBaoZhangJinController?.dispose();
    gongZuoBaoXianJinEController?.dispose();
    luYueBaoZhengJinController?.dispose();
    yuQiLiRunController?.dispose();
    beiZhuController?.dispose();
  }

  /// Additional helper methods are added here.
}
