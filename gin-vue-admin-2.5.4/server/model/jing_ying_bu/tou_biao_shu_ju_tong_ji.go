// 自动生成模板TouBiaoShuJuTongJi
package jing_ying_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// TouBiaoShuJuTongJi 结构体
type TouBiaoShuJuTongJi struct {
      global.GVA_MODEL
      BeiZhu  string `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
      DengJiYaoQiu  *int `json:"dengJiYaoQiu" form:"dengJiYaoQiu" gorm:"column:deng_ji_yao_qiu;comment:等级要求;size:10;"`
      DiErBaoJia  string `json:"diErBaoJia" form:"diErBaoJia" gorm:"column:di_er_bao_jia;comment:第二报价;size:20;"`
      DiErXiaFuLu  string `json:"diErXiaFuLu" form:"diErXiaFuLu" gorm:"column:di_er_xia_fu_lu;comment:第二下浮率;size:10;"`
      DiErZhongBiaoHouXuanRen  string `json:"diErZhongBiaoHouXuanRen" form:"diErZhongBiaoHouXuanRen" gorm:"column:di_er_zhong_biao_hou_xuan_ren;comment:第二中标候选人;size:100;"`
      DiSanBaoJia  string `json:"diSanBaoJia" form:"diSanBaoJia" gorm:"column:di_san_bao_jia;comment:第三报价;size:20;"`
      DiSanXiaFuLu  string `json:"diSanXiaFuLu" form:"diSanXiaFuLu" gorm:"column:di_san_xia_fu_lu;comment:第三下浮率;size:10;"`
      DiSanZhongBiaoHouXuanRen  string `json:"diSanZhongBiaoHouXuanRen" form:"diSanZhongBiaoHouXuanRen" gorm:"column:di_san_zhong_biao_hou_xuan_ren;comment:第三中标候选人;size:100;"`
      DiYiBaoJia  string `json:"diYiBaoJia" form:"diYiBaoJia" gorm:"column:di_yi_bao_jia;comment:第一报价;size:20;"`
      DiYiXiaFuLu  string `json:"diYiXiaFuLu" form:"diYiXiaFuLu" gorm:"column:di_yi_xia_fu_lu;comment:第一下浮率;size:10;"`
      DiYiZhongBiaoHouXuanRen  string `json:"diYiZhongBiaoHouXuanRen" form:"diYiZhongBiaoHouXuanRen" gorm:"column:di_yi_zhong_biao_hou_xuan_ren;comment:第一中标候选人;size:100;"`
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      KaiBiaoShiJian  *time.Time `json:"kaiBiaoShiJian" form:"kaiBiaoShiJian" gorm:"column:kai_biao_shi_jian;comment:开标时间;"`
      SuoShuDiQu  *int `json:"suoShuDiQu" form:"suoShuDiQu" gorm:"column:suo_shu_di_qu;comment:所属地区;size:10;"`
      TouBiaoDanWeiMingCheng  string `json:"touBiaoDanWeiMingCheng" form:"touBiaoDanWeiMingCheng" gorm:"column:tou_biao_dan_wei_ming_cheng;comment:投标单位名称;size:100;"`
      XiangMuMingCheng  string `json:"xiangMuMingCheng" form:"xiangMuMingCheng" gorm:"column:xiang_mu_ming_cheng;comment:项目名称;size:100;"`
      ZhaoBiaoGuiMo  string `json:"zhaoBiaoGuiMo" form:"zhaoBiaoGuiMo" gorm:"column:zhao_biao_gui_mo;comment:招标规模;size:20;"`
      ZiZhiYaoQiu  *int `json:"ziZhiYaoQiu" form:"ziZhiYaoQiu" gorm:"column:zi_zhi_yao_qiu;comment:资质要求;size:10;"`
}


// TableName TouBiaoShuJuTongJi 表名
func (TouBiaoShuJuTongJi) TableName() string {
  return "tou_biao_shu_ju_tong_ji"
}

