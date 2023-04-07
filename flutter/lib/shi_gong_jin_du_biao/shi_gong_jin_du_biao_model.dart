import '../components/page_search_and_next_model.dart';
import '../components/web_nav_model.dart';
import '/backend/api_requests/api_calls.dart';
import '/components/shi_gong_jin_du_item/shi_gong_jin_du_item_widget.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';

class ShiGongJinDuBiaoModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this page.

  // Model for webNav component.
  late WebNavModel webNavModel;
  // State field(s) for Switch widget.
  bool? switchValue;
  // Stores action output result for [Backend Call - API (getMeiRiKaoQinList)] action in Row widget.
  ApiCallResponse? getRes;
  // Model for shiGongJinDuItem component.
  late ShiGongJinDuItemModel shiGongJinDuItemModel1;
  // Models for shiGongJinDuItem dynamic component.
  late FlutterFlowDynamicModels<ShiGongJinDuItemModel> shiGongJinDuItemModels2;
  // Model for pageSearchAndNext component.
  late PageSearchAndNextModel pageSearchAndNextModel;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {
    webNavModel = createModel(context, () => WebNavModel());
    shiGongJinDuItemModel1 =
        createModel(context, () => ShiGongJinDuItemModel());
    shiGongJinDuItemModels2 =
        FlutterFlowDynamicModels(() => ShiGongJinDuItemModel());
    pageSearchAndNextModel =
        createModel(context, () => PageSearchAndNextModel());
  }

  void dispose() {
    webNavModel.dispose();
    shiGongJinDuItemModel1.dispose();
    shiGongJinDuItemModels2.dispose();
    pageSearchAndNextModel.dispose();
  }

  /// Additional helper methods are added here.
}
