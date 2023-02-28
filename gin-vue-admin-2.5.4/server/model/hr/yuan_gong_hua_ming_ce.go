// 自动生成模板YuanGongHuaMingCe
package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// YuanGongHuaMingCe 结构体
type YuanGongHuaMingCe struct {
      global.GVA_MODEL
      BeiZhu  string `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
      BuMen  string `json:"buMen" form:"buMen" gorm:"column:bu_men;comment:部门;size:255;"`
      ChuShengRiQi  string `json:"chuShengRiQi" form:"chuShengRiQi" gorm:"column:chu_sheng_ri_qi;comment:出生日期;size:255;"`
      CongShiGangWeiHuoGongZhong  string `json:"congShiGangWeiHuoGongZhong" form:"congShiGangWeiHuoGongZhong" gorm:"column:cong_shi_gang_wei_huo_gong_zhong;comment:从事岗位或工种;size:255;"`
      GeRenSheBaoMiMa  string `json:"geRenSheBaoMiMa" form:"geRenSheBaoMiMa" gorm:"column:ge_ren_she_bao_mi_ma;comment:个人社保密码;size:255;"`
      GeRenSheBaoZhangHao  string `json:"geRenSheBaoZhangHao" form:"geRenSheBaoZhangHao" gorm:"column:ge_ren_she_bao_zhang_hao;comment:个人社保账号;size:255;"`
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      HeTongYueDingGongZi  string `json:"heTongYueDingGongZi" form:"heTongYueDingGongZi" gorm:"column:he_tong_yue_ding_gong_zi;comment:合同约定工资;size:255;"`
      HuKouSuoZaiDi  string `json:"huKouSuoZaiDi" form:"huKouSuoZaiDi" gorm:"column:hu_kou_suo_zai_di;comment:户口所在地;size:255;"`
      LaoDongHeTongJieZhiRiQi  string `json:"laoDongHeTongJieZhiRiQi" form:"laoDongHeTongJieZhiRiQi" gorm:"column:lao_dong_he_tong_jie_zhi_ri_qi;comment:劳动合同截止日期;size:255;"`
      LaoDongHeTongKaiShiRiQi  string `json:"laoDongHeTongKaiShiRiQi" form:"laoDongHeTongKaiShiRiQi" gorm:"column:lao_dong_he_tong_kai_shi_ri_qi;comment:劳动合同开始日期;size:255;"`
      LiZhiRiQi  string `json:"liZhiRiQi" form:"liZhiRiQi" gorm:"column:li_zhi_ri_qi;comment:离职日期;size:255;"`
      LianXiDianHua  string `json:"lianXiDianHua" form:"lianXiDianHua" gorm:"column:lian_xi_dian_hua;comment:联系电话;size:255;"`
      MenJinBianHao  string `json:"menJinBianHao" form:"menJinBianHao" gorm:"column:men_jin_bian_hao;comment:门禁编号;size:5;"`
      NianLing  string `json:"nianLing" form:"nianLing" gorm:"column:nian_ling;comment:年龄;size:255;"`
      RuZhiShiJian  string `json:"ruZhiShiJian" form:"ruZhiShiJian" gorm:"column:ru_zhi_shi_jian;comment:入职时间;size:255;"`
      ShenFenZhengHaoMa  string `json:"shenFenZhengHaoMa" form:"shenFenZhengHaoMa" gorm:"column:shen_fen_zheng_hao_ma;comment:身份证号码;size:255;"`
      WenHuaChengDu  string `json:"wenHuaChengDu" form:"wenHuaChengDu" gorm:"column:wen_hua_cheng_du;comment:文化程度;size:255;"`
      XianJuZhuDiZhi  string `json:"xianJuZhuDiZhi" form:"xianJuZhuDiZhi" gorm:"column:xian_ju_zhu_di_zhi;comment:现居住地址;size:255;"`
      XingBie  string `json:"xingBie" form:"xingBie" gorm:"column:xing_bie;comment:性别;size:255;"`
      XingMing  string `json:"xingMing" form:"xingMing" gorm:"column:xing_ming;comment:姓名;size:255;"`
}


// TableName YuanGongHuaMingCe 表名
func (YuanGongHuaMingCe) TableName() string {
  return "yuan_gong_hua_ming_ce"
}

