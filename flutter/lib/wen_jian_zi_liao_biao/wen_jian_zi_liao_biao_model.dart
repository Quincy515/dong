import '../components/page_search_and_next_model.dart';
import '../components/web_nav_model.dart';
import '/backend/api_requests/api_calls.dart';
import '/components/wen_jian_zi_liao_item/wen_jian_zi_liao_item_widget.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';

class WenJianZiLiaoBiaoModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this page.

  // Model for webNav component.
  late WebNavModel webNavModel;
  // State field(s) for Switch widget.
  bool? switchValue;
  // Stores action output result for [Backend Call - API (getMeiRiKaoQinList)] action in Row widget.
  ApiCallResponse? getRes;
  // Model for wenJianZiLiaoItem component.
  late WenJianZiLiaoItemModel wenJianZiLiaoItemModel1;
  // Models for wenJianZiLiaoItem dynamic component.
  late FlutterFlowDynamicModels<WenJianZiLiaoItemModel>
      wenJianZiLiaoItemModels2;
  // Model for pageSearchAndNext component.
  late PageSearchAndNextModel pageSearchAndNextModel;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {
    webNavModel = createModel(context, () => WebNavModel());
    wenJianZiLiaoItemModel1 =
        createModel(context, () => WenJianZiLiaoItemModel());
    wenJianZiLiaoItemModels2 =
        FlutterFlowDynamicModels(() => WenJianZiLiaoItemModel());
    pageSearchAndNextModel =
        createModel(context, () => PageSearchAndNextModel());
  }

  void dispose() {
    webNavModel.dispose();
    wenJianZiLiaoItemModel1.dispose();
    wenJianZiLiaoItemModels2.dispose();
    pageSearchAndNextModel.dispose();
  }

  /// Additional helper methods are added here.
}
