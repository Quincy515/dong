import '../components/he_tong_wen_jian_item/he_tong_wen_jian_item_model.dart';
import '../components/page_search_and_next_model.dart';
import '../components/web_nav_model.dart';
import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';

class HeTongWenJianBiaoModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this page.

  // Model for webNav component.
  late WebNavModel webNavModel;
  // State field(s) for Switch widget.
  bool? switchValue;
  // Stores action output result for [Backend Call - API (getMeiRiKaoQinList)] action in Row widget.
  ApiCallResponse? getRes;
  // Model for heTongWenJianItem component.
  late HeTongWenJianItemModel heTongWenJianItemModel1;
  // Models for heTongWenJianItem dynamic component.
  late FlutterFlowDynamicModels<HeTongWenJianItemModel>
      heTongWenJianItemModels2;
  // Model for pageSearchAndNext component.
  late PageSearchAndNextModel pageSearchAndNextModel;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {
    webNavModel = createModel(context, () => WebNavModel());
    heTongWenJianItemModel1 =
        createModel(context, () => HeTongWenJianItemModel());
    heTongWenJianItemModels2 =
        FlutterFlowDynamicModels(() => HeTongWenJianItemModel());
    pageSearchAndNextModel =
        createModel(context, () => PageSearchAndNextModel());
  }

  void dispose() {
    webNavModel.dispose();
    heTongWenJianItemModel1.dispose();
    heTongWenJianItemModels2.dispose();
    pageSearchAndNextModel.dispose();
  }

  /// Additional helper methods are added here.
}
