// 自动生成模板XiangMuBiao
package gong_cheng_bu

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// XiangMuBiao 结构体
type XiangMuBiao struct {
	global.GVA_MODEL
	Group                        string     `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
	XiangMuBianHao               *int       `json:"xiangMuBianHao" form:"xiangMuBianHao" gorm:"column:xiang_mu_bian_hao;comment:项目编号;size:10;"`
	JiHuaKaiGongRiQi             *time.Time `json:"jiHuaKaiGongRiQi" form:"jiHuaKaiGongRiQi" gorm:"column:ji_hua_kai_gong_ri_qi;comment:计划开工日期;"`
	JiHuaJieShuRiQi              *time.Time `json:"jiHuaJieShuRiQi" form:"jiHuaJieShuRiQi" gorm:"column:ji_hua_jie_shu_ri_qi;comment:计划结束日期;"`
	GongChengZaoJia              *float64   `json:"gongChengZaoJia" form:"gongChengZaoJia" gorm:"column:gong_cheng_zao_jia;comment:工程造价;size:10;"`
	NongMinGongGongZiBaoZhangJin *float64   `json:"nongMinGongGongZiBaoZhangJin" form:"nongMinGongGongZiBaoZhangJin" gorm:"column:nong_min_gong_gong_zi_bao_zhang_jin;comment:农民工工资保障金;size:10;"`
	BaoZhangJinZhuangTai         *int       `json:"baoZhangJinZhuangTai" form:"baoZhangJinZhuangTai" gorm:"column:bao_zhang_jin_zhuang_tai;comment:保障金状态;size:10;"`
	LuYueBaoZhengJin             *float64   `json:"luYueBaoZhengJin" form:"luYueBaoZhengJin" gorm:"column:lu_yue_bao_zheng_jin;comment:履约保证金;size:10;"`
	BaoZhengJinZhuangTai         *int       `json:"baoZhengJinZhuangTai" form:"baoZhengJinZhuangTai" gorm:"column:bao_zheng_jin_zhuang_tai;comment:保证金状态;size:10;"`
	GongZuoBaoXianJinE           *float64   `json:"gongZuoBaoXianJinE" form:"gongZuoBaoXianJinE" gorm:"column:gong_zuo_bao_xian_jin_e;comment:工作保险金额;size:10;"`
	FuKuanFangShi                *int       `json:"fuKuanFangShi" form:"fuKuanFangShi" gorm:"column:fu_kuan_fang_shi;comment:付款方式;size:10;"`
	YuQiLiRun                    *float64   `json:"yuQiLiRun" form:"yuQiLiRun" gorm:"column:yu_qi_li_run;comment:预期利润;size:10;"`
	ShiJiKaiGongRiQi             *time.Time `json:"shiJiKaiGongRiQi" form:"shiJiKaiGongRiQi" gorm:"column:shi_ji_kai_gong_ri_qi;comment:实际开工日期;"`
	ShiJiJieShuRiQi              *time.Time `json:"shiJiJieShuRiQi" form:"shiJiJieShuRiQi" gorm:"column:shi_ji_jie_shu_ri_qi;comment:实际结束日期;"`
	JieSuanZhuangTai             *int       `json:"jieSuanZhuangTai" form:"jieSuanZhuangTai" gorm:"column:jie_suan_zhuang_tai;comment:结算状态 1.未结算 2.结算完成;size:10;"`
	XiangMuZhuangTai             *int       `json:"xiangMuZhuangTai" form:"xiangMuZhuangTai" gorm:"column:xiang_mu_zhuang_tai;comment:项目状态 1.建设中 2.未完成 3.竣工;size:10;"`
	XiangMuJingLi                *int       `json:"xiangMuJingLi" form:"xiangMuJingLi" gorm:"column:xiang_mu_jing_li;comment:项目经理;size:10;"`
	XiangMuZhuGuan               *int       `json:"xiangMuZhuGuan" form:"xiangMuZhuGuan" gorm:"column:xiang_mu_zhu_guan;comment:项目主管;size:10;"`
	BeiZhu                       string     `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
	CreatedBy                    uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy                    uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy                    uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName XiangMuBiao 表名
func (XiangMuBiao) TableName() string {
	return "xiang_mu_biao"
}
