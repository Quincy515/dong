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

class KaoQinTongJiModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this page.

  final formKey = GlobalKey<FormState>();
  // Model for webNav component.
  late WebNavModel webNavModel;
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for xing_ming widget.
  TextEditingController? xingMingController;
  String? Function(BuildContext, String?)? xingMingControllerValidator;
  // State field(s) for chu_qin_tian_shu widget.
  TextEditingController? chuQinTianShuController;
  String? Function(BuildContext, String?)? chuQinTianShuControllerValidator;
  // State field(s) for qing_jia_tian_shu widget.
  TextEditingController? qingJiaTianShuController;
  String? Function(BuildContext, String?)? qingJiaTianShuControllerValidator;
  // State field(s) for xiang_qing widget.
  TextEditingController? xiangQingController;
  String? Function(BuildContext, String?)? xiangQingControllerValidator;
  DateTime? datePicked1;
  DateTime? datePicked2;
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
    xingMingController?.dispose();
    chuQinTianShuController?.dispose();
    qingJiaTianShuController?.dispose();
    xiangQingController?.dispose();
    beiZhuController?.dispose();
    pageSearchAndNextModel.dispose();
  }

  /// Additional helper methods are added here.
}
