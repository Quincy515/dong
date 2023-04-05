package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu"
	gong_cheng_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu/request"
)

type XiangMuShiGongBiaoService struct {
}

// CreateXiangMuShiGongBiao 创建XiangMuShiGongBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiangMuShiGongBiaoService *XiangMuShiGongBiaoService) CreateXiangMuShiGongBiao(xiangMuShiGongBiao gong_cheng_bu.XiangMuShiGongBiao) (err error) {
	err = global.GVA_DB.Create(&xiangMuShiGongBiao).Error
	return err
}

// DeleteXiangMuShiGongBiao 删除XiangMuShiGongBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiangMuShiGongBiaoService *XiangMuShiGongBiaoService) DeleteXiangMuShiGongBiao(xiangMuShiGongBiao gong_cheng_bu.XiangMuShiGongBiao) (err error) {
	err = global.GVA_DB.Delete(&xiangMuShiGongBiao).Error
	return err
}

// DeleteXiangMuShiGongBiaoByIds 批量删除XiangMuShiGongBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiangMuShiGongBiaoService *XiangMuShiGongBiaoService) DeleteXiangMuShiGongBiaoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]gong_cheng_bu.XiangMuShiGongBiao{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateXiangMuShiGongBiao 更新XiangMuShiGongBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiangMuShiGongBiaoService *XiangMuShiGongBiaoService) UpdateXiangMuShiGongBiao(xiangMuShiGongBiao gong_cheng_bu.XiangMuShiGongBiao) (err error) {
	err = global.GVA_DB.Save(&xiangMuShiGongBiao).Error
	return err
}

// GetXiangMuShiGongBiao 根据id获取XiangMuShiGongBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiangMuShiGongBiaoService *XiangMuShiGongBiaoService) GetXiangMuShiGongBiao(id uint) (xiangMuShiGongBiao gong_cheng_bu.XiangMuShiGongBiao, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&xiangMuShiGongBiao).Error
	return
}

// GetXiangMuShiGongBiaoInfoList 分页获取XiangMuShiGongBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiangMuShiGongBiaoService *XiangMuShiGongBiaoService) GetXiangMuShiGongBiaoInfoList(info gong_cheng_buReq.XiangMuShiGongBiaoSearch) (list []gong_cheng_bu.XiangMuShiGongBiao, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&gong_cheng_bu.XiangMuShiGongBiao{})
	var xiangMuShiGongBiaos []gong_cheng_bu.XiangMuShiGongBiao
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.GongZuoMingCheng != "" {
		db = db.Where("gong_zuo_ming_cheng LIKE ?", "%"+info.GongZuoMingCheng+"%")
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
	err = db.Limit(limit).Offset(offset).Find(&xiangMuShiGongBiaos).Error
	return xiangMuShiGongBiaos, total, err
}
