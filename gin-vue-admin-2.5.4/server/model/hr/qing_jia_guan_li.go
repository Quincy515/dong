// 自动生成模板QingJiaGuanLi
package hr

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// QingJiaGuanLi 结构体
type QingJiaGuanLi struct {
	global.GVA_MODEL
	BeiZhu          string     `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
	Group           string     `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
	JieShuRiQi      *time.Time `json:"jieShuRiQi" form:"jieShuRiQi" gorm:"column:jie_shu_ri_qi;comment:结束日期;"`
	JinDu           string     `json:"jinDu" form:"jinDu" gorm:"column:jin_du;comment:进度;size:10;"`
	KaiShiRiQi      *time.Time `json:"kaiShiRiQi" form:"kaiShiRiQi" gorm:"column:kai_shi_ri_qi;comment:开始日期;"`
	LeiXing         *int       `json:"leiXing" form:"leiXing" gorm:"column:lei_xing;comment:请假类型;size:10;"`
	ShenHeRen       string     `json:"shenHeRen" form:"shenHeRen" gorm:"column:shen_he_ren;comment:审核人;size:100;"`
	ShenHeShiJian   *time.Time `json:"shenHeShiJian" form:"shenHeShiJian" gorm:"column:shen_he_shi_jian;comment:审核时间;"`
	ShenHeXiangQing string     `json:"shenHeXiangQing" form:"shenHeXiangQing" gorm:"column:shen_he_xiang_qing;comment:审核详情;"`
	ShenHeZhuangTai string     `json:"shenHeZhuangTai" form:"shenHeZhuangTai" gorm:"column:shen_he_zhuang_tai;comment:审核状态;size:10;"`
	ShiYou          string     `json:"shiYou" form:"shiYou" gorm:"column:shi_you;comment:事由;"`
	TianShu         *int       `json:"tianShu" form:"tianShu" gorm:"column:tian_shu;comment:天数;size:10;"`
	XingMing        string     `json:"xingMing" form:"xingMing" gorm:"column:xing_ming;comment:姓名;size:100;"`
}

// TableName QingJiaGuanLi 表名
func (QingJiaGuanLi) TableName() string {
	return "qing_jia_guan_li"
}
