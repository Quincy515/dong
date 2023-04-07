package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu"
	gong_cheng_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu/request"
)

type ChuRuKuDanJuBiaoService struct {
}

// CreateChuRuKuDanJuBiao 创建ChuRuKuDanJuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (chuRuKuDanJuBiaoService *ChuRuKuDanJuBiaoService) CreateChuRuKuDanJuBiao(chuRuKuDanJuBiao gong_cheng_bu.ChuRuKuDanJuBiao) (err error) {
	err = global.GVA_DB.Create(&chuRuKuDanJuBiao).Error
	return err
}

// DeleteChuRuKuDanJuBiao 删除ChuRuKuDanJuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (chuRuKuDanJuBiaoService *ChuRuKuDanJuBiaoService) DeleteChuRuKuDanJuBiao(chuRuKuDanJuBiao gong_cheng_bu.ChuRuKuDanJuBiao) (err error) {
	err = global.GVA_DB.Delete(&chuRuKuDanJuBiao).Error
	return err
}

// DeleteChuRuKuDanJuBiaoByIds 批量删除ChuRuKuDanJuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (chuRuKuDanJuBiaoService *ChuRuKuDanJuBiaoService) DeleteChuRuKuDanJuBiaoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]gong_cheng_bu.ChuRuKuDanJuBiao{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateChuRuKuDanJuBiao 更新ChuRuKuDanJuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (chuRuKuDanJuBiaoService *ChuRuKuDanJuBiaoService) UpdateChuRuKuDanJuBiao(chuRuKuDanJuBiao gong_cheng_bu.ChuRuKuDanJuBiao) (err error) {
	err = global.GVA_DB.Save(&chuRuKuDanJuBiao).Error
	return err
}

// GetChuRuKuDanJuBiao 根据id获取ChuRuKuDanJuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (chuRuKuDanJuBiaoService *ChuRuKuDanJuBiaoService) GetChuRuKuDanJuBiao(id uint) (chuRuKuDanJuBiao gong_cheng_bu.ChuRuKuDanJuBiao, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&chuRuKuDanJuBiao).Error
	return
}

// GetChuRuKuDanJuBiaoInfoList 分页获取ChuRuKuDanJuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (chuRuKuDanJuBiaoService *ChuRuKuDanJuBiaoService) GetChuRuKuDanJuBiaoInfoList(info gong_cheng_buReq.ChuRuKuDanJuBiaoSearch) (list []gong_cheng_bu.ChuRuKuDanJuBiao, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&gong_cheng_bu.ChuRuKuDanJuBiao{})
	var chuRuKuDanJuBiaos []gong_cheng_bu.ChuRuKuDanJuBiao
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CaiGouJiHuaId != nil {
		db = db.Where("cai_gou_ji_hua_id = ?", info.CaiGouJiHuaId)
	}
	if info.ChuRuKuDanJuLeiXing != nil {
		db = db.Where("chu_ru_ku_dan_ju_lei_xing = ?", info.ChuRuKuDanJuLeiXing)
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
	err = db.Limit(limit).Offset(offset).Find(&chuRuKuDanJuBiaos).Error
	return chuRuKuDanJuBiaos, total, err
}
