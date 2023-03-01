import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/flutter_flow_widgets.dart';
import 'package:styled_divider/styled_divider.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:provider/provider.dart';

class TouBiaoShuJuTongJiItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for xiang_mu_ming_cheng widget.
  TextEditingController? xiangMuMingChengController;
  String? Function(BuildContext, String?)? xiangMuMingChengControllerValidator;
  // State field(s) for tou_biao_dan_wei_ming_cheng widget.
  TextEditingController? touBiaoDanWeiMingChengController;
  String? Function(BuildContext, String?)?
      touBiaoDanWeiMingChengControllerValidator;
  DateTime? datePicked;
  // State field(s) for zi_zhi_yao_qiu widget.
  TextEditingController? ziZhiYaoQiuController;
  String? Function(BuildContext, String?)? ziZhiYaoQiuControllerValidator;
  // State field(s) for suo_shu_di_qu widget.
  TextEditingController? suoShuDiQuController;
  String? Function(BuildContext, String?)? suoShuDiQuControllerValidator;
  // State field(s) for zhao_biao_gui_mo widget.
  TextEditingController? zhaoBiaoGuiMoController;
  String? Function(BuildContext, String?)? zhaoBiaoGuiMoControllerValidator;
  // State field(s) for deng_ji_yao_qiu widget.
  TextEditingController? dengJiYaoQiuController;
  String? Function(BuildContext, String?)? dengJiYaoQiuControllerValidator;
  // State field(s) for di_yi_zhong_biao_hou_xuan_ren widget.
  TextEditingController? diYiZhongBiaoHouXuanRenController;
  String? Function(BuildContext, String?)?
      diYiZhongBiaoHouXuanRenControllerValidator;
  // State field(s) for di_yi_bao_jia widget.
  TextEditingController? diYiBaoJiaController;
  String? Function(BuildContext, String?)? diYiBaoJiaControllerValidator;
  // State field(s) for di_yi_xia_fu_lu widget.
  TextEditingController? diYiXiaFuLuController;
  String? Function(BuildContext, String?)? diYiXiaFuLuControllerValidator;
  // State field(s) for di_er_zhong_biao_hou_xuan_ren widget.
  TextEditingController? diErZhongBiaoHouXuanRenController;
  String? Function(BuildContext, String?)?
      diErZhongBiaoHouXuanRenControllerValidator;
  // State field(s) for di_er_bao_jia widget.
  TextEditingController? diErBaoJiaController;
  String? Function(BuildContext, String?)? diErBaoJiaControllerValidator;
  // State field(s) for di_er_xia_fu_lu widget.
  TextEditingController? diErXiaFuLuController;
  String? Function(BuildContext, String?)? diErXiaFuLuControllerValidator;
  // State field(s) for di_san_zhong_biao_hou_xuan_ren widget.
  TextEditingController? diSanZhongBiaoHouXuanRenController;
  String? Function(BuildContext, String?)?
      diSanZhongBiaoHouXuanRenControllerValidator;
  // State field(s) for di_san_bao_jia widget.
  TextEditingController? diSanBaoJiaController;
  String? Function(BuildContext, String?)? diSanBaoJiaControllerValidator;
  // State field(s) for di_san_xia_fu_lu widget.
  TextEditingController? diSanXiaFuLuController;
  String? Function(BuildContext, String?)? diSanXiaFuLuControllerValidator;
  // State field(s) for bei_zhu widget.
  TextEditingController? beiZhuController;
  String? Function(BuildContext, String?)? beiZhuControllerValidator;
  // Stores action output result for [Backend Call - API (updateTouBiaoShuJuTongJi)] action in Button widget.
  ApiCallResponse? updateRes;
  // Stores action output result for [Backend Call - API (deleteTouBiaoShuJuTongJi)] action in Button widget.
  ApiCallResponse? deleteRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    groupController?.dispose();
    xiangMuMingChengController?.dispose();
    touBiaoDanWeiMingChengController?.dispose();
    ziZhiYaoQiuController?.dispose();
    suoShuDiQuController?.dispose();
    zhaoBiaoGuiMoController?.dispose();
    dengJiYaoQiuController?.dispose();
    diYiZhongBiaoHouXuanRenController?.dispose();
    diYiBaoJiaController?.dispose();
    diYiXiaFuLuController?.dispose();
    diErZhongBiaoHouXuanRenController?.dispose();
    diErBaoJiaController?.dispose();
    diErXiaFuLuController?.dispose();
    diSanZhongBiaoHouXuanRenController?.dispose();
    diSanBaoJiaController?.dispose();
    diSanXiaFuLuController?.dispose();
    beiZhuController?.dispose();
  }

  /// Additional helper methods are added here.

}
