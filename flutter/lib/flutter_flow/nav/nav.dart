// ignore_for_file: unused_import

import 'dart:async';

import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:page_transition/page_transition.dart';
import '../../cai_gou_ji_hua_biao/cai_gou_ji_hua_biao_widget.dart';
import '../../cai_gou_ji_hua_ming_xi_biao/cai_gou_ji_hua_ming_xi_biao_widget.dart';
import '../../chu_ru_ku_dan_ju_biao/chu_ru_ku_dan_ju_biao_widget.dart';
import '../../chu_ru_ku_ming_xi_biao/chu_ru_ku_ming_xi_biao_widget.dart';
import '../../fen_bao_shang_biao/fen_bao_shang_biao_widget.dart';
import '../../fen_bao_shang_yu_xiang_mu_guan_lian_biao/fen_bao_shang_yu_xiang_mu_guan_lian_biao_widget.dart';
import '../../gong_ying_shang_biao/gong_ying_shang_biao_widget.dart';
import '../../he_tong_wen_jian_biao/he_tong_wen_jian_biao_widget.dart';
import '../../she_bei_ming_xi_biao/she_bei_ming_xi_biao_widget.dart';
import '../../she_bei_zu_lin_biao/she_bei_zu_lin_biao_widget.dart';
import '../../shi_gong_jin_du_biao/shi_gong_jin_du_biao_widget.dart';
import '../../wen_jian_zi_liao_biao/wen_jian_zi_liao_biao_widget.dart';
import '../../wen_jian_zi_liao_jie_dian_biao/wen_jian_zi_liao_jie_dian_biao_widget.dart';
import '../../xiang_mu_shi_gong_biao/xiang_mu_shi_gong_biao_widget.dart';
import '../../xiang_mu_zhu_cai_biao/xiang_mu_zhu_cai_biao_widget.dart';
import '../flutter_flow_theme.dart';

import '../../index.dart';
import '../../main.dart';
import '../lat_lng.dart';
import '../place.dart';
import 'serialization_util.dart';

export 'package:go_router/go_router.dart';
export 'serialization_util.dart';

const kTransitionInfoKey = '__transition_info__';

class AppStateNotifier extends ChangeNotifier {
  bool showSplashImage = true;

  void stopShowingSplashImage() {
    showSplashImage = false;
    notifyListeners();
  }
}

GoRouter createRouter(AppStateNotifier appStateNotifier) => GoRouter(
      initialLocation: '/',
      debugLogDiagnostics: true,
      refreshListenable: appStateNotifier,
      errorBuilder: (context, _) => XiangMuDengJiBiaoWidget(),
      routes: [
        FFRoute(
          name: '_initialize',
          path: '/',
          builder: (context, _) => XiangMuDengJiBiaoWidget(),
          routes: [
            FFRoute(
              name: 'HomePage',
              path: 'homePage',
              builder: (context, params) => HomePageWidget(),
            ),
            FFRoute(
              name: 'LoginPage',
              path: 'login',
              builder: (context, params) => LoginPageWidget(),
            ),
            FFRoute(
              name: 'UserProfile',
              path: 'userProfile',
              builder: (context, params) => UserProfileWidget(),
            ),
            FFRoute(
              name: 'Document',
              path: 'document',
              builder: (context, params) => DocumentWidget(),
            ),
            FFRoute(
              name: 'xiangMuDengJiBiao',
              path: 'xiangMuDengJiBiao',
              builder: (context, params) => XiangMuDengJiBiaoWidget(),
            ),
            FFRoute(
              name: 'touBiaoShuJuTongJi',
              path: 'touBiaoShuJuTongJi',
              builder: (context, params) => TouBiaoShuJuTongJiWidget(),
            ),
            FFRoute(
              name: 'yuanGongHuaMingCe',
              path: 'yuanGongHuaMingCe',
              builder: (context, params) => YuanGongHuaMingCeWidget(),
            ),
            FFRoute(
              name: 'meiRiKaoQin',
              path: 'meiRiKaoQin',
              builder: (context, params) => MeiRiKaoQinWidget(),
            ),
            FFRoute(
              name: 'kaoQinTongJi',
              path: 'kaoQinTongJi',
              builder: (context, params) => KaoQinTongJiWidget(),
            ),
            FFRoute(
              name: 'qingJiaGuanLi',
              path: 'qingJiaGuanLi',
              builder: (context, params) => QingJiaGuanLiWidget(),
            ),
            FFRoute(
              name: 'huiYuanGuanLi',
              path: 'huiYuanGuanLi',
              builder: (context, params) => HuiYuanGuanLiWidget(),
            ),
            FFRoute(
              name: 'gongSiZhangHaoMiMa',
              path: 'gongSiZhangHaoMiMa',
              builder: (context, params) => GongSiZhangHaoMiMaWidget(),
            ),
            FFRoute(
              name: 'gongZhangDengJi',
              path: 'gongZhangDengJi',
              builder: (context, params) => GongZhangDengJiWidget(),
            ),
            FFRoute(
              name: 'zhengShuBiao',
              path: 'zhengShuBiao',
              builder: (context, params) => ZhengShuBiaoWidget(),
            ),
            FFRoute(
              name: 'zhengShuJieChuJiLuBiao',
              path: 'zhengShuJieChuJiLuBiao',
              builder: (context, params) => ZhengShuJieChuJiLWidget(),
            ),
            FFRoute(
              name: 'dictionary',
              path: 'dictionary',
              builder: (context, params) => DictionaryWidget(),
            ),
            FFRoute(
              name: 'dictionaryDetail',
              path: 'dictionaryDetail',
              builder: (context, params) => DictionaryDetailWidget(
                sysDictionaryID:
                    params.getParam('sysDictionaryID', ParamType.int),
              ),
            ),
            FFRoute(
              name: 'xiangMuBiao',
              path: 'xiangMuBiao',
              builder: (context, params) => XiangMuBiaoWidget(),
            ),
            FFRoute(
              name: 'xiangMuShiGongBiao',
              path: 'xiangMuShiGongBiao',
              builder: (context, params) => XiangMuShiGongBiaoWidget(),
            ),
            FFRoute(
              name: 'heTongWenJianBiao',
              path: 'heTongWenJianBiao',
              builder: (context, params) => HeTongWenJianBiaoWidget(),
            ),
            FFRoute(
              name: 'fenBaoShangBiao',
              path: 'fenBaoShangBiao',
              builder: (context, params) => FenBaoShangBiaoWidget(),
            ),
            FFRoute(
              name: 'fenBaoShangYuXiangMuGuanLianBiao',
              path: 'fenBaoShangYuXiangMuGuanLianBiao',
              builder: (context, params) =>
                  FenBaoShangYuXiangMuGuanLianBiaoWidget(),
            ),
            FFRoute(
              name: 'gongYingShangBiao',
              path: 'gongYingShangBiao',
              builder: (context, params) => GongYingShangBiaoWidget(),
            ),
            FFRoute(
              name: 'xiangMuZhuCaiBiao',
              path: 'xiangMuZhuCaiBiao',
              builder: (context, params) => XiangMuZhuCaiBiaoWidget(),
            ),
            FFRoute(
              name: 'caiGouJiHuaBiao',
              path: 'caiGouJiHuaBiao',
              builder: (context, params) => CaiGouJiHuaBiaoWidget(),
            ),
            FFRoute(
              name: 'caiGouJiHuaMingXiBiao',
              path: 'caiGouJiHuaMingXiBiao',
              builder: (context, params) => CaiGouJiHuaMingXiBiaoWidget(),
            ),
            FFRoute(
              name: 'chuRuKuMingXiBiao',
              path: 'chuRuKuMingXiBiao',
              builder: (context, params) => ChuRuKuMingXiBiaoWidget(),
            ),
            FFRoute(
              name: 'chuRuKuDanJuBiao',
              path: 'chuRuKuDanJuBiao',
              builder: (context, params) => ChuRuKuDanJuBiaoWidget(),
            ),
            FFRoute(
              name: 'sheBeiZuLinBiao',
              path: 'sheBeiZuLinBiao',
              builder: (context, params) => SheBeiZuLinBiaoWidget(),
            ),
            FFRoute(
              name: 'sheBeiMingXiBiao',
              path: 'sheBeiMingXiBiao',
              builder: (context, params) => SheBeiMingXiBiaoWidget(),
            ),
            FFRoute(
              name: 'shiGongJinDuBiao',
              path: 'shiGongJinDuBiao',
              builder: (context, params) => ShiGongJinDuBiaoWidget(),
            ),
            FFRoute(
              name: 'wenJianZiLiaoJieDianBiao',
              path: 'wenJianZiLiaoJieDianBiao',
              builder: (context, params) => WenJianZiLiaoJieDianBiaoWidget(),
            ),
            FFRoute(
              name: 'wenJianZiLiaoBiao',
              path: 'wenJianZiLiaoBiao',
              builder: (context, params) => WenJianZiLiaoBiaoWidget(),
            ),
          ].map((r) => r.toRoute(appStateNotifier)).toList(),
        ).toRoute(appStateNotifier),
      ],
      urlPathStrategy: UrlPathStrategy.path,
    );

extension NavParamExtensions on Map<String, String?> {
  Map<String, String> get withoutNulls => Map.fromEntries(
        entries
            .where((e) => e.value != null)
            .map((e) => MapEntry(e.key, e.value!)),
      );
}

extension _GoRouterStateExtensions on GoRouterState {
  Map<String, dynamic> get extraMap =>
      extra != null ? extra as Map<String, dynamic> : {};
  Map<String, dynamic> get allParams => <String, dynamic>{}
    ..addAll(params)
    ..addAll(queryParams)
    ..addAll(extraMap);
  TransitionInfo get transitionInfo => extraMap.containsKey(kTransitionInfoKey)
      ? extraMap[kTransitionInfoKey] as TransitionInfo
      : TransitionInfo.appDefault();
}

class FFParameters {
  FFParameters(this.state, [this.asyncParams = const {}]);

  final GoRouterState state;
  final Map<String, Future<dynamic> Function(String)> asyncParams;

  Map<String, dynamic> futureParamValues = {};

  // Parameters are empty if the params map is empty or if the only parameter
  // present is the special extra parameter reserved for the transition info.
  bool get isEmpty =>
      state.allParams.isEmpty ||
      (state.extraMap.length == 1 &&
          state.extraMap.containsKey(kTransitionInfoKey));
  bool isAsyncParam(MapEntry<String, dynamic> param) =>
      asyncParams.containsKey(param.key) && param.value is String;
  bool get hasFutures => state.allParams.entries.any(isAsyncParam);
  Future<bool> completeFutures() => Future.wait(
        state.allParams.entries.where(isAsyncParam).map(
          (param) async {
            final doc = await asyncParams[param.key]!(param.value)
                .onError((_, __) => null);
            if (doc != null) {
              futureParamValues[param.key] = doc;
              return true;
            }
            return false;
          },
        ),
      ).onError((_, __) => [false]).then((v) => v.every((e) => e));

  dynamic getParam<T>(
    String paramName,
    ParamType type, [
    bool isList = false,
  ]) {
    if (futureParamValues.containsKey(paramName)) {
      return futureParamValues[paramName];
    }
    if (!state.allParams.containsKey(paramName)) {
      return null;
    }
    final param = state.allParams[paramName];
    // Got parameter from `extras`, so just directly return it.
    if (param is! String) {
      return param;
    }
    // Return serialized value.
    return deserializeParam<T>(
      param,
      type,
      isList,
    );
  }
}

class FFRoute {
  const FFRoute({
    required this.name,
    required this.path,
    required this.builder,
    this.requireAuth = false,
    this.asyncParams = const {},
    this.routes = const [],
  });

  final String name;
  final String path;
  final bool requireAuth;
  final Map<String, Future<dynamic> Function(String)> asyncParams;
  final Widget Function(BuildContext, FFParameters) builder;
  final List<GoRoute> routes;

  GoRoute toRoute(AppStateNotifier appStateNotifier) => GoRoute(
        name: name,
        path: path,
        pageBuilder: (context, state) {
          final ffParams = FFParameters(state, asyncParams);
          final page = ffParams.hasFutures
              ? FutureBuilder(
                  future: ffParams.completeFutures(),
                  builder: (context, _) => builder(context, ffParams),
                )
              : builder(context, ffParams);
          final child = page;

          final transitionInfo = state.transitionInfo;
          return transitionInfo.hasTransition
              ? CustomTransitionPage(
                  key: state.pageKey,
                  child: child,
                  transitionDuration: transitionInfo.duration,
                  transitionsBuilder: PageTransition(
                    type: transitionInfo.transitionType,
                    duration: transitionInfo.duration,
                    reverseDuration: transitionInfo.duration,
                    alignment: transitionInfo.alignment,
                    child: child,
                  ).transitionsBuilder,
                )
              : MaterialPage(key: state.pageKey, child: child);
        },
        routes: routes,
      );
}

class TransitionInfo {
  const TransitionInfo({
    required this.hasTransition,
    this.transitionType = PageTransitionType.fade,
    this.duration = const Duration(milliseconds: 300),
    this.alignment,
  });

  final bool hasTransition;
  final PageTransitionType transitionType;
  final Duration duration;
  final Alignment? alignment;

  static TransitionInfo appDefault() => TransitionInfo(hasTransition: false);
}
