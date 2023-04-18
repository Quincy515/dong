// 自动生成模板GongSiZhangHaoMiMa
package hr

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// GongSiZhangHaoMiMa 结构体
type GongSiZhangHaoMiMa struct {
	global.GVA_MODEL
	BeiZhu          string     `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
	Group           string     `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
	JieShuYouXiaoQi *time.Time `json:"jieShuYouXiaoQi" form:"jieShuYouXiaoQi" gorm:"column:jie_shu_you_xiao_qi;comment:结束有效期;"`
	KaiShiYouXiaoQi *time.Time `json:"kaiShiYouXiaoQi" form:"kaiShiYouXiaoQi" gorm:"column:kai_shi_you_xiao_qi;comment:开始有效期;"`
	LianJie         string     `json:"lianJie" form:"lianJie" gorm:"column:lian_jie;comment:链接;size:200;"`
	MiMa            string     `json:"miMa" form:"miMa" gorm:"column:mi_ma;comment:密码;size:20;"`
	NianFei         string     `json:"nianFei" form:"nianFei" gorm:"column:nian_fei;comment:年费;size:50;"`
	XiTongMingCheng string     `json:"xiTongMingCheng" form:"xiTongMingCheng" gorm:"column:xi_tong_ming_cheng;comment:系统名称;size:100;"`
	YongHuMing      string     `json:"yongHuMing" form:"yongHuMing" gorm:"column:yong_hu_ming;comment:用户名;size:20;"`
	YongTu          string     `json:"yongTu" form:"yongTu" gorm:"column:yong_tu;comment:用途;"`
	ZhaoHuiMiMaJiLu string     `json:"zhaoHuiMiMaJiLu" form:"zhaoHuiMiMaJiLu" gorm:"column:zhao_hui_mi_ma_ji_lu;comment:找回密码记录;"`
}

// TableName GongSiZhangHaoMiMa 表名
func (GongSiZhangHaoMiMa) TableName() string {
	return "gong_si_zhang_hao_mi_ma"
}
