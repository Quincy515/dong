// 自动生成模板ShiGongJinDuBiao
package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// ShiGongJinDuBiao 结构体
type ShiGongJinDuBiao struct {
      global.GVA_MODEL
      BeiZhu  string `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;size:200;"`
      GongZuoMingChengId  *int `json:"gongZuoMingChengId" form:"gongZuoMingChengId" gorm:"column:gong_zuo_ming_cheng_id;comment:工作名称ID;size:10;"`
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      ShiJiKaiShiRiQi  *time.Time `json:"shiJiKaiShiRiQi" form:"shiJiKaiShiRiQi" gorm:"column:shi_ji_kai_shi_ri_qi;comment:实际开始日期;"`
      ShiJiTianShu  *int `json:"shiJiTianShu" form:"shiJiTianShu" gorm:"column:shi_ji_tian_shu;comment:实际天数;size:10;"`
      ShiJiWanChengRiQi  *time.Time `json:"shiJiWanChengRiQi" form:"shiJiWanChengRiQi" gorm:"column:shi_ji_wan_cheng_ri_qi;comment:实际完成日期;"`
      TianJiaRen  *int `json:"tianJiaRen" form:"tianJiaRen" gorm:"column:tian_jia_ren;comment:添加人;size:10;"`
      TianJiaShiJian  *time.Time `json:"tianJiaShiJian" form:"tianJiaShiJian" gorm:"column:tian_jia_shi_jian;comment:添加时间;"`
      WanChengBaiFenBi  *float64 `json:"wanChengBaiFenBi" form:"wanChengBaiFenBi" gorm:"column:wan_cheng_bai_fen_bi;comment:完成百分比;"`
      XiangMuId  *int `json:"xiangMuId" form:"xiangMuId" gorm:"column:xiang_mu_id;comment:项目ID;size:10;"`
      ZhuangTai  *int `json:"zhuangTai" form:"zhuangTai" gorm:"column:zhuang_tai;comment:状态;size:10;"`
}


// TableName ShiGongJinDuBiao 表名
func (ShiGongJinDuBiao) TableName() string {
  return "shi_gong_jin_du_biao"
}

