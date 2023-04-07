package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu"
	gong_cheng_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu/request"
)

type SheBeiZuLinBiaoService struct {
}

// CreateSheBeiZuLinBiao 创建SheBeiZuLinBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (sheBeiZuLinBiaoService *SheBeiZuLinBiaoService) CreateSheBeiZuLinBiao(sheBeiZuLinBiao gong_cheng_bu.SheBeiZuLinBiao) (err error) {
	err = global.GVA_DB.Create(&sheBeiZuLinBiao).Error
	return err
}

// DeleteSheBeiZuLinBiao 删除SheBeiZuLinBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (sheBeiZuLinBiaoService *SheBeiZuLinBiaoService) DeleteSheBeiZuLinBiao(sheBeiZuLinBiao gong_cheng_bu.SheBeiZuLinBiao) (err error) {
	err = global.GVA_DB.Delete(&sheBeiZuLinBiao).Error
	return err
}

// DeleteSheBeiZuLinBiaoByIds 批量删除SheBeiZuLinBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (sheBeiZuLinBiaoService *SheBeiZuLinBiaoService) DeleteSheBeiZuLinBiaoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]gong_cheng_bu.SheBeiZuLinBiao{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSheBeiZuLinBiao 更新SheBeiZuLinBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (sheBeiZuLinBiaoService *SheBeiZuLinBiaoService) UpdateSheBeiZuLinBiao(sheBeiZuLinBiao gong_cheng_bu.SheBeiZuLinBiao) (err error) {
	err = global.GVA_DB.Save(&sheBeiZuLinBiao).Error
	return err
}

// GetSheBeiZuLinBiao 根据id获取SheBeiZuLinBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (sheBeiZuLinBiaoService *SheBeiZuLinBiaoService) GetSheBeiZuLinBiao(id uint) (sheBeiZuLinBiao gong_cheng_bu.SheBeiZuLinBiao, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sheBeiZuLinBiao).Error
	return
}

// GetSheBeiZuLinBiaoInfoList 分页获取SheBeiZuLinBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (sheBeiZuLinBiaoService *SheBeiZuLinBiaoService) GetSheBeiZuLinBiaoInfoList(info gong_cheng_buReq.SheBeiZuLinBiaoSearch) (list []gong_cheng_bu.SheBeiZuLinBiao, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&gong_cheng_bu.SheBeiZuLinBiao{})
	var sheBeiZuLinBiaos []gong_cheng_bu.SheBeiZuLinBiao
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.XiangMuId != nil {
		db = db.Where("xiang_mu_id = ?", info.XiangMuId)
	}
	if info.ZhuangTai != nil {
		db = db.Where("zhuang_tai = ?", info.ZhuangTai)
	}
	if info.ZuLinId != nil {
		db = db.Where("zu_lin_id = ?", info.ZuLinId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&sheBeiZuLinBiaos).Error
	return sheBeiZuLinBiaos, total, err
}
