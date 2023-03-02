package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hr"
	hrReq "github.com/flipped-aurora/gin-vue-admin/server/model/hr/request"
)

type ZhengShuBiaoService struct {
}

// CreateZhengShuBiao 创建ZhengShuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (zhengShuBiaoService *ZhengShuBiaoService) CreateZhengShuBiao(zhengShuBiao hr.ZhengShuBiao) (err error) {
	err = global.GVA_DB.Create(&zhengShuBiao).Error
	return err
}

// DeleteZhengShuBiao 删除ZhengShuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (zhengShuBiaoService *ZhengShuBiaoService) DeleteZhengShuBiao(zhengShuBiao hr.ZhengShuBiao) (err error) {
	err = global.GVA_DB.Delete(&zhengShuBiao).Error
	return err
}

// DeleteZhengShuBiaoByIds 批量删除ZhengShuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (zhengShuBiaoService *ZhengShuBiaoService) DeleteZhengShuBiaoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]hr.ZhengShuBiao{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateZhengShuBiao 更新ZhengShuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (zhengShuBiaoService *ZhengShuBiaoService) UpdateZhengShuBiao(zhengShuBiao hr.ZhengShuBiao) (err error) {
	if zhengShuBiao.CreatedAt.IsZero() && zhengShuBiao.ID != 0 {
		var temp hr.ZhengShuBiao
		err = global.GVA_DB.Where("id = ?", zhengShuBiao.ID).First(&temp).Error
		if err != nil {
			return err
		}
		zhengShuBiao.CreatedAt = temp.CreatedAt
	}
	err = global.GVA_DB.Save(&zhengShuBiao).Error
	return err
}

// GetZhengShuBiao 根据id获取ZhengShuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (zhengShuBiaoService *ZhengShuBiaoService) GetZhengShuBiao(id uint) (zhengShuBiao hr.ZhengShuBiao, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&zhengShuBiao).Error
	return
}

// GetZhengShuBiaoInfoList 分页获取ZhengShuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (zhengShuBiaoService *ZhengShuBiaoService) GetZhengShuBiaoInfoList(info hrReq.ZhengShuBiaoSearch) (list []hr.ZhengShuBiao, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hr.ZhengShuBiao{})
	var zhengShuBiaos []hr.ZhengShuBiao
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
	if info.ZhengShuLeiBie != nil {
		db = db.Where("zheng_shu_lei_bie = ?", info.ZhengShuLeiBie)
	}
	if info.SuoYouRenXingMing != "" {
		db = db.Where("suo_you_ren_xing_ming LIKE ?", "%"+info.SuoYouRenXingMing+"%")
	}
	if info.ZhengShuSuoShuChengShi != nil {
		db = db.Where("zheng_shu_suo_shu_cheng_shi = ?", info.ZhengShuSuoShuChengShi)
	}
	if info.ZhengShuZhuangTai != nil {
		db = db.Where("zheng_shu_zhuang_tai = ?", info.ZhengShuZhuangTai)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id DESC").Limit(limit).Offset(offset).Find(&zhengShuBiaos).Error
	return zhengShuBiaos, total, err
}
