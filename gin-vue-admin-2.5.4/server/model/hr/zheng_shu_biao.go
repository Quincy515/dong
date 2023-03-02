// 自动生成模板ZhengShuBiao
package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// ZhengShuBiao 结构体
type ZhengShuBiao struct {
      global.GVA_MODEL
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      ZhengShuMingCheng  string `json:"zhengShuMingCheng" form:"zhengShuMingCheng" gorm:"column:zheng_shu_ming_cheng;comment:证书名称;size:100;"`
      ZhengShuBianHao  string `json:"zhengShuBianHao" form:"zhengShuBianHao" gorm:"column:zheng_shu_bian_hao;comment:证书编号;size:100;"`
      ZhengShuLeiBie  *int `json:"zhengShuLeiBie" form:"zhengShuLeiBie" gorm:"column:zheng_shu_lei_bie;comment:证书类别;size:10;"`
      SuoYouRenXingMing  string `json:"suoYouRenXingMing" form:"suoYouRenXingMing" gorm:"column:suo_you_ren_xing_ming;comment:所有人姓名;size:100;"`
      ShenFenZhengHao  string `json:"shenFenZhengHao" form:"shenFenZhengHao" gorm:"column:shen_fen_zheng_hao;comment:身份证号;size:100;"`
      ShouJiHaoMa  string `json:"shouJiHaoMa" form:"shouJiHaoMa" gorm:"column:shou_ji_hao_ma;comment:手机号码;size:100;"`
      FaZhengBuMen  string `json:"faZhengBuMen" form:"faZhengBuMen" gorm:"column:fa_zheng_bu_men;comment:发证部门;size:100;"`
      ZhengShuSuoShuChengShi  *int `json:"zhengShuSuoShuChengShi" form:"zhengShuSuoShuChengShi" gorm:"column:zheng_shu_suo_shu_cheng_shi;comment:证书所属城市;size:10;"`
      ZhengShuZhuangTai  *int `json:"zhengShuZhuangTai" form:"zhengShuZhuangTai" gorm:"column:zheng_shu_zhuang_tai;comment:证书状态;size:10;"`
      ZhengShuDaoQiRiQi  *time.Time `json:"zhengShuDaoQiRiQi" form:"zhengShuDaoQiRiQi" gorm:"column:zheng_shu_dao_qi_ri_qi;comment:证书到期日期;"`
      ZhengShuNianShiYongFei  string `json:"zhengShuNianShiYongFei" form:"zhengShuNianShiYongFei" gorm:"column:zheng_shu_nian_shi_yong_fei;comment:证书年使用费;size:100;"`
      ZhiFuRiQi  *time.Time `json:"zhiFuRiQi" form:"zhiFuRiQi" gorm:"column:zhi_fu_ri_qi;comment:支付日期;"`
      ShiFouJieChu  *bool `json:"shiFouJieChu" form:"shiFouJieChu" gorm:"column:shi_fou_jie_chu;comment:是否借出;"`
      TianJiaRen  *int `json:"tianJiaRen" form:"tianJiaRen" gorm:"column:tian_jia_ren;comment:添加人;size:10;"`
      BeiZhu  string `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
}


// TableName ZhengShuBiao 表名
func (ZhengShuBiao) TableName() string {
  return "zheng_shu_biao"
}

