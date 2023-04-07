// 自动生成模板ChuRuKuMingXiBiao
package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// ChuRuKuMingXiBiao 结构体
type ChuRuKuMingXiBiao struct {
      global.GVA_MODEL
      BeiZhu  string `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
      CaiGouLiang  *float64 `json:"caiGouLiang" form:"caiGouLiang" gorm:"column:cai_gou_liang;comment:采购量;"`
      CaiLiaoId  *int `json:"caiLiaoId" form:"caiLiaoId" gorm:"column:cai_liao_id;comment:材料ID;size:10;"`
      ChuRuKuDanJuId  *int `json:"chuRuKuDanJuId" form:"chuRuKuDanJuId" gorm:"column:chu_ru_ku_dan_ju_id;comment:出入库单据ID;size:10;"`
      DanJia  *float64 `json:"danJia" form:"danJia" gorm:"column:dan_jia;comment:单价;"`
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      RuKuLiang  *float64 `json:"ruKuLiang" form:"ruKuLiang" gorm:"column:ru_ku_liang;comment:入库量;"`
      ShiHouKuCunLiang  *float64 `json:"shiHouKuCunLiang" form:"shiHouKuCunLiang" gorm:"column:shi_hou_ku_cun_liang;comment:事后库存量;"`
      ZhuangTai  *int `json:"zhuangTai" form:"zhuangTai" gorm:"column:zhuang_tai;comment:状态;size:10;"`
      ZongJia  *float64 `json:"zongJia" form:"zongJia" gorm:"column:zong_jia;comment:总价;size:10;"`
}


// TableName ChuRuKuMingXiBiao 表名
func (ChuRuKuMingXiBiao) TableName() string {
  return "chu_ru_ku_ming_xi_biao"
}

