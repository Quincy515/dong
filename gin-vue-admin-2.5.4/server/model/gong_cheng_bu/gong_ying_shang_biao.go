// 自动生成模板GongYingShangBiao
package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// GongYingShangBiao 结构体
type GongYingShangBiao struct {
      global.GVA_MODEL
      BeiZhu  string `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
      FaPiaoLeiXing  *int `json:"faPiaoLeiXing" form:"faPiaoLeiXing" gorm:"column:fa_piao_lei_xing;comment:发票类型;size:10;"`
      FaPiaoTaiTou  string `json:"faPiaoTaiTou" form:"faPiaoTaiTou" gorm:"column:fa_piao_tai_tou;comment:发票抬头;size:100;"`
      FaRenShenFenZheng  string `json:"faRenShenFenZheng" form:"faRenShenFenZheng" gorm:"column:fa_ren_shen_fen_zheng;comment:法人身份证;size:100;"`
      GongSiDiZhi  string `json:"gongSiDiZhi" form:"gongSiDiZhi" gorm:"column:gong_si_di_zhi;comment:公司地址;size:100;"`
      GongYingShangMingCheng  *int `json:"gongYingShangMingCheng" form:"gongYingShangMingCheng" gorm:"column:gong_ying_shang_ming_cheng;comment:供应商名称;size:10;"`
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      KaiHuHang  string `json:"kaiHuHang" form:"kaiHuHang" gorm:"column:kai_hu_hang;comment:开户行;size:100;"`
      LianXiRen  string `json:"lianXiRen" form:"lianXiRen" gorm:"column:lian_xi_ren;comment:联系人;size:100;"`
      LianXiRenShouJi  string `json:"lianXiRenShouJi" form:"lianXiRenShouJi" gorm:"column:lian_xi_ren_shou_ji;comment:联系人手机;size:100;"`
      SheHuiXinYongDaiMa  string `json:"sheHuiXinYongDaiMa" form:"sheHuiXinYongDaiMa" gorm:"column:she_hui_xin_yong_dai_ma;comment:社会信用代码;size:100;"`
      TianJiaRen  *int `json:"tianJiaRen" form:"tianJiaRen" gorm:"column:tian_jia_ren;comment:添加人;size:10;"`
      TianJiaRiQi  *time.Time `json:"tianJiaRiQi" form:"tianJiaRiQi" gorm:"column:tian_jia_ri_qi;comment:添加日期;"`
      YinHangMingCheng  string `json:"yinHangMingCheng" form:"yinHangMingCheng" gorm:"column:yin_hang_ming_cheng;comment:银行名称;size:100;"`
      ZhangHao  string `json:"zhangHao" form:"zhangHao" gorm:"column:zhang_hao;comment:账号;size:100;"`
      ZhuangTai  *int `json:"zhuangTai" form:"zhuangTai" gorm:"column:zhuang_tai;comment:状态;size:10;"`
      ZiZhiZhengZhao  string `json:"ziZhiZhengZhao" form:"ziZhiZhengZhao" gorm:"column:zi_zhi_zheng_zhao;comment:资质证照;size:100;"`
}


// TableName GongYingShangBiao 表名
func (GongYingShangBiao) TableName() string {
  return "gong_ying_shang_biao"
}

