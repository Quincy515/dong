
-- model/system/sys_user.go 用户登录信息
-- 员工花名册
create table `dong`.yuan_gong_hua_ming_ce
(
    `xu_hao`                           varchar(255) NULL comment '序号',
    `bu_men`                           varchar(255) NULL comment '部门',
    `xing_ming`                        varchar(255) NULL comment '姓名',
    `xing_bie`                         varchar(255) NULL comment '性别',
    `nian_ling`                        varchar(255) NULL comment '年龄',
    `hu_kou_suo_zai_di`                varchar(255) NULL comment '户口所在地',
    `ru_zhi_shi_jian`                  varchar(255) NULL comment '入职时间',
    `lao_dong_he_tong_kai_shi_ri_qi`   varchar(255) NULL comment '劳动合同开始日期',
    `lao_dong_he_tong_jie_zhi_ri_qi`   varchar(255) NULL comment '劳动合同截止日期',
    `wen_hua_cheng_du`                 varchar(255) NULL comment '文化程度',
    `shen_fen_zheng_hao_ma`            varchar(255) NULL comment '身份证号码',
    `chu_sheng_ri_qi`                  varchar(255) NULL comment '出生日期',
    `he_tong_yue_ding_gong_zi`         varchar(255) NULL comment '合同约定工资',
    `cong_shi_gang_wei_huo_gong_zhong` varchar(255) NULL comment '从事岗位或工种',
    `lian_xi_dian_hua`                 varchar(255) NULL comment '联系电话',
    `xian_ju_zhu_di_zhi`               varchar(255) NULL comment '现居住地址',
    `li_zhi_ri_qi`                     varchar(255) NULL comment '离职日期',
    `ge_ren_she_bao_zhang_hao`         varchar(255) NULL comment '个人社保账号',
    `ge_ren_she_bao_mi_ma`             varchar(255) NULL comment '个人社保密码',
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci comment '员工花名册';
