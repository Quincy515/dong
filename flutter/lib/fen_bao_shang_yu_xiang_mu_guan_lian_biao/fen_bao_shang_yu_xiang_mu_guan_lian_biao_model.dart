import '../components/page_search_and_next_model.dart';
import '../components/web_nav_model.dart';
import '/backend/api_requests/api_calls.dart';
import '/components/fen_bao_shang_yu_xiang_mu_guan_lian_item/fen_bao_shang_yu_xiang_mu_guan_lian_item_widget.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';

class FenBaoShangYuXiangMuGuanLianBiaoModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this page.

  // Model for webNav component.
  late WebNavModel webNavModel;
  // State field(s) for Switch widget.
  bool? switchValue;
  // Stores action output result for [Backend Call - API (getMeiRiKaoQinList)] action in Row widget.
  ApiCallResponse? getRes;
  // Model for fenBaoShangYuXiangMuGuanLianItem component.
  late FenBaoShangYuXiangMuGuanLianItemModel
      fenBaoShangYuXiangMuGuanLianItemModel1;
  // Models for fenBaoShangYuXiangMuGuanLianItem dynamic component.
  late FlutterFlowDynamicModels<FenBaoShangYuXiangMuGuanLianItemModel>
      fenBaoShangYuXiangMuGuanLianItemModels2;
  // Model for pageSearchAndNext component.
  late PageSearchAndNextModel pageSearchAndNextModel;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {
    webNavModel = createModel(context, () => WebNavModel());
    fenBaoShangYuXiangMuGuanLianItemModel1 =
        createModel(context, () => FenBaoShangYuXiangMuGuanLianItemModel());
    fenBaoShangYuXiangMuGuanLianItemModels2 =
        FlutterFlowDynamicModels(() => FenBaoShangYuXiangMuGuanLianItemModel());
    pageSearchAndNextModel =
        createModel(context, () => PageSearchAndNextModel());
  }

  void dispose() {
    webNavModel.dispose();
    fenBaoShangYuXiangMuGuanLianItemModel1.dispose();
    fenBaoShangYuXiangMuGuanLianItemModels2.dispose();
    pageSearchAndNextModel.dispose();
  }

  /// Additional helper methods are added here.
}
