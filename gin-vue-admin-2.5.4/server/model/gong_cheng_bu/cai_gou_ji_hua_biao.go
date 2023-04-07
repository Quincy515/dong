// 自动生成模板CaiGouJiHuaBiao
package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// CaiGouJiHuaBiao 结构体
type CaiGouJiHuaBiao struct {
      global.GVA_MODEL
      BeiZhu  string `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
      CaiGouFuZeRen  *int `json:"caiGouFuZeRen" form:"caiGouFuZeRen" gorm:"column:cai_gou_fu_ze_ren;comment:采购负责人;size:10;"`
      CaiGouJiHuaMingCheng  string `json:"caiGouJiHuaMingCheng" form:"caiGouJiHuaMingCheng" gorm:"column:cai_gou_ji_hua_ming_cheng;comment:采购计划名称;size:100;"`
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      LeiJiCaiGouJinE  *float64 `json:"leiJiCaiGouJinE" form:"leiJiCaiGouJinE" gorm:"column:lei_ji_cai_gou_jin_e;comment:累计采购金额;size:10;"`
      ShiJiZhiFuJinE  *float64 `json:"shiJiZhiFuJinE" form:"shiJiZhiFuJinE" gorm:"column:shi_ji_zhi_fu_jin_e;comment:实际支付金额;size:10;"`
      TianJiaRiQi  *time.Time `json:"tianJiaRiQi" form:"tianJiaRiQi" gorm:"column:tian_jia_ri_qi;comment:添加日期;"`
      XiangMuId  *int `json:"xiangMuId" form:"xiangMuId" gorm:"column:xiang_mu_id;comment:项目id;size:10;"`
      ZhuangTai  *int `json:"zhuangTai" form:"zhuangTai" gorm:"column:zhuang_tai;comment:状态;size:10;"`
}


// TableName CaiGouJiHuaBiao 表名
func (CaiGouJiHuaBiao) TableName() string {
  return "cai_gou_ji_hua_biao"
}

