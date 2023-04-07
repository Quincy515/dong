// 自动生成模板SheBeiZuLinBiao
package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// SheBeiZuLinBiao 结构体
type SheBeiZuLinBiao struct {
      global.GVA_MODEL
      BeiZhu  string `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      TianJiaRen  *int `json:"tianJiaRen" form:"tianJiaRen" gorm:"column:tian_jia_ren;comment:添加人;size:10;"`
      TianJiaRiQi  *time.Time `json:"tianJiaRiQi" form:"tianJiaRiQi" gorm:"column:tian_jia_ri_qi;comment:添加日期;"`
      XiangMuId  *int `json:"xiangMuId" form:"xiangMuId" gorm:"column:xiang_mu_id;comment:项目ID;size:10;"`
      ZhuangTai  *int `json:"zhuangTai" form:"zhuangTai" gorm:"column:zhuang_tai;comment:状态;size:10;"`
      ZuLinId  *int `json:"zuLinId" form:"zuLinId" gorm:"column:zu_lin_id;comment:租赁ID;size:10;"`
      ZuLinYongTu  string `json:"zuLinYongTu" form:"zuLinYongTu" gorm:"column:zu_lin_yong_tu;comment:租赁用途;size:500;"`
}


// TableName SheBeiZuLinBiao 表名
func (SheBeiZuLinBiao) TableName() string {
  return "she_bei_zu_lin_biao"
}

