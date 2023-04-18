// 自动生成模板HuiYuanGuanLi
package hr

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HuiYuanGuanLi 结构体
type HuiYuanGuanLi struct {
	global.GVA_MODEL
	BeiZhu           string     `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
	Group            string     `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
	HuiFei           string     `json:"huiFei" form:"huiFei" gorm:"column:hui_fei;comment:会费;size:50;"`
	HuiYuanMingCheng string     `json:"huiYuanMingCheng" form:"huiYuanMingCheng" gorm:"column:hui_yuan_ming_cheng;comment:会员名称;size:200;"`
	JiBie            string     `json:"jiBie" form:"jiBie" gorm:"column:ji_bie;comment:级别;size:50;"`
	JiaoFeiShiJian   *time.Time `json:"jiaoFeiShiJian" form:"jiaoFeiShiJian" gorm:"column:jiao_fei_shi_jian;comment:缴费时间;"`
	JieShuYouXiaoQi  *time.Time `json:"jieShuYouXiaoQi" form:"jieShuYouXiaoQi" gorm:"column:jie_shu_you_xiao_qi;comment:结束有效期;"`
	KaiShiYouXiaoQi  *time.Time `json:"kaiShiYouXiaoQi" form:"kaiShiYouXiaoQi" gorm:"column:kai_shi_you_xiao_qi;comment:开始有效期;"`
	RuHuiLianXi      string     `json:"ruHuiLianXi" form:"ruHuiLianXi" gorm:"column:ru_hui_lian_xi;comment:入会联系;size:20;"`
	WangZhi          string     `json:"wangZhi" form:"wangZhi" gorm:"column:wang_zhi;comment:网址;size:200;"`
	HuiYuanLeiBie    *int       `json:"huiYuanLeiBie" form:"huiYuanLeiBie" gorm:"column:hui_yuan_lei_bie;comment:会员类别;size:10;"`
	HuiYuanJiBie     *int       `json:"huiYuanJiBie" form:"huiYuanJiBie" gorm:"column:hui_yuan_ji_bie;comment:会员级别;size:10;"`
	GuoQiTiXing      *int       `json:"guoQiTiXing" form:"guoQiTiXing" gorm:"column:guo_qi_ti_xing;comment:过期提醒;size:10;"`
	GuoQiShiJian     *time.Time `json:"guoQiShiJian" form:"guoQiShiJian" gorm:"column:guo_qi_shi_jian;comment:过期时间;"`
}

// TableName HuiYuanGuanLi 表名
func (HuiYuanGuanLi) TableName() string {
	return "hui_yuan_guan_li"
}
