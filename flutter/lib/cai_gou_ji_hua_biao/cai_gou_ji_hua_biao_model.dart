import '../components/page_search_and_next_model.dart';
import '../components/web_nav_model.dart';
import '/backend/api_requests/api_calls.dart';
import '/components/cai_gou_ji_hua_item/cai_gou_ji_hua_item_widget.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';

class CaiGouJiHuaBiaoModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this page.

  // Model for webNav component.
  late WebNavModel webNavModel;
  // State field(s) for Switch widget.
  bool? switchValue;
  // Stores action output result for [Backend Call - API (getMeiRiKaoQinList)] action in Row widget.
  ApiCallResponse? getRes;
  // Model for caiGouJiHuaItem component.
  late CaiGouJiHuaItemModel caiGouJiHuaItemModel1;
  // Models for caiGouJiHuaItem dynamic component.
  late FlutterFlowDynamicModels<CaiGouJiHuaItemModel> caiGouJiHuaItemModels2;
  // Model for pageSearchAndNext component.
  late PageSearchAndNextModel pageSearchAndNextModel;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {
    webNavModel = createModel(context, () => WebNavModel());
    caiGouJiHuaItemModel1 = createModel(context, () => CaiGouJiHuaItemModel());
    caiGouJiHuaItemModels2 =
        FlutterFlowDynamicModels(() => CaiGouJiHuaItemModel());
    pageSearchAndNextModel =
        createModel(context, () => PageSearchAndNextModel());
  }

  void dispose() {
    webNavModel.dispose();
    caiGouJiHuaItemModel1.dispose();
    caiGouJiHuaItemModels2.dispose();
    pageSearchAndNextModel.dispose();
  }

  /// Additional helper methods are added here.
}
