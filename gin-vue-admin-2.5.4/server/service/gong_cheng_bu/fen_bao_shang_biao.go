package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu"
	gong_cheng_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu/request"
)

type FenBaoShangBiaoService struct {
}

// CreateFenBaoShangBiao 创建FenBaoShangBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (fenBaoShangBiaoService *FenBaoShangBiaoService) CreateFenBaoShangBiao(fenBaoShangBiao gong_cheng_bu.FenBaoShangBiao) (err error) {
	err = global.GVA_DB.Create(&fenBaoShangBiao).Error
	return err
}

// DeleteFenBaoShangBiao 删除FenBaoShangBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (fenBaoShangBiaoService *FenBaoShangBiaoService) DeleteFenBaoShangBiao(fenBaoShangBiao gong_cheng_bu.FenBaoShangBiao) (err error) {
	err = global.GVA_DB.Delete(&fenBaoShangBiao).Error
	return err
}

// DeleteFenBaoShangBiaoByIds 批量删除FenBaoShangBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (fenBaoShangBiaoService *FenBaoShangBiaoService) DeleteFenBaoShangBiaoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]gong_cheng_bu.FenBaoShangBiao{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateFenBaoShangBiao 更新FenBaoShangBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (fenBaoShangBiaoService *FenBaoShangBiaoService) UpdateFenBaoShangBiao(fenBaoShangBiao gong_cheng_bu.FenBaoShangBiao) (err error) {
	err = global.GVA_DB.Save(&fenBaoShangBiao).Error
	return err
}

// GetFenBaoShangBiao 根据id获取FenBaoShangBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (fenBaoShangBiaoService *FenBaoShangBiaoService) GetFenBaoShangBiao(id uint) (fenBaoShangBiao gong_cheng_bu.FenBaoShangBiao, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&fenBaoShangBiao).Error
	return
}

// GetFenBaoShangBiaoInfoList 分页获取FenBaoShangBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (fenBaoShangBiaoService *FenBaoShangBiaoService) GetFenBaoShangBiaoInfoList(info gong_cheng_buReq.FenBaoShangBiaoSearch) (list []gong_cheng_bu.FenBaoShangBiao, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&gong_cheng_bu.FenBaoShangBiao{})
	var fenBaoShangBiaos []gong_cheng_bu.FenBaoShangBiao
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.FaRen != "" {
		db = db.Where("fa_ren LIKE ?", "%"+info.FaRen+"%")
	}
	if info.FenBaoShangLeiXing != nil {
		db = db.Where("fen_bao_shang_lei_xing = ?", info.FenBaoShangLeiXing)
	}
	if info.FenBaoShangMingCheng != "" {
		db = db.Where("fen_bao_shang_ming_cheng LIKE ?", "%"+info.FenBaoShangMingCheng+"%")
	}
	if info.TianJiaRen != nil {
		db = db.Where("tian_jia_ren = ?", info.TianJiaRen)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&fenBaoShangBiaos).Error
	return fenBaoShangBiaos, total, err
}
