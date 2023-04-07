package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu"
	gong_cheng_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu/request"
)

type ChuRuKuMingXiBiaoService struct {
}

// CreateChuRuKuMingXiBiao 创建ChuRuKuMingXiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (chuRuKuMingXiBiaoService *ChuRuKuMingXiBiaoService) CreateChuRuKuMingXiBiao(chuRuKuMingXiBiao gong_cheng_bu.ChuRuKuMingXiBiao) (err error) {
	err = global.GVA_DB.Create(&chuRuKuMingXiBiao).Error
	return err
}

// DeleteChuRuKuMingXiBiao 删除ChuRuKuMingXiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (chuRuKuMingXiBiaoService *ChuRuKuMingXiBiaoService) DeleteChuRuKuMingXiBiao(chuRuKuMingXiBiao gong_cheng_bu.ChuRuKuMingXiBiao) (err error) {
	err = global.GVA_DB.Delete(&chuRuKuMingXiBiao).Error
	return err
}

// DeleteChuRuKuMingXiBiaoByIds 批量删除ChuRuKuMingXiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (chuRuKuMingXiBiaoService *ChuRuKuMingXiBiaoService) DeleteChuRuKuMingXiBiaoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]gong_cheng_bu.ChuRuKuMingXiBiao{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateChuRuKuMingXiBiao 更新ChuRuKuMingXiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (chuRuKuMingXiBiaoService *ChuRuKuMingXiBiaoService) UpdateChuRuKuMingXiBiao(chuRuKuMingXiBiao gong_cheng_bu.ChuRuKuMingXiBiao) (err error) {
	err = global.GVA_DB.Save(&chuRuKuMingXiBiao).Error
	return err
}

// GetChuRuKuMingXiBiao 根据id获取ChuRuKuMingXiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (chuRuKuMingXiBiaoService *ChuRuKuMingXiBiaoService) GetChuRuKuMingXiBiao(id uint) (chuRuKuMingXiBiao gong_cheng_bu.ChuRuKuMingXiBiao, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&chuRuKuMingXiBiao).Error
	return
}

// GetChuRuKuMingXiBiaoInfoList 分页获取ChuRuKuMingXiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (chuRuKuMingXiBiaoService *ChuRuKuMingXiBiaoService) GetChuRuKuMingXiBiaoInfoList(info gong_cheng_buReq.ChuRuKuMingXiBiaoSearch) (list []gong_cheng_bu.ChuRuKuMingXiBiao, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&gong_cheng_bu.ChuRuKuMingXiBiao{})
	var chuRuKuMingXiBiaos []gong_cheng_bu.ChuRuKuMingXiBiao
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CaiLiaoId != nil {
		db = db.Where("cai_liao_id = ?", info.CaiLiaoId)
	}
	if info.ChuRuKuDanJuId != nil {
		db = db.Where("chu_ru_ku_dan_ju_id = ?", info.ChuRuKuDanJuId)
	}
	if info.ZhuangTai != nil {
		db = db.Where("zhuang_tai = ?", info.ZhuangTai)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&chuRuKuMingXiBiaos).Error
	return chuRuKuMingXiBiaos, total, err
}
