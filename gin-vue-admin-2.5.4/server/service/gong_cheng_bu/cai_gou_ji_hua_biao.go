package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu"
	gong_cheng_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu/request"
)

type CaiGouJiHuaBiaoService struct {
}

// CreateCaiGouJiHuaBiao 创建CaiGouJiHuaBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (caiGouJiHuaBiaoService *CaiGouJiHuaBiaoService) CreateCaiGouJiHuaBiao(caiGouJiHuaBiao gong_cheng_bu.CaiGouJiHuaBiao) (err error) {
	err = global.GVA_DB.Create(&caiGouJiHuaBiao).Error
	return err
}

// DeleteCaiGouJiHuaBiao 删除CaiGouJiHuaBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (caiGouJiHuaBiaoService *CaiGouJiHuaBiaoService) DeleteCaiGouJiHuaBiao(caiGouJiHuaBiao gong_cheng_bu.CaiGouJiHuaBiao) (err error) {
	err = global.GVA_DB.Delete(&caiGouJiHuaBiao).Error
	return err
}

// DeleteCaiGouJiHuaBiaoByIds 批量删除CaiGouJiHuaBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (caiGouJiHuaBiaoService *CaiGouJiHuaBiaoService) DeleteCaiGouJiHuaBiaoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]gong_cheng_bu.CaiGouJiHuaBiao{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCaiGouJiHuaBiao 更新CaiGouJiHuaBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (caiGouJiHuaBiaoService *CaiGouJiHuaBiaoService) UpdateCaiGouJiHuaBiao(caiGouJiHuaBiao gong_cheng_bu.CaiGouJiHuaBiao) (err error) {
	err = global.GVA_DB.Save(&caiGouJiHuaBiao).Error
	return err
}

// GetCaiGouJiHuaBiao 根据id获取CaiGouJiHuaBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (caiGouJiHuaBiaoService *CaiGouJiHuaBiaoService) GetCaiGouJiHuaBiao(id uint) (caiGouJiHuaBiao gong_cheng_bu.CaiGouJiHuaBiao, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&caiGouJiHuaBiao).Error
	return
}

// GetCaiGouJiHuaBiaoInfoList 分页获取CaiGouJiHuaBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (caiGouJiHuaBiaoService *CaiGouJiHuaBiaoService) GetCaiGouJiHuaBiaoInfoList(info gong_cheng_buReq.CaiGouJiHuaBiaoSearch) (list []gong_cheng_bu.CaiGouJiHuaBiao, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&gong_cheng_bu.CaiGouJiHuaBiao{})
	var caiGouJiHuaBiaos []gong_cheng_bu.CaiGouJiHuaBiao
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CaiGouFuZeRen != nil {
		db = db.Where("cai_gou_fu_ze_ren = ?", info.CaiGouFuZeRen)
	}
	if info.XiangMuId != nil {
		db = db.Where("xiang_mu_id = ?", info.XiangMuId)
	}
	if info.ZhuangTai != nil {
		db = db.Where("zhuang_tai = ?", info.ZhuangTai)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&caiGouJiHuaBiaos).Error
	return caiGouJiHuaBiaos, total, err
}
