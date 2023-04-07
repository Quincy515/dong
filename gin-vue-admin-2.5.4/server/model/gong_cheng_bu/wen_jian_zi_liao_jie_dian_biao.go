// 自动生成模板WenJianZiLiaoJieDianBiao
package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// WenJianZiLiaoJieDianBiao 结构体
type WenJianZiLiaoJieDianBiao struct {
      global.GVA_MODEL
      BeiZhu  string `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
      FuJieDianId  *int `json:"fuJieDianId" form:"fuJieDianId" gorm:"column:fu_jie_dian_id;comment:父节点多层级迭代寻祖;size:10;"`
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      PaiXu  *int `json:"paiXu" form:"paiXu" gorm:"column:pai_xu;comment:排序;size:10;"`
      ShiFouCunDang  *int `json:"shiFouCunDang" form:"shiFouCunDang" gorm:"column:shi_fou_cun_dang;comment:是否存档-存档后不可以变更文件;size:10;"`
      TianJiaRen  *int `json:"tianJiaRen" form:"tianJiaRen" gorm:"column:tian_jia_ren;comment:添加人;size:10;"`
      TianJiaShiJian  *time.Time `json:"tianJiaShiJian" form:"tianJiaShiJian" gorm:"column:tian_jia_shi_jian;comment:添加时间;"`
      WenJianJiaMingCheng  string `json:"wenJianJiaMingCheng" form:"wenJianJiaMingCheng" gorm:"column:wen_jian_jia_ming_cheng;comment:文件夹名称;size:200;"`
      XiangMuId  *int `json:"xiangMuId" form:"xiangMuId" gorm:"column:xiang_mu_id;comment:项目ID;size:10;"`
      ZhuangTai  *int `json:"zhuangTai" form:"zhuangTai" gorm:"column:zhuang_tai;comment:状态;size:10;"`
}


// TableName WenJianZiLiaoJieDianBiao 表名
func (WenJianZiLiaoJieDianBiao) TableName() string {
  return "wen_jian_zi_liao_jie_dian_biao"
}

