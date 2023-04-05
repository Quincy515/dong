import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/form_field_controller.dart';
import 'package:flutter/material.dart';

class HeTongWenJianItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for xiangMuId widget.
  TextEditingController? xiangMuIdController;
  String? Function(BuildContext, String?)? xiangMuIdControllerValidator;
  // State field(s) for heTongBianHao widget.
  TextEditingController? heTongBianHaoController;
  String? Function(BuildContext, String?)? heTongBianHaoControllerValidator;
  // State field(s) for heTongMingCheng widget.
  TextEditingController? heTongMingChengController;
  String? Function(BuildContext, String?)? heTongMingChengControllerValidator;
  DateTime? datePicked1;
  // State field(s) for qianDingRen widget.
  int? qianDingRenValue;
  FormFieldController<int>? qianDingRenController;
  // State field(s) for heTongLeiBie widget.
  int? heTongLeiBieValue;
  FormFieldController<int>? heTongLeiBieController;
  // State field(s) for heTongLeiXing widget.
  int? heTongLeiXingValue;
  FormFieldController<int>? heTongLeiXingController;
  // State field(s) for heTongJinE widget.
  TextEditingController? heTongJinEController;
  String? Function(BuildContext, String?)? heTongJinEControllerValidator;
  // State field(s) for jiaFangGongSi widget.
  int? jiaFangGongSiValue;
  FormFieldController<int>? jiaFangGongSiController;
  // State field(s) for yiFangGongSi widget.
  int? yiFangGongSiValue;
  FormFieldController<int>? yiFangGongSiController;
  // State field(s) for fuKuanFangShi widget.
  int? fuKuanFangShiValue;
  FormFieldController<int>? fuKuanFangShiController;
  // State field(s) for jieSuanFangShi widget.
  int? jieSuanFangShiValue;
  FormFieldController<int>? jieSuanFangShiController;
  // State field(s) for shouFuKuanTiaoJian widget.
  TextEditingController? shouFuKuanTiaoJianController;
  String? Function(BuildContext, String?)?
      shouFuKuanTiaoJianControllerValidator;
  // State field(s) for tianJiaRen widget.
  int? tianJiaRenValue;
  FormFieldController<int>? tianJiaRenController;
  // State field(s) for zhuYaoTiaoKuan widget.
  TextEditingController? zhuYaoTiaoKuanController;
  String? Function(BuildContext, String?)? zhuYaoTiaoKuanControllerValidator;
  DateTime? datePicked2;
  // State field(s) for heTongXiaZaiDiZhi widget.
  TextEditingController? heTongXiaZaiDiZhiController;
  String? Function(BuildContext, String?)? heTongXiaZaiDiZhiControllerValidator;
  // State field(s) for heTongZhuangTai widget.
  int? heTongZhuangTaiValue;
  FormFieldController<int>? heTongZhuangTaiController;
  // State field(s) for faPiaoLeiXing widget.
  int? faPiaoLeiXingValue;
  FormFieldController<int>? faPiaoLeiXingController;
  // State field(s) for shouKuanZhangHao widget.
  TextEditingController? shouKuanZhangHaoController;
  String? Function(BuildContext, String?)? shouKuanZhangHaoControllerValidator;
  // State field(s) for yinHangMingCheng widget.
  TextEditingController? yinHangMingChengController;
  String? Function(BuildContext, String?)? yinHangMingChengControllerValidator;
  // State field(s) for kaiHuHang widget.
  TextEditingController? kaiHuHangController;
  String? Function(BuildContext, String?)? kaiHuHangControllerValidator;
  // State field(s) for kaiPiaoDanWei widget.
  TextEditingController? kaiPiaoDanWeiController;
  String? Function(BuildContext, String?)? kaiPiaoDanWeiControllerValidator;
  // State field(s) for bei_zhu widget.
  TextEditingController? beiZhuController;
  String? Function(BuildContext, String?)? beiZhuControllerValidator;
  // Stores action output result for [Backend Call - API (deleteHeTongWenJianBiao)] action in Button widget.
  ApiCallResponse? deleteRes;
  // Stores action output result for [Backend Call - API (updateHeTongWenJianBiao)] action in Button widget.
  ApiCallResponse? updateRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    groupController?.dispose();
    xiangMuIdController?.dispose();
    heTongBianHaoController?.dispose();
    heTongMingChengController?.dispose();
    heTongJinEController?.dispose();
    shouFuKuanTiaoJianController?.dispose();
    zhuYaoTiaoKuanController?.dispose();
    heTongXiaZaiDiZhiController?.dispose();
    shouKuanZhangHaoController?.dispose();
    yinHangMingChengController?.dispose();
    kaiHuHangController?.dispose();
    kaiPiaoDanWeiController?.dispose();
    beiZhuController?.dispose();
  }

  /// Additional helper methods are added here.
}
