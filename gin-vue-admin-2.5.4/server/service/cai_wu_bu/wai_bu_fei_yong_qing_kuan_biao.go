package cai_wu_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cai_wu_bu"
	cai_wu_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/cai_wu_bu/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type WaiBuFeiYongQingKuanBiaoService struct {
}

// CreateWaiBuFeiYongQingKuanBiao 创建WaiBuFeiYongQingKuanBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (waiBuFeiYongQingKuanBiaoService *WaiBuFeiYongQingKuanBiaoService) CreateWaiBuFeiYongQingKuanBiao(waiBuFeiYongQingKuanBiao cai_wu_bu.WaiBuFeiYongQingKuanBiao) (err error) {
	err = global.GVA_DB.Create(&waiBuFeiYongQingKuanBiao).Error
	return err
}

// DeleteWaiBuFeiYongQingKuanBiao 删除WaiBuFeiYongQingKuanBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (waiBuFeiYongQingKuanBiaoService *WaiBuFeiYongQingKuanBiaoService) DeleteWaiBuFeiYongQingKuanBiao(waiBuFeiYongQingKuanBiao cai_wu_bu.WaiBuFeiYongQingKuanBiao) (err error) {
	err = global.GVA_DB.Delete(&waiBuFeiYongQingKuanBiao).Error
	return err
}

// DeleteWaiBuFeiYongQingKuanBiaoByIds 批量删除WaiBuFeiYongQingKuanBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (waiBuFeiYongQingKuanBiaoService *WaiBuFeiYongQingKuanBiaoService) DeleteWaiBuFeiYongQingKuanBiaoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]cai_wu_bu.WaiBuFeiYongQingKuanBiao{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateWaiBuFeiYongQingKuanBiao 更新WaiBuFeiYongQingKuanBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (waiBuFeiYongQingKuanBiaoService *WaiBuFeiYongQingKuanBiaoService) UpdateWaiBuFeiYongQingKuanBiao(waiBuFeiYongQingKuanBiao cai_wu_bu.WaiBuFeiYongQingKuanBiao) (err error) {
	err = global.GVA_DB.Save(&waiBuFeiYongQingKuanBiao).Error
	return err
}

// GetWaiBuFeiYongQingKuanBiao 根据id获取WaiBuFeiYongQingKuanBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (waiBuFeiYongQingKuanBiaoService *WaiBuFeiYongQingKuanBiaoService) GetWaiBuFeiYongQingKuanBiao(id uint) (waiBuFeiYongQingKuanBiao cai_wu_bu.WaiBuFeiYongQingKuanBiao, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&waiBuFeiYongQingKuanBiao).Error
	return
}

// GetWaiBuFeiYongQingKuanBiaoInfoList 分页获取WaiBuFeiYongQingKuanBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (waiBuFeiYongQingKuanBiaoService *WaiBuFeiYongQingKuanBiaoService) GetWaiBuFeiYongQingKuanBiaoInfoList(info cai_wu_buReq.WaiBuFeiYongQingKuanBiaoSearch) (list []cai_wu_bu.WaiBuFeiYongQingKuanBiao, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&cai_wu_bu.WaiBuFeiYongQingKuanBiao{})
	var waiBuFeiYongQingKuanBiaos []cai_wu_bu.WaiBuFeiYongQingKuanBiao
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.BaoXiaoLeiXing != nil {
		db = db.Where("bao_xiao_lei_xing = ?", info.BaoXiaoLeiXing)
	}
	if info.BaoXiaoRenYuan != nil {
		db = db.Where("bao_xiao_ren_yuan = ?", info.BaoXiaoRenYuan)
	}
	if info.BaoXiaoShiXiang != nil {
		db = db.Where("bao_xiao_shi_xiang = ?", info.BaoXiaoShiXiang)
	}
	if info.HeTongId != nil {
		db = db.Where("he_tong_id = ?", info.HeTongId)
	}
	if info.ShouCiBaoXiaoId != nil {
		db = db.Where("shou_ci_bao_xiao_id = ?", info.ShouCiBaoXiaoId)
	}
	if info.XiangMuMingCheng != nil {
		db = db.Where("xiang_mu_ming_cheng = ?", info.XiangMuMingCheng)
	}
	if info.ZhuangTai != nil {
		db = db.Where("zhuang_tai = ?", info.ZhuangTai)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&waiBuFeiYongQingKuanBiaos).Error
	return waiBuFeiYongQingKuanBiaos, total, err
}
