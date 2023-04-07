package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu"
	gong_cheng_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu/request"
)

type ShiGongJinDuBiaoService struct {
}

// CreateShiGongJinDuBiao 创建ShiGongJinDuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (shiGongJinDuBiaoService *ShiGongJinDuBiaoService) CreateShiGongJinDuBiao(shiGongJinDuBiao gong_cheng_bu.ShiGongJinDuBiao) (err error) {
	err = global.GVA_DB.Create(&shiGongJinDuBiao).Error
	return err
}

// DeleteShiGongJinDuBiao 删除ShiGongJinDuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (shiGongJinDuBiaoService *ShiGongJinDuBiaoService) DeleteShiGongJinDuBiao(shiGongJinDuBiao gong_cheng_bu.ShiGongJinDuBiao) (err error) {
	err = global.GVA_DB.Delete(&shiGongJinDuBiao).Error
	return err
}

// DeleteShiGongJinDuBiaoByIds 批量删除ShiGongJinDuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (shiGongJinDuBiaoService *ShiGongJinDuBiaoService) DeleteShiGongJinDuBiaoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]gong_cheng_bu.ShiGongJinDuBiao{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateShiGongJinDuBiao 更新ShiGongJinDuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (shiGongJinDuBiaoService *ShiGongJinDuBiaoService) UpdateShiGongJinDuBiao(shiGongJinDuBiao gong_cheng_bu.ShiGongJinDuBiao) (err error) {
	err = global.GVA_DB.Save(&shiGongJinDuBiao).Error
	return err
}

// GetShiGongJinDuBiao 根据id获取ShiGongJinDuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (shiGongJinDuBiaoService *ShiGongJinDuBiaoService) GetShiGongJinDuBiao(id uint) (shiGongJinDuBiao gong_cheng_bu.ShiGongJinDuBiao, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&shiGongJinDuBiao).Error
	return
}

// GetShiGongJinDuBiaoInfoList 分页获取ShiGongJinDuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (shiGongJinDuBiaoService *ShiGongJinDuBiaoService) GetShiGongJinDuBiaoInfoList(info gong_cheng_buReq.ShiGongJinDuBiaoSearch) (list []gong_cheng_bu.ShiGongJinDuBiao, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&gong_cheng_bu.ShiGongJinDuBiao{})
	var shiGongJinDuBiaos []gong_cheng_bu.ShiGongJinDuBiao
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.GongZuoMingChengId != nil {
		db = db.Where("gong_zuo_ming_cheng_id = ?", info.GongZuoMingChengId)
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
	err = db.Limit(limit).Offset(offset).Find(&shiGongJinDuBiaos).Error
	return shiGongJinDuBiaos, total, err
}
