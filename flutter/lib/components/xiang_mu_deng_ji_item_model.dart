// ignore_for_file: unused_import

import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/flutter_flow_widgets.dart';
import 'package:styled_divider/styled_divider.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:provider/provider.dart';

class XiangMuDengJiItemModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this component.

  final formKey = GlobalKey<FormState>();
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for xiang_mu_ming_cheng widget.
  TextEditingController? xiangMuMingChengController;
  String? Function(BuildContext, String?)? xiangMuMingChengControllerValidator;
  // State field(s) for xiang_mu_jian_jie widget.
  TextEditingController? xiangMuJianJieController;
  String? Function(BuildContext, String?)? xiangMuJianJieControllerValidator;
  // State field(s) for xiang_mu_deng_ji_jin_du widget.
  TextEditingController? xiangMuDengJiJinDuController;
  String? Function(BuildContext, String?)?
      xiangMuDengJiJinDuControllerValidator;
  // State field(s) for zi_zhi_yao_qiu widget.
  TextEditingController? ziZhiYaoQiuController;
  String? Function(BuildContext, String?)? ziZhiYaoQiuControllerValidator;
  // State field(s) for dengJiYaoQiu widget.
  TextEditingController? dengJiYaoQiuController;
  String? Function(BuildContext, String?)? dengJiYaoQiuControllerValidator;
  // State field(s) for tou_biao_dan_wei_ming_cheng widget.
  TextEditingController? touBiaoDanWeiMingChengController;
  String? Function(BuildContext, String?)?
      touBiaoDanWeiMingChengControllerValidator;
  // State field(s) for kaiBiaoShiJian widget.
  TextEditingController? kaiBiaoShiJianController;
  String? Function(BuildContext, String?)? kaiBiaoShiJianControllerValidator;
  // State field(s) for kai_biao_di_dian widget.
  TextEditingController? kaiBiaoDiDianController;
  String? Function(BuildContext, String?)? kaiBiaoDiDianControllerValidator;
  // State field(s) for suo_shu_di_qu widget.
  TextEditingController? suoShuDiQuController;
  String? Function(BuildContext, String?)? suoShuDiQuControllerValidator;
  // State field(s) for zhao_biao_gui_mo widget.
  TextEditingController? zhaoBiaoGuiMoController;
  String? Function(BuildContext, String?)? zhaoBiaoGuiMoControllerValidator;
  // State field(s) for kai_biao_ren_yuan widget.
  TextEditingController? kaiBiaoRenYuanController;
  String? Function(BuildContext, String?)? kaiBiaoRenYuanControllerValidator;
  // State field(s) for ye_wu_fu_ze_ren widget.
  TextEditingController? yeWuFuZeRenController;
  String? Function(BuildContext, String?)? yeWuFuZeRenControllerValidator;
  // State field(s) for bao_zheng_jin_xing_shi widget.
  TextEditingController? baoZhengJinXingShiController;
  String? Function(BuildContext, String?)?
      baoZhengJinXingShiControllerValidator;
  // State field(s) for bao_zheng_jin_jin_e widget.
  TextEditingController? baoZhengJinJinEController;
  String? Function(BuildContext, String?)? baoZhengJinJinEControllerValidator;
  // State field(s) for ping_biao_ban_fa widget.
  TextEditingController? pingBiaoBanFaController;
  String? Function(BuildContext, String?)? pingBiaoBanFaControllerValidator;
  // State field(s) for biao_shu_lei_xing widget.
  TextEditingController? biaoShuLeiXingController;
  String? Function(BuildContext, String?)? biaoShuLeiXingControllerValidator;
  // State field(s) for xiangMuDengJiShiJian widget.
  TextEditingController? xiangMuDengJiShiJianController;
  String? Function(BuildContext, String?)?
      xiangMuDengJiShiJianControllerValidator;
  // State field(s) for bu_men_shen_he widget.
  TextEditingController? buMenShenHeController;
  String? Function(BuildContext, String?)? buMenShenHeControllerValidator;
  // State field(s) for buMenShenHeShiJian widget.
  TextEditingController? buMenShenHeShiJianController;
  String? Function(BuildContext, String?)?
      buMenShenHeShiJianControllerValidator;
  // State field(s) for ce_oshen_he widget.
  TextEditingController? ceOshenHeController;
  String? Function(BuildContext, String?)? ceOshenHeControllerValidator;
  // State field(s) for ceOshenHeShiJian widget.
  TextEditingController? ceOshenHeShiJianController;
  String? Function(BuildContext, String?)? ceOshenHeShiJianControllerValidator;
  // State field(s) for tou_biao_han_zhi_zuo_ren widget.
  TextEditingController? touBiaoHanZhiZuoRenController;
  String? Function(BuildContext, String?)?
      touBiaoHanZhiZuoRenControllerValidator;
  // State field(s) for tou_biao_han_ti_cheng_jin_e widget.
  TextEditingController? touBiaoHanTiChengJinEController;
  String? Function(BuildContext, String?)?
      touBiaoHanTiChengJinEControllerValidator;
  // State field(s) for tou_biao_han_zhi_zuo_ren_que_ren widget.
  TextEditingController? touBiaoHanZhiZuoRenQueRenController;
  String? Function(BuildContext, String?)?
      touBiaoHanZhiZuoRenQueRenControllerValidator;
  // State field(s) for touBiaoHanZhiZuoRenQueRenShiJian widget.
  TextEditingController? touBiaoHanZhiZuoRenQueRenShiJianController;
  String? Function(BuildContext, String?)?
      touBiaoHanZhiZuoRenQueRenShiJianControllerValidator;
  // State field(s) for shang_wu_biao_zhi_zuo_ren widget.
  TextEditingController? shangWuBiaoZhiZuoRenController;
  String? Function(BuildContext, String?)?
      shangWuBiaoZhiZuoRenControllerValidator;
  // State field(s) for shang_wu_biao_ti_cheng_jin_e widget.
  TextEditingController? shangWuBiaoTiChengJinEController;
  String? Function(BuildContext, String?)?
      shangWuBiaoTiChengJinEControllerValidator;
  // State field(s) for shang_wu_biao_zhi_zuo_ren_que_ren widget.
  TextEditingController? shangWuBiaoZhiZuoRenQueRenController;
  String? Function(BuildContext, String?)?
      shangWuBiaoZhiZuoRenQueRenControllerValidator;
  // State field(s) for shangWuBiaoZhiZuoRenQueRenShiJian widget.
  TextEditingController? shangWuBiaoZhiZuoRenQueRenShiJianController;
  String? Function(BuildContext, String?)?
      shangWuBiaoZhiZuoRenQueRenShiJianControllerValidator;
  // State field(s) for ji_shu_biao_zhi_zuo_ren widget.
  TextEditingController? jiShuBiaoZhiZuoRenController;
  String? Function(BuildContext, String?)?
      jiShuBiaoZhiZuoRenControllerValidator;
  // State field(s) for ji_shu_biao_ti_cheng_jin_e widget.
  TextEditingController? jiShuBiaoTiChengJinEController;
  String? Function(BuildContext, String?)?
      jiShuBiaoTiChengJinEControllerValidator;
  // State field(s) for ji_shu_biao_zhi_zuo_ren_que_ren widget.
  TextEditingController? jiShuBiaoZhiZuoRenQueRenController;
  String? Function(BuildContext, String?)?
      jiShuBiaoZhiZuoRenQueRenControllerValidator;
  // State field(s) for jiShuBiaoZhiZuoRenQueRenShiJian widget.
  TextEditingController? jiShuBiaoZhiZuoRenQueRenShiJianController;
  String? Function(BuildContext, String?)?
      jiShuBiaoZhiZuoRenQueRenShiJianControllerValidator;
  // State field(s) for bu_men_shen_qing_xia_fu_dian_wei widget.
  TextEditingController? buMenShenQingXiaFuDianWeiController;
  String? Function(BuildContext, String?)?
      buMenShenQingXiaFuDianWeiControllerValidator;
  // State field(s) for buMenXiaFuDianWeiShenQingShiJian widget.
  TextEditingController? buMenXiaFuDianWeiShenQingShiJianController;
  String? Function(BuildContext, String?)?
      buMenXiaFuDianWeiShenQingShiJianControllerValidator;
  // State field(s) for ce_oshen_he_xia_fu_dian_wei widget.
  TextEditingController? ceOshenHeXiaFuDianWeiController;
  String? Function(BuildContext, String?)?
      ceOshenHeXiaFuDianWeiControllerValidator;
  // State field(s) for ceOxiaFuDianWeiShenHeShiJian widget.
  TextEditingController? ceOxiaFuDianWeiShenHeShiJianController;
  String? Function(BuildContext, String?)?
      ceOxiaFuDianWeiShenHeShiJianControllerValidator;
  // State field(s) for xiang_mu_zhuang_tai widget.
  TextEditingController? xiangMuZhuangTaiController;
  String? Function(BuildContext, String?)? xiangMuZhuangTaiControllerValidator;
  // State field(s) for zhong_biao_jin_e widget.
  TextEditingController? zhongBiaoJinEController;
  String? Function(BuildContext, String?)? zhongBiaoJinEControllerValidator;
  // State field(s) for bei_zhu widget.
  TextEditingController? beiZhuController;
  String? Function(BuildContext, String?)? beiZhuControllerValidator;
  // Stores action output result for [Backend Call - API (deleteXiangMuDengJiBiao)] action in Button widget.
  ApiCallResponse? deleteRes;
  // Stores action output result for [Backend Call - API (updateXiangMuDengJiBiao)] action in Button widget.
  ApiCallResponse? updateRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    groupController?.dispose();
    xiangMuMingChengController?.dispose();
    xiangMuJianJieController?.dispose();
    xiangMuDengJiJinDuController?.dispose();
    ziZhiYaoQiuController?.dispose();
    dengJiYaoQiuController?.dispose();
    touBiaoDanWeiMingChengController?.dispose();
    kaiBiaoShiJianController?.dispose();
    kaiBiaoDiDianController?.dispose();
    suoShuDiQuController?.dispose();
    zhaoBiaoGuiMoController?.dispose();
    kaiBiaoRenYuanController?.dispose();
    yeWuFuZeRenController?.dispose();
    baoZhengJinXingShiController?.dispose();
    baoZhengJinJinEController?.dispose();
    pingBiaoBanFaController?.dispose();
    biaoShuLeiXingController?.dispose();
    xiangMuDengJiShiJianController?.dispose();
    buMenShenHeController?.dispose();
    buMenShenHeShiJianController?.dispose();
    ceOshenHeController?.dispose();
    ceOshenHeShiJianController?.dispose();
    touBiaoHanZhiZuoRenController?.dispose();
    touBiaoHanTiChengJinEController?.dispose();
    touBiaoHanZhiZuoRenQueRenController?.dispose();
    touBiaoHanZhiZuoRenQueRenShiJianController?.dispose();
    shangWuBiaoZhiZuoRenController?.dispose();
    shangWuBiaoTiChengJinEController?.dispose();
    shangWuBiaoZhiZuoRenQueRenController?.dispose();
    shangWuBiaoZhiZuoRenQueRenShiJianController?.dispose();
    jiShuBiaoZhiZuoRenController?.dispose();
    jiShuBiaoTiChengJinEController?.dispose();
    jiShuBiaoZhiZuoRenQueRenController?.dispose();
    jiShuBiaoZhiZuoRenQueRenShiJianController?.dispose();
    buMenShenQingXiaFuDianWeiController?.dispose();
    buMenXiaFuDianWeiShenQingShiJianController?.dispose();
    ceOshenHeXiaFuDianWeiController?.dispose();
    ceOxiaFuDianWeiShenHeShiJianController?.dispose();
    xiangMuZhuangTaiController?.dispose();
    zhongBiaoJinEController?.dispose();
    beiZhuController?.dispose();
  }

  /// Additional helper methods are added here.
}
