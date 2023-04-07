// 自动生成模板NeiBuBaoXiaoBiao
package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// NeiBuBaoXiaoBiao 结构体
type NeiBuBaoXiaoBiao struct {
      global.GVA_MODEL
      BaoXiaoBuMen  *int `json:"baoXiaoBuMen" form:"baoXiaoBuMen" gorm:"column:bao_xiao_bu_men;comment:报销部门;size:10;"`
      BaoXiaoJinE  *float64 `json:"baoXiaoJinE" form:"baoXiaoJinE" gorm:"column:bao_xiao_jin_e;comment:报销金额;size:10;"`
      BaoXiaoNeiRong  string `json:"baoXiaoNeiRong" form:"baoXiaoNeiRong" gorm:"column:bao_xiao_nei_rong;comment:报销内容;"`
      BaoXiaoRenYuan  *int `json:"baoXiaoRenYuan" form:"baoXiaoRenYuan" gorm:"column:bao_xiao_ren_yuan;comment:报销人员;size:10;"`
      BaoXiaoShiXiang  *int `json:"baoXiaoShiXiang" form:"baoXiaoShiXiang" gorm:"column:bao_xiao_shi_xiang;comment:报销事项;size:10;"`
      BeiZhuShuoMing  string `json:"beiZhuShuoMing" form:"beiZhuShuoMing" gorm:"column:bei_zhu_shuo_ming;comment:备注说明;size:200;"`
      FeiYongFaShengRiQi  *time.Time `json:"feiYongFaShengRiQi" form:"feiYongFaShengRiQi" gorm:"column:fei_yong_fa_sheng_ri_qi;comment:费用发生日期;"`
      GongSiMingCheng  *int `json:"gongSiMingCheng" form:"gongSiMingCheng" gorm:"column:gong_si_ming_cheng;comment:公司名称;size:10;"`
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      HuiDanDiZhi  string `json:"huiDanDiZhi" form:"huiDanDiZhi" gorm:"column:hui_dan_di_zhi;comment:回单地址-可多图片上传;size:500;"`
      BaoXiaoLeiXing  *int `json:"baoXiaoLeiXing" form:"baoXiaoLeiXing" gorm:"column:bao_xiao_lei_xing;comment:报销类型 1. 内部报销 2. 项目经理报销;size:10;"`
      PiaoJuDiZhi  string `json:"piaoJuDiZhi" form:"piaoJuDiZhi" gorm:"column:piao_ju_di_zhi;comment:票据地址-可多图片上传;size:500;"`
      ShenQingRiQi  *time.Time `json:"shenQingRiQi" form:"shenQingRiQi" gorm:"column:shen_qing_ri_qi;comment:申请日期;"`
      ShouCiBaoXiaoId  *int `json:"shouCiBaoXiaoId" form:"shouCiBaoXiaoId" gorm:"column:shou_ci_bao_xiao_id;comment:首次报销ID-再次申请时关联的ID;size:10;"`
      XiangMuMingCheng  *int `json:"xiangMuMingCheng" form:"xiangMuMingCheng" gorm:"column:xiang_mu_ming_cheng;comment:项目名称;size:10;"`
      ZhuangTai  *int `json:"zhuangTai" form:"zhuangTai" gorm:"column:zhuang_tai;comment:状态;size:10;"`
}


// TableName NeiBuBaoXiaoBiao 表名
func (NeiBuBaoXiaoBiao) TableName() string {
  return "nei_bu_bao_xiao_biao"
}

