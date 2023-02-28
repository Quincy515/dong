// ignore_for_file: unused_import

import '/components/web_nav_widget.dart';
import '/flutter_flow/flutter_flow_icon_button.dart';
import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:provider/provider.dart';

class UserProfileModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this page.

  final formKey = GlobalKey<FormState>();
  // Model for webNav component.
  late WebNavModel webNavModel;
  // State field(s) for xingbie widget.
  TextEditingController? xingbieController;
  String? Function(BuildContext, String?)? xingbieControllerValidator;
  // State field(s) for nianling widget.
  TextEditingController? nianlingController;
  String? Function(BuildContext, String?)? nianlingControllerValidator;
  // State field(s) for lianxidianhua widget.
  TextEditingController? lianxidianhuaController;
  String? Function(BuildContext, String?)? lianxidianhuaControllerValidator;
  // State field(s) for lizhishijian widget.
  TextEditingController? lizhishijianController;
  String? Function(BuildContext, String?)? lizhishijianControllerValidator;
  // State field(s) for bumen widget.
  TextEditingController? bumenController;
  String? Function(BuildContext, String?)? bumenControllerValidator;
  // State field(s) for ruzhishijian widget.
  TextEditingController? ruzhishijianController;
  String? Function(BuildContext, String?)? ruzhishijianControllerValidator;
  // State field(s) for laodonghetongkaishishijian widget.
  TextEditingController? laodonghetongkaishishijianController;
  String? Function(BuildContext, String?)?
      laodonghetongkaishishijianControllerValidator;
  // State field(s) for laodonghetongjiezjiezhishijian widget.
  TextEditingController? laodonghetongjiezjiezhishijianController;
  String? Function(BuildContext, String?)?
      laodonghetongjiezjiezhishijianControllerValidator;
  // State field(s) for wenhuachengdu widget.
  TextEditingController? wenhuachengduController;
  String? Function(BuildContext, String?)? wenhuachengduControllerValidator;
  // State field(s) for hukousuozaidi widget.
  TextEditingController? hukousuozaidiController;
  String? Function(BuildContext, String?)? hukousuozaidiControllerValidator;
  // State field(s) for xianjuzhudizhi widget.
  TextEditingController? xianjuzhudizhiController;
  String? Function(BuildContext, String?)? xianjuzhudizhiControllerValidator;
  // State field(s) for shenfenzhenghaoma widget.
  TextEditingController? shenfenzhenghaomaController;
  String? Function(BuildContext, String?)? shenfenzhenghaomaControllerValidator;
  // State field(s) for chushengriqi widget.
  TextEditingController? chushengriqiController;
  String? Function(BuildContext, String?)? chushengriqiControllerValidator;
  // State field(s) for hetongyuedinggongzi widget.
  TextEditingController? hetongyuedinggongziController;
  String? Function(BuildContext, String?)?
      hetongyuedinggongziControllerValidator;
  // State field(s) for congshigangweihuogongzhong widget.
  TextEditingController? congshigangweihuogongzhongController;
  String? Function(BuildContext, String?)?
      congshigangweihuogongzhongControllerValidator;
  // State field(s) for zhiyezige widget.
  TextEditingController? zhiyezigeController;
  String? Function(BuildContext, String?)? zhiyezigeControllerValidator;
  // State field(s) for zhucebianhao widget.
  TextEditingController? zhucebianhaoController;
  String? Function(BuildContext, String?)? zhucebianhaoControllerValidator;
  // State field(s) for zhuanye widget.
  TextEditingController? zhuanyeController;
  String? Function(BuildContext, String?)? zhuanyeControllerValidator;
  // State field(s) for nianfei widget.
  TextEditingController? nianfeiController;
  String? Function(BuildContext, String?)? nianfeiControllerValidator;
  // State field(s) for zhengshushijian widget.
  TextEditingController? zhengshushijianController;
  String? Function(BuildContext, String?)? zhengshushijianControllerValidator;
  // State field(s) for youxiaoqi widget.
  TextEditingController? youxiaoqiController;
  String? Function(BuildContext, String?)? youxiaoqiControllerValidator;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {
    webNavModel = createModel(context, () => WebNavModel());
  }

  void dispose() {
    webNavModel.dispose();
    xingbieController?.dispose();
    nianlingController?.dispose();
    lianxidianhuaController?.dispose();
    lizhishijianController?.dispose();
    bumenController?.dispose();
    ruzhishijianController?.dispose();
    laodonghetongkaishishijianController?.dispose();
    laodonghetongjiezjiezhishijianController?.dispose();
    wenhuachengduController?.dispose();
    hukousuozaidiController?.dispose();
    xianjuzhudizhiController?.dispose();
    shenfenzhenghaomaController?.dispose();
    chushengriqiController?.dispose();
    hetongyuedinggongziController?.dispose();
    congshigangweihuogongzhongController?.dispose();
    zhiyezigeController?.dispose();
    zhucebianhaoController?.dispose();
    zhuanyeController?.dispose();
    nianfeiController?.dispose();
    zhengshushijianController?.dispose();
    youxiaoqiController?.dispose();
  }

  /// Additional helper methods are added here.
}
