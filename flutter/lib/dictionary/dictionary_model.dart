// ignore_for_file: unused_import

import '/backend/api_requests/api_calls.dart';
import '/components/dictionary_item_widget.dart';
import '/components/page_search_and_next_widget.dart';
import '/components/web_nav_widget.dart';
import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:provider/provider.dart';

class DictionaryModel extends FlutterFlowModel {
  ///  State fields for stateful widgets in this page.

  // Model for webNav component.
  late WebNavModel webNavModel;
  // State field(s) for Switch widget.
  bool? switchValue;
  // Stores action output result for [Backend Call - API (getSysDictionaryList)] action in Row widget.
  ApiCallResponse? getRes;
  // Model for dictionaryItem component.
  late DictionaryItemModel dictionaryItemModel1;
  // Model for pageSearchAndNext component.
  late PageSearchAndNextModel pageSearchAndNextModel;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {
    webNavModel = createModel(context, () => WebNavModel());
    dictionaryItemModel1 = createModel(context, () => DictionaryItemModel());
    pageSearchAndNextModel =
        createModel(context, () => PageSearchAndNextModel());
  }

  void dispose() {
    webNavModel.dispose();
    dictionaryItemModel1.dispose();
    pageSearchAndNextModel.dispose();
  }

  /// Additional helper methods are added here.
}
