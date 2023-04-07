import '../components/page_search_and_next_model.dart';
import '../components/web_nav_model.dart';
import '/backend/api_requests/api_calls.dart';
import '/components/xiang_mu_zhu_ca_item/xiang_mu_zhu_ca_item_widget.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';

class XiangMuZhuCaiBiaoModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this page.

  // Model for webNav component.
  late WebNavModel webNavModel;
  // State field(s) for Switch widget.
  bool? switchValue;
  // Stores action output result for [Backend Call - API (getMeiRiKaoQinList)] action in Row widget.
  ApiCallResponse? getRes;
  // Model for xiangMuZhuCaItem component.
  late XiangMuZhuCaItemModel xiangMuZhuCaItemModel1;
  // Models for xiangMuZhuCaItem dynamic component.
  late FlutterFlowDynamicModels<XiangMuZhuCaItemModel> xiangMuZhuCaItemModels2;
  // Model for pageSearchAndNext component.
  late PageSearchAndNextModel pageSearchAndNextModel;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {
    webNavModel = createModel(context, () => WebNavModel());
    xiangMuZhuCaItemModel1 =
        createModel(context, () => XiangMuZhuCaItemModel());
    xiangMuZhuCaItemModels2 =
        FlutterFlowDynamicModels(() => XiangMuZhuCaItemModel());
    pageSearchAndNextModel =
        createModel(context, () => PageSearchAndNextModel());
  }

  void dispose() {
    webNavModel.dispose();
    xiangMuZhuCaItemModel1.dispose();
    xiangMuZhuCaItemModels2.dispose();
    pageSearchAndNextModel.dispose();
  }

  /// Additional helper methods are added here.
}
