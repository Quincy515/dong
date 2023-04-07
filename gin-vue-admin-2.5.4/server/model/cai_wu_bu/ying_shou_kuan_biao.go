// 自动生成模板YingShouKuanBiao
package cai_wu_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// YingShouKuanBiao 结构体
type YingShouKuanBiao struct {
      global.GVA_MODEL
      BeiZhuShuoMing  string `json:"beiZhuShuoMing" form:"beiZhuShuoMing" gorm:"column:bei_zhu_shuo_ming;comment:备注说明;size:500;"`
      DianZiKuan  *float64 `json:"dianZiKuan" form:"dianZiKuan" gorm:"column:dian_zi_kuan;comment:垫资款;size:10;"`
      FaPiaoLeiXing  *int `json:"faPiaoLeiXing" form:"faPiaoLeiXing" gorm:"column:fa_piao_lei_xing;comment:发票类型;size:10;"`
      FaQiRenYuan  *int `json:"faQiRenYuan" form:"faQiRenYuan" gorm:"column:fa_qi_ren_yuan;comment:发起人员;size:10;"`
      FeiYongFaShengRiQi  *time.Time `json:"feiYongFaShengRiQi" form:"feiYongFaShengRiQi" gorm:"column:fei_yong_fa_sheng_ri_qi;comment:费用发生日期;"`
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      HeTongId  *int `json:"heTongId" form:"heTongId" gorm:"column:he_tong_id;comment:合同ID;size:10;"`
      HuiDanDiZhi  string `json:"huiDanDiZhi" form:"huiDanDiZhi" gorm:"column:hui_dan_di_zhi;comment:回单地址;size:500;"`
      KaiHuHang  string `json:"kaiHuHang" form:"kaiHuHang" gorm:"column:kai_hu_hang;comment:开户行;size:100;"`
      KaiPiaoDanWei  string `json:"kaiPiaoDanWei" form:"kaiPiaoDanWei" gorm:"column:kai_piao_dan_wei;comment:开票单位;size:50;"`
      KaiPiaoJinE  *float64 `json:"kaiPiaoJinE" form:"kaiPiaoJinE" gorm:"column:kai_piao_jin_e;comment:开票金额;size:10;"`
      LuYueBaoZhengJin  *float64 `json:"luYueBaoZhengJin" form:"luYueBaoZhengJin" gorm:"column:lu_yue_bao_zheng_jin;comment:履约保证金;size:10;"`
      ShenQingRiQi  *time.Time `json:"shenQingRiQi" form:"shenQingRiQi" gorm:"column:shen_qing_ri_qi;comment:申请日期;"`
      ShiFuJinE  *float64 `json:"shiFuJinE" form:"shiFuJinE" gorm:"column:shi_fu_jin_e;comment:实付金额;size:10;"`
      ShiShouJinE  *float64 `json:"shiShouJinE" form:"shiShouJinE" gorm:"column:shi_shou_jin_e;comment:实收金额;size:10;"`
      ShouCiShouKuanId  *int `json:"shouCiShouKuanId" form:"shouCiShouKuanId" gorm:"column:shou_ci_shou_kuan_id;comment:首次收款ID;size:10;"`
      ShouKuanMingCheng  string `json:"shouKuanMingCheng" form:"shouKuanMingCheng" gorm:"column:shou_kuan_ming_cheng;comment:收款名称;size:200;"`
      ShouKuanZhangHao  string `json:"shouKuanZhangHao" form:"shouKuanZhangHao" gorm:"column:shou_kuan_zhang_hao;comment:收款账号;size:25;"`
      TouBiaoBaoZhengJin  *float64 `json:"touBiaoBaoZhengJin" form:"touBiaoBaoZhengJin" gorm:"column:tou_biao_bao_zheng_jin;comment:投标保证金;size:10;"`
      XiangMuMingCheng  *int `json:"xiangMuMingCheng" form:"xiangMuMingCheng" gorm:"column:xiang_mu_ming_cheng;comment:项目名称;size:10;"`
      YinHangMingCheng  string `json:"yinHangMingCheng" form:"yinHangMingCheng" gorm:"column:yin_hang_ming_cheng;comment:银行名称;size:20;"`
      YingShouJinE  *float64 `json:"yingShouJinE" form:"yingShouJinE" gorm:"column:ying_shou_jin_e;comment:应收金额;size:10;"`
      YingShouLeiXing  *int `json:"yingShouLeiXing" form:"yingShouLeiXing" gorm:"column:ying_shou_lei_xing;comment:应收类型;size:10;"`
      YingShouShiXiang  *int `json:"yingShouShiXiang" form:"yingShouShiXiang" gorm:"column:ying_shou_shi_xiang;comment:应收事项;size:10;"`
      YuJiDaoZhangRiQi  *time.Time `json:"yuJiDaoZhangRiQi" form:"yuJiDaoZhangRiQi" gorm:"column:yu_ji_dao_zhang_ri_qi;comment:预计到账日期;"`
      ZhiBaoJin  *float64 `json:"zhiBaoJin" form:"zhiBaoJin" gorm:"column:zhi_bao_jin;comment:质保金;size:10;"`
      ZhuangTai  *int `json:"zhuangTai" form:"zhuangTai" gorm:"column:zhuang_tai;comment:状态;size:10;"`
}


// TableName YingShouKuanBiao 表名
func (YingShouKuanBiao) TableName() string {
  return "ying_shou_kuan_biao"
}

