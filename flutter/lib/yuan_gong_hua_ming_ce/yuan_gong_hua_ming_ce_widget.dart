import '/backend/api_requests/api_calls.dart';
import '/components/page_search_and_next_widget.dart';
import '/components/web_nav_widget.dart';
import '/components/yuan_gong_hua_ming_ce_item_widget.dart';
import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'yuan_gong_hua_ming_ce_model.dart';
export 'yuan_gong_hua_ming_ce_model.dart';

class YuanGongHuaMingCeWidget extends StatefulWidget {
  const YuanGongHuaMingCeWidget({Key? key}) : super(key: key);

  @override
  _YuanGongHuaMingCeWidgetState createState() =>
      _YuanGongHuaMingCeWidgetState();
}

class _YuanGongHuaMingCeWidgetState extends State<YuanGongHuaMingCeWidget> {
  late YuanGongHuaMingCeModel _model;

  final scaffoldKey = GlobalKey<ScaffoldState>();
  final _unfocusNode = FocusNode();

  @override
  void initState() {
    super.initState();
    _model = createModel(context, () => YuanGongHuaMingCeModel());
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
                                '员工花名册',
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
                            Padding(
                              padding: EdgeInsetsDirectional.fromSTEB(
                                  16.0, 0.0, 0.0, 0.0),
                              child: InkWell(
                                onTap: () async {
                                  _model.getRes = await BanGongShiGroup
                                      .getYuanGongHuaMingCeListCall
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
                                      padding: EdgeInsetsDirectional.fromSTEB(
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
                        wrapWithModel(
                          model: _model.yuanGongHuaMingCeItemModel1,
                          updateCallback: () => setState(() {}),
                          child: YuanGongHuaMingCeItemWidget(
                            item: null,
                          ),
                        ),
                        if (!_model.switchValue!)
                          FutureBuilder<ApiCallResponse>(
                            future: BanGongShiGroup.getYuanGongHuaMingCeListCall
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
                              final columnGetYuanGongHuaMingCeListResponse =
                                  snapshot.data!;
                              return Builder(
                                builder: (context) {
                                  final list = BanGongShiGroup
                                          .getYuanGongHuaMingCeListCall
                                          .list(
                                            columnGetYuanGongHuaMingCeListResponse
                                                .jsonBody,
                                          )
                                          ?.toList() ??
                                      [];
                                  return SingleChildScrollView(
                                    child: Column(
                                      mainAxisSize: MainAxisSize.max,
                                      children: List.generate(list.length,
                                          (listIndex) {
                                        final listItem = list[listIndex];
                                        return YuanGongHuaMingCeItemWidget(
                                          key: Key(
                                              'Keybs4_${listIndex}_of_${list.length}'),
                                          item: listItem,
                                        );
                                      }),
                                    ),
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
            ],
          ),
        ),
      ),
    );
  }
}
