import '../components/page_search_and_next_model.dart';
import '../components/web_nav_model.dart';
import '/backend/api_requests/api_calls.dart';
import '/components/chu_ru_ku_ming_xiitem/chu_ru_ku_ming_xiitem_widget.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';

class ChuRuKuMingXiBiaoModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this page.

  // Model for webNav component.
  late WebNavModel webNavModel;
  // State field(s) for Switch widget.
  bool? switchValue;
  // Stores action output result for [Backend Call - API (getMeiRiKaoQinList)] action in Row widget.
  ApiCallResponse? getRes;
  // Model for chuRuKuMingXiitem component.
  late ChuRuKuMingXiitemModel chuRuKuMingXiitemModel1;
  // Models for chuRuKuMingXiitem dynamic component.
  late FlutterFlowDynamicModels<ChuRuKuMingXiitemModel>
      chuRuKuMingXiitemModels2;
  // Model for pageSearchAndNext component.
  late PageSearchAndNextModel pageSearchAndNextModel;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {
    webNavModel = createModel(context, () => WebNavModel());
    chuRuKuMingXiitemModel1 =
        createModel(context, () => ChuRuKuMingXiitemModel());
    chuRuKuMingXiitemModels2 =
        FlutterFlowDynamicModels(() => ChuRuKuMingXiitemModel());
    pageSearchAndNextModel =
        createModel(context, () => PageSearchAndNextModel());
  }

  void dispose() {
    webNavModel.dispose();
    chuRuKuMingXiitemModel1.dispose();
    chuRuKuMingXiitemModels2.dispose();
    pageSearchAndNextModel.dispose();
  }

  /// Additional helper methods are added here.
}
