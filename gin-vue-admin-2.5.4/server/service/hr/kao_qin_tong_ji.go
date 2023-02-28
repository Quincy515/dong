package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hr"
	hrReq "github.com/flipped-aurora/gin-vue-admin/server/model/hr/request"
)

type KaoQinTongJiService struct {
}

// CreateKaoQinTongJi 创建KaoQinTongJi记录
// Author [piexlmax](https://github.com/piexlmax)
func (kaoQinTongJiService *KaoQinTongJiService) CreateKaoQinTongJi(kaoQinTongJi hr.KaoQinTongJi) (err error) {
	err = global.GVA_DB.Create(&kaoQinTongJi).Error
	return err
}

// DeleteKaoQinTongJi 删除KaoQinTongJi记录
// Author [piexlmax](https://github.com/piexlmax)
func (kaoQinTongJiService *KaoQinTongJiService) DeleteKaoQinTongJi(kaoQinTongJi hr.KaoQinTongJi) (err error) {
	err = global.GVA_DB.Delete(&kaoQinTongJi).Error
	return err
}

// DeleteKaoQinTongJiByIds 批量删除KaoQinTongJi记录
// Author [piexlmax](https://github.com/piexlmax)
func (kaoQinTongJiService *KaoQinTongJiService) DeleteKaoQinTongJiByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]hr.KaoQinTongJi{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateKaoQinTongJi 更新KaoQinTongJi记录
// Author [piexlmax](https://github.com/piexlmax)
func (kaoQinTongJiService *KaoQinTongJiService) UpdateKaoQinTongJi(kaoQinTongJi hr.KaoQinTongJi) (err error) {
	err = global.GVA_DB.Save(&kaoQinTongJi).Error
	return err
}

// GetKaoQinTongJi 根据id获取KaoQinTongJi记录
// Author [piexlmax](https://github.com/piexlmax)
func (kaoQinTongJiService *KaoQinTongJiService) GetKaoQinTongJi(id uint) (kaoQinTongJi hr.KaoQinTongJi, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&kaoQinTongJi).Error
	return
}

// GetKaoQinTongJiInfoList 分页获取KaoQinTongJi记录
// Author [piexlmax](https://github.com/piexlmax)
func (kaoQinTongJiService *KaoQinTongJiService) GetKaoQinTongJiInfoList(info hrReq.KaoQinTongJiSearch) (list []hr.KaoQinTongJi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hr.KaoQinTongJi{})
	var kaoQinTongJis []hr.KaoQinTongJi
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Group != "" {
		db = db.Where("group = ?", info.Group)
	}
	if info.XingMing != "" {
		db = db.Where("xing_ming LIKE ?", "%"+info.XingMing+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("ID DESC").Limit(limit).Offset(offset).Find(&kaoQinTongJis).Error
	return kaoQinTongJis, total, err
}
