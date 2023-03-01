package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hr"
	hrReq "github.com/flipped-aurora/gin-vue-admin/server/model/hr/request"
)

type GongZhangDengJiService struct {
}

// CreateGongZhangDengJi 创建GongZhangDengJi记录
// Author [piexlmax](https://github.com/piexlmax)
func (gongZhangDengJiService *GongZhangDengJiService) CreateGongZhangDengJi(gongZhangDengJi hr.GongZhangDengJi) (err error) {
	err = global.GVA_DB.Create(&gongZhangDengJi).Error
	return err
}

// DeleteGongZhangDengJi 删除GongZhangDengJi记录
// Author [piexlmax](https://github.com/piexlmax)
func (gongZhangDengJiService *GongZhangDengJiService) DeleteGongZhangDengJi(gongZhangDengJi hr.GongZhangDengJi) (err error) {
	err = global.GVA_DB.Delete(&gongZhangDengJi).Error
	return err
}

// DeleteGongZhangDengJiByIds 批量删除GongZhangDengJi记录
// Author [piexlmax](https://github.com/piexlmax)
func (gongZhangDengJiService *GongZhangDengJiService) DeleteGongZhangDengJiByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]hr.GongZhangDengJi{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateGongZhangDengJi 更新GongZhangDengJi记录
// Author [piexlmax](https://github.com/piexlmax)
func (gongZhangDengJiService *GongZhangDengJiService) UpdateGongZhangDengJi(gongZhangDengJi hr.GongZhangDengJi) (err error) {
	if gongZhangDengJi.CreatedAt.IsZero() && gongZhangDengJi.ID != 0 {
		var temp hr.GongZhangDengJi
		err = global.GVA_DB.Where("id = ?", gongZhangDengJi.ID).First(&temp).Error
		if err != nil {
			return err
		}
		gongZhangDengJi.CreatedAt = temp.CreatedAt
	}
	err = global.GVA_DB.Save(&gongZhangDengJi).Error
	return err
}

// GetGongZhangDengJi 根据id获取GongZhangDengJi记录
// Author [piexlmax](https://github.com/piexlmax)
func (gongZhangDengJiService *GongZhangDengJiService) GetGongZhangDengJi(id uint) (gongZhangDengJi hr.GongZhangDengJi, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&gongZhangDengJi).Error
	return
}

// GetGongZhangDengJiInfoList 分页获取GongZhangDengJi记录
// Author [piexlmax](https://github.com/piexlmax)
func (gongZhangDengJiService *GongZhangDengJiService) GetGongZhangDengJiInfoList(info hrReq.GongZhangDengJiSearch) (list []hr.GongZhangDengJi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hr.GongZhangDengJi{})
	var gongZhangDengJis []hr.GongZhangDengJi
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.GongZhangMingCheng != "" {
		db = db.Where("gong_zhang_ming_cheng LIKE ?", "%"+info.GongZhangMingCheng+"%")
	}
	if info.Group != "" {
		db = db.Where("group = ?", info.Group)
	}
	if info.ShenQingRen != "" {
		db = db.Where("shen_qing_ren LIKE ?", "%"+info.ShenQingRen+"%")
	}
	if info.ShiXiang != "" {
		db = db.Where("shi_xiang LIKE ?", "%"+info.ShiXiang+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("ID DESC").Limit(limit).Offset(offset).Find(&gongZhangDengJis).Error
	return gongZhangDengJis, total, err
}
