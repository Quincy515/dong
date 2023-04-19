import '../components/web_nav_model.dart';
import '../flutter_flow/form_field_controller.dart';
import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';

class HomePageModel extends FlutterFlowModel {
  ///  Local state fields for this page.

  bool qingJia = false;

  ///  State fields for stateful widgets in this page.

  // Model for webNav component.
  late WebNavModel webNavModel;
  // Stores action output result for [Backend Call - API (updateQingJiaGuanLi)] action in Button widget.
  ApiCallResponse? updateRes;
  // State field(s) for lei_xing widget.
  TextEditingController? leiXingController;
  String? Function(BuildContext, String?)? leiXingControllerValidator;
  // State field(s) for tian_shu widget.
  TextEditingController? tianShuController;
  String? Function(BuildContext, String?)? tianShuControllerValidator;
  DateTime? datePicked1;
  DateTime? datePicked2;
  // State field(s) for DropDown widget.
  int? dropDownValue;
  FormFieldController<int>? dropDownController;
  int? leiXingValue;
  FormFieldController<int>? leiXingValueController;
  // State field(s) for shi_you widget.
  TextEditingController? shiYouController;
  String? Function(BuildContext, String?)? shiYouControllerValidator;
  // Stores action output result for [Backend Call - API (updateXiaoXiTongZhi)] action in Column widget.
  ApiCallResponse? makeReadRes;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {
    webNavModel = createModel(context, () => WebNavModel());
  }

  void dispose() {
    webNavModel.dispose();
    leiXingController?.dispose();
    tianShuController?.dispose();
    shiYouController?.dispose();
  }

  /// Additional helper methods are added here.
}
