// 自动生成模板SheBeiMingXiBiao
package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// SheBeiMingXiBiao 结构体
type SheBeiMingXiBiao struct {
      global.GVA_MODEL
      DanJia  *float64 `json:"danJia" form:"danJia" gorm:"column:dan_jia;comment:单价;"`
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      JiHuaTuiZuRiQi  *time.Time `json:"jiHuaTuiZuRiQi" form:"jiHuaTuiZuRiQi" gorm:"column:ji_hua_tui_zu_ri_qi;comment:计划退租日期;"`
      JiHuaZuLinRiQi  *time.Time `json:"jiHuaZuLinRiQi" form:"jiHuaZuLinRiQi" gorm:"column:ji_hua_zu_lin_ri_qi;comment:计划租赁日期;"`
      SheBeiId  *int `json:"sheBeiId" form:"sheBeiId" gorm:"column:she_bei_id;comment:设备id;size:10;"`
      ShiJiTuiZuRiQi  *time.Time `json:"shiJiTuiZuRiQi" form:"shiJiTuiZuRiQi" gorm:"column:shi_ji_tui_zu_ri_qi;comment:实际退租日期;"`
      ShiJiZuLinRiQi  *time.Time `json:"shiJiZuLinRiQi" form:"shiJiZuLinRiQi" gorm:"column:shi_ji_zu_lin_ri_qi;comment:实际租赁日期;"`
      ZhuangTai  *int `json:"zhuangTai" form:"zhuangTai" gorm:"column:zhuang_tai;comment:状态;size:10;"`
      ZuLinDanJuId  *int `json:"zuLinDanJuId" form:"zuLinDanJuId" gorm:"column:zu_lin_dan_ju_id;comment:租赁单据ID;size:10;"`
      ZuLinShuLiang  *int `json:"zuLinShuLiang" form:"zuLinShuLiang" gorm:"column:zu_lin_shu_liang;comment:租赁数量;size:10;"`
      ZuLinTianShu  *int `json:"zuLinTianShu" form:"zuLinTianShu" gorm:"column:zu_lin_tian_shu;comment:租赁天数;size:10;"`
}


// TableName SheBeiMingXiBiao 表名
func (SheBeiMingXiBiao) TableName() string {
  return "she_bei_ming_xi_biao"
}

