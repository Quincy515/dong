// 自动生成模板HuiYuanGuanLi
package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// HuiYuanGuanLi 结构体
type HuiYuanGuanLi struct {
      global.GVA_MODEL
      BeiZhu  string `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      HuiFei  string `json:"huiFei" form:"huiFei" gorm:"column:hui_fei;comment:会费;size:50;"`
      HuiYuanMingCheng  string `json:"huiYuanMingCheng" form:"huiYuanMingCheng" gorm:"column:hui_yuan_ming_cheng;comment:会员名称;size:200;"`
      JiBie  string `json:"jiBie" form:"jiBie" gorm:"column:ji_bie;comment:级别;size:50;"`
      JiaoFeiShiJian  *time.Time `json:"jiaoFeiShiJian" form:"jiaoFeiShiJian" gorm:"column:jiao_fei_shi_jian;comment:缴费时间;"`
      JieShuYouXiaoQi  *time.Time `json:"jieShuYouXiaoQi" form:"jieShuYouXiaoQi" gorm:"column:jie_shu_you_xiao_qi;comment:结束有效期;"`
      KaiShiYouXiaoQi  *time.Time `json:"kaiShiYouXiaoQi" form:"kaiShiYouXiaoQi" gorm:"column:kai_shi_you_xiao_qi;comment:开始有效期;"`
      RuHuiLianXi  string `json:"ruHuiLianXi" form:"ruHuiLianXi" gorm:"column:ru_hui_lian_xi;comment:入会联系;size:20;"`
      WangZhi  string `json:"wangZhi" form:"wangZhi" gorm:"column:wang_zhi;comment:网址;size:200;"`
}


// TableName HuiYuanGuanLi 表名
func (HuiYuanGuanLi) TableName() string {
  return "hui_yuan_guan_li"
}

