// 自动生成模板GongZhangDengJi
package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// GongZhangDengJi 结构体
type GongZhangDengJi struct {
      global.GVA_MODEL
      BeiZhu  string `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
      GongZhangMingCheng  string `json:"gongZhangMingCheng" form:"gongZhangMingCheng" gorm:"column:gong_zhang_ming_cheng;comment:公章名称;size:100;"`
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      JingShouRen  string `json:"jingShouRen" form:"jingShouRen" gorm:"column:jing_shou_ren;comment:经手人;size:100;"`
      JingShouShiJian  *time.Time `json:"jingShouShiJian" form:"jingShouShiJian" gorm:"column:jing_shou_shi_jian;comment:经手时间;"`
      ShenQingRen  string `json:"shenQingRen" form:"shenQingRen" gorm:"column:shen_qing_ren;comment:申请人;size:100;"`
      ShenQingShiJian  *time.Time `json:"shenQingShiJian" form:"shenQingShiJian" gorm:"column:shen_qing_shi_jian;comment:申请时间;"`
      ShiXiang  string `json:"shiXiang" form:"shiXiang" gorm:"column:shi_xiang;comment:事项;size:100;"`
      ShuLiang  *int `json:"shuLiang" form:"shuLiang" gorm:"column:shu_liang;comment:数量;size:10;"`
      XiangQing  string `json:"xiangQing" form:"xiangQing" gorm:"column:xiang_qing;comment:详情;"`
}


// TableName GongZhangDengJi 表名
func (GongZhangDengJi) TableName() string {
  return "gong_zhang_deng_ji"
}

