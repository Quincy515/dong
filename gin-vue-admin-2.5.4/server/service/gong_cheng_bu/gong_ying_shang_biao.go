package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu"
	gong_cheng_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu/request"
)

type GongYingShangBiaoService struct {
}

// CreateGongYingShangBiao 创建GongYingShangBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (gongYingShangBiaoService *GongYingShangBiaoService) CreateGongYingShangBiao(gongYingShangBiao gong_cheng_bu.GongYingShangBiao) (err error) {
	err = global.GVA_DB.Create(&gongYingShangBiao).Error
	return err
}

// DeleteGongYingShangBiao 删除GongYingShangBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (gongYingShangBiaoService *GongYingShangBiaoService) DeleteGongYingShangBiao(gongYingShangBiao gong_cheng_bu.GongYingShangBiao) (err error) {
	err = global.GVA_DB.Delete(&gongYingShangBiao).Error
	return err
}

// DeleteGongYingShangBiaoByIds 批量删除GongYingShangBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (gongYingShangBiaoService *GongYingShangBiaoService) DeleteGongYingShangBiaoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]gong_cheng_bu.GongYingShangBiao{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateGongYingShangBiao 更新GongYingShangBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (gongYingShangBiaoService *GongYingShangBiaoService) UpdateGongYingShangBiao(gongYingShangBiao gong_cheng_bu.GongYingShangBiao) (err error) {
	err = global.GVA_DB.Save(&gongYingShangBiao).Error
	return err
}

// GetGongYingShangBiao 根据id获取GongYingShangBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (gongYingShangBiaoService *GongYingShangBiaoService) GetGongYingShangBiao(id uint) (gongYingShangBiao gong_cheng_bu.GongYingShangBiao, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&gongYingShangBiao).Error
	return
}

// GetGongYingShangBiaoInfoList 分页获取GongYingShangBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (gongYingShangBiaoService *GongYingShangBiaoService) GetGongYingShangBiaoInfoList(info gong_cheng_buReq.GongYingShangBiaoSearch) (list []gong_cheng_bu.GongYingShangBiao, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&gong_cheng_bu.GongYingShangBiao{})
	var gongYingShangBiaos []gong_cheng_bu.GongYingShangBiao
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.FaPiaoLeiXing != nil {
		db = db.Where("fa_piao_lei_xing = ?", info.FaPiaoLeiXing)
	}
	if info.GongYingShangMingCheng != nil {
		db = db.Where("gong_ying_shang_ming_cheng = ?", info.GongYingShangMingCheng)
	}
	if info.ZhuangTai != nil {
		db = db.Where("zhuang_tai = ?", info.ZhuangTai)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&gongYingShangBiaos).Error
	return gongYingShangBiaos, total, err
}
