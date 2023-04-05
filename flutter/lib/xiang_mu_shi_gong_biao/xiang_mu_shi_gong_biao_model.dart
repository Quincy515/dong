import '../components/page_search_and_next_model.dart';
import '../components/web_nav_model.dart';
import '/backend/api_requests/api_calls.dart';
import '/components/xiang_mu_shi_gong_item/xiang_mu_shi_gong_item_widget.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';

class XiangMuShiGongBiaoModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this page.

  // Model for webNav component.
  late WebNavModel webNavModel;
  // State field(s) for Switch widget.
  bool? switchValue;
  // Stores action output result for [Backend Call - API (getMeiRiKaoQinList)] action in Row widget.
  ApiCallResponse? getRes;
  // Model for xiangMuShiGongItem component.
  late XiangMuShiGongItemModel xiangMuShiGongItemModel1;
  // Models for xiangMuShiGongItem dynamic component.
  late FlutterFlowDynamicModels<XiangMuShiGongItemModel>
      xiangMuShiGongItemModels2;
  // Model for pageSearchAndNext component.
  late PageSearchAndNextModel pageSearchAndNextModel;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {
    webNavModel = createModel(context, () => WebNavModel());
    xiangMuShiGongItemModel1 =
        createModel(context, () => XiangMuShiGongItemModel());
    xiangMuShiGongItemModels2 =
        FlutterFlowDynamicModels(() => XiangMuShiGongItemModel());
    pageSearchAndNextModel =
        createModel(context, () => PageSearchAndNextModel());
  }

  void dispose() {
    webNavModel.dispose();
    xiangMuShiGongItemModel1.dispose();
    xiangMuShiGongItemModels2.dispose();
    pageSearchAndNextModel.dispose();
  }

  /// Additional helper methods are added here.
}
