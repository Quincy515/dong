// ignore_for_file: unused_import

import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_drop_down.dart';
import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/flutter_flow_widgets.dart';
import 'package:styled_divider/styled_divider.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:provider/provider.dart';

class ZhengShuBiaoItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for zhengShuMingCheng widget.
  TextEditingController? zhengShuMingChengController;
  String? Function(BuildContext, String?)? zhengShuMingChengControllerValidator;
  // State field(s) for zhengShuLeiBie widget.
  int? zhengShuLeiBieValue;
  // State field(s) for zhengShuBianHao widget.
  TextEditingController? zhengShuBianHaoController;
  String? Function(BuildContext, String?)? zhengShuBianHaoControllerValidator;
  // State field(s) for suoYouRenXingMing widget.
  TextEditingController? suoYouRenXingMingController;
  String? Function(BuildContext, String?)? suoYouRenXingMingControllerValidator;
  // State field(s) for shenFenZhengHao widget.
  TextEditingController? shenFenZhengHaoController;
  String? Function(BuildContext, String?)? shenFenZhengHaoControllerValidator;
  // State field(s) for shouJiHaoMa widget.
  TextEditingController? shouJiHaoMaController;
  String? Function(BuildContext, String?)? shouJiHaoMaControllerValidator;
  // State field(s) for faZhengBuMen widget.
  TextEditingController? faZhengBuMenController;
  String? Function(BuildContext, String?)? faZhengBuMenControllerValidator;
  // State field(s) for zhengShuSuoShuChengShi widget.
  int? zhengShuSuoShuChengShiValue;
  // State field(s) for zhengShuZhuangTai widget.
  int? zhengShuZhuangTaiValue;
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
    zhengShuMingChengController?.dispose();
    zhengShuBianHaoController?.dispose();
    suoYouRenXingMingController?.dispose();
    shenFenZhengHaoController?.dispose();
    shouJiHaoMaController?.dispose();
    faZhengBuMenController?.dispose();
    zhengShuNianShiYongFeiController?.dispose();
    beiZhuController?.dispose();
  }

  /// Additional helper methods are added here.
}
