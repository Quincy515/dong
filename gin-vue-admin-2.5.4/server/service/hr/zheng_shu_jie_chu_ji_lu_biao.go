package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hr"
	hrReq "github.com/flipped-aurora/gin-vue-admin/server/model/hr/request"
)

type ZhengShuJieChuJiLuBiaoService struct {
}

// CreateZhengShuJieChuJiLuBiao 创建ZhengShuJieChuJiLuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (zhengShuJieChuJiLuBiaoService *ZhengShuJieChuJiLuBiaoService) CreateZhengShuJieChuJiLuBiao(zhengShuJieChuJiLuBiao hr.ZhengShuJieChuJiLuBiao) (err error) {
	err = global.GVA_DB.Create(&zhengShuJieChuJiLuBiao).Error
	return err
}

// DeleteZhengShuJieChuJiLuBiao 删除ZhengShuJieChuJiLuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (zhengShuJieChuJiLuBiaoService *ZhengShuJieChuJiLuBiaoService) DeleteZhengShuJieChuJiLuBiao(zhengShuJieChuJiLuBiao hr.ZhengShuJieChuJiLuBiao) (err error) {
	err = global.GVA_DB.Delete(&zhengShuJieChuJiLuBiao).Error
	return err
}

// DeleteZhengShuJieChuJiLuBiaoByIds 批量删除ZhengShuJieChuJiLuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (zhengShuJieChuJiLuBiaoService *ZhengShuJieChuJiLuBiaoService) DeleteZhengShuJieChuJiLuBiaoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]hr.ZhengShuJieChuJiLuBiao{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateZhengShuJieChuJiLuBiao 更新ZhengShuJieChuJiLuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (zhengShuJieChuJiLuBiaoService *ZhengShuJieChuJiLuBiaoService) UpdateZhengShuJieChuJiLuBiao(zhengShuJieChuJiLuBiao hr.ZhengShuJieChuJiLuBiao) (err error) {
	if zhengShuJieChuJiLuBiao.CreatedAt.IsZero() && zhengShuJieChuJiLuBiao.ID != 0 {
		var temp hr.ZhengShuJieChuJiLuBiao
		err = global.GVA_DB.Where("id = ?", zhengShuJieChuJiLuBiao.ID).First(&temp).Error
		if err != nil {
			return err
		}
		zhengShuJieChuJiLuBiao.CreatedAt = temp.CreatedAt
	}
	err = global.GVA_DB.Save(&zhengShuJieChuJiLuBiao).Error
	return err
}

// GetZhengShuJieChuJiLuBiao 根据id获取ZhengShuJieChuJiLuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (zhengShuJieChuJiLuBiaoService *ZhengShuJieChuJiLuBiaoService) GetZhengShuJieChuJiLuBiao(id uint) (zhengShuJieChuJiLuBiao hr.ZhengShuJieChuJiLuBiao, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&zhengShuJieChuJiLuBiao).Error
	return
}

// GetZhengShuJieChuJiLuBiaoInfoList 分页获取ZhengShuJieChuJiLuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (zhengShuJieChuJiLuBiaoService *ZhengShuJieChuJiLuBiaoService) GetZhengShuJieChuJiLuBiaoInfoList(info hrReq.ZhengShuJieChuJiLuBiaoSearch) (list []hr.ZhengShuJieChuJiLuBiao, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hr.ZhengShuJieChuJiLuBiao{})
	var zhengShuJieChuJiLuBiaos []hr.ZhengShuJieChuJiLuBiao
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ZhengShuMingCheng != "" {
		db = db.Where("zheng_shu_ming_cheng LIKE ?", "%"+info.ZhengShuMingCheng+"%")
	}
	if info.ZhengShuBianHao != "" {
		db = db.Where("zheng_shu_bian_hao LIKE ?", "%"+info.ZhengShuBianHao+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id DESC").Limit(limit).Offset(offset).Find(&zhengShuJieChuJiLuBiaos).Error
	return zhengShuJieChuJiLuBiaos, total, err
}
