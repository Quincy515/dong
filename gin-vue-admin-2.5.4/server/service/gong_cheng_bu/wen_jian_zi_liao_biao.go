package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu"
	gong_cheng_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu/request"
)

type WenJianZiLiaoBiaoService struct {
}

// CreateWenJianZiLiaoBiao 创建WenJianZiLiaoBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (wenJianZiLiaoBiaoService *WenJianZiLiaoBiaoService) CreateWenJianZiLiaoBiao(wenJianZiLiaoBiao gong_cheng_bu.WenJianZiLiaoBiao) (err error) {
	err = global.GVA_DB.Create(&wenJianZiLiaoBiao).Error
	return err
}

// DeleteWenJianZiLiaoBiao 删除WenJianZiLiaoBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (wenJianZiLiaoBiaoService *WenJianZiLiaoBiaoService) DeleteWenJianZiLiaoBiao(wenJianZiLiaoBiao gong_cheng_bu.WenJianZiLiaoBiao) (err error) {
	err = global.GVA_DB.Delete(&wenJianZiLiaoBiao).Error
	return err
}

// DeleteWenJianZiLiaoBiaoByIds 批量删除WenJianZiLiaoBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (wenJianZiLiaoBiaoService *WenJianZiLiaoBiaoService) DeleteWenJianZiLiaoBiaoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]gong_cheng_bu.WenJianZiLiaoBiao{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateWenJianZiLiaoBiao 更新WenJianZiLiaoBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (wenJianZiLiaoBiaoService *WenJianZiLiaoBiaoService) UpdateWenJianZiLiaoBiao(wenJianZiLiaoBiao gong_cheng_bu.WenJianZiLiaoBiao) (err error) {
	err = global.GVA_DB.Save(&wenJianZiLiaoBiao).Error
	return err
}

// GetWenJianZiLiaoBiao 根据id获取WenJianZiLiaoBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (wenJianZiLiaoBiaoService *WenJianZiLiaoBiaoService) GetWenJianZiLiaoBiao(id uint) (wenJianZiLiaoBiao gong_cheng_bu.WenJianZiLiaoBiao, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&wenJianZiLiaoBiao).Error
	return
}

// GetWenJianZiLiaoBiaoInfoList 分页获取WenJianZiLiaoBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (wenJianZiLiaoBiaoService *WenJianZiLiaoBiaoService) GetWenJianZiLiaoBiaoInfoList(info gong_cheng_buReq.WenJianZiLiaoBiaoSearch) (list []gong_cheng_bu.WenJianZiLiaoBiao, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&gong_cheng_bu.WenJianZiLiaoBiao{})
	var wenJianZiLiaoBiaos []gong_cheng_bu.WenJianZiLiaoBiao
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.JieDianId != nil {
		db = db.Where("jie_dian_id = ?", info.JieDianId)
	}
	if info.ShiFouShenHe != nil {
		db = db.Where("shi_fou_shen_he = ?", info.ShiFouShenHe)
	}
	if info.WenJianId != nil {
		db = db.Where("wen_jian_id = ?", info.WenJianId)
	}
	if info.WenJianMingCheng != "" {
		db = db.Where("wen_jian_ming_cheng LIKE ?", "%"+info.WenJianMingCheng+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&wenJianZiLiaoBiaos).Error
	return wenJianZiLiaoBiaos, total, err
}
