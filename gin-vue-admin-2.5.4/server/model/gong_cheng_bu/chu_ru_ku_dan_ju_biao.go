// 自动生成模板ChuRuKuDanJuBiao
package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// ChuRuKuDanJuBiao 结构体
type ChuRuKuDanJuBiao struct {
      global.GVA_MODEL
      BeiZhu  string `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
      CaiGouJiHuaId  *int `json:"caiGouJiHuaId" form:"caiGouJiHuaId" gorm:"column:cai_gou_ji_hua_id;comment:采购计划ID;size:10;"`
      ChuRuKuDanJuLeiXing  *int `json:"chuRuKuDanJuLeiXing" form:"chuRuKuDanJuLeiXing" gorm:"column:chu_ru_ku_dan_ju_lei_xing;comment:类型 1. 入库 2. 出库;size:10;"`
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      TianJiaRen  *int `json:"tianJiaRen" form:"tianJiaRen" gorm:"column:tian_jia_ren;comment:添加人;size:10;"`
      TianJiaRiQi  *time.Time `json:"tianJiaRiQi" form:"tianJiaRiQi" gorm:"column:tian_jia_ri_qi;comment:添加日期;"`
      XiangMuId  *int `json:"xiangMuId" form:"xiangMuId" gorm:"column:xiang_mu_id;comment:项目ID;size:10;"`
      ZhuangTai  *int `json:"zhuangTai" form:"zhuangTai" gorm:"column:zhuang_tai;comment:状态;size:10;"`
}


// TableName ChuRuKuDanJuBiao 表名
func (ChuRuKuDanJuBiao) TableName() string {
  return "chu_ru_ku_dan_ju_biao"
}

