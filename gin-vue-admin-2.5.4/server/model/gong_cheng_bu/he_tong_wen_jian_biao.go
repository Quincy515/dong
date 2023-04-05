// 自动生成模板HeTongWenJianBiao
package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// HeTongWenJianBiao 结构体
type HeTongWenJianBiao struct {
      global.GVA_MODEL
      BeiZhu  string `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
      FaPiaoLeiXing  *int `json:"faPiaoLeiXing" form:"faPiaoLeiXing" gorm:"column:fa_piao_lei_xing;comment:发票类型;size:10;"`
      FuKuanFangShi  *int `json:"fuKuanFangShi" form:"fuKuanFangShi" gorm:"column:fu_kuan_fang_shi;comment:付款方式;size:10;"`
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      HeTongBianHao  string `json:"heTongBianHao" form:"heTongBianHao" gorm:"column:he_tong_bian_hao;comment:合同编号;size:100;"`
      HeTongJinE  *float64 `json:"heTongJinE" form:"heTongJinE" gorm:"column:he_tong_jin_e;comment:合同金额;size:10;"`
      HeTongLeiBie  *int `json:"heTongLeiBie" form:"heTongLeiBie" gorm:"column:he_tong_lei_bie;comment:合同类别 1.主合同 2.补充合同;size:10;"`
      HeTongLeiXing  *int `json:"heTongLeiXing" form:"heTongLeiXing" gorm:"column:he_tong_lei_xing;comment:合同类型 1.承包 2.分包 3.采购 4.租聘;size:10;"`
      HeTongMingCheng  string `json:"heTongMingCheng" form:"heTongMingCheng" gorm:"column:he_tong_ming_cheng;comment:合同名称;size:100;"`
      HeTongXiaZaiDiZhi  string `json:"heTongXiaZaiDiZhi" form:"heTongXiaZaiDiZhi" gorm:"column:he_tong_xia_zai_di_zhi;comment:合同下载地址;"`
      HeTongZhuangTai  *int `json:"heTongZhuangTai" form:"heTongZhuangTai" gorm:"column:he_tong_zhuang_tai;comment:合同状态;size:10;"`
      JiaFangGongSi  *int `json:"jiaFangGongSi" form:"jiaFangGongSi" gorm:"column:jia_fang_gong_si;comment:甲方公司;size:10;"`
      JieSuanFangShi  *int `json:"jieSuanFangShi" form:"jieSuanFangShi" gorm:"column:jie_suan_fang_shi;comment:结算方式;size:10;"`
      KaiHuHang  string `json:"kaiHuHang" form:"kaiHuHang" gorm:"column:kai_hu_hang;comment:开户行;size:100;"`
      KaiPiaoDanWei  string `json:"kaiPiaoDanWei" form:"kaiPiaoDanWei" gorm:"column:kai_piao_dan_wei;comment:开票单位;size:100;"`
      QianDingRen  *int `json:"qianDingRen" form:"qianDingRen" gorm:"column:qian_ding_ren;comment:签订人;size:10;"`
      QianDingRiQi  *time.Time `json:"qianDingRiQi" form:"qianDingRiQi" gorm:"column:qian_ding_ri_qi;comment:签订日期;"`
      ShouFuKuanTiaoJian  string `json:"shouFuKuanTiaoJian" form:"shouFuKuanTiaoJian" gorm:"column:shou_fu_kuan_tiao_jian;comment:收付款条件;size:100;"`
      ShouKuanZhangHao  string `json:"shouKuanZhangHao" form:"shouKuanZhangHao" gorm:"column:shou_kuan_zhang_hao;comment:收款账号;size:100;"`
      TianJiaRen  *int `json:"tianJiaRen" form:"tianJiaRen" gorm:"column:tian_jia_ren;comment:添加人;size:10;"`
      TianJiaShiJian  *time.Time `json:"tianJiaShiJian" form:"tianJiaShiJian" gorm:"column:tian_jia_shi_jian;comment:添加时间;"`
      XiangMuId  *int `json:"xiangMuId" form:"xiangMuId" gorm:"column:xiang_mu_id;comment:项目ID;size:10;"`
      YiFangGongSi  *int `json:"yiFangGongSi" form:"yiFangGongSi" gorm:"column:yi_fang_gong_si;comment:乙方公司;size:10;"`
      YinHangMingCheng  string `json:"yinHangMingCheng" form:"yinHangMingCheng" gorm:"column:yin_hang_ming_cheng;comment:银行名称;size:100;"`
      ZhuYaoTiaoKuan  string `json:"zhuYaoTiaoKuan" form:"zhuYaoTiaoKuan" gorm:"column:zhu_yao_tiao_kuan;comment:主要条款;"`
}


// TableName HeTongWenJianBiao 表名
func (HeTongWenJianBiao) TableName() string {
  return "he_tong_wen_jian_biao"
}

