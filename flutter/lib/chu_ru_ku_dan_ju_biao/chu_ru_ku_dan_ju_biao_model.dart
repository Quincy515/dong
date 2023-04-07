import '../components/page_search_and_next_model.dart';
import '../components/web_nav_model.dart';
import '/backend/api_requests/api_calls.dart';
import '/components/chu_ru_ku_dan_j_item/chu_ru_ku_dan_j_item_widget.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';

class ChuRuKuDanJuBiaoModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this page.

  // Model for webNav component.
  late WebNavModel webNavModel;
  // State field(s) for Switch widget.
  bool? switchValue;
  // Stores action output result for [Backend Call - API (getMeiRiKaoQinList)] action in Row widget.
  ApiCallResponse? getRes;
  // Model for chuRuKuDanJItem component.
  late ChuRuKuDanJItemModel chuRuKuDanJItemModel1;
  // Models for chuRuKuDanJItem dynamic component.
  late FlutterFlowDynamicModels<ChuRuKuDanJItemModel> chuRuKuDanJItemModels2;
  // Model for pageSearchAndNext component.
  late PageSearchAndNextModel pageSearchAndNextModel;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {
    webNavModel = createModel(context, () => WebNavModel());
    chuRuKuDanJItemModel1 = createModel(context, () => ChuRuKuDanJItemModel());
    chuRuKuDanJItemModels2 =
        FlutterFlowDynamicModels(() => ChuRuKuDanJItemModel());
    pageSearchAndNextModel =
        createModel(context, () => PageSearchAndNextModel());
  }

  void dispose() {
    webNavModel.dispose();
    chuRuKuDanJItemModel1.dispose();
    chuRuKuDanJItemModels2.dispose();
    pageSearchAndNextModel.dispose();
  }

  /// Additional helper methods are added here.
}
