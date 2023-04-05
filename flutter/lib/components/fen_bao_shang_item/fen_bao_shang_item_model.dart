import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/form_field_controller.dart';
import 'package:flutter/material.dart';

class FenBaoShangItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for fenBaoShangMingCheng widget.
  TextEditingController? fenBaoShangMingChengController;
  String? Function(BuildContext, String?)?
      fenBaoShangMingChengControllerValidator;
  // State field(s) for fenBaoShangLeiXing widget.
  int? fenBaoShangLeiXingValue;
  FormFieldController<int>? fenBaoShangLeiXingController;
  // State field(s) for faRen widget.
  TextEditingController? faRenController;
  String? Function(BuildContext, String?)? faRenControllerValidator;
  // State field(s) for sheHuiXinYongDaiMa widget.
  TextEditingController? sheHuiXinYongDaiMaController;
  String? Function(BuildContext, String?)?
      sheHuiXinYongDaiMaControllerValidator;
  // State field(s) for ziZhiZhengZhao widget.
  TextEditingController? ziZhiZhengZhaoController;
  String? Function(BuildContext, String?)? ziZhiZhengZhaoControllerValidator;
  // State field(s) for faRenShouJi widget.
  TextEditingController? faRenShouJiController;
  String? Function(BuildContext, String?)? faRenShouJiControllerValidator;
  // State field(s) for tianJiaRen widget.
  int? tianJiaRenValue;
  FormFieldController<int>? tianJiaRenController;
  DateTime? datePicked;
  // State field(s) for zhuangTai widget.
  bool? zhuangTaiValue;
  // State field(s) for yinHangMingCheng widget.
  TextEditingController? yinHangMingChengController;
  String? Function(BuildContext, String?)? yinHangMingChengControllerValidator;
  // State field(s) for kaiHuHang widget.
  TextEditingController? kaiHuHangController;
  String? Function(BuildContext, String?)? kaiHuHangControllerValidator;
  // State field(s) for zhangHao widget.
  TextEditingController? zhangHaoController;
  String? Function(BuildContext, String?)? zhangHaoControllerValidator;
  // State field(s) for bei_zhu widget.
  TextEditingController? beiZhuController;
  String? Function(BuildContext, String?)? beiZhuControllerValidator;
  // Stores action output result for [Backend Call - API (deleteFenBaoShangBiao)] action in Button widget.
  ApiCallResponse? deleteRes;
  // Stores action output result for [Backend Call - API (updateFenBaoShangBiao)] action in Button widget.
  ApiCallResponse? updateRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    groupController?.dispose();
    fenBaoShangMingChengController?.dispose();
    faRenController?.dispose();
    sheHuiXinYongDaiMaController?.dispose();
    ziZhiZhengZhaoController?.dispose();
    faRenShouJiController?.dispose();
    yinHangMingChengController?.dispose();
    kaiHuHangController?.dispose();
    zhangHaoController?.dispose();
    beiZhuController?.dispose();
  }

  /// Additional helper methods are added here.
}
