// 自动生成模板WaiBuFeiYongQingKuanBiao
package cai_wu_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// WaiBuFeiYongQingKuanBiao 结构体
type WaiBuFeiYongQingKuanBiao struct {
      global.GVA_MODEL
      BaoXiaoJinE  *float64 `json:"baoXiaoJinE" form:"baoXiaoJinE" gorm:"column:bao_xiao_jin_e;comment:报销金额;size:10;"`
      BaoXiaoLeiXing  *int `json:"baoXiaoLeiXing" form:"baoXiaoLeiXing" gorm:"column:bao_xiao_lei_xing;comment:报销类型;size:10;"`
      BaoXiaoMingCheng  string `json:"baoXiaoMingCheng" form:"baoXiaoMingCheng" gorm:"column:bao_xiao_ming_cheng;comment:报销名称;size:200;"`
      BaoXiaoNeiRong  string `json:"baoXiaoNeiRong" form:"baoXiaoNeiRong" gorm:"column:bao_xiao_nei_rong;comment:报销内容;"`
      BaoXiaoRenYuan  *int `json:"baoXiaoRenYuan" form:"baoXiaoRenYuan" gorm:"column:bao_xiao_ren_yuan;comment:报销人员;size:10;"`
      BaoXiaoShiXiang  *int `json:"baoXiaoShiXiang" form:"baoXiaoShiXiang" gorm:"column:bao_xiao_shi_xiang;comment:报销事项;size:10;"`
      BeiZhuShuoMing  string `json:"beiZhuShuoMing" form:"beiZhuShuoMing" gorm:"column:bei_zhu_shuo_ming;comment:备注说明;size:200;"`
      FeiYongFaShengRiQi  *time.Time `json:"feiYongFaShengRiQi" form:"feiYongFaShengRiQi" gorm:"column:fei_yong_fa_sheng_ri_qi;comment:费用发生日期;"`
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      HeTongId  *int `json:"heTongId" form:"heTongId" gorm:"column:he_tong_id;comment:合同ID;size:10;"`
      HuiDanDiZhi  string `json:"huiDanDiZhi" form:"huiDanDiZhi" gorm:"column:hui_dan_di_zhi;comment:回单地址-可多图片上传NULL;size:500;"`
      KaiHuHang  string `json:"kaiHuHang" form:"kaiHuHang" gorm:"column:kai_hu_hang;comment:开户行;size:100;"`
      ShenQingRiQi  *time.Time `json:"shenQingRiQi" form:"shenQingRiQi" gorm:"column:shen_qing_ri_qi;comment:申请日期;"`
      ShiFuJinE  *float64 `json:"shiFuJinE" form:"shiFuJinE" gorm:"column:shi_fu_jin_e;comment:实付金额;size:10;"`
      ShouCiBaoXiaoId  *int `json:"shouCiBaoXiaoId" form:"shouCiBaoXiaoId" gorm:"column:shou_ci_bao_xiao_id;comment:首次报销ID;size:10;"`
      ShouKuanZhangHao  string `json:"shouKuanZhangHao" form:"shouKuanZhangHao" gorm:"column:shou_kuan_zhang_hao;comment:收款账号;size:25;"`
      XiangMuMingCheng  *int `json:"xiangMuMingCheng" form:"xiangMuMingCheng" gorm:"column:xiang_mu_ming_cheng;comment:项目名称;size:10;"`
      YinHangMingCheng  string `json:"yinHangMingCheng" form:"yinHangMingCheng" gorm:"column:yin_hang_ming_cheng;comment:银行名称;size:20;"`
      ZhuangTai  *int `json:"zhuangTai" form:"zhuangTai" gorm:"column:zhuang_tai;comment:状态;size:10;"`
      ZuiWanFuKuanRiQi  *time.Time `json:"zuiWanFuKuanRiQi" form:"zuiWanFuKuanRiQi" gorm:"column:zui_wan_fu_kuan_ri_qi;comment:最晚付款日期;"`
}


// TableName WaiBuFeiYongQingKuanBiao 表名
func (WaiBuFeiYongQingKuanBiao) TableName() string {
  return "wai_bu_fei_yong_qing_kuan_biao"
}

