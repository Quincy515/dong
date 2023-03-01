package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hr"
	hrReq "github.com/flipped-aurora/gin-vue-admin/server/model/hr/request"
)

type MeiRiKaoQinService struct {
}

// CreateMeiRiKaoQin 创建MeiRiKaoQin记录
// Author [piexlmax](https://github.com/piexlmax)
func (meiRiKaoQinService *MeiRiKaoQinService) CreateMeiRiKaoQin(meiRiKaoQin hr.MeiRiKaoQin) (err error) {
	err = global.GVA_DB.Create(&meiRiKaoQin).Error
	return err
}

// DeleteMeiRiKaoQin 删除MeiRiKaoQin记录
// Author [piexlmax](https://github.com/piexlmax)
func (meiRiKaoQinService *MeiRiKaoQinService) DeleteMeiRiKaoQin(meiRiKaoQin hr.MeiRiKaoQin) (err error) {
	err = global.GVA_DB.Delete(&meiRiKaoQin).Error
	return err
}

// DeleteMeiRiKaoQinByIds 批量删除MeiRiKaoQin记录
// Author [piexlmax](https://github.com/piexlmax)
func (meiRiKaoQinService *MeiRiKaoQinService) DeleteMeiRiKaoQinByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]hr.MeiRiKaoQin{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMeiRiKaoQin 更新MeiRiKaoQin记录
// Author [piexlmax](https://github.com/piexlmax)
func (meiRiKaoQinService *MeiRiKaoQinService) UpdateMeiRiKaoQin(meiRiKaoQin hr.MeiRiKaoQin) (err error) {
	if meiRiKaoQin.CreatedAt.IsZero() && meiRiKaoQin.ID != 0 {
		var temp hr.MeiRiKaoQin
		err = global.GVA_DB.Where("id = ?", meiRiKaoQin.ID).First(&temp).Error
		if err != nil {
			return err
		}
		meiRiKaoQin.CreatedAt = temp.CreatedAt
	}
	err = global.GVA_DB.Save(&meiRiKaoQin).Error
	return err
}

// GetMeiRiKaoQin 根据id获取MeiRiKaoQin记录
// Author [piexlmax](https://github.com/piexlmax)
func (meiRiKaoQinService *MeiRiKaoQinService) GetMeiRiKaoQin(id uint) (meiRiKaoQin hr.MeiRiKaoQin, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&meiRiKaoQin).Error
	return
}

// GetMeiRiKaoQinInfoList 分页获取MeiRiKaoQin记录
// Author [piexlmax](https://github.com/piexlmax)
func (meiRiKaoQinService *MeiRiKaoQinService) GetMeiRiKaoQinInfoList(info hrReq.MeiRiKaoQinSearch) (list []hr.MeiRiKaoQin, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hr.MeiRiKaoQin{})
	var meiRiKaoQins []hr.MeiRiKaoQin
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.XingMing != "" {
		db = db.Where("xing_ming LIKE ?", "%"+info.XingMing+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("ID DESC").Limit(limit).Offset(offset).Find(&meiRiKaoQins).Error
	return meiRiKaoQins, total, err
}
