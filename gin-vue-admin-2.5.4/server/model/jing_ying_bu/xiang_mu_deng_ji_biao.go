// 自动生成模板XiangMuDengJiBiao
package jing_ying_bu

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// XiangMuDengJiBiao 结构体
type XiangMuDengJiBiao struct {
	global.GVA_MODEL
	Group                             string     `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
	XiangMuMingCheng                  string     `json:"xiangMuMingCheng" form:"xiangMuMingCheng" gorm:"column:xiang_mu_ming_cheng;comment:项目名称;size:100;"`
	XiangMuJianJie                    string     `json:"xiangMuJianJie" form:"xiangMuJianJie" gorm:"column:xiang_mu_jian_jie;comment:项目简介;"`
	XiangMuDengJiJinDu                string     `json:"xiangMuDengJiJinDu" form:"xiangMuDengJiJinDu" gorm:"column:xiang_mu_deng_ji_jin_du;comment:项目登记进度;size:10;"`
	TouBiaoDanWeiMingCheng            string     `json:"touBiaoDanWeiMingCheng" form:"touBiaoDanWeiMingCheng" gorm:"column:tou_biao_dan_wei_ming_cheng;comment:投标单位名称;size:100;"`
	ZiZhiYaoQiu                       *int       `json:"ziZhiYaoQiu" form:"ziZhiYaoQiu" gorm:"column:zi_zhi_yao_qiu;comment:资质要求;size:10;"`
	DengJiYaoQiu                      *int       `json:"dengJiYaoQiu" form:"dengJiYaoQiu" gorm:"column:deng_ji_yao_qiu;comment:等级要求;size:10;"`
	KaiBiaoShiJian                    *time.Time `json:"kaiBiaoShiJian" form:"kaiBiaoShiJian" gorm:"column:kai_biao_shi_jian;comment:开标时间;"`
	KaiBiaoDiDian                     string     `json:"kaiBiaoDiDian" form:"kaiBiaoDiDian" gorm:"column:kai_biao_di_dian;comment:开标地点;size:200;"`
	SuoShuDiQu                        *int       `json:"suoShuDiQu" form:"suoShuDiQu" gorm:"column:suo_shu_di_qu;comment:所属地区;size:10;"`
	ZhaoBiaoGuiMo                     string     `json:"zhaoBiaoGuiMo" form:"zhaoBiaoGuiMo" gorm:"column:zhao_biao_gui_mo;comment:招标规模;size:20;"`
	KaiBiaoRenYuan                    string     `json:"kaiBiaoRenYuan" form:"kaiBiaoRenYuan" gorm:"column:kai_biao_ren_yuan;comment:开标人员;size:10;"`
	YeWuFuZeRen                       string     `json:"yeWuFuZeRen" form:"yeWuFuZeRen" gorm:"column:ye_wu_fu_ze_ren;comment:业务负责人;size:10;"`
	BaoZhengJinXingShi                *int       `json:"baoZhengJinXingShi" form:"baoZhengJinXingShi" gorm:"column:bao_zheng_jin_xing_shi;comment:保证金形式;size:10;"`
	BaoZhengJinJinE                   string     `json:"baoZhengJinJinE" form:"baoZhengJinJinE" gorm:"column:bao_zheng_jin_jin_e;comment:保证金金额;size:20;"`
	PingBiaoBanFa                     *int       `json:"pingBiaoBanFa" form:"pingBiaoBanFa" gorm:"column:ping_biao_ban_fa;comment:评标办法;size:10;"`
	XiangMuWangZhi                    string     `json:"xiangMuWangZhi" form:"xiangMuWangZhi" gorm:"column:xiang_mu_wang_zhi;comment:项目网址;size:200;"`
	BiaoShuLeiXing                    *int       `json:"biaoShuLeiXing" form:"biaoShuLeiXing" gorm:"column:biao_shu_lei_xing;comment:标书类型;size:10;"`
	XiangMuDengJiShiJian              *time.Time `json:"xiangMuDengJiShiJian" form:"xiangMuDengJiShiJian" gorm:"column:xiang_mu_deng_ji_shi_jian;comment:项目登记时间;"`
	BuMenShenHe                       string     `json:"buMenShenHe" form:"buMenShenHe" gorm:"column:bu_men_shen_he;comment:部门审核;size:10;"`
	BuMenShenHeShiJian                *time.Time `json:"buMenShenHeShiJian" form:"buMenShenHeShiJian" gorm:"column:bu_men_shen_he_shi_jian;comment:部门审核时间;"`
	CeOshenHe                         string     `json:"ceOshenHe" form:"ceOshenHe" gorm:"column:ce_oshen_he;comment:CEO审核;size:10;"`
	CeOshenHeShiJian                  *time.Time `json:"ceOshenHeShiJian" form:"ceOshenHeShiJian" gorm:"column:ce_oshen_he_shi_jian;comment:CEO审核时间;"`
	TouBiaoHanZhiZuoRen               string     `json:"touBiaoHanZhiZuoRen" form:"touBiaoHanZhiZuoRen" gorm:"column:tou_biao_han_zhi_zuo_ren;comment:投标函制作人;size:10;"`
	TouBiaoHanTiChengJinE             string     `json:"touBiaoHanTiChengJinE" form:"touBiaoHanTiChengJinE" gorm:"column:tou_biao_han_ti_cheng_jin_e;comment:投标函提成金额;size:20;"`
	TouBiaoHanZhiZuoRenQueRen         string     `json:"touBiaoHanZhiZuoRenQueRen" form:"touBiaoHanZhiZuoRenQueRen" gorm:"column:tou_biao_han_zhi_zuo_ren_que_ren;comment:投标函制作人确认;size:10;"`
	TouBiaoHanZhiZuoRenQueRenShiJian  *time.Time `json:"touBiaoHanZhiZuoRenQueRenShiJian" form:"touBiaoHanZhiZuoRenQueRenShiJian" gorm:"column:tou_biao_han_zhi_zuo_ren_que_ren_shi_jian;comment:投标函制作人确认时间;"`
	ShangWuBiaoZhiZuoRen              string     `json:"shangWuBiaoZhiZuoRen" form:"shangWuBiaoZhiZuoRen" gorm:"column:shang_wu_biao_zhi_zuo_ren;comment:商务标制作人;size:10;"`
	ShangWuBiaoTiChengJinE            string     `json:"shangWuBiaoTiChengJinE" form:"shangWuBiaoTiChengJinE" gorm:"column:shang_wu_biao_ti_cheng_jin_e;comment:商务标提成金额;size:20;"`
	ShangWuBiaoZhiZuoRenQueRen        string     `json:"shangWuBiaoZhiZuoRenQueRen" form:"shangWuBiaoZhiZuoRenQueRen" gorm:"column:shang_wu_biao_zhi_zuo_ren_que_ren;comment:商务标制作人确认;size:10;"`
	ShangWuBiaoZhiZuoRenQueRenShiJian *time.Time `json:"shangWuBiaoZhiZuoRenQueRenShiJian" form:"shangWuBiaoZhiZuoRenQueRenShiJian" gorm:"column:shang_wu_biao_zhi_zuo_ren_que_ren_shi_jian;comment:商务标制作人确认时间;"`
	JiShuBiaoZhiZuoRen                string     `json:"jiShuBiaoZhiZuoRen" form:"jiShuBiaoZhiZuoRen" gorm:"column:ji_shu_biao_zhi_zuo_ren;comment:技术标制作人;size:10;"`
	JiShuBiaoTiChengJinE              string     `json:"jiShuBiaoTiChengJinE" form:"jiShuBiaoTiChengJinE" gorm:"column:ji_shu_biao_ti_cheng_jin_e;comment:技术标提成金额;size:20;"`
	JiShuBiaoZhiZuoRenQueRen          string     `json:"jiShuBiaoZhiZuoRenQueRen" form:"jiShuBiaoZhiZuoRenQueRen" gorm:"column:ji_shu_biao_zhi_zuo_ren_que_ren;comment:技术标制作人确认;size:10;"`
	JiShuBiaoZhiZuoRenQueRenShiJian   *time.Time `json:"jiShuBiaoZhiZuoRenQueRenShiJian" form:"jiShuBiaoZhiZuoRenQueRenShiJian" gorm:"column:ji_shu_biao_zhi_zuo_ren_que_ren_shi_jian;comment:技术标制作人确认时间;"`
	XiangMuZhuangTai                  string     `json:"xiangMuZhuangTai" form:"xiangMuZhuangTai" gorm:"column:xiang_mu_zhuang_tai;comment:项目状态;size:5;"`
	BuMenShenQingXiaFuDianWei         string     `json:"buMenShenQingXiaFuDianWei" form:"buMenShenQingXiaFuDianWei" gorm:"column:bu_men_shen_qing_xia_fu_dian_wei;comment:部门申请下浮点位;size:10;"`
	BuMenXiaFuDianWeiShenQingShiJian  *time.Time `json:"buMenXiaFuDianWeiShenQingShiJian" form:"buMenXiaFuDianWeiShenQingShiJian" gorm:"column:bu_men_xia_fu_dian_wei_shen_qing_shi_jian;comment:部门下浮点位申请时间;"`
	CeOshenHeXiaFuDianWei             string     `json:"ceOshenHeXiaFuDianWei" form:"ceOshenHeXiaFuDianWei" gorm:"column:ce_oshen_he_xia_fu_dian_wei;comment:CEO审核下浮点位;size:10;"`
	CeOxiaFuDianWeiShenHeShiJian      *time.Time `json:"ceOxiaFuDianWeiShenHeShiJian" form:"ceOxiaFuDianWeiShenHeShiJian" gorm:"column:ce_oxia_fu_dian_wei_shen_he_shi_jian;comment:CEO下浮点位审核时间;"`
	ZhongBiaoJinE                     string     `json:"zhongBiaoJinE" form:"zhongBiaoJinE" gorm:"column:zhong_biao_jin_e;comment:中标金额;size:20;"`
	BeiZhu                            string     `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
}

// TableName XiangMuDengJiBiao 表名
func (XiangMuDengJiBiao) TableName() string {
	return "xiang_mu_deng_ji_biao"
}
