package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hr"
	hrReq "github.com/flipped-aurora/gin-vue-admin/server/model/hr/request"
)

type XiaoXiTongZhiService struct {
}

// CreateXiaoXiTongZhi 创建XiaoXiTongZhi记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiaoXiTongZhiService *XiaoXiTongZhiService) CreateXiaoXiTongZhi(xiaoXiTongZhi hr.XiaoXiTongZhi) (err error) {
	err = global.GVA_DB.Create(&xiaoXiTongZhi).Error
	return err
}

// DeleteXiaoXiTongZhi 删除XiaoXiTongZhi记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiaoXiTongZhiService *XiaoXiTongZhiService) DeleteXiaoXiTongZhi(xiaoXiTongZhi hr.XiaoXiTongZhi) (err error) {
	err = global.GVA_DB.Delete(&xiaoXiTongZhi).Error
	return err
}

// DeleteXiaoXiTongZhiByIds 批量删除XiaoXiTongZhi记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiaoXiTongZhiService *XiaoXiTongZhiService) DeleteXiaoXiTongZhiByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]hr.XiaoXiTongZhi{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateXiaoXiTongZhi 更新XiaoXiTongZhi记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiaoXiTongZhiService *XiaoXiTongZhiService) UpdateXiaoXiTongZhi(xiaoXiTongZhi hr.XiaoXiTongZhi) (err error) {
	err = global.GVA_DB.Save(&xiaoXiTongZhi).Error
	return err
}

// GetXiaoXiTongZhi 根据id获取XiaoXiTongZhi记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiaoXiTongZhiService *XiaoXiTongZhiService) GetXiaoXiTongZhi(id uint) (xiaoXiTongZhi hr.XiaoXiTongZhi, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&xiaoXiTongZhi).Error
	return
}

// GetXiaoXiTongZhiInfoList 分页获取XiaoXiTongZhi记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiaoXiTongZhiService *XiaoXiTongZhiService) GetXiaoXiTongZhiInfoList(info hrReq.XiaoXiTongZhiSearch) (list []hr.XiaoXiTongZhi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hr.XiaoXiTongZhi{})
	var xiaoXiTongZhis []hr.XiaoXiTongZhi
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ReceiverId != nil {
		db = db.Where("receiver_id = ?", info.ReceiverId)
	}
	if info.SenderId != nil {
		db = db.Where("sender_id = ?", info.SenderId)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&xiaoXiTongZhis).Error
	return xiaoXiTongZhis, total, err
}
