// 自动生成模板KaoQinTongJi
package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// KaoQinTongJi 结构体
type KaoQinTongJi struct {
	global.GVA_MODEL
	BeiZhu         string `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;size:100;"`
	ChuQinTianShu  *int   `json:"chuQinTianShu" form:"chuQinTianShu" gorm:"column:chu_qin_tian_shu;comment:出勤天数;size:10;"`
	Group          string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
	JieShuRiQi     string `json:"jieShuRiQi" form:"jieShuRiQi" gorm:"column:jie_shu_ri_qi;comment:结束日期;size:100;"`
	KaiShiRiQi     string `json:"kaiShiRiQi" form:"kaiShiRiQi" gorm:"column:kai_shi_ri_qi;comment:开始日期;size:100;"`
	QingJiaTianShu *int   `json:"qingJiaTianShu" form:"qingJiaTianShu" gorm:"column:qing_jia_tian_shu;comment:请假天数;size:10;"`
	XiangQing      string `json:"xiangQing" form:"xiangQing" gorm:"column:xiang_qing;comment:详情;size:100;"`
	XingMing       string `json:"xingMing" form:"xingMing" gorm:"column:xing_ming;comment:姓名;size:100;"`
}

// TableName KaoQinTongJi 表名
func (KaoQinTongJi) TableName() string {
	return "kao_qin_tong_ji"
}
