// 自动生成模板MeiRiKaoQin
package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// MeiRiKaoQin 结构体
type MeiRiKaoQin struct {
      global.GVA_MODEL
      BeiZhu  string `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
      ChuQin  *bool `json:"chuQin" form:"chuQin" gorm:"column:chu_qin;comment:出勤;"`
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      RiQi  *time.Time `json:"riQi" form:"riQi" gorm:"column:ri_qi;comment:日期;"`
      XiangQing  string `json:"xiangQing" form:"xiangQing" gorm:"column:xiang_qing;comment:详情;"`
      XingMing  string `json:"xingMing" form:"xingMing" gorm:"column:xing_ming;comment:姓名;size:100;"`
      XiuJiaQingJiaChuCha  *bool `json:"xiuJiaQingJiaChuCha" form:"xiuJiaQingJiaChuCha" gorm:"column:xiu_jia_qing_jia_chu_cha;comment:休假请假出差;"`
}


// TableName MeiRiKaoQin 表名
func (MeiRiKaoQin) TableName() string {
  return "mei_ri_kao_qin"
}

