// 自动生成模板ZhengShuJieChuJiLuBiao
package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// ZhengShuJieChuJiLuBiao 结构体
type ZhengShuJieChuJiLuBiao struct {
      global.GVA_MODEL
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      ZhengShuMingCheng  string `json:"zhengShuMingCheng" form:"zhengShuMingCheng" gorm:"column:zheng_shu_ming_cheng;comment:证书名称;size:100;"`
      ZhengShuBianHao  string `json:"zhengShuBianHao" form:"zhengShuBianHao" gorm:"column:zheng_shu_bian_hao;comment:证书编号;size:100;"`
      ZhengShuId  *int `json:"zhengShuId" form:"zhengShuId" gorm:"column:zhengShuId;comment:证书ID;size:10;"`
      JieChuRen  *int `json:"jieChuRen" form:"jieChuRen" gorm:"column:jieChuRen;comment:借出人ID;size:10;"`
      JieChuRiQi  *time.Time `json:"jieChuRiQi" form:"jieChuRiQi" gorm:"column:jieChuRiQi;comment:借出日期;"`
      JieChuShiYou  string `json:"jieChuShiYou" form:"jieChuShiYou" gorm:"column:jieChuShiYou;comment:借出事由;size:100;"`
      YuJiGuiHaiRiQi  *time.Time `json:"yuJiGuiHaiRiQi" form:"yuJiGuiHaiRiQi" gorm:"column:yuJiGuiHaiRiQi;comment:预计归还日期;"`
      ShiJiGuiHaiRiQi  *time.Time `json:"shiJiGuiHaiRiQi" form:"shiJiGuiHaiRiQi" gorm:"column:shiJiGuiHaiRiQi;comment:实际归还日期;"`
      TianJiaRen  *int `json:"tianJiaRen" form:"tianJiaRen" gorm:"column:tianJiaRen;comment:添加人;size:10;"`
      JieChuBeiZhu  string `json:"jieChuBeiZhu" form:"jieChuBeiZhu" gorm:"column:jieChuBeiZhu;comment:借出备注;"`
      GuiHaiBeiZhu  string `json:"guiHaiBeiZhu" form:"guiHaiBeiZhu" gorm:"column:guiHaiBeiZhu;comment:归还备注;"`
}


// TableName ZhengShuJieChuJiLuBiao 表名
func (ZhengShuJieChuJiLuBiao) TableName() string {
  return "zheng_shu_jie_chu_ji_lu_biao"
}

