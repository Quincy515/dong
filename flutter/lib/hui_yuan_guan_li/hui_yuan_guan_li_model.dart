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

class HuiYuanGuanLiModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this page.

  final formKey = GlobalKey<FormState>();
  // Model for webNav component.
  late WebNavModel webNavModel;
  // State field(s) for group widget.
  TextEditingController? groupController;
  String? Function(BuildContext, String?)? groupControllerValidator;
  // State field(s) for hui_yuan_ming_cheng widget.
  TextEditingController? huiYuanMingChengController;
  String? Function(BuildContext, String?)? huiYuanMingChengControllerValidator;
  // State field(s) for ji_bie widget.
  TextEditingController? jiBieController;
  String? Function(BuildContext, String?)? jiBieControllerValidator;
  // State field(s) for hui_fei widget.
  TextEditingController? huiFeiController;
  String? Function(BuildContext, String?)? huiFeiControllerValidator;
  // State field(s) for wang_zhi widget.
  TextEditingController? wangZhiController;
  String? Function(BuildContext, String?)? wangZhiControllerValidator;
  // State field(s) for ru_hui_lian_xi widget.
  TextEditingController? ruHuiLianXiController;
  String? Function(BuildContext, String?)? ruHuiLianXiControllerValidator;
  DateTime? datePicked1;
  DateTime? datePicked2;
  DateTime? datePicked3;
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
    huiYuanMingChengController?.dispose();
    jiBieController?.dispose();
    huiFeiController?.dispose();
    wangZhiController?.dispose();
    ruHuiLianXiController?.dispose();
    beiZhuController?.dispose();
    pageSearchAndNextModel.dispose();
  }

  /// Additional helper methods are added here.
}
