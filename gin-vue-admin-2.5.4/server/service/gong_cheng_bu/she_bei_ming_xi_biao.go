package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu"
	gong_cheng_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu/request"
)

type SheBeiMingXiBiaoService struct {
}

// CreateSheBeiMingXiBiao 创建SheBeiMingXiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (sheBeiMingXiBiaoService *SheBeiMingXiBiaoService) CreateSheBeiMingXiBiao(sheBeiMingXiBiao gong_cheng_bu.SheBeiMingXiBiao) (err error) {
	err = global.GVA_DB.Create(&sheBeiMingXiBiao).Error
	return err
}

// DeleteSheBeiMingXiBiao 删除SheBeiMingXiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (sheBeiMingXiBiaoService *SheBeiMingXiBiaoService) DeleteSheBeiMingXiBiao(sheBeiMingXiBiao gong_cheng_bu.SheBeiMingXiBiao) (err error) {
	err = global.GVA_DB.Delete(&sheBeiMingXiBiao).Error
	return err
}

// DeleteSheBeiMingXiBiaoByIds 批量删除SheBeiMingXiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (sheBeiMingXiBiaoService *SheBeiMingXiBiaoService) DeleteSheBeiMingXiBiaoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]gong_cheng_bu.SheBeiMingXiBiao{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSheBeiMingXiBiao 更新SheBeiMingXiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (sheBeiMingXiBiaoService *SheBeiMingXiBiaoService) UpdateSheBeiMingXiBiao(sheBeiMingXiBiao gong_cheng_bu.SheBeiMingXiBiao) (err error) {
	err = global.GVA_DB.Save(&sheBeiMingXiBiao).Error
	return err
}

// GetSheBeiMingXiBiao 根据id获取SheBeiMingXiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (sheBeiMingXiBiaoService *SheBeiMingXiBiaoService) GetSheBeiMingXiBiao(id uint) (sheBeiMingXiBiao gong_cheng_bu.SheBeiMingXiBiao, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sheBeiMingXiBiao).Error
	return
}

// GetSheBeiMingXiBiaoInfoList 分页获取SheBeiMingXiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (sheBeiMingXiBiaoService *SheBeiMingXiBiaoService) GetSheBeiMingXiBiaoInfoList(info gong_cheng_buReq.SheBeiMingXiBiaoSearch) (list []gong_cheng_bu.SheBeiMingXiBiao, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&gong_cheng_bu.SheBeiMingXiBiao{})
	var sheBeiMingXiBiaos []gong_cheng_bu.SheBeiMingXiBiao
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ZhuangTai != nil {
		db = db.Where("zhuang_tai = ?", info.ZhuangTai)
	}
	if info.ZuLinDanJuId != nil {
		db = db.Where("zu_lin_dan_ju_id = ?", info.ZuLinDanJuId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&sheBeiMingXiBiaos).Error
	return sheBeiMingXiBiaos, total, err
}
