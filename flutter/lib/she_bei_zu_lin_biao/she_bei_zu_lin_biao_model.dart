import '../components/page_search_and_next_model.dart';
import '../components/web_nav_model.dart';
import '/backend/api_requests/api_calls.dart';
import '/components/she_bei_zu_lin_item/she_bei_zu_lin_item_widget.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';

class SheBeiZuLinBiaoModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this page.

  // Model for webNav component.
  late WebNavModel webNavModel;
  // State field(s) for Switch widget.
  bool? switchValue;
  // Stores action output result for [Backend Call - API (getMeiRiKaoQinList)] action in Row widget.
  ApiCallResponse? getRes;
  // Model for sheBeiZuLinItem component.
  late SheBeiZuLinItemModel sheBeiZuLinItemModel1;
  // Models for sheBeiZuLinItem dynamic component.
  late FlutterFlowDynamicModels<SheBeiZuLinItemModel> sheBeiZuLinItemModels2;
  // Model for pageSearchAndNext component.
  late PageSearchAndNextModel pageSearchAndNextModel;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {
    webNavModel = createModel(context, () => WebNavModel());
    sheBeiZuLinItemModel1 = createModel(context, () => SheBeiZuLinItemModel());
    sheBeiZuLinItemModels2 =
        FlutterFlowDynamicModels(() => SheBeiZuLinItemModel());
    pageSearchAndNextModel =
        createModel(context, () => PageSearchAndNextModel());
  }

  void dispose() {
    webNavModel.dispose();
    sheBeiZuLinItemModel1.dispose();
    sheBeiZuLinItemModels2.dispose();
    pageSearchAndNextModel.dispose();
  }

  /// Additional helper methods are added here.
}
