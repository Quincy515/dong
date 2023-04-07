package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu"
	gong_cheng_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu/request"
)

type XiangMuZhuCaiBiaoService struct {
}

// CreateXiangMuZhuCaiBiao 创建XiangMuZhuCaiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiangMuZhuCaiBiaoService *XiangMuZhuCaiBiaoService) CreateXiangMuZhuCaiBiao(xiangMuZhuCaiBiao gong_cheng_bu.XiangMuZhuCaiBiao) (err error) {
	err = global.GVA_DB.Create(&xiangMuZhuCaiBiao).Error
	return err
}

// DeleteXiangMuZhuCaiBiao 删除XiangMuZhuCaiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiangMuZhuCaiBiaoService *XiangMuZhuCaiBiaoService) DeleteXiangMuZhuCaiBiao(xiangMuZhuCaiBiao gong_cheng_bu.XiangMuZhuCaiBiao) (err error) {
	err = global.GVA_DB.Delete(&xiangMuZhuCaiBiao).Error
	return err
}

// DeleteXiangMuZhuCaiBiaoByIds 批量删除XiangMuZhuCaiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiangMuZhuCaiBiaoService *XiangMuZhuCaiBiaoService) DeleteXiangMuZhuCaiBiaoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]gong_cheng_bu.XiangMuZhuCaiBiao{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateXiangMuZhuCaiBiao 更新XiangMuZhuCaiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiangMuZhuCaiBiaoService *XiangMuZhuCaiBiaoService) UpdateXiangMuZhuCaiBiao(xiangMuZhuCaiBiao gong_cheng_bu.XiangMuZhuCaiBiao) (err error) {
	err = global.GVA_DB.Save(&xiangMuZhuCaiBiao).Error
	return err
}

// GetXiangMuZhuCaiBiao 根据id获取XiangMuZhuCaiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiangMuZhuCaiBiaoService *XiangMuZhuCaiBiaoService) GetXiangMuZhuCaiBiao(id uint) (xiangMuZhuCaiBiao gong_cheng_bu.XiangMuZhuCaiBiao, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&xiangMuZhuCaiBiao).Error
	return
}

// GetXiangMuZhuCaiBiaoInfoList 分页获取XiangMuZhuCaiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiangMuZhuCaiBiaoService *XiangMuZhuCaiBiaoService) GetXiangMuZhuCaiBiaoInfoList(info gong_cheng_buReq.XiangMuZhuCaiBiaoSearch) (list []gong_cheng_bu.XiangMuZhuCaiBiao, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&gong_cheng_bu.XiangMuZhuCaiBiao{})
	var xiangMuZhuCaiBiaos []gong_cheng_bu.XiangMuZhuCaiBiao
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.XiangMuMingCheng != nil {
		db = db.Where("xiang_mu_ming_cheng = ?", info.XiangMuMingCheng)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&xiangMuZhuCaiBiaos).Error
	return xiangMuZhuCaiBiaos, total, err
}
