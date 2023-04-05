// 自动生成模板XiangMuShiGongBiao
package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// XiangMuShiGongBiao 结构体
type XiangMuShiGongBiao struct {
      global.GVA_MODEL
      BeiZhu  string `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
      GongZuoMingCheng  string `json:"gongZuoMingCheng" form:"gongZuoMingCheng" gorm:"column:gong_zuo_ming_cheng;comment:工作名称;size:100;"`
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      JiHuaKaiShiRiQi  *time.Time `json:"jiHuaKaiShiRiQi" form:"jiHuaKaiShiRiQi" gorm:"column:ji_hua_kai_shi_ri_qi;comment:计划开始日期;"`
      JiHuaTianShu  *int `json:"jiHuaTianShu" form:"jiHuaTianShu" gorm:"column:ji_hua_tian_shu;comment:计划天数;size:10;"`
      JiHuaWanChengRiQi  *time.Time `json:"jiHuaWanChengRiQi" form:"jiHuaWanChengRiQi" gorm:"column:ji_hua_wan_cheng_ri_qi;comment:计划完成日期;"`
      ShunXu  *int `json:"shunXu" form:"shunXu" gorm:"column:shun_xu;comment:顺序;size:10;"`
      WanChengBi  *float64 `json:"wanChengBi" form:"wanChengBi" gorm:"column:wan_cheng_bi;comment:完成比;"`
      XiangMuId  *int `json:"xiangMuId" form:"xiangMuId" gorm:"column:xiang_mu_id;comment:项目ID;size:10;"`
      ZhanZongGongZuoLiangBaiFenBi  *float64 `json:"zhanZongGongZuoLiangBaiFenBi" form:"zhanZongGongZuoLiangBaiFenBi" gorm:"column:zhan_zong_gong_zuo_liang_bai_fen_bi;comment:占总工作量百分比;"`
      ZhuangTai  *int `json:"zhuangTai" form:"zhuangTai" gorm:"column:zhuang_tai;comment:状态;size:10;"`
}


// TableName XiangMuShiGongBiao 表名
func (XiangMuShiGongBiao) TableName() string {
  return "xiang_mu_shi_gong_biao"
}

