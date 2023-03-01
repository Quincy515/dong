package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hr"
	hrReq "github.com/flipped-aurora/gin-vue-admin/server/model/hr/request"
)

type QingJiaGuanLiService struct {
}

// CreateQingJiaGuanLi 创建QingJiaGuanLi记录
// Author [piexlmax](https://github.com/piexlmax)
func (qingJiaGuanLiService *QingJiaGuanLiService) CreateQingJiaGuanLi(qingJiaGuanLi hr.QingJiaGuanLi) (err error) {
	err = global.GVA_DB.Create(&qingJiaGuanLi).Error
	return err
}

// DeleteQingJiaGuanLi 删除QingJiaGuanLi记录
// Author [piexlmax](https://github.com/piexlmax)
func (qingJiaGuanLiService *QingJiaGuanLiService) DeleteQingJiaGuanLi(qingJiaGuanLi hr.QingJiaGuanLi) (err error) {
	err = global.GVA_DB.Delete(&qingJiaGuanLi).Error
	return err
}

// DeleteQingJiaGuanLiByIds 批量删除QingJiaGuanLi记录
// Author [piexlmax](https://github.com/piexlmax)
func (qingJiaGuanLiService *QingJiaGuanLiService) DeleteQingJiaGuanLiByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]hr.QingJiaGuanLi{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateQingJiaGuanLi 更新QingJiaGuanLi记录
// Author [piexlmax](https://github.com/piexlmax)
func (qingJiaGuanLiService *QingJiaGuanLiService) UpdateQingJiaGuanLi(qingJiaGuanLi hr.QingJiaGuanLi) (err error) {
	if qingJiaGuanLi.CreatedAt.IsZero() && qingJiaGuanLi.ID != 0 {
		var temp hr.QingJiaGuanLi
		err = global.GVA_DB.Where("id = ?", qingJiaGuanLi.ID).First(&temp).Error
		if err != nil {
			return err
		}
		qingJiaGuanLi.CreatedAt = temp.CreatedAt
	}
	err = global.GVA_DB.Save(&qingJiaGuanLi).Error
	return err
}

// GetQingJiaGuanLi 根据id获取QingJiaGuanLi记录
// Author [piexlmax](https://github.com/piexlmax)
func (qingJiaGuanLiService *QingJiaGuanLiService) GetQingJiaGuanLi(id uint) (qingJiaGuanLi hr.QingJiaGuanLi, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&qingJiaGuanLi).Error
	return
}

// GetQingJiaGuanLiInfoList 分页获取QingJiaGuanLi记录
// Author [piexlmax](https://github.com/piexlmax)
func (qingJiaGuanLiService *QingJiaGuanLiService) GetQingJiaGuanLiInfoList(info hrReq.QingJiaGuanLiSearch) (list []hr.QingJiaGuanLi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hr.QingJiaGuanLi{})
	var qingJiaGuanLis []hr.QingJiaGuanLi
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Group != "" {
		db = db.Where("group = ?", info.Group)
	}
	if info.JinDu != "" {
		db = db.Where("jin_du = ?", info.JinDu)
	}
	if info.LeiXing != nil {
		db = db.Where("lei_xing = ?", info.LeiXing)
	}
	if info.XingMing != "" {
		db = db.Where("xing_ming = ?", info.XingMing)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("ID DESC").Limit(limit).Offset(offset).Find(&qingJiaGuanLis).Error
	return qingJiaGuanLis, total, err
}
