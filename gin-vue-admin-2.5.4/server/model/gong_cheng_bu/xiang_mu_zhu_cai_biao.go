// 自动生成模板XiangMuZhuCaiBiao
package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// XiangMuZhuCaiBiao 结构体
type XiangMuZhuCaiBiao struct {
      global.GVA_MODEL
      BeiZhu  string `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
      CaiLiaoBianHao  string `json:"caiLiaoBianHao" form:"caiLiaoBianHao" gorm:"column:cai_liao_bian_hao;comment:材料编号;size:50;"`
      CaiLiaoMingCheng  string `json:"caiLiaoMingCheng" form:"caiLiaoMingCheng" gorm:"column:cai_liao_ming_cheng;comment:材料名称;size:50;"`
      DanJia  *float64 `json:"danJia" form:"danJia" gorm:"column:dan_jia;comment:单价;"`
      DanWei  string `json:"danWei" form:"danWei" gorm:"column:dan_wei;comment:单位;size:10;"`
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      HeJi  *float64 `json:"heJi" form:"heJi" gorm:"column:he_ji;comment:合计;size:10;"`
      JiaChaHeJi  *float64 `json:"jiaChaHeJi" form:"jiaChaHeJi" gorm:"column:jia_cha_he_ji;comment:价差合计;size:10;"`
      XianHangJiaGe  *float64 `json:"xianHangJiaGe" form:"xianHangJiaGe" gorm:"column:xian_hang_jia_ge;comment:现行价格;size:10;"`
      XiangMuMingCheng  *int `json:"xiangMuMingCheng" form:"xiangMuMingCheng" gorm:"column:xiang_mu_ming_cheng;comment:项目名称;size:10;"`
      YongLiang  *float64 `json:"yongLiang" form:"yongLiang" gorm:"column:yong_liang;comment:用量;"`
}


// TableName XiangMuZhuCaiBiao 表名
func (XiangMuZhuCaiBiao) TableName() string {
  return "xiang_mu_zhu_cai_biao"
}

