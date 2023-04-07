import '../components/page_search_and_next_model.dart';
import '../components/web_nav_model.dart';
import '/backend/api_requests/api_calls.dart';
import '/components/gong_ying_shang_item/gong_ying_shang_item_widget.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';

class GongYingShangBiaoModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this page.

  // Model for webNav component.
  late WebNavModel webNavModel;
  // State field(s) for Switch widget.
  bool? switchValue;
  // Stores action output result for [Backend Call - API (getMeiRiKaoQinList)] action in Row widget.
  ApiCallResponse? getRes;
  // Model for gongYingShangItem component.
  late GongYingShangItemModel gongYingShangItemModel1;
  // Models for gongYingShangItem dynamic component.
  late FlutterFlowDynamicModels<GongYingShangItemModel>
      gongYingShangItemModels2;
  // Model for pageSearchAndNext component.
  late PageSearchAndNextModel pageSearchAndNextModel;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {
    webNavModel = createModel(context, () => WebNavModel());
    gongYingShangItemModel1 =
        createModel(context, () => GongYingShangItemModel());
    gongYingShangItemModels2 =
        FlutterFlowDynamicModels(() => GongYingShangItemModel());
    pageSearchAndNextModel =
        createModel(context, () => PageSearchAndNextModel());
  }

  void dispose() {
    webNavModel.dispose();
    gongYingShangItemModel1.dispose();
    gongYingShangItemModels2.dispose();
    pageSearchAndNextModel.dispose();
  }

  /// Additional helper methods are added here.
}
