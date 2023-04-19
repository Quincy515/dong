import '../components/page_search_and_next_model.dart';
import '../components/qing_jia_guan_li_item_model.dart';
import '../components/web_nav_model.dart';
import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';

class QingJiaGuanLiModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this page.

  // Model for webNav component.
  late WebNavModel webNavModel;
  // State field(s) for Switch widget.
  bool? switchValue;
  // Stores action output result for [Backend Call - API (getQingJiaGuanLiList)] action in Row widget.
  ApiCallResponse? getRes;
  // Model for qingJiaGuanLiItem component.
  late QingJiaGuanLiItemModel qingJiaGuanLiItemModel1;
  // Model for pageSearchAndNext component.
  late PageSearchAndNextModel pageSearchAndNextModel;
  // Model for qingJiaGuanLiItem component.
  late QingJiaGuanLiItemModel qingJiaGuanLiItemModel3;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {
    webNavModel = createModel(context, () => WebNavModel());
    qingJiaGuanLiItemModel1 =
        createModel(context, () => QingJiaGuanLiItemModel());
    pageSearchAndNextModel =
        createModel(context, () => PageSearchAndNextModel());
    qingJiaGuanLiItemModel3 =
        createModel(context, () => QingJiaGuanLiItemModel());
  }

  void dispose() {
    webNavModel.dispose();
    qingJiaGuanLiItemModel1.dispose();
    pageSearchAndNextModel.dispose();
    qingJiaGuanLiItemModel3.dispose();
  }

  /// Additional helper methods are added here.
}
