// ignore_for_file: unused_import

import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/flutter_flow_widgets.dart';
import 'package:styled_divider/styled_divider.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:provider/provider.dart';

class YuanGongHuaMingCeItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for bu_men widget.
  TextEditingController? buMenController;
  String? Function(BuildContext, String?)? buMenControllerValidator;
  // State field(s) for xing_ming widget.
  TextEditingController? xingMingController;
  String? Function(BuildContext, String?)? xingMingControllerValidator;
  // State field(s) for xing_bie widget.
  TextEditingController? xingBieController;
  String? Function(BuildContext, String?)? xingBieControllerValidator;
  // State field(s) for nian_ling widget.
  TextEditingController? nianLingController;
  String? Function(BuildContext, String?)? nianLingControllerValidator;
  // State field(s) for hu_kou_suo_zai_di widget.
  TextEditingController? huKouSuoZaiDiController;
  String? Function(BuildContext, String?)? huKouSuoZaiDiControllerValidator;
  // State field(s) for ru_zhi_shi_jian widget.
  TextEditingController? ruZhiShiJianController;
  String? Function(BuildContext, String?)? ruZhiShiJianControllerValidator;
  // State field(s) for wen_hua_cheng_du widget.
  TextEditingController? wenHuaChengDuController;
  String? Function(BuildContext, String?)? wenHuaChengDuControllerValidator;
  // State field(s) for lao_dong_he_tong_kai_shi_ri_qi widget.
  TextEditingController? laoDongHeTongKaiShiRiQiController;
  String? Function(BuildContext, String?)?
      laoDongHeTongKaiShiRiQiControllerValidator;
  // State field(s) for lao_dong_he_tong_jie_zhi_ri_qi widget.
  TextEditingController? laoDongHeTongJieZhiRiQiController;
  String? Function(BuildContext, String?)?
      laoDongHeTongJieZhiRiQiControllerValidator;
  // State field(s) for shen_fen_zheng_hao_ma widget.
  TextEditingController? shenFenZhengHaoMaController;
  String? Function(BuildContext, String?)? shenFenZhengHaoMaControllerValidator;
  // State field(s) for chu_sheng_ri_qi widget.
  TextEditingController? chuShengRiQiController;
  String? Function(BuildContext, String?)? chuShengRiQiControllerValidator;
  // State field(s) for he_tong_yue_ding_gong_zi widget.
  TextEditingController? heTongYueDingGongZiController;
  String? Function(BuildContext, String?)?
      heTongYueDingGongZiControllerValidator;
  // State field(s) for cong_shi_gang_wei_huo_gong_zhong widget.
  TextEditingController? congShiGangWeiHuoGongZhongController;
  String? Function(BuildContext, String?)?
      congShiGangWeiHuoGongZhongControllerValidator;
  // State field(s) for lian_xi_dian_hua widget.
  TextEditingController? lianXiDianHuaController;
  String? Function(BuildContext, String?)? lianXiDianHuaControllerValidator;
  // State field(s) for xian_ju_zhu_di_zhi widget.
  TextEditingController? xianJuZhuDiZhiController;
  String? Function(BuildContext, String?)? xianJuZhuDiZhiControllerValidator;
  // State field(s) for li_zhi_ri_qi widget.
  TextEditingController? liZhiRiQiController;
  String? Function(BuildContext, String?)? liZhiRiQiControllerValidator;
  // State field(s) for ge_ren_she_bao_zhang_hao widget.
  TextEditingController? geRenSheBaoZhangHaoController;
  String? Function(BuildContext, String?)?
      geRenSheBaoZhangHaoControllerValidator;
  // State field(s) for ge_ren_she_bao_mi_ma widget.
  TextEditingController? geRenSheBaoMiMaController;
  String? Function(BuildContext, String?)? geRenSheBaoMiMaControllerValidator;
  // State field(s) for men_jin_bian_hao widget.
  TextEditingController? menJinBianHaoController;
  String? Function(BuildContext, String?)? menJinBianHaoControllerValidator;
  // State field(s) for bei_zhu widget.
  TextEditingController? beiZhuController;
  String? Function(BuildContext, String?)? beiZhuControllerValidator;
  // Stores action output result for [Backend Call - API (deleteYuanGongHuaMingCe)] action in Button widget.
  ApiCallResponse? deleteRes;
  // Stores action output result for [Backend Call - API (updateYuanGongHuaMingCe)] action in Button widget.
  ApiCallResponse? updateRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    groupController?.dispose();
    buMenController?.dispose();
    xingMingController?.dispose();
    xingBieController?.dispose();
    nianLingController?.dispose();
    huKouSuoZaiDiController?.dispose();
    ruZhiShiJianController?.dispose();
    wenHuaChengDuController?.dispose();
    laoDongHeTongKaiShiRiQiController?.dispose();
    laoDongHeTongJieZhiRiQiController?.dispose();
    shenFenZhengHaoMaController?.dispose();
    chuShengRiQiController?.dispose();
    heTongYueDingGongZiController?.dispose();
    congShiGangWeiHuoGongZhongController?.dispose();
    lianXiDianHuaController?.dispose();
    xianJuZhuDiZhiController?.dispose();
    liZhiRiQiController?.dispose();
    geRenSheBaoZhangHaoController?.dispose();
    geRenSheBaoMiMaController?.dispose();
    menJinBianHaoController?.dispose();
    beiZhuController?.dispose();
  }

  /// Additional helper methods are added here.
}
