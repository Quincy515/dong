// 自动生成模板WenJianZiLiaoBiao
package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// WenJianZiLiaoBiao 结构体
type WenJianZiLiaoBiao struct {
      global.GVA_MODEL
      BeiZhu  string `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      JieDianId  *int `json:"jieDianId" form:"jieDianId" gorm:"column:jie_dian_id;comment:节点ID只保存最近的节点;size:10;"`
      ShangChuanRen  *int `json:"shangChuanRen" form:"shangChuanRen" gorm:"column:shang_chuan_ren;comment:上传人;size:10;"`
      ShangChuanRiQi  *time.Time `json:"shangChuanRiQi" form:"shangChuanRiQi" gorm:"column:shang_chuan_ri_qi;comment:上传日期;"`
      ShenHeRen  *int `json:"shenHeRen" form:"shenHeRen" gorm:"column:shen_he_ren;comment:审核人;size:10;"`
      ShenHeRiQi  *time.Time `json:"shenHeRiQi" form:"shenHeRiQi" gorm:"column:shen_he_ri_qi;comment:审核日期;"`
      ShiFouShenHe  *int `json:"shiFouShenHe" form:"shiFouShenHe" gorm:"column:shi_fou_shen_he;comment:是否审核;size:10;"`
      WenJianId  *int `json:"wenJianId" form:"wenJianId" gorm:"column:wen_jian_id;comment:文件id;size:10;"`
      WenJianLeiXing  *int `json:"wenJianLeiXing" form:"wenJianLeiXing" gorm:"column:wen_jian_lei_xing;comment:文件类型;size:10;"`
      WenJianMingCheng  string `json:"wenJianMingCheng" form:"wenJianMingCheng" gorm:"column:wen_jian_ming_cheng;comment:文件名称;size:100;"`
      YueDuCiShu  *int `json:"yueDuCiShu" form:"yueDuCiShu" gorm:"column:yue_du_ci_shu;comment:阅读次数;size:10;"`
      ZhuangTai  *int `json:"zhuangTai" form:"zhuangTai" gorm:"column:zhuang_tai;comment:状态;size:10;"`
}


// TableName WenJianZiLiaoBiao 表名
func (WenJianZiLiaoBiao) TableName() string {
  return "wen_jian_zi_liao_biao"
}

