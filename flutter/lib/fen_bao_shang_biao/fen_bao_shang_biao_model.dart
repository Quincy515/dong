import '../components/page_search_and_next_model.dart';
import '../components/web_nav_model.dart';
import '/backend/api_requests/api_calls.dart';
import '/components/fen_bao_shang_item/fen_bao_shang_item_widget.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';

class FenBaoShangBiaoModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this page.

  // Model for webNav component.
  late WebNavModel webNavModel;
  // State field(s) for Switch widget.
  bool? switchValue;
  // Stores action output result for [Backend Call - API (getMeiRiKaoQinList)] action in Row widget.
  ApiCallResponse? getRes;
  // Model for fenBaoShangItem component.
  late FenBaoShangItemModel fenBaoShangItemModel1;
  // Models for fenBaoShangItem dynamic component.
  late FlutterFlowDynamicModels<FenBaoShangItemModel> fenBaoShangItemModels2;
  // Model for pageSearchAndNext component.
  late PageSearchAndNextModel pageSearchAndNextModel;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {
    webNavModel = createModel(context, () => WebNavModel());
    fenBaoShangItemModel1 = createModel(context, () => FenBaoShangItemModel());
    fenBaoShangItemModels2 =
        FlutterFlowDynamicModels(() => FenBaoShangItemModel());
    pageSearchAndNextModel =
        createModel(context, () => PageSearchAndNextModel());
  }

  void dispose() {
    webNavModel.dispose();
    fenBaoShangItemModel1.dispose();
    fenBaoShangItemModels2.dispose();
    pageSearchAndNextModel.dispose();
  }

  /// Additional helper methods are added here.
}
