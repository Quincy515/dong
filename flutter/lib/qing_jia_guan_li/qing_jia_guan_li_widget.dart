import 'dart:html';

import '../components/page_search_and_next_widget.dart';
import '../components/qing_jia_guan_li_item/qing_jia_guan_li_item_widget.dart';
import '../components/web_nav_widget.dart';
import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'qing_jia_guan_li_model.dart';
export 'qing_jia_guan_li_model.dart';

class QingJiaGuanLiWidget extends StatefulWidget {
  const QingJiaGuanLiWidget({
    Key? key,
    this.id,
    this.detail,
    // bool? detail,
    this.xiaoXiId,
  }) : //this.detail = detail ?? false,
        super(key: key);

  final String? id;
  final bool? detail;
  // final bool detail;
  final String? xiaoXiId;

  @override
  _QingJiaGuanLiWidgetState createState() => _QingJiaGuanLiWidgetState();
}

class _QingJiaGuanLiWidgetState extends State<QingJiaGuanLiWidget> {
  late QingJiaGuanLiModel _model;

  final scaffoldKey = GlobalKey<ScaffoldState>();
  final _unfocusNode = FocusNode();
  String? id;
  bool? detail;
  String? xiaoXiId;
  @override
  void initState() {
    super.initState();
    _model = createModel(context, () => QingJiaGuanLiModel());

    var uri = Uri.dataFromString(window.location.href);
    var qp = uri.queryParameters;
    id = qp['id'] ?? widget.id;
    detail = qp['detail'] == 'true' ? true : false;
    xiaoXiId = qp['xiaoXiId'] ?? widget.xiaoXiId;
  }

  @override
  void dispose() {
    _model.dispose();

    _unfocusNode.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    context.watch<FFAppState>();
    print(
        '>>>url: ${window.location.href}, id: $id detail: $detail, xiaoXiId: $xiaoXiId');
    print(
        '>>> id: ${widget.id} detail: ${widget.detail}, xiaoXiId: ${widget.xiaoXiId}');
    return GestureDetector(
      onTap: () => FocusScope.of(context).requestFocus(_unfocusNode),
      child: Scaffold(
        key: scaffoldKey,
        backgroundColor: FlutterFlowTheme.of(context).primaryBackground,
        body: SafeArea(
          child: Row(
            mainAxisSize: MainAxisSize.max,
            children: [
              wrapWithModel(
                model: _model.webNavModel,
                updateCallback: () => setState(() {}),
                child: WebNavWidget(
                  navBgOne: FlutterFlowTheme.of(context).primaryBackground,
                  navColorOne: FlutterFlowTheme.of(context).secondaryBackground,
                  navBgTwo: FlutterFlowTheme.of(context).secondaryBackground,
                  navColorTwo: FlutterFlowTheme.of(context).secondaryText,
                  navBgThree: FlutterFlowTheme.of(context).secondaryBackground,
                  navColorThree: FlutterFlowTheme.of(context).secondaryText,
                  navBgFour: FlutterFlowTheme.of(context).secondaryBackground,
                  navColorFour: FlutterFlowTheme.of(context).secondaryText,
                  navBgFive: FlutterFlowTheme.of(context).secondaryBackground,
                  navColorFive: FlutterFlowTheme.of(context).secondaryText,
                ),
              ),
              if (widget.detail == false || detail == false)
                Expanded(
                  child: Padding(
                    padding: EdgeInsetsDirectional.fromSTEB(8.0, 8.0, 8.0, 8.0),
                    child: Container(
                      decoration: BoxDecoration(
                        color: FlutterFlowTheme.of(context).secondaryBackground,
                        borderRadius: BorderRadius.circular(20.0),
                      ),
                      child: SingleChildScrollView(
                        child: Column(
                          mainAxisSize: MainAxisSize.max,
                          mainAxisAlignment: MainAxisAlignment.spaceBetween,
                          children: [
                            Padding(
                              padding: EdgeInsetsDirectional.fromSTEB(
                                  8.0, 8.0, 8.0, 8.0),
                              child: Row(
                                mainAxisSize: MainAxisSize.max,
                                mainAxisAlignment: MainAxisAlignment.center,
                                children: [
                                  Text(
                                    '请假管理',
                                    style: FlutterFlowTheme.of(context)
                                        .title3
                                        .override(
                                          fontFamily: 'Poppins',
                                          fontWeight: FontWeight.normal,
                                        ),
                                  ),
                                ],
                              ),
                            ),
                            Row(
                              mainAxisSize: MainAxisSize.max,
                              children: [
                                Switch(
                                  value: _model.switchValue ??= false,
                                  onChanged: (newValue) async {
                                    setState(
                                        () => _model.switchValue = newValue);
                                  },
                                ),
                                Text(
                                  '新增',
                                  style: FlutterFlowTheme.of(context)
                                      .bodyText1
                                      .override(
                                        fontFamily: 'Poppins',
                                        fontWeight: FontWeight.normal,
                                      ),
                                ),
                                Padding(
                                  padding: EdgeInsetsDirectional.fromSTEB(
                                      16.0, 0.0, 0.0, 0.0),
                                  child: InkWell(
                                    onTap: () async {
                                      _model.getRes = await BanGongShiGroup
                                          .getQingJiaGuanLiListCall
                                          .call(
                                        page: 1,
                                        pageSize: 10,
                                      );

                                      setState(() {});
                                    },
                                    child: Row(
                                      mainAxisSize: MainAxisSize.max,
                                      children: [
                                        Icon(
                                          Icons.refresh,
                                          color: FlutterFlowTheme.of(context)
                                              .secondaryText,
                                          size: 24.0,
                                        ),
                                        Padding(
                                          padding:
                                              EdgeInsetsDirectional.fromSTEB(
                                                  8.0, 0.0, 0.0, 0.0),
                                          child: Text(
                                            '刷新',
                                            style: FlutterFlowTheme.of(context)
                                                .bodyText1
                                                .override(
                                                  fontFamily: 'Poppins',
                                                  fontWeight: FontWeight.normal,
                                                ),
                                          ),
                                        ),
                                      ],
                                    ),
                                  ),
                                ),
                              ],
                            ),
                            if (_model.switchValue ?? true)
                              wrapWithModel(
                                model: _model.qingJiaGuanLiItemModel1,
                                updateCallback: () => setState(() {}),
                                child: QingJiaGuanLiItemWidget(
                                  item: jsonDecode('''
{
  "beiZhu": "",
  "group": "dong",
  "jieShuRiQi": "${DateTime.now()}",
  "jinDu": "提交申请",
  "kaiShiRiQi": "${DateTime.now()}",
  "leiXing": 1,
  "shenHeRenId": 0,
  "shenHeShiJian": "${DateTime.now()}",
  "shenHeXiangQing": "",
  "shenHeZhuangTai": "",
  "shiYou": "",
  "tianShu": 0,
  "xingMingId": ${FFAppState().userID}
}
  '''),
                                ),
                              ),
                            if (!_model.switchValue!)
                              FutureBuilder<ApiCallResponse>(
                                future: BanGongShiGroup.getQingJiaGuanLiListCall
                                    .call(
                                  page: 1,
                                  pageSize: 10,
                                ),
                                builder: (context, snapshot) {
                                  // Customize what your widget looks like when it's loading.
                                  if (!snapshot.hasData) {
                                    return Center(
                                      child: SizedBox(
                                        width: 50.0,
                                        height: 50.0,
                                        child: CircularProgressIndicator(
                                          color: FlutterFlowTheme.of(context)
                                              .primaryColor,
                                        ),
                                      ),
                                    );
                                  }
                                  final columnGetQingJiaGuanLiListResponse =
                                      snapshot.data!;
                                  return Builder(
                                    builder: (context) {
                                      final list = BanGongShiGroup
                                              .getQingJiaGuanLiListCall
                                              .list(
                                                columnGetQingJiaGuanLiListResponse
                                                    .jsonBody,
                                              )
                                              ?.toList() ??
                                          [];
                                      return Column(
                                        mainAxisSize: MainAxisSize.max,
                                        children: List.generate(list.length,
                                            (listIndex) {
                                          final listItem = list[listIndex];
                                          return QingJiaGuanLiItemWidget(
                                            key: Key(
                                                'Keyppo_${listIndex}_of_${list.length}'),
                                            item: listItem,
                                          );
                                        }),
                                      );
                                    },
                                  );
                                },
                              ),
                            wrapWithModel(
                              model: _model.pageSearchAndNextModel,
                              updateCallback: () => setState(() {}),
                              child: PageSearchAndNextWidget(),
                            ),
                          ],
                        ),
                      ),
                    ),
                  ),
                ),
              if (widget.detail == true || detail == true)
                Expanded(
                  child: Padding(
                    padding: EdgeInsetsDirectional.fromSTEB(8, 8, 8, 8),
                    child: Container(
                      decoration: BoxDecoration(
                        color: FlutterFlowTheme.of(context).secondaryBackground,
                        borderRadius: BorderRadius.circular(20),
                      ),
                      child: FutureBuilder<ApiCallResponse>(
                        future: BanGongShiGroup.findQingJiaGuanLiCall.call(
                          id: int.parse(id ?? widget.id!),
                        ),
                        builder: (context, snapshot) {
                          // Customize what your widget looks like when it's loading.
                          if (!snapshot.hasData) {
                            return Center(
                              child: SizedBox(
                                width: 50,
                                height: 50,
                                child: CircularProgressIndicator(
                                  color:
                                      FlutterFlowTheme.of(context).primaryColor,
                                ),
                              ),
                            );
                          }
                          final res = snapshot.data!;
                          final item = BanGongShiGroup.findQingJiaGuanLiCall
                              .item(res.jsonBody);
                          print(
                              '>>>item: $item, shiYou: ${getJsonField(item, r'''$.shiYou''')} ');
                          return SingleChildScrollView(
                            child: Column(
                              mainAxisSize: MainAxisSize.max,
                              mainAxisAlignment: MainAxisAlignment.spaceBetween,
                              children: [
                                Padding(
                                  padding: EdgeInsetsDirectional.fromSTEB(
                                      8, 8, 8, 8),
                                  child: Row(
                                    mainAxisSize: MainAxisSize.max,
                                    mainAxisAlignment: MainAxisAlignment.center,
                                    children: [
                                      Text(
                                        '请假管理',
                                        style: FlutterFlowTheme.of(context)
                                            .title3
                                            .override(
                                              fontFamily: 'Poppins',
                                              fontWeight: FontWeight.normal,
                                            ),
                                      ),
                                    ],
                                  ),
                                ),
                                wrapWithModel(
                                  model: _model.qingJiaGuanLiItemModel3,
                                  updateCallback: () => setState(() {}),
                                  child: QingJiaGuanLiItemWidget(
                                    item: jsonDecode('''
{
  "detail": true,
  "ID": ${BanGongShiGroup.findQingJiaGuanLiCall.id(res.jsonBody)},
  "beiZhu": "${BanGongShiGroup.findQingJiaGuanLiCall.beiZhu(res.jsonBody)}",
  "group": "${BanGongShiGroup.findQingJiaGuanLiCall.group(res.jsonBody)}",
  "jieShuRiQi": "${BanGongShiGroup.findQingJiaGuanLiCall.jieShuRiQi(res.jsonBody)}",
  "jinDu": "${BanGongShiGroup.findQingJiaGuanLiCall.jinDu(res.jsonBody)}",
  "kaiShiRiQi": "${BanGongShiGroup.findQingJiaGuanLiCall.kaiShiRiQi(res.jsonBody)}",
  "leiXing": ${BanGongShiGroup.findQingJiaGuanLiCall.leiXing(res.jsonBody)},
  "shenHeRenId": ${BanGongShiGroup.findQingJiaGuanLiCall.shenHeRenId(res.jsonBody)},
  "shenHeShiJian": "${BanGongShiGroup.findQingJiaGuanLiCall.shenHeShiJian(res.jsonBody) ?? DateTime.now()}",
  "shenHeXiangQing": "${BanGongShiGroup.findQingJiaGuanLiCall.shenHeXiangQing(res.jsonBody)}",
  "shenHeZhuangTai": "${BanGongShiGroup.findQingJiaGuanLiCall.shenHeZhuangTai(res.jsonBody)}",
  "shiYou": "${BanGongShiGroup.findQingJiaGuanLiCall.shiYou(res.jsonBody)}",
  "tianShu": ${BanGongShiGroup.findQingJiaGuanLiCall.tianShu(res.jsonBody)},
  "xingMingId": ${BanGongShiGroup.findQingJiaGuanLiCall.xingMingId(res.jsonBody)}
}
  '''),
                                  ),
                                ),
                              ],
                            ),
                          );
                        },
                      ),
                    ),
                  ),
                )
            ],
          ),
        ),
      ),
    );
  }
}
