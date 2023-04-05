// 自动生成模板FenBaoShangYuXiangMuGuanLianBiao
package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// FenBaoShangYuXiangMuGuanLianBiao 结构体
type FenBaoShangYuXiangMuGuanLianBiao struct {
      global.GVA_MODEL
      BeiZhu  string `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
      FenBaoShangId  *int `json:"fenBaoShangId" form:"fenBaoShangId" gorm:"column:fen_bao_shang_id;comment:分包商ID;size:10;"`
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      HeTongBianHao  string `json:"heTongBianHao" form:"heTongBianHao" gorm:"column:he_tong_bian_hao;comment:合同编号;size:100;"`
      HeTongMingCheng  string `json:"heTongMingCheng" form:"heTongMingCheng" gorm:"column:he_tong_ming_cheng;comment:合同名称;size:100;"`
      JinChangRiQi  *time.Time `json:"jinChangRiQi" form:"jinChangRiQi" gorm:"column:jin_chang_ri_qi;comment:进场日期;"`
      TuiChangRiQi  *time.Time `json:"tuiChangRiQi" form:"tuiChangRiQi" gorm:"column:tui_chang_ri_qi;comment:退场日期;"`
      WeiTuoRen  string `json:"weiTuoRen" form:"weiTuoRen" gorm:"column:wei_tuo_ren;comment:委托人;size:100;"`
      WeiTuoRenShouJi  string `json:"weiTuoRenShouJi" form:"weiTuoRenShouJi" gorm:"column:wei_tuo_ren_shou_ji;comment:委托人手机;size:100;"`
      XiangMuId  *int `json:"xiangMuId" form:"xiangMuId" gorm:"column:xiang_mu_id;comment:项目ID;size:10;"`
      ZhuangTai  *int `json:"zhuangTai" form:"zhuangTai" gorm:"column:zhuang_tai;comment:状态;size:10;"`
}


// TableName FenBaoShangYuXiangMuGuanLianBiao 表名
func (FenBaoShangYuXiangMuGuanLianBiao) TableName() string {
  return "fen_bao_shang_yu_xiang_mu_guan_lian_biao"
}

