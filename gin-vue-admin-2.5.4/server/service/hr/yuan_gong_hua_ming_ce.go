package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hr"
	hrReq "github.com/flipped-aurora/gin-vue-admin/server/model/hr/request"
)

type YuanGongHuaMingCeService struct {
}

// CreateYuanGongHuaMingCe 创建YuanGongHuaMingCe记录
// Author [piexlmax](https://github.com/piexlmax)
func (yuanGongHuaMingCeService *YuanGongHuaMingCeService) CreateYuanGongHuaMingCe(yuanGongHuaMingCe hr.YuanGongHuaMingCe) (err error) {
	err = global.GVA_DB.Create(&yuanGongHuaMingCe).Error
	return err
}

// DeleteYuanGongHuaMingCe 删除YuanGongHuaMingCe记录
// Author [piexlmax](https://github.com/piexlmax)
func (yuanGongHuaMingCeService *YuanGongHuaMingCeService) DeleteYuanGongHuaMingCe(yuanGongHuaMingCe hr.YuanGongHuaMingCe) (err error) {
	err = global.GVA_DB.Delete(&yuanGongHuaMingCe).Error
	return err
}

// DeleteYuanGongHuaMingCeByIds 批量删除YuanGongHuaMingCe记录
// Author [piexlmax](https://github.com/piexlmax)
func (yuanGongHuaMingCeService *YuanGongHuaMingCeService) DeleteYuanGongHuaMingCeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]hr.YuanGongHuaMingCe{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateYuanGongHuaMingCe 更新YuanGongHuaMingCe记录
// Author [piexlmax](https://github.com/piexlmax)
func (yuanGongHuaMingCeService *YuanGongHuaMingCeService) UpdateYuanGongHuaMingCe(yuanGongHuaMingCe hr.YuanGongHuaMingCe) (err error) {
	err = global.GVA_DB.Save(&yuanGongHuaMingCe).Error
	return err
}

// GetYuanGongHuaMingCe 根据id获取YuanGongHuaMingCe记录
// Author [piexlmax](https://github.com/piexlmax)
func (yuanGongHuaMingCeService *YuanGongHuaMingCeService) GetYuanGongHuaMingCe(id uint) (yuanGongHuaMingCe hr.YuanGongHuaMingCe, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&yuanGongHuaMingCe).Error
	return
}

// GetYuanGongHuaMingCeInfoList 分页获取YuanGongHuaMingCe记录
// Author [piexlmax](https://github.com/piexlmax)
func (yuanGongHuaMingCeService *YuanGongHuaMingCeService) GetYuanGongHuaMingCeInfoList(info hrReq.YuanGongHuaMingCeSearch) (list []hr.YuanGongHuaMingCe, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hr.YuanGongHuaMingCe{})
	var yuanGongHuaMingCes []hr.YuanGongHuaMingCe
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
	err = db.Order("id DESC").Limit(limit).Offset(offset).Find(&yuanGongHuaMingCes).Error
	return yuanGongHuaMingCes, total, err
}
