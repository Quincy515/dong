-- model/system/sys_user.go 用户登录信息
-- 经营部
-- 1. 项目登记表
create table `dong`.xiang_mu_deng_ji_biao (
  `group` varchar(100) NULL comment '所属',
  `xiang_mu_ming_cheng` varchar(100) NULL comment '项目名称',
  `xiang_mu_jian_jie` text NULL comment '项目简介',
  `xiang_mu_deng_ji_jin_du` varchar(10) NULL comment '项目登记进度',
  `tou_biao_dan_wei_ming_cheng` varchar(100) NULL comment '投标单位名称',
  `zi_zhi_yao_qiu` INT NULL comment '资质要求',
  `deng_ji_yao_qiu` INT NULL comment '等级要求',
  `kai_biao_shi_jian` datetime NULL comment '开标时间',
  `kai_biao_di_dian` varchar(200) NULL comment '开标地点',
  `suo_shu_di_qu` INT NULL comment '所属地区',
  `zhao_biao_gui_mo` varchar(20) NULL comment '招标规模',
  `kai_biao_ren_yuan` varchar(10) NULL comment '开标人员',
  `ye_wu_fu_ze_ren` varchar(10) NULL comment '业务负责人',
  `bao_zheng_jin_xing_shi` INT NULL comment '保证金形式',
  `bao_zheng_jin_jin_e` varchar(20) NULL comment '保证金金额',
  `ping_biao_ban_fa` INT NULL comment '评标办法',
  `xiang_mu_wang_zhi` varchar(200) NULL comment '项目网址',
  `biao_shu_lei_xing` INT NULL comment '标书类型',
  `xiang_mu_deng_ji_shi_jian` datetime NULL comment '项目登记时间',
  `bu_men_shen_he` varchar(10) NULL comment '部门审核',
  `bu_men_shen_he_shi_jian` datetime NULL comment '部门审核时间',
  `ce_oshen_he` varchar(10) NULL comment 'CEO审核',
  `ce_oshen_he_shi_jian` datetime NULL comment 'CEO审核时间',
  `tou_biao_han_zhi_zuo_ren` varchar(10) NULL comment '投标函制作人',
  `tou_biao_han_ti_cheng_jin_e` varchar(20) NULL comment '投标函提成金额',
  `tou_biao_han_zhi_zuo_ren_que_ren` varchar(10) NULL comment '投标函制作人确认',
  `tou_biao_han_zhi_zuo_ren_que_ren_shi_jian` datetime NULL comment '投标函制作人确认时间',
  `shang_wu_biao_zhi_zuo_ren` varchar(10) NULL comment '商务标制作人',
  `shang_wu_biao_ti_cheng_jin_e` varchar(20) NULL comment '商务标提成金额',
  `shang_wu_biao_zhi_zuo_ren_que_ren` varchar(10) NULL comment '商务标制作人确认',
  `shang_wu_biao_zhi_zuo_ren_que_ren_shi_jian` datetime NULL comment '商务标制作人确认时间',
  `ji_shu_biao_zhi_zuo_ren` varchar(10) NULL comment '技术标制作人',
  `ji_shu_biao_ti_cheng_jin_e` varchar(20) NULL comment '技术标提成金额',
  `ji_shu_biao_zhi_zuo_ren_que_ren` varchar(10) NULL comment '技术标制作人确认',
  `ji_shu_biao_zhi_zuo_ren_que_ren_shi_jian` datetime NULL comment '技术标制作人确认时间',
  `xiang_mu_zhuang_tai` varchar(5) NULL comment '项目状态',
  `bu_men_shen_qing_xia_fu_dian_wei` varchar(10) NULL comment '部门申请下浮点位',
  `bu_men_xia_fu_dian_wei_shen_qing_shi_jian` datetime NULL comment '部门下浮点位申请时间',
  `ce_oshen_he_xia_fu_dian_wei` varchar(10) NULL comment 'CEO审核下浮点位',
  `ce_oxia_fu_dian_wei_shen_he_shi_jian` datetime NULL comment 'CEO下浮点位审核时间',
  `zhong_biao_jin_e` varchar(20) NULL comment '中标金额',
  `bei_zhu` text NULL comment '备注'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '项目登记表';
-- 1.1 资质要求 市政3、建筑3、水利3
create table `dong`.zi_zhi_yao_qiu (
  `group` varchar(100) NULL comment '所属',
  `content` varchar(50) NULL comment '内容',
  `sort` int DEFAULT 0 comment '排序'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '资质要求';
-- 1.2 等级要求 三级、二级、一级
create table `dong`.deng_ji_yao_qiu (
  `group` varchar(100) NULL comment '所属',
  `content` varchar(50) NULL comment '内容',
  `sort` int DEFAULT 0 comment '排序'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '等级要求';
-- 1.3 所属地区
create table `dong`.suo_shu_di_qu (
  `group` varchar(100) NULL comment '所属',
  `content` varchar(50) NULL comment '内容',
  `sort` int DEFAULT 0 comment '排序'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '所属地区';
-- 1.4 保证金形式 转账、保函
create table `dong`.bao_zheng_jin_xing_shi (
  `group` varchar(100) NULL comment '所属',
  `content` varchar(50) NULL comment '内容',
  `sort` int DEFAULT 0 comment '排序'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '保证金形式';
-- 1.5 评标办法 低价法、综评法、抽取法
create table `dong`.ping_biao_ban_fa (
  `group` varchar(100) NULL comment '所属',
  `content` varchar(50) NULL comment '内容',
  `sort` int DEFAULT 0 comment '排序'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '评标办法';
-- 1.6 标书类型 纸质标、电子标
create table `dong`.biao_shu_lei_xing (
  `group` varchar(100) NULL comment '所属',
  `content` varchar(50) NULL comment '内容',
  `sort` int DEFAULT 0 comment '排序'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '标书类型';
-- 2. 投标数据统计
create table `dong`.tou_biao_shu_ju_tong_ji (
  `group` varchar(100) NULL comment '所属',
  `xiang_mu_ming_cheng` varchar(100) NULL comment '项目名称',
  `tou_biao_dan_wei_ming_cheng` varchar(100) NULL comment '投标单位名称',
  `zi_zhi_yao_qiu` INT NULL comment '资质要求',
  `deng_ji_yao_qiu` INT NULL comment '等级要求',
  `kai_biao_shi_jian` datetime NULL comment '开标时间',
  `suo_shu_di_qu` INT NULL comment '所属地区',
  `zhao_biao_gui_mo` varchar(20) NULL comment '招标规模',
  `di_yi_zhong_biao_hou_xuan_ren` varchar(100) NULL comment '第一中标候选人',
  `di_yi_bao_jia` varchar(20) NULL comment '第一报价',
  `di_yi_xia_fu_lu` varchar(10) NULL comment '第一下浮率',
  `di_er_zhong_biao_hou_xuan_ren` varchar(100) NULL comment '第二中标候选人',
  `di_er_bao_jia` varchar(20) NULL comment '第二报价',
  `di_er_xia_fu_lu` varchar(10) NULL comment '第二下浮率',
  `di_san_zhong_biao_hou_xuan_ren` varchar(100) NULL comment '第三中标候选人',
  `di_san_bao_jia` varchar(20) NULL comment '第三报价',
  `di_san_xia_fu_lu` varchar(10) NULL comment '第三下浮率',
  `bei_zhu` text NULL comment '备注'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '投标数据统计';
-- 人事部
-- 1. 员工花名册
create table `dong`.yuan_gong_hua_ming_ce (
  `group` varchar(100) NULL comment '所属',
  `bu_men` varchar(255) NULL comment '部门',
  `xing_ming` varchar(255) NULL comment '姓名',
  `xing_bie` varchar(255) NULL comment '性别',
  `nian_ling` varchar(255) NULL comment '年龄',
  `hu_kou_suo_zai_di` varchar(255) NULL comment '户口所在地',
  `ru_zhi_shi_jian` varchar(255) NULL comment '入职时间',
  `lao_dong_he_tong_kai_shi_ri_qi` varchar(255) NULL comment '劳动合同开始日期',
  `lao_dong_he_tong_jie_zhi_ri_qi` varchar(255) NULL comment '劳动合同截止日期',
  `wen_hua_cheng_du` varchar(255) NULL comment '文化程度',
  `shen_fen_zheng_hao_ma` varchar(255) NULL comment '身份证号码',
  `chu_sheng_ri_qi` varchar(255) NULL comment '出生日期',
  `he_tong_yue_ding_gong_zi` varchar(255) NULL comment '合同约定工资',
  `cong_shi_gang_wei_huo_gong_zhong` varchar(255) NULL comment '从事岗位或工种',
  `lian_xi_dian_hua` varchar(255) NULL comment '联系电话',
  `xian_ju_zhu_di_zhi` varchar(255) NULL comment '现居住地址',
  `li_zhi_ri_qi` varchar(255) NULL comment '离职日期',
  `ge_ren_she_bao_zhang_hao` varchar(255) NULL comment '个人社保账号',
  `ge_ren_she_bao_mi_ma` varchar(255) NULL comment '个人社保密码',
  `men_jin_bian_hao` varchar(5) NULL comment '门禁编号',
  `bei_zhu` text NULL comment '备注'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '员工花名册';
-- 2. 每日考勤
create table `dong`.mei_ri_kao_qin (
  `group` varchar(100) NULL comment '所属',
  `xing_ming` varchar(100) NULL comment '姓名',
  `chu_qin` boolean DEFAULT TRUE NOT NULL comment '出勤',
  `xiu_jia_qing_jia_chu_cha` boolean DEFAULT FALSE NOT NULL comment '休假\请假\出差',
  `xiang_qing` text NULL comment '详情',
  `bei_zhu` text NULL comment '备注',
  `ri_qi` datetime NULL comment '日期'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '每日考勤';
-- 3. 考勤统计
create table `dong`.kao_qin_tong_ji (
  `group` varchar(100) NULL comment '所属',
  `xing_ming` varchar(100) NULL comment '姓名',
  `chu_qin_tian_shu` int DEFAULT 0 NOT NULL comment '出勤天数',
  `qing_jia_tian_shu` int DEFAULT 0 NOT NULL comment '请假天数',
  `xiang_qing` varchar(100) NULL comment '详情',
  `bei_zhu` varchar(100) NULL comment '备注',
  `kai_shi_ri_qi` varchar(100) NULL comment '开始日期',
  `jie_shu_ri_qi` varchar(100) NULL comment '结束日期'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '考勤统计';
-- 4. 请假管理
create table `dong`.qing_jia_guan_li (
  `group` varchar(100) NULL comment '所属',
  `xing_ming` varchar(100) NULL comment '姓名',
  `lei_xing` INT NULL comment '请假类型',
  `kai_shi_ri_qi` datetime NULL comment '开始日期',
  `jie_shu_ri_qi` datetime NULL comment '结束日期',
  `tian_shu` int DEFAULT 0 NOT NULL comment '天数',
  `shi_you` text NULL comment '事由',
  `shen_he_zhuang_tai` varchar(10) NULL comment '审核状态',
  `shen_he_ren` varchar(100) NULL comment '审核人',
  `shen_he_xiang_qing` text NULL comment '审核详情',
  `shen_he_shi_jian` datetime NULL comment '审核时间',
  `jin_du` varchar(10) NULL comment '进度',
  `bei_zhu` text NULL comment '备注'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '请假管理';
-- 4.1 请假类型
create table `dong`.qing_jia_lei_xing (
  `group` varchar(100) NULL comment '所属',
  `content` varchar(50) NULL comment '内容',
  `sort` int DEFAULT 0 comment '排序'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '请假类型';
-- 5. 会员管理
create table `dong`.hui_yuan_guan_li (
  `group` varchar(100) NULL comment '所属',
  `hui_yuan_ming_cheng` varchar(200) NULL comment '会员名称',
  `ji_bie` varchar(50) NULL comment '级别',
  `hui_fei` varchar(50) NULL comment '会费',
  `kai_shi_you_xiao_qi` datetime NULL comment '开始有效期',
  `jie_shu_you_xiao_qi` datetime NULL comment '结束有效期',
  `ru_hui_lian_xi` varchar(20) NULL comment '入会联系',
  `wang_zhi` varchar(200) NULL comment '网址',
  `jiao_fei_shi_jian` datetime NULL comment '缴费时间',
  `bei_zhu` text NULL comment '备注'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '会员管理';
-- 6. 公司账号密码
create table `dong`.gong_si_zhang_hao_mi_ma (
  `group` varchar(100) NULL comment '所属',
  `xi_tong_ming_cheng` varchar(100) NULL comment '系统名称',
  `yong_hu_ming` varchar(20) NULL comment '用户名',
  `mi_ma` varchar(20) NULL comment '密码',
  `yong_tu` text NULL comment '用途',
  `lian_jie` varchar(200) NULL comment '链接',
  `kai_shi_you_xiao_qi` datetime NULL comment '开始有效期',
  `jie_shu_you_xiao_qi` datetime NULL comment '结束有效期',
  `nian_fei` varchar(50) NULL comment '年费',
  `bei_zhu` text NULL comment '备注'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '公司账号密码';
-- 7. 公章登记
create table `dong`.gong_zhang_deng_ji (
  `group` varchar(100) NULL comment '所属',
  `gong_zhang_ming_cheng` varchar(100) NULL comment '公章名称',
  `shi_xiang` varchar(100) NULL comment '事项',
  `shen_qing_ren` varchar(100) NULL comment '申请人',
  `shen_qing_shi_jian` DATETIME NULL comment '申请时间',
  `jing_shou_ren` varchar(100) NULL comment '经手人',
  `jing_shou_shi_jian` DATETIME NULL comment '经手时间',
  `shu_liang` INT DEFAULT 1 NOT NULL comment '数量',
  `xiang_qing` text NULL comment '详情',
  `bei_zhu` text NULL comment '备注'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '公章登记';
-- 8. 证书
create table `dong`.zheng_shu_biao (
  `group` varchar(100) NULL comment '所属',
  `zheng_shu_ming_cheng` varchar(100) NULL comment '证书名称',
  `zheng_shu_bian_hao` varchar(100) NULL comment '证书编号',
  `zheng_shu_lei_bie` INT NULL comment '证书类别',
  `suo_you_ren_xing_ming` varchar(100) NULL comment '所有人姓名',
  `shen_fen_zheng_hao` varchar(100) NULL comment '身份证号',
  `shou_ji_hao_ma` varchar(100) NULL comment '手机号码',
  `fa_zheng_bu_men` varchar(100) NULL comment '发证部门',
  `zheng_shu_suo_shu_cheng_shi` INT NULL comment '证书所属城市',
  `zheng_shu_zhuang_tai` INT NULL comment '证书状态',
  `zheng_shu_dao_qi_ri_qi` DATETIME NULL comment '证书到期日期',
  `zheng_shu_nian_shi_yong_fei` varchar(100) NULL comment '证书年使用费',
  `zhi_fu_ri_qi` DATETIME NULL comment '支付日期',
  `shi_fou_jie_chu` boolean NULL comment '是否借出',
  `tian_jia_ren` INT NULL comment '添加人',
  `bei_zhu` text NULL comment '备注'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '证书表';
-- 8.1 证书类别 外聘、内证
create table `dong`.zheng_shu_biao (
  `group` varchar(100) NULL comment '所属',
  `lei_bie_ming_cheng` varchar(100) NULL comment '类别名称',
  `zhuang_tai` INT NULL comment '状态'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '证书类别';
-- 9. 证书借出记录表
create table `dong`.zheng_shu_jie_chu_ji_lu_biao (
  `group` varchar(100) NULL comment '所属',
  `zheng_shu_ming_cheng` varchar(100) NULL comment '证书名称',
  `zheng_shu_bian_hao` varchar(100) NULL comment '证书编号',
  `zheng_shu_id` INT NULL comment '证书ID',
  `jie_chu_ren` INT NULL comment '借出人ID',
  `jie_chu_ri_qi` DATETIME NULL comment '借出日期',
  `jie_chu_shi_you` VARCHAR(100) NULL comment '借出事由',
  `yu_ji_gui_hai_ri_qi` DATETIME NULL comment '预计归还日期',
  `shi_ji_gui_hai_ri_qi` DATETIME NULL comment '实际归还日期',
  `tian_jia_ren` INT NULL comment '添加人',
  `jie_chu_bei_zhu` text NULL comment '借出备注',
  `gui_hai_bei_zhu` text NULL comment '归还备注'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '证书借出记录表';
-- 工程部
-- 1. 项目管理
create table `dong`.xiang_mu_biao (
  `group` varchar(100) NULL comment '所属',
  `xiang_mu_bian_hao` INT NULL COMMENT '项目编号',
  `ji_hua_kai_gong_ri_qi` DATETIME NULL COMMENT '计划开工日期',
  `ji_hua_jie_shu_ri_qi` DATETIME NULL COMMENT '计划结束日期',
  `gong_cheng_zao_jia` Decimal(10, 2) NULL COMMENT '工程造价',
  `nong_min_gong_gong_zi_bao_zhang_jin` Decimal(10, 2) NULL COMMENT '农民工工资保障金',
  `bao_zhang_jin_zhuang_tai` INT NULL COMMENT '保障金状态',
  `lu_yue_bao_zheng_jin` Decimal(10, 2) NULL COMMENT '履约保证金',
  `bao_zheng_jin_zhuang_tai` INT NULL COMMENT '保证金状态',
  `gong_zuo_bao_xian_jin_e` Decimal(10, 2) NULL COMMENT '工作保险金额',
  `fu_kuan_fang_shi` INT NULL COMMENT '付款方式',
  `yu_qi_li_run` Decimal(10, 2) NULL COMMENT '预期利润',
  `shi_ji_kai_gong_ri_qi` DATETIME NULL COMMENT '实际开工日期',
  `shi_ji_jie_shu_ri_qi` DATETIME NULL COMMENT '实际结束日期',
  `jie_suan_zhuang_tai` INT NULL COMMENT '结算状态 1.未结算 2.结算完成',
  `xiang_mu_zhuang_tai` INT NULL COMMENT '项目状态 1.建设中 2.未完成 3.竣工',
  `xiang_mu_jing_li` INT NULL COMMENT '项目经理',
  `xiang_mu_zhu_guan` INT NULL COMMENT '项目主管'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '项目表';
-- 2. 项目施工表
create table `dong`.xiang_mu_shi_gong_biao (
  `group` varchar(100) NULL comment '所属',
  `xiang_mu_id` INT NULL comment '项目ID',
  `gong_zuo_ming_cheng` VARCHAR(100) NULL comment '工作名称',
  `ji_hua_kai_shi_ri_qi` DATETIME NULL comment '计划开始日期',
  `ji_hua_wan_cheng_ri_qi` DATETIME NULL comment '计划完成日期',
  `ji_hua_tian_shu` INT NULL comment '计划天数',
  `zhan_zong_gong_zuo_liang_bai_fen_bi` FLOAT NULL comment '占总工作量百分比',
  `wan_cheng_bi` FLOAT NULL comment '完成比',
  `bei_zhu` TEXT NULL comment '备注',
  `shun_xu` INT NULL comment '顺序',
  `zhuang_tai` INT NULL comment '状态',
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '项目施工表';
-- 3. 合同文件表
create table `dong`.he_tong_wen_jian_biao (
  `group` varchar(100) NULL comment '所属',
  `xiang_mu_ming_cheng` INT NULL comment '项目名称',
  `he_tong_bian_hao` varchar(100) NULL comment '合同编号',
  `he_tong_ming_cheng` varchar(100) NULL comment '合同名称',
  `qian_ding_ri_qi` DATETIME NULL comment '签订日期',
  `qian_ding_ren` INT NULL comment '签订人',
  `he_tong_lei_bie` INT NULL comment '合同类别 1.主合同 2.补充合同',
  `he_tong_lei_xing` INT NULL comment '合同类型 1.承包 2.分包 3.采购 4.租聘',
  `he_tong_jin_e` Decimal(10, 2) NULL comment '合同金额',
  `jia_fang_gong_si` INT NULL comment '甲方公司',
  `yi_fang_gong_si` INT NULL comment '乙方公司',
  `fu_kuan_fang_shi` INT NULL comment '付款方式',
  `jie_suan_fang_shi` INT NULL comment '结算方式',
  `shou_fu_kuan_tiao_jian` varchar(100) NULL comment '收付款条件',
  `zhu_yao_tiao_kuan` TEXT NULL comment '主要条款',
  `bei_zhu` TEXT NULL comment '备注',
  `he_tong_xia_zai_di_zhi` TEXT NULL comment '合同下载地址',
  `tian_jia_ren` INT NULL comment '添加人',
  `tian_jia_shi_jian` DATETIME NULL comment '添加时间',
  `he_tong_zhuang_tai` INT NULL comment '合同状态',
  `shou_kuan_zhang_hao` varchar(100) NULL comment '收款账号',
  `yin_hang_ming_cheng` varchar(100) NULL comment '银行名称',
  `kai_hu_hang` varchar(100) NULL comment '开户行',
  `kai_piao_dan_wei` varchar(100) NULL comment '开票单位',
  `fa_piao_lei_xing` INT NULL comment '发票类型'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '合同文件表';
-- 4. 分包商表
create table `dong`.fen_bao_shang_biao (
  `group` varchar(100) NULL comment '所属',
  `fen_bao_shang_ming_cheng` VARCHAR(100) NULL comment '分包商名称',
  `fen_bao_shang_lei_xing` INT NULL comment '分包商类型',
  `fa_ren` VARCHAR(100) NULL comment '法人',
  `she_hui_xin_yong_dai_ma` VARCHAR(100) NULL comment '社会信用代码',
  `zi_zhi_zheng_zhao` VARCHAR(100) NULL comment '资质证照',
  `fa_ren_shou_ji` VARCHAR(100) NULL comment '法人手机',
  `yin_hang_ming_cheng` VARCHAR(100) NULL comment '银行名称',
  `kai_hu_hang` VARCHAR(100) NULL comment '开户行',
  `zhang_hao` VARCHAR(100) NULL comment '账号',
  `tian_jia_shi_jian` DATETIME NULL comment '添加时间',
  `tian_jia_ren` INT NULL comment '添加人',
  `zhuang_tai` INT NULL comment '状态'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '分包商表';
-- 5. 分包商与项目关联表
create table `dong`.fen_bao_shang_biao (
  `group` varchar(100) NULL comment '所属',
  `xiang_mu_id` VARCHAR(100) NULL comment '项目ID',
  `fen_bao_shang_id` VARCHAR(100) NULL comment '分包商ID',
  `he_tong_bian_hao` VARCHAR(100) NULL comment '合同编号',
  `he_tong_ming_cheng` VARCHAR(100) NULL comment '合同名称',
  `jin_chang_ri_qi` VARCHAR(100) NULL comment '进场日期',
  `tui_chang_ri_qi` VARCHAR(100) NULL comment '退场日期',
  `wei_tuo_ren` VARCHAR(100) NULL comment '委托人',
  `wei_tuo_ren_shou_ji` VARCHAR(100) NULL comment '委托人手机',
  `zhuang_tai` VARCHAR(100) NULL comment '状态'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '分包商与项目关联表';
-- 6. 工程班组表
create table `dong`.gong_cheng_ban_zu_biao (
  `group` varchar(100) NULL comment '所属',
  `xiang_mu_id` INT NULL comment '项目ID',
  `fen_bao_shang_id` INT NULL comment '分包商ID',
  `he_tong_bian_hao` VARCHAR(100) NULL comment '合同编号',
  `ban_zu_ming_cheng` VARCHAR(100) NULL comment '班组名称',
  `ban_zu_zhang_xing_ming` VARCHAR(100) NULL comment '班组长姓名',
  `shen_fen_zheng_hao` VARCHAR(100) NULL comment '身份证号',
  `shou_ji_hao_ma` VARCHAR(100) NULL comment '手机号码',
  `ban_zu_gong_zhong` INT NULL comment '班组工种',
  `bei_zhu_shuo_ming` TEXT NULL comment '备注说明',
  `jin_chang_ri_qi` DATETIME NULL comment '进场日期',
  `tui_chang_ri_qi` DATETIME NULL comment '退场日期',
  `zhuang_tai` INT NULL comment '状态',
  `tui_chang_ping_fen` INT NULL comment '退场评分',
  `tui_chang_bei_zhu` TEXT NULL comment '退场备注'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '工程班组表';
-- 7. 工程班组人员表
create table `dong`.gong_cheng_ban_zu_ren_yuan_biao (
  `group` varchar(100) NULL comment '所属',
  `ban_zu_id` INT NULL comment '班组ID',
  `xiang_mu_id` INT NULL comment '项目ID',
  `fen_bao_shang_id` INT NULL comment '分包商ID',
  `ren_yuan_xing_ming` VARCHAR(100) NULL comment '人员姓名',
  `shen_fen_zheng_hao` VARCHAR(100) NULL comment '身份证号',
  `gong_zhong` INT NULL comment '工种',
  `zi_ge_zheng_shu` VARCHAR(100) NULL comment '资格证书',
  `jin_chang_ri_qi` DATETIME NULL comment '进场日期',
  `tui_chang_ri_qi` DATETIME NULL comment '退场日期',
  `shi_fou_ban_zu_zhang` INT NULL comment '是否班组长',
  `shou_ji_hao_ma` VARCHAR(100) NULL comment '手机号码',
  `he_tong_bian_hao` VARCHAR(100) NULL comment '合同编号',
  `zhuang_tai` INT NULL comment '状态'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '工程班组人员表';
-- 8. 人员考勤汇总表
create table `dong`.ren_yuan_kao_qin_hui_zong_biao (
  `group` varchar(100) NULL comment '所属',
  `xiang_mu_id` INT NULL comment '项目ID',
  `fen_bao_shang_id` INT NULL comment '分包商ID',
  `ban_zu_id` INT NULL comment '班组ID',
  `nian_du_yue_fen` DATETIME NULL comment '年度月份',
  `ying_dao_ren_shu` INT NULL comment '应到人数',
  `shi_dao_ren_shu` INT NULL comment '实到人数',
  `zong_gong_shi` FLOAT NULL comment '总工时',
  `zong_gong_zi` Decimal(10, 2) NULL comment '总工资',
  `kao_qin_ren` VARCHAR(100) NULL comment '考勤人',
  `zhuang_tai` INT NULL comment '状态'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '人员考勤汇总表';
-- 9. 人员考勤表
create table `dong`.ren_yuan_kao_qin_hui_zong_biao (
  `group` varchar(100) NULL comment '所属',
  `ren_yuan_kao_qin_biao` INT NULL comment '人员考勤表',
  `xiang_mu_id` INT NULL comment '项目ID',
  `fen_bao_shang_id` INT NULL comment '分包商ID',
  `ban_zu_id` INT NULL comment '班组ID',
  `ren_yuan_id` INT NULL comment '人员ID',
  `nian_du_yue_fen` DATETIME NULL comment '年度月份',
  `zai_gang_tian_shu` FLOAT NULL comment '在岗天数',
  `dang_yue_gong_zi` Decimal(10, 2) NULL comment '当月工资',
  `yi_fa_gong_zi` Decimal(10, 2) NULL comment '已发工资',
  `wei_fa_gong_ci` Decimal(10, 2) NULL comment '未发工次',
  `zhuang_tai` INT NULL comment '状态'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '人员考勤汇总表';
-- 10. 安全管理表
create table `dong`.ren_yuan_kao_qin_hui_zong_biao (
  `group` varchar(100) NULL comment '所属',
  `xiang_mu_ming_cheng` INT NULL comment '项目名称',
  `jin_du_jie_dian_id` INT NULL comment '进度节点ID',
  `an_quan_biao_ti` VARCHAR(100) NULL comment '安全标题',
  `an_quan_lei_xing` INT NULL comment '安全类型',
  `wen_ti_miao_shu` TEXT NULL comment '问题描述',
  `tu_pian_zi_liao` VARCHAR(100) NULL comment '图片资料',
  `shi_pin_zi_liao` VARCHAR(100) NULL comment '视频资料',
  `zheng_gai_ren` INT NULL comment '整改人',
  `shang_bao_ren` INT NULL comment '上报人',
  `shang_bao_ri_qi` DATETIME NULL comment '上报日期',
  `zhuang_tai` INT NULL comment '状态'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '安全管理表';
-- 11. 安全管理表
create table `dong`.an_quan_zheng_gai_biao (
  `group` varchar(100) NULL comment '所属',
  `an_quan_bian_hao` INT NULL comment '安全编号',
  `zheng_gai_miao_shu` TEXT NULL comment '整改描述',
  `tu_pian_zi_liao` VARCHAR(100) NULL comment '图片资料',
  `shi_pin_zi_liao` VARCHAR(100) NULL comment '视频资料',
  `fu_cha_zhuang_tai` INT NULL comment '复查状态'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '安全整改表';
-- 12. 质量管理表
create table `dong`.zhi_liang_guan_li_biao (
  `group` varchar(100) NULL comment '所属',
  `xiang_mu_ming_cheng` INT NULL comment '项目名称',
  `jin_du_jie_dian` INT NULL comment '进度节点',
  `zhi_liang_biao_ti` VARCHAR(100) NULL comment '质量标题',
  `zhi_liang_lei_xing` INT NULL comment '质量类型',
  `wen_ti_miao_shu` TEXT NULL comment '问题描述',
  `tu_pian_zi_liao` VARCHAR(100) NULL comment '图片资料',
  `shi_pin_zi_liao` VARCHAR(100) NULL comment '视频资料',
  `zheng_gai_ren` INT NULL comment '整改人',
  `shang_bao_ren` INT NULL comment '上报人',
  `shang_bao_ri_qi` DATETIME NULL comment '上报日期',
  `zhuang_tai` INT NULL comment '状态'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '质量管理表';
-- 13. 安全整改表
create table `dong`.an_quan_zheng_gai_biao (
  `group` varchar(100) NULL comment '所属',
  `an_quan_bian_hao` INT NULL comment '安全编号',
  `zheng_gai_miao_shu` TEXT NULL comment '整改描述',
  `tu_pian_zi_liao` VARCHAR(100) NULL comment '图片资料',
  `shi_pin_zi_liao` VARCHAR(100) NULL comment '视频资料',
  `fu_cha_zhuang_tai` INT NULL comment '复查状态'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '安全整改表';
-- 14. 供应商表
create table `dong`.gong_ying_shang_biao (
  `group` varchar(100) NULL comment '所属',
  `gong_ying_shang_ming_cheng` INT NULL comment '供应商名称',
  `fa_ren_shen_fen_zheng` VARCHAR(100) NULL comment '法人身份证',
  `she_hui_xin_yong_dai_ma` VARCHAR(100) NULL comment '社会信用代码',
  `zi_zhi_zheng_zhao` VARCHAR(100) NULL comment '资质证照',
  `gong_si_di_zhi` VARCHAR(100) NULL comment '公司地址',
  `lian_xi_ren` VARCHAR(100) NULL comment '联系人',
  `lian_xi_ren_shou_ji` VARCHAR(100) NULL comment '联系人手机',
  `fa_piao_tai_tou` VARCHAR(100) NULL comment '发票抬头',
  `fa_piao_lei_xing` INT NULL comment '发票类型',
  `zhang_hao` VARCHAR(100) NULL comment '账号',
  `yin_hang_ming_cheng` VARCHAR(100) NULL comment '银行名称',
  `kai_hu_hang` VARCHAR(100) NULL comment '开户行',
  `tian_jia_ren` INT NULL comment '添加人',
  `tian_jia_ri_qi` DATETIME NULL comment '添加日期',
  `zhuang_tai` INT NULL comment '状态'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '供应商表';
-- 15. 项目主材表
create table `dong`.xiang_mu_ming_cheng (
  `group` varchar(100) NULL comment '所属',
  `xiang_mu_ming_cheng` INT NULL comment '项目名称',
  `cai_liao_ming_cheng` Varchar(50) NULL comment '材料名称',
  `cai_liao_bian_hao` Varchar(50) NULL comment '材料编号',
  `dan_wei` Varchar(10) NULL comment '单位',
  `yong_liang` Float NULL comment '用量',
  `dan_jia` Float NULL comment '单价',
  `xian_hang_jia_ge` Decimal(10, 2) NULL comment '现行价格',
  `jia_cha_he_ji` Decimal(10, 2) NULL comment '价差合计',
  `he_ji` Decimal(10, 2) NULL comment '合计'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '项目主材表';
-- 16. 采购计划表
create table `dong`.cai_gou_ji_hua_biao (
  `group` varchar(100) NULL comment '所属',
  `xiang_mu_id` INT NULL comment '项目id',
  `cai_gou_ji_hua_ming_cheng` Varchar(100) NULL comment '采购计划名称',
  `cai_gou_fu_ze_ren` INT NULL comment '采购负责人',
  `tian_jia_ri_qi` Datetime NULL comment '添加日期',
  `zhuang_tai` INT NULL comment '状态',
  `lei_ji_cai_gou_jin_e` Decimal(10, 2) NULL comment '累计采购金额',
  `shi_ji_zhi_fu_jin_e` Decimal(10, 2) NULL comment '实际支付金额'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '采购计划表';
-- 17. 采购计划明细表
create table `dong`.cai_gou_ji_hua_ming_xi_biao (
  `group` varchar(100) NULL comment '所属',
  `cai_liao_id` int NULL comment '材料ID',
  `cai_gou_ji hua_id` int NULL comment '采购计划ID',
  `gong_ying_shang_id` int NULL comment '供应商ID',
  `dan_jia` Varchar(10) NULL comment '单价',
  `cai_gou_liang` Float NULL comment '采购量',
  `zong_jia` Decimal(10, 2) NULL comment '总价',
  `ru_ku_liang` Float NULL comment '入库量',
  `zhuang_tai` int NULL comment '状态'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '采购计划明细表';
-- 18. 出入库单据表
create table `dong`.chu_ru_ku_dan_ju_biao (
  `group` varchar(100) NULL comment '所属',
  `xiang_mu_id` INT NULL ccomment '项目ID',
  `cai_gou_ji_hua_id` INT NULL ccomment '采购计划ID',
  `lei_xing` INT NULL ccomment '类型 1. 入库 2. 出库',
  `tian_jia_ren` INT NULL ccomment '添加人',
  `tian_jia_ri_qi` Datetime NULL ccomment '添加日期',
  `zhuang_tai` INT NULL ccomment '状态'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '出入库单据表';
-- 19. 出入库明细表
create table `dong`.chu_ru_ku_ming_xi_biao (
  `group` varchar(100) NULL comment '所属',
  `cai_liao_id` int NULL comment '材料ID',
  `chu_ru_ku_dan_ju_id` int NULL comment '出入库单据ID',
  `dan_jia` Float NULL comment '单价',
  `cai_gou_liang` Float NULL comment '采购量',
  `zong_jia` Decimal(10, 2) NULL comment '总价',
  `ru_ku_liang` Float NULL comment '入库量',
  `shi_hou_ku_cun_liang` Float NULL comment '事后库存量',
  `zhuang_tai` int NULL comment '状态'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '出入库明细表';
-- 20. 设备租赁表
create table `dong`.she_bei_zu_lin_biao (
  `group` varchar(100) NULL comment '所属',
  `xiang_mu_id` Int NULL comment '项目ID',
  `zu_lin_ming_cheng` Int NULL comment '租赁名称',
  `zu_lin_yong_tu` Varchar(500) NULL comment '租赁用途',
  `tian_jia_ren` Int NULL comment '添加人',
  `tian_jia_ri_qi` Datetime NULL comment '添加日期',
  `zhuang_tai` Int NULL comment '状态'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '设备租赁表';
-- 21. 设备明细表
create table `dong`.she_bei_ming_xi_biao (
  `group` varchar(100) NULL comment '所属',
  `she_bei_ming_cheng` int NULL comment '设备名称',
  `zu_lin_dan_ju_id` int NULL comment '租赁单据ID',
  `dan_jia` Float NULL comment '单价',
  `zu_lin_shu_liang` Float NULL comment '租赁数量',
  `ji_hua_zu_lin_ri_qi` Decimal(10, 2) NULL comment '计划租赁日期',
  `ji_hua_tui_zu_ri_qi` Float NULL comment '计划退租日期',
  `shi_ji_zu_lin_ri_qi` Float NULL comment '实际租赁日期',
  `shi_ji_tui_zu_ri_qi` Datetime NULL comment '实际退租日期',
  `zu_lin_tian_shu` Datetime NULL comment '租赁天数',
  `zhuang_tai` int NULL comment '状态'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '设备明细表';
-- 22. 施工进度表
create table `dong`.shi_gong_jin_du_biao (
  `group` varchar(100) NULL comment '所属',
  `xiang_mu_id` Int NULL comment '项目ID',
  `gong_zuo_ming_cheng_id` int NULL comment '工作名称ID',
  `wan_cheng_bai_fen_bi` Float NULL comment '完成百分比',
  `shi_ji_kai_shi_ri_qi` Datetime NULL comment '实际开始日期',
  `shi_ji_wan_cheng_ri_qi` Datetime NULL comment '实际完成日期',
  `shi_ji_tian_shu` Int NULL comment '实际天数',
  `bei_zhu` Varchar(200) NULL comment '备注',
  `zhuang_tai` Int NULL comment '状态',
  `tian_jia_ren` Int NULL comment '添加人',
  `tian_jia_shi_jian` Datetime NULL comment '添加时间'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '施工进度表';
-- 23. 文件资料节点表
create table `dong`.wen_jian_zi_liao_jie_dian_biao (
  `group` varchar(100) NULL comment '所属',
  `xiang_mu_ming_cheng` Int NULL comment '项目名称',
  `wen_jian_jia_ming_cheng` Varchar(200) NULL comment '文件夹名称',
  `fu_jie_dian` Int NULL comment '父节点多层级迭代寻祖',
  `tian_jia_shi_jian` Datetime NULL comment '添加时间',
  `tian_jia_ren` Int NULL comment '添加人',
  `shi_fou_cun_dang` Int NULL comment '是否存档-存档后不可以变更文件',
  `pai_xu` Int NULL comment '排序',
  `zhuang_tai` Int NULL comment '状态',
  `bei_zhu` TEXT NULL comment '备注'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '文件资料节点表';
-- 24. 文件资料表
create table `dong`.wen_jian_zi_liao_biao (
  `group` varchar(100) NULL comment '所属',
  `wen_jian_ming_cheng` Int NULL comment '文件名称',
  `jie_dian_i` Int NULL comment '节点ID只保存最近的节点',
  `shang_chuan_ren` Int NULL comment '上传人',
  `shang_chuan_ri_qi` Datetime NULL comment '上传日期',
  `shi_fou_shen_he` Int NULL comment '是否审核',
  `shen_he_ren` Int NULL comment '审核人',
  `shen_he_ri_qi` Datetime NULL comment '审核日期',
  `wen_jian_lei_xing` Int NULL comment '文件类型',
  `yue_du_ci_shu` Int NULL comment '阅读次数',
  `zhuang_tai` Int NULL comment '状态',
  `bei_zhu` TEXT NULL comment '备注'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '文件资料表';
-- 财务部
-- 1. 内部报销表
create table `dong`.nei_bu_bao_xiao_biao (
  `group` varchar(100) NULL comment '所属',
  `xiang_mu_ming_cheng` Int NULL comment '项目名称',
  `shou_ci_bao_xiao_id` Int NULL comment '首次报销ID-再次申请时关联的ID',
  `bao_xiao_ren_yuan` Int NULL comment '报销人员',
  `bao_xiao_shi_xiang` Int NULL comment '报销事项',
  `bao_xiao_jin_e` Decimal(10, 2) NULL comment '报销金额',
  `piao_ju_di_zhi` Varchar(500) NULL comment '票据地址-可多图片上传',
  `fei_yong_fa_sheng_ri_qi` Datetime NULL comment '费用发生日期',
  `gong_si_ming_cheng` Int NULL comment '公司名称',
  `bao_xiao_bu_men` Int NULL comment '报销部门',
  `lei_xing` Int NULL comment '类型 1. 内部报销 2. 项目经理报销',
  `bei_zhu_shuo_ming` Varchar(200) NULL comment '备注说明',
  `shen_qing_ri_qi` Datetime NULL comment '申请日期',
  `zhuang_tai` Int NULL comment '状态',
  `hui_dan_di_zhi` Varchar(500) NULL comment '回单地址-可多图片上传'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '内部报销表';
-- 2. 外部费用请款表（应付款表）
create table `dong`.wai_bu_fei_yong_qing_kuan_biao (
  `group` varchar(100) NULL comment '所属',
  `bao_xiao_ming_cheng` Varchar(200) NULL comment '报销名称',
  `shou_ci_bao_xiao_id` Int NULL comment '首次报销ID',
  `xiang_mu_ming_cheng` Int NULL comment '项目名称',
  `bao_xiao_ren_yuan` Int NULL comment '报销人员',
  `bao_xiao_shi_xiang` Int NULL comment '报销事项',
  `bao_xiao_jin_e` Decimal(10, 2) NULL comment '报销金额',
  `shi_fu_jin_e` Decimal(10, 2) NULL comment '实付金额',
  `bao_xiao_lei_xing` Int NULL comment '报销类型',
  `shou_kuan_zhang_hao` Varchar(25) NULL comment '收款账号',
  `yin_hang_ming_cheng` Varchar(20) NULL comment '银行名称',
  `kai_hu_hang` Varchar(100) NULL comment '开户行',
  `he_tong_id` int NULL comment '合同ID',
  `fei_yong_fa_sheng_ri_qi` Datetime NULL comment '费用发生日期',
  `zui_wan_fu_kuan_ri_qi` Datetime NULL comment '最晚付款日期',
  `bei_zhu_shuo_ming` Varchar(200) NULL comment '备注说明',
  `shen_qing_ri_qi` Datetime NULL comment '申请日期',
  `zhuang_tai` Int NULL comment '状态',
  `hui_dan_di_zhi` Varchar(500) comment '回单地址-可多图片上传NULL'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '外部费用请款表（应付款表）';
-- 3. 应收款表
create table `dong`.ying_shou_kuan_biao (
  `group` varchar(100) NULL comment '所属',
  `shou_kuan_ming_cheng` Varchar(200) NULL comment '收款名称',
  `shou_ci_shou_kuan_id` Int NULL comment '首次收款ID',
  `xiang_mu_ming_cheng` Int NULL comment '项目名称',
  `fa_qi_ren_yuan` Int NULL comment '发起人员',
  `ying_shou_shi_xiang` Int NULL comment '应收事项',
  `kai_piao_jin_e` Decimal(10, 2) NULL comment '开票金额',
  `ying_shou_jin_e` Decimal(10, 2) NULL comment '应收金额',
  `shi_shou_jin_e` Decimal(10, 2) NULL comment '实收金额',
  `shi_fu_jin_e` Decimal(10, 2) NULL comment '实付金额',
  `ying_shou_lei_xing` Int 工程款 NULL comment '应收类型',
  `dian_zi_kuan` Decimal(10, 2) NULL comment '垫资款',
  `tou_biao_bao_zheng_jin` Decimal(10, 2) NULL comment '投标保证金',
  `lu_yue_bao_zheng_jin` Decimal(10, 2) NULL comment '履约保证金',
  `zhi_bao_jin` Decimal(10, 2) NULL comment '质保金',
  `shou_kuan_zhang_hao` Varchar(25) NULL comment '收款账号',
  `yin_hang_ming_cheng` Varchar(20) NULL comment '银行名称',
  `kai_hu_hang` Varchar(100) NULL comment '开户行',
  `he_tong_id` Int NULL comment '合同ID',
  `fei_yong_fa_sheng_ri_qi` Datetime NULL comment '费用发生日期',
  `yu_ji_dao_zhang_ri_qi` Datetime NULL comment '预计到账日期',
  `kai_piao_dan_wei` Varchar(50) NULL comment '开票单位',
  `fa_piao_lei_xing` Int NULL comment '发票类型',
  `bei_zhu_shuo_ming` Varchar(500) NULL comment '备注说明',
  `shen_qing_ri_qi` Datetime NULL comment '申请日期',
  `zhuang_tai` Int NULL comment '状态',
  `hui_dan_di_zhi` Varchar(500) 可多图片上传 NULL comment '回单地址'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci comment '应收款表';