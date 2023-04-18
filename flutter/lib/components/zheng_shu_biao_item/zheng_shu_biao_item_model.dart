import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/form_field_controller.dart';
import 'package:flutter/material.dart';

class ZhengShuBiaoItemModel extends FlutterFlowModel {
  ///  Local state fields for this component.

  bool visible = false;

  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for suoYouRenXingMing widget.
  TextEditingController? suoYouRenXingMingController;
  String? Function(BuildContext, String?)? suoYouRenXingMingControllerValidator;
  // State field(s) for zhengShuMingChengOpt widget.
  int? zhengShuMingChengOptValue;
  FormFieldController<int>? zhengShuMingChengOptController;
  // State field(s) for zhengShuLeiBie widget.
  int? zhengShuLeiBieValue;
  FormFieldController<int>? zhengShuLeiBieController;
  // State field(s) for zhengShuBianHao widget.
  TextEditingController? zhengShuBianHaoController;
  String? Function(BuildContext, String?)? zhengShuBianHaoControllerValidator;
  // State field(s) for shenFenZhengHao widget.
  TextEditingController? shenFenZhengHaoController;
  String? Function(BuildContext, String?)? shenFenZhengHaoControllerValidator;
  // State field(s) for shouJiHaoMa widget.
  TextEditingController? shouJiHaoMaController;
  String? Function(BuildContext, String?)? shouJiHaoMaControllerValidator;
  // State field(s) for sheBaoChaXunZhangHao widget.
  TextEditingController? sheBaoChaXunZhangHaoController;
  String? Function(BuildContext, String?)?
      sheBaoChaXunZhangHaoControllerValidator;
  // State field(s) for sheBaoChaXunMiMa widget.
  TextEditingController? sheBaoChaXunMiMaController;
  String? Function(BuildContext, String?)? sheBaoChaXunMiMaControllerValidator;
  // State field(s) for faZhengBuMen widget.
  TextEditingController? faZhengBuMenController;
  String? Function(BuildContext, String?)? faZhengBuMenControllerValidator;
  // State field(s) for zhengShuSuoShuChengShi widget.
  int? zhengShuSuoShuChengShiValue;
  FormFieldController<int>? zhengShuSuoShuChengShiController;
  // State field(s) for zhengShuZhuangTai widget.
  int? zhengShuZhuangTaiValue;
  FormFieldController<int>? zhengShuZhuangTaiController;
  // State field(s) for shiFouJieChu widget.
  bool? shiFouJieChuValue;
  // State field(s) for zhengShuNianShiYongFei widget.
  TextEditingController? zhengShuNianShiYongFeiController;
  String? Function(BuildContext, String?)?
      zhengShuNianShiYongFeiControllerValidator;
  // State field(s) for bei_zhu widget.
  TextEditingController? beiZhuController;
  String? Function(BuildContext, String?)? beiZhuControllerValidator;
  DateTime? datePicked1;
  DateTime? datePicked2;
  // State field(s) for DropDown widget.
  int? dropDownValue;
  FormFieldController<int>? dropDownController;
  // Stores action output result for [Backend Call - API (updateZhengShuJieChuJiLuBiao)] action in Button widget.
  ApiCallResponse? createRes;
  // Stores action output result for [Backend Call - API (deleteZhengShuBiao)] action in Button widget.
  ApiCallResponse? deleteRes;
  // Stores action output result for [Backend Call - API (updateZhengShuBiao)] action in Button widget.
  ApiCallResponse? updateRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    groupController?.dispose();
    suoYouRenXingMingController?.dispose();
    zhengShuBianHaoController?.dispose();
    shenFenZhengHaoController?.dispose();
    shouJiHaoMaController?.dispose();
    sheBaoChaXunZhangHaoController?.dispose();
    sheBaoChaXunMiMaController?.dispose();
    faZhengBuMenController?.dispose();
    zhengShuNianShiYongFeiController?.dispose();
    beiZhuController?.dispose();
  }

  /// Additional helper methods are added here.
}
