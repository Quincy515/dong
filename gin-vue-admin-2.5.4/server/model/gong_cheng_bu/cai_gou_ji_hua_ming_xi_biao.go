// 自动生成模板CaiGouJiHuaMingXiBiao
package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// CaiGouJiHuaMingXiBiao 结构体
type CaiGouJiHuaMingXiBiao struct {
      global.GVA_MODEL
      BeiZhu  string `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
      CaiGouJiHuaId  *int `json:"caiGouJiHuaId" form:"caiGouJiHuaId" gorm:"column:cai_gou_ji_hua_id;comment:采购计划ID;size:10;"`
      CaiGouLiang  *float64 `json:"caiGouLiang" form:"caiGouLiang" gorm:"column:cai_gou_liang;comment:采购量;"`
      CaiLiaoId  *int `json:"caiLiaoId" form:"caiLiaoId" gorm:"column:cai_liao_id;comment:材料ID;size:10;"`
      DanJia  string `json:"danJia" form:"danJia" gorm:"column:dan_jia;comment:单价;size:10;"`
      GongYingShangId  *int `json:"gongYingShangId" form:"gongYingShangId" gorm:"column:gong_ying_shang_id;comment:供应商ID;size:10;"`
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      RuKuLiang  *float64 `json:"ruKuLiang" form:"ruKuLiang" gorm:"column:ru_ku_liang;comment:入库量;"`
      ZhuangTai  *int `json:"zhuangTai" form:"zhuangTai" gorm:"column:zhuang_tai;comment:状态;size:10;"`
      ZongJia  *float64 `json:"zongJia" form:"zongJia" gorm:"column:zong_jia;comment:总价;size:10;"`
}


// TableName CaiGouJiHuaMingXiBiao 表名
func (CaiGouJiHuaMingXiBiao) TableName() string {
  return "cai_gou_ji_hua_ming_xi_biao"
}

