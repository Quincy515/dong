package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hr"
	hrReq "github.com/flipped-aurora/gin-vue-admin/server/model/hr/request"
)

type HuiYuanGuanLiService struct {
}

// CreateHuiYuanGuanLi 创建HuiYuanGuanLi记录
// Author [piexlmax](https://github.com/piexlmax)
func (huiYuanGuanLiService *HuiYuanGuanLiService) CreateHuiYuanGuanLi(huiYuanGuanLi hr.HuiYuanGuanLi) (err error) {
	err = global.GVA_DB.Create(&huiYuanGuanLi).Error
	return err
}

// DeleteHuiYuanGuanLi 删除HuiYuanGuanLi记录
// Author [piexlmax](https://github.com/piexlmax)
func (huiYuanGuanLiService *HuiYuanGuanLiService) DeleteHuiYuanGuanLi(huiYuanGuanLi hr.HuiYuanGuanLi) (err error) {
	err = global.GVA_DB.Delete(&huiYuanGuanLi).Error
	return err
}

// DeleteHuiYuanGuanLiByIds 批量删除HuiYuanGuanLi记录
// Author [piexlmax](https://github.com/piexlmax)
func (huiYuanGuanLiService *HuiYuanGuanLiService) DeleteHuiYuanGuanLiByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]hr.HuiYuanGuanLi{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHuiYuanGuanLi 更新HuiYuanGuanLi记录
// Author [piexlmax](https://github.com/piexlmax)
func (huiYuanGuanLiService *HuiYuanGuanLiService) UpdateHuiYuanGuanLi(huiYuanGuanLi hr.HuiYuanGuanLi) (err error) {
	err = global.GVA_DB.Save(&huiYuanGuanLi).Error
	return err
}

// GetHuiYuanGuanLi 根据id获取HuiYuanGuanLi记录
// Author [piexlmax](https://github.com/piexlmax)
func (huiYuanGuanLiService *HuiYuanGuanLiService) GetHuiYuanGuanLi(id uint) (huiYuanGuanLi hr.HuiYuanGuanLi, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&huiYuanGuanLi).Error
	return
}

// GetHuiYuanGuanLiInfoList 分页获取HuiYuanGuanLi记录
// Author [piexlmax](https://github.com/piexlmax)
func (huiYuanGuanLiService *HuiYuanGuanLiService) GetHuiYuanGuanLiInfoList(info hrReq.HuiYuanGuanLiSearch) (list []hr.HuiYuanGuanLi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hr.HuiYuanGuanLi{})
	var huiYuanGuanLis []hr.HuiYuanGuanLi
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Group != "" {
		db = db.Where("group = ?", info.Group)
	}
	if info.HuiYuanMingCheng != "" {
		db = db.Where("hui_yuan_ming_cheng LIKE ?", "%"+info.HuiYuanMingCheng+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("ID DESC").Limit(limit).Offset(offset).Find(&huiYuanGuanLis).Error
	return huiYuanGuanLis, total, err
}
