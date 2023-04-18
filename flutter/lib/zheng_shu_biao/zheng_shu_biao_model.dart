import '../components/page_search_and_next_model.dart';
import '../components/web_nav_model.dart';
import '../components/zheng_shu_biao_item_model.dart';
import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';

class ZhengShuBiaoModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this page.

  // Model for webNav component.
  late WebNavModel webNavModel;
  // State field(s) for Switch widget.
  bool? switchValue;
  // Stores action output result for [Backend Call - API (getZhengShuBiaoList)] action in Row widget.
  ApiCallResponse? getRes;
  // Model for zhengShuBiaoItem component.
  late ZhengShuBiaoItemModel zhengShuBiaoItemModel1;
  // Model for pageSearchAndNext component.
  late PageSearchAndNextModel pageSearchAndNextModel;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {
    webNavModel = createModel(context, () => WebNavModel());
    zhengShuBiaoItemModel1 =
        createModel(context, () => ZhengShuBiaoItemModel());
    pageSearchAndNextModel =
        createModel(context, () => PageSearchAndNextModel());
  }

  void dispose() {
    webNavModel.dispose();
    zhengShuBiaoItemModel1.dispose();
    pageSearchAndNextModel.dispose();
  }

  /// Additional helper methods are added here.
}
