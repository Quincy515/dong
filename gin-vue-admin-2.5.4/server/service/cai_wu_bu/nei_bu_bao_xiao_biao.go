package cai_wu_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cai_wu_bu"
	cai_wu_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/cai_wu_bu/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type NeiBuBaoXiaoBiaoService struct {
}

// CreateNeiBuBaoXiaoBiao 创建NeiBuBaoXiaoBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (neiBuBaoXiaoBiaoService *NeiBuBaoXiaoBiaoService) CreateNeiBuBaoXiaoBiao(neiBuBaoXiaoBiao cai_wu_bu.NeiBuBaoXiaoBiao) (err error) {
	err = global.GVA_DB.Create(&neiBuBaoXiaoBiao).Error
	return err
}

// DeleteNeiBuBaoXiaoBiao 删除NeiBuBaoXiaoBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (neiBuBaoXiaoBiaoService *NeiBuBaoXiaoBiaoService) DeleteNeiBuBaoXiaoBiao(neiBuBaoXiaoBiao cai_wu_bu.NeiBuBaoXiaoBiao) (err error) {
	err = global.GVA_DB.Delete(&neiBuBaoXiaoBiao).Error
	return err
}

// DeleteNeiBuBaoXiaoBiaoByIds 批量删除NeiBuBaoXiaoBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (neiBuBaoXiaoBiaoService *NeiBuBaoXiaoBiaoService) DeleteNeiBuBaoXiaoBiaoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]cai_wu_bu.NeiBuBaoXiaoBiao{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateNeiBuBaoXiaoBiao 更新NeiBuBaoXiaoBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (neiBuBaoXiaoBiaoService *NeiBuBaoXiaoBiaoService) UpdateNeiBuBaoXiaoBiao(neiBuBaoXiaoBiao cai_wu_bu.NeiBuBaoXiaoBiao) (err error) {
	err = global.GVA_DB.Save(&neiBuBaoXiaoBiao).Error
	return err
}

// GetNeiBuBaoXiaoBiao 根据id获取NeiBuBaoXiaoBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (neiBuBaoXiaoBiaoService *NeiBuBaoXiaoBiaoService) GetNeiBuBaoXiaoBiao(id uint) (neiBuBaoXiaoBiao cai_wu_bu.NeiBuBaoXiaoBiao, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&neiBuBaoXiaoBiao).Error
	return
}

// GetNeiBuBaoXiaoBiaoInfoList 分页获取NeiBuBaoXiaoBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (neiBuBaoXiaoBiaoService *NeiBuBaoXiaoBiaoService) GetNeiBuBaoXiaoBiaoInfoList(info cai_wu_buReq.NeiBuBaoXiaoBiaoSearch) (list []cai_wu_bu.NeiBuBaoXiaoBiao, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&cai_wu_bu.NeiBuBaoXiaoBiao{})
	var neiBuBaoXiaoBiaos []cai_wu_bu.NeiBuBaoXiaoBiao
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.BaoXiaoBuMen != nil {
		db = db.Where("bao_xiao_bu_men = ?", info.BaoXiaoBuMen)
	}
	if info.BaoXiaoShiXiang != nil {
		db = db.Where("bao_xiao_shi_xiang = ?", info.BaoXiaoShiXiang)
	}
	if info.GongSiMingCheng != nil {
		db = db.Where("gong_si_ming_cheng = ?", info.GongSiMingCheng)
	}
	if info.BaoXiaoLeiXing != nil {
		db = db.Where("bao_xiao_lei_xing = ?", info.BaoXiaoLeiXing)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&neiBuBaoXiaoBiaos).Error
	return neiBuBaoXiaoBiaos, total, err
}
