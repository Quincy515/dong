// ignore_for_file: unused_import

import '/components/page_search_and_next_widget.dart';
import '/components/web_nav_widget.dart';
import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/flutter_flow_widgets.dart';
import 'package:styled_divider/styled_divider.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:provider/provider.dart';

class GongSiZhangHaoMiMaModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this page.

  final formKey = GlobalKey<FormState>();
  // Model for webNav component.
  late WebNavModel webNavModel;
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for xi_tong_ming_cheng widget.
  TextEditingController? xiTongMingChengController;
  String? Function(BuildContext, String?)? xiTongMingChengControllerValidator;
  // State field(s) for yong_hu_ming widget.
  TextEditingController? yongHuMingController;
  String? Function(BuildContext, String?)? yongHuMingControllerValidator;
  // State field(s) for mi_ma widget.
  TextEditingController? miMaController;
  String? Function(BuildContext, String?)? miMaControllerValidator;
  // State field(s) for yong_tu widget.
  TextEditingController? yongTuController;
  String? Function(BuildContext, String?)? yongTuControllerValidator;
  // State field(s) for lian_jie widget.
  TextEditingController? lianJieController;
  String? Function(BuildContext, String?)? lianJieControllerValidator;
  DateTime? datePicked1;
  DateTime? datePicked2;
  // State field(s) for nian_fei widget.
  TextEditingController? nianFeiController;
  String? Function(BuildContext, String?)? nianFeiControllerValidator;
  // State field(s) for bei_zhu widget.
  TextEditingController? beiZhuController;
  String? Function(BuildContext, String?)? beiZhuControllerValidator;
  // Model for pageSearchAndNext component.
  late PageSearchAndNextModel pageSearchAndNextModel;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {
    webNavModel = createModel(context, () => WebNavModel());
    pageSearchAndNextModel =
        createModel(context, () => PageSearchAndNextModel());
  }

  void dispose() {
    webNavModel.dispose();
    groupController?.dispose();
    xiTongMingChengController?.dispose();
    yongHuMingController?.dispose();
    miMaController?.dispose();
    yongTuController?.dispose();
    lianJieController?.dispose();
    nianFeiController?.dispose();
    beiZhuController?.dispose();
    pageSearchAndNextModel.dispose();
  }

  /// Additional helper methods are added here.
}
