package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hr"
	hrReq "github.com/flipped-aurora/gin-vue-admin/server/model/hr/request"
)

type GongSiZhangHaoMiMaService struct {
}

// CreateGongSiZhangHaoMiMa 创建GongSiZhangHaoMiMa记录
// Author [piexlmax](https://github.com/piexlmax)
func (gongSiZhangHaoMiMaService *GongSiZhangHaoMiMaService) CreateGongSiZhangHaoMiMa(gongSiZhangHaoMiMa hr.GongSiZhangHaoMiMa) (err error) {
	err = global.GVA_DB.Create(&gongSiZhangHaoMiMa).Error
	return err
}

// DeleteGongSiZhangHaoMiMa 删除GongSiZhangHaoMiMa记录
// Author [piexlmax](https://github.com/piexlmax)
func (gongSiZhangHaoMiMaService *GongSiZhangHaoMiMaService) DeleteGongSiZhangHaoMiMa(gongSiZhangHaoMiMa hr.GongSiZhangHaoMiMa) (err error) {
	err = global.GVA_DB.Delete(&gongSiZhangHaoMiMa).Error
	return err
}

// DeleteGongSiZhangHaoMiMaByIds 批量删除GongSiZhangHaoMiMa记录
// Author [piexlmax](https://github.com/piexlmax)
func (gongSiZhangHaoMiMaService *GongSiZhangHaoMiMaService) DeleteGongSiZhangHaoMiMaByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]hr.GongSiZhangHaoMiMa{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateGongSiZhangHaoMiMa 更新GongSiZhangHaoMiMa记录
// Author [piexlmax](https://github.com/piexlmax)
func (gongSiZhangHaoMiMaService *GongSiZhangHaoMiMaService) UpdateGongSiZhangHaoMiMa(gongSiZhangHaoMiMa hr.GongSiZhangHaoMiMa) (err error) {
	if gongSiZhangHaoMiMa.CreatedAt.IsZero() && gongSiZhangHaoMiMa.ID != 0 {
		var temp hr.GongSiZhangHaoMiMa
		err = global.GVA_DB.Where("id = ?", gongSiZhangHaoMiMa.ID).First(&temp).Error
		if err != nil {
			return err
		}
		gongSiZhangHaoMiMa.CreatedAt = temp.CreatedAt
	}
	err = global.GVA_DB.Save(&gongSiZhangHaoMiMa).Error
	return err
}

// GetGongSiZhangHaoMiMa 根据id获取GongSiZhangHaoMiMa记录
// Author [piexlmax](https://github.com/piexlmax)
func (gongSiZhangHaoMiMaService *GongSiZhangHaoMiMaService) GetGongSiZhangHaoMiMa(id uint) (gongSiZhangHaoMiMa hr.GongSiZhangHaoMiMa, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&gongSiZhangHaoMiMa).Error
	return
}

// GetGongSiZhangHaoMiMaInfoList 分页获取GongSiZhangHaoMiMa记录
// Author [piexlmax](https://github.com/piexlmax)
func (gongSiZhangHaoMiMaService *GongSiZhangHaoMiMaService) GetGongSiZhangHaoMiMaInfoList(info hrReq.GongSiZhangHaoMiMaSearch) (list []hr.GongSiZhangHaoMiMa, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hr.GongSiZhangHaoMiMa{})
	var gongSiZhangHaoMiMas []hr.GongSiZhangHaoMiMa
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Group != "" {
		db = db.Where("group = ?", info.Group)
	}
	if info.YongHuMing != "" {
		db = db.Where("yong_hu_ming LIKE ?", "%"+info.YongHuMing+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("ID DESC").Limit(limit).Offset(offset).Find(&gongSiZhangHaoMiMas).Error
	return gongSiZhangHaoMiMas, total, err
}
