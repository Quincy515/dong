package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu"
	gong_cheng_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu/request"
)

type WenJianZiLiaoJieDianBiaoService struct {
}

// CreateWenJianZiLiaoJieDianBiao 创建WenJianZiLiaoJieDianBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (wenJianZiLiaoJieDianBiaoService *WenJianZiLiaoJieDianBiaoService) CreateWenJianZiLiaoJieDianBiao(wenJianZiLiaoJieDianBiao gong_cheng_bu.WenJianZiLiaoJieDianBiao) (err error) {
	err = global.GVA_DB.Create(&wenJianZiLiaoJieDianBiao).Error
	return err
}

// DeleteWenJianZiLiaoJieDianBiao 删除WenJianZiLiaoJieDianBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (wenJianZiLiaoJieDianBiaoService *WenJianZiLiaoJieDianBiaoService) DeleteWenJianZiLiaoJieDianBiao(wenJianZiLiaoJieDianBiao gong_cheng_bu.WenJianZiLiaoJieDianBiao) (err error) {
	err = global.GVA_DB.Delete(&wenJianZiLiaoJieDianBiao).Error
	return err
}

// DeleteWenJianZiLiaoJieDianBiaoByIds 批量删除WenJianZiLiaoJieDianBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (wenJianZiLiaoJieDianBiaoService *WenJianZiLiaoJieDianBiaoService) DeleteWenJianZiLiaoJieDianBiaoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]gong_cheng_bu.WenJianZiLiaoJieDianBiao{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateWenJianZiLiaoJieDianBiao 更新WenJianZiLiaoJieDianBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (wenJianZiLiaoJieDianBiaoService *WenJianZiLiaoJieDianBiaoService) UpdateWenJianZiLiaoJieDianBiao(wenJianZiLiaoJieDianBiao gong_cheng_bu.WenJianZiLiaoJieDianBiao) (err error) {
	err = global.GVA_DB.Save(&wenJianZiLiaoJieDianBiao).Error
	return err
}

// GetWenJianZiLiaoJieDianBiao 根据id获取WenJianZiLiaoJieDianBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (wenJianZiLiaoJieDianBiaoService *WenJianZiLiaoJieDianBiaoService) GetWenJianZiLiaoJieDianBiao(id uint) (wenJianZiLiaoJieDianBiao gong_cheng_bu.WenJianZiLiaoJieDianBiao, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&wenJianZiLiaoJieDianBiao).Error
	return
}

// GetWenJianZiLiaoJieDianBiaoInfoList 分页获取WenJianZiLiaoJieDianBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (wenJianZiLiaoJieDianBiaoService *WenJianZiLiaoJieDianBiaoService) GetWenJianZiLiaoJieDianBiaoInfoList(info gong_cheng_buReq.WenJianZiLiaoJieDianBiaoSearch) (list []gong_cheng_bu.WenJianZiLiaoJieDianBiao, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&gong_cheng_bu.WenJianZiLiaoJieDianBiao{})
	var wenJianZiLiaoJieDianBiaos []gong_cheng_bu.WenJianZiLiaoJieDianBiao
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.FuJieDianId != nil {
		db = db.Where("fu_jie_dian_id = ?", info.FuJieDianId)
	}
	if info.ShiFouCunDang != nil {
		db = db.Where("shi_fou_cun_dang = ?", info.ShiFouCunDang)
	}
	if info.WenJianJiaMingCheng != "" {
		db = db.Where("wen_jian_jia_ming_cheng LIKE ?", "%"+info.WenJianJiaMingCheng+"%")
	}
	if info.XiangMuId != nil {
		db = db.Where("xiang_mu_id = ?", info.XiangMuId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&wenJianZiLiaoJieDianBiaos).Error
	return wenJianZiLiaoJieDianBiaos, total, err
}
