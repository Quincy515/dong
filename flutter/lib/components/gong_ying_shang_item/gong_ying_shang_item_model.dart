import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/form_field_controller.dart';
import 'package:flutter/material.dart';

class GongYingShangItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for gongYingShangMingCheng widget.
  int? gongYingShangMingChengValue;
  FormFieldController<int>? gongYingShangMingChengController;
  // State field(s) for faRenShenFenZheng widget.
  TextEditingController? faRenShenFenZhengController;
  String? Function(BuildContext, String?)? faRenShenFenZhengControllerValidator;
  // State field(s) for sheHuiXinYongDaiMa widget.
  TextEditingController? sheHuiXinYongDaiMaController;
  String? Function(BuildContext, String?)?
      sheHuiXinYongDaiMaControllerValidator;
  // State field(s) for ziZhiZhengZhao widget.
  TextEditingController? ziZhiZhengZhaoController;
  String? Function(BuildContext, String?)? ziZhiZhengZhaoControllerValidator;
  // State field(s) for gongSiDiZhi widget.
  TextEditingController? gongSiDiZhiController;
  String? Function(BuildContext, String?)? gongSiDiZhiControllerValidator;
  // State field(s) for lianXiRenlianXiRen widget.
  TextEditingController? lianXiRenlianXiRenController;
  String? Function(BuildContext, String?)?
      lianXiRenlianXiRenControllerValidator;
  // State field(s) for lianXiRenShouJi widget.
  TextEditingController? lianXiRenShouJiController;
  String? Function(BuildContext, String?)? lianXiRenShouJiControllerValidator;
  // State field(s) for faPiaoTaiTou widget.
  TextEditingController? faPiaoTaiTouController;
  String? Function(BuildContext, String?)? faPiaoTaiTouControllerValidator;
  // State field(s) for faPiaoLeiXing widget.
  int? faPiaoLeiXingValue;
  FormFieldController<int>? faPiaoLeiXingController;
  // State field(s) for yinHangMingCheng widget.
  TextEditingController? yinHangMingChengController;
  String? Function(BuildContext, String?)? yinHangMingChengControllerValidator;
  // State field(s) for kaiHuHang widget.
  TextEditingController? kaiHuHangController;
  String? Function(BuildContext, String?)? kaiHuHangControllerValidator;
  // State field(s) for zhangHao widget.
  TextEditingController? zhangHaoController;
  String? Function(BuildContext, String?)? zhangHaoControllerValidator;
  // State field(s) for tianJiaRen widget.
  int? tianJiaRenValue;
  FormFieldController<int>? tianJiaRenController;
  DateTime? datePicked;
  // State field(s) for zhuangTai widget.
  bool? zhuangTaiValue;
  // State field(s) for bei_zhu widget.
  TextEditingController? beiZhuController;
  String? Function(BuildContext, String?)? beiZhuControllerValidator;
  // Stores action output result for [Backend Call - API (deleteGongYingShangBiao)] action in Button widget.
  ApiCallResponse? deleteRes;
  // Stores action output result for [Backend Call - API (updateGongYingShangBiao)] action in Button widget.
  ApiCallResponse? updateRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    groupController?.dispose();
    faRenShenFenZhengController?.dispose();
    sheHuiXinYongDaiMaController?.dispose();
    ziZhiZhengZhaoController?.dispose();
    gongSiDiZhiController?.dispose();
    lianXiRenlianXiRenController?.dispose();
    lianXiRenShouJiController?.dispose();
    faPiaoTaiTouController?.dispose();
    yinHangMingChengController?.dispose();
    kaiHuHangController?.dispose();
    zhangHaoController?.dispose();
    beiZhuController?.dispose();
  }

  /// Additional helper methods are added here.
}
