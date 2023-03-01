import '/backend/api_requests/api_calls.dart';
import '/components/page_search_and_next_widget.dart';
import '/components/web_nav_widget.dart';
import '/components/xiang_mu_deng_ji_item_widget.dart';
import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'xiang_mu_deng_ji_biao_model.dart';
export 'xiang_mu_deng_ji_biao_model.dart';

class XiangMuDengJiBiaoWidget extends StatefulWidget {
  const XiangMuDengJiBiaoWidget({Key? key}) : super(key: key);

  @override
  _XiangMuDengJiBiaoWidgetState createState() =>
      _XiangMuDengJiBiaoWidgetState();
}

class _XiangMuDengJiBiaoWidgetState extends State<XiangMuDengJiBiaoWidget> {
  late XiangMuDengJiBiaoModel _model;

  final scaffoldKey = GlobalKey<ScaffoldState>();
  final _unfocusNode = FocusNode();

  @override
  void initState() {
    super.initState();
    _model = createModel(context, () => XiangMuDengJiBiaoModel());
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

    return Scaffold(
      key: scaffoldKey,
      backgroundColor: FlutterFlowTheme.of(context).primaryBackground,
      body: SafeArea(
        child: GestureDetector(
          onTap: () => FocusScope.of(context).requestFocus(_unfocusNode),
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
                                  '项目登记表',
                                  style: FlutterFlowTheme.of(context)
                                      .title1
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
                                  setState(() => _model.switchValue = newValue);
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
                            ],
                          ),
                          _model.switchValue!
                              ? XiangMuDengJiItemWidget(
                                  key: Key('new'),
                                  item: _model,
                                )
                              : FutureBuilder<ApiCallResponse>(
                                  future: JingYingBuGroup
                                      .getXiangMuDengJiBiaoListCall
                                      .call(
                                    page: 1,
                                    pageSize: 10,
                                  ),
                                  builder: (context, snapshot) {
                                    // Customize what your widget looks like when it's loading.
                                    if (!snapshot.hasData) {
                                      return Center(
                                        child: SizedBox(
                                          width: 100.0,
                                          height: 50.0,
                                          child: Text('暂无数据'),
                                        ),
                                      );
                                    }
                                    final columnGetXiangMuDengJiBiaoListResponse =
                                        snapshot.data!;
                                    return Builder(
                                      builder: (context) {
                                        final list = JingYingBuGroup
                                                .getXiangMuDengJiBiaoListCall
                                                .list(
                                                  columnGetXiangMuDengJiBiaoListResponse
                                                      .jsonBody,
                                                )
                                                ?.toList() ??
                                            [];
                                        return Column(
                                          mainAxisSize: MainAxisSize.max,
                                          children: List.generate(list.length,
                                              (listIndex) {
                                            final listItem = list[listIndex];
                                            return XiangMuDengJiItemWidget(
                                              key: Key(
                                                  'Key3qa_${listIndex}_of_${list.length}'),
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
            ],
          ),
        ),
      ),
    );
  }
}
