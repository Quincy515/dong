// ignore_for_file: unused_import

import '/backend/api_requests/api_calls.dart';
import '/components/page_search_and_next_widget.dart';
import '/components/web_nav_widget.dart';
import '/components/zheng_shu_jie_chu_ji_lu_item_widget.dart';
import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:provider/provider.dart';

class ZhengShuJieChuJiLModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this page.

  // Model for webNav component.
  late WebNavModel webNavModel;
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
    pageSearchAndNextModel.dispose();
  }

  /// Additional helper methods are added here.
}
