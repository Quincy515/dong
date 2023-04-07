import '../components/page_search_and_next_model.dart';
import '../components/web_nav_model.dart';
import '/backend/api_requests/api_calls.dart';
import '/components/she_bei_ming_xi_item/she_bei_ming_xi_item_widget.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';

class SheBeiMingXiBiaoModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this page.

  // Model for webNav component.
  late WebNavModel webNavModel;
  // State field(s) for Switch widget.
  bool? switchValue;
  // Stores action output result for [Backend Call - API (getMeiRiKaoQinList)] action in Row widget.
  ApiCallResponse? getRes;
  // Model for sheBeiMingXiItem component.
  late SheBeiMingXiItemModel sheBeiMingXiItemModel1;
  // Models for sheBeiMingXiItem dynamic component.
  late FlutterFlowDynamicModels<SheBeiMingXiItemModel> sheBeiMingXiItemModels2;
  // Model for pageSearchAndNext component.
  late PageSearchAndNextModel pageSearchAndNextModel;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {
    webNavModel = createModel(context, () => WebNavModel());
    sheBeiMingXiItemModel1 =
        createModel(context, () => SheBeiMingXiItemModel());
    sheBeiMingXiItemModels2 =
        FlutterFlowDynamicModels(() => SheBeiMingXiItemModel());
    pageSearchAndNextModel =
        createModel(context, () => PageSearchAndNextModel());
  }

  void dispose() {
    webNavModel.dispose();
    sheBeiMingXiItemModel1.dispose();
    sheBeiMingXiItemModels2.dispose();
    pageSearchAndNextModel.dispose();
  }

  /// Additional helper methods are added here.
}
