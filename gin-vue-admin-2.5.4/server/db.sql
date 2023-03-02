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