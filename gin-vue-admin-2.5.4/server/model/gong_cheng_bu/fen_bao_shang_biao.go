// 自动生成模板FenBaoShangBiao
package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// FenBaoShangBiao 结构体
type FenBaoShangBiao struct {
      global.GVA_MODEL
      BeiZhu  string `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
      FaRen  string `json:"faRen" form:"faRen" gorm:"column:fa_ren;comment:法人;size:100;"`
      FaRenShouJi  string `json:"faRenShouJi" form:"faRenShouJi" gorm:"column:fa_ren_shou_ji;comment:法人手机;size:100;"`
      FenBaoShangLeiXing  *int `json:"fenBaoShangLeiXing" form:"fenBaoShangLeiXing" gorm:"column:fen_bao_shang_lei_xing;comment:分包商类型;size:10;"`
      FenBaoShangMingCheng  string `json:"fenBaoShangMingCheng" form:"fenBaoShangMingCheng" gorm:"column:fen_bao_shang_ming_cheng;comment:分包商名称;size:100;"`
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      KaiHuHang  string `json:"kaiHuHang" form:"kaiHuHang" gorm:"column:kai_hu_hang;comment:开户行;size:100;"`
      SheHuiXinYongDaiMa  string `json:"sheHuiXinYongDaiMa" form:"sheHuiXinYongDaiMa" gorm:"column:she_hui_xin_yong_dai_ma;comment:社会信用代码;size:100;"`
      TianJiaRen  *int `json:"tianJiaRen" form:"tianJiaRen" gorm:"column:tian_jia_ren;comment:添加人;size:10;"`
      TianJiaShiJian  *time.Time `json:"tianJiaShiJian" form:"tianJiaShiJian" gorm:"column:tian_jia_shi_jian;comment:添加时间;"`
      YinHangMingCheng  string `json:"yinHangMingCheng" form:"yinHangMingCheng" gorm:"column:yin_hang_ming_cheng;comment:银行名称;size:100;"`
      ZhangHao  string `json:"zhangHao" form:"zhangHao" gorm:"column:zhang_hao;comment:账号;size:100;"`
      ZhuangTai  *int `json:"zhuangTai" form:"zhuangTai" gorm:"column:zhuang_tai;comment:状态;size:10;"`
      ZiZhiZhengZhao  string `json:"ziZhiZhengZhao" form:"ziZhiZhengZhao" gorm:"column:zi_zhi_zheng_zhao;comment:资质证照;size:100;"`
}


// TableName FenBaoShangBiao 表名
func (FenBaoShangBiao) TableName() string {
  return "fen_bao_shang_biao"
}

