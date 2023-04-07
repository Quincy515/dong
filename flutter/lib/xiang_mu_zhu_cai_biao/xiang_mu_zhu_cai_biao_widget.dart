import '../components/page_search_and_next_widget.dart';
import '../components/web_nav_widget.dart';
import '/backend/api_requests/api_calls.dart';
import '/components/xiang_mu_zhu_ca_item/xiang_mu_zhu_ca_item_widget.dart';
import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'xiang_mu_zhu_cai_biao_model.dart';
export 'xiang_mu_zhu_cai_biao_model.dart';

class XiangMuZhuCaiBiaoWidget extends StatefulWidget {
  const XiangMuZhuCaiBiaoWidget({Key? key}) : super(key: key);

  @override
  _XiangMuZhuCaiBiaoWidgetState createState() =>
      _XiangMuZhuCaiBiaoWidgetState();
}

class _XiangMuZhuCaiBiaoWidgetState extends State<XiangMuZhuCaiBiaoWidget> {
  late XiangMuZhuCaiBiaoModel _model;

  final scaffoldKey = GlobalKey<ScaffoldState>();
  final _unfocusNode = FocusNode();

  @override
  void initState() {
    super.initState();
    _model = createModel(context, () => XiangMuZhuCaiBiaoModel());
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
                                  '项目主材表',
                                  style: FlutterFlowTheme.of(context)
                                      .bodyText1
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
                                        .getMeiRiKaoQinListCall
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
                          if (_model.switchValue ?? true)
                            wrapWithModel(
                              model: _model.xiangMuZhuCaItemModel1,
                              updateCallback: () => setState(() {}),
                              child: XiangMuZhuCaItemWidget(
                                item: jsonDecode('''
{
  "beiZhu": "",
  "caiLiaoBianHao": "",
  "caiLiaoMingCheng": "",
  "danJia": 0,
  "danWei": "",
  "group": "",
  "heJi": 0,
  "jiaChaHeJi": 0,
  "xianHangJiaGe": 0,
  "xiangMuMingCheng": 0,
  "yongLiang": 0
}
  '''),
                              ),
                            ),
                          if (!_model.switchValue!)
                            FutureBuilder<ApiCallResponse>(
                              future: XiangMuZhuCaiBiaoGroup
                                  .getXiangMuZhuCaiBiaoListCall
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
                                final columnGetXiangMuZhuCaiBiaoListResponse =
                                    snapshot.data!;
                                return Builder(
                                  builder: (context) {
                                    final list = XiangMuZhuCaiBiaoGroup
                                            .getXiangMuZhuCaiBiaoListCall
                                            .list(
                                              columnGetXiangMuZhuCaiBiaoListResponse
                                                  .jsonBody,
                                            )
                                            ?.toList() ??
                                        [];
                                    return Column(
                                      mainAxisSize: MainAxisSize.max,
                                      children: List.generate(list.length,
                                          (listIndex) {
                                        final listItem = list[listIndex];
                                        return wrapWithModel(
                                          model: _model.xiangMuZhuCaItemModels2
                                              .getModel(
                                            listIndex.toString(),
                                            listIndex,
                                          ),
                                          updateCallback: () => setState(() {}),
                                          child: XiangMuZhuCaItemWidget(
                                            key: Key(
                                              'Key0gx_${listIndex.toString()}',
                                            ),
                                            item: listItem,
                                          ),
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
