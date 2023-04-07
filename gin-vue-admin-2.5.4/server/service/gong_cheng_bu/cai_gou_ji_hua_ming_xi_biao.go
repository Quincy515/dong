package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu"
	gong_cheng_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu/request"
)

type CaiGouJiHuaMingXiBiaoService struct {
}

// CreateCaiGouJiHuaMingXiBiao 创建CaiGouJiHuaMingXiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (caiGouJiHuaMingXiBiaoService *CaiGouJiHuaMingXiBiaoService) CreateCaiGouJiHuaMingXiBiao(caiGouJiHuaMingXiBiao gong_cheng_bu.CaiGouJiHuaMingXiBiao) (err error) {
	err = global.GVA_DB.Create(&caiGouJiHuaMingXiBiao).Error
	return err
}

// DeleteCaiGouJiHuaMingXiBiao 删除CaiGouJiHuaMingXiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (caiGouJiHuaMingXiBiaoService *CaiGouJiHuaMingXiBiaoService) DeleteCaiGouJiHuaMingXiBiao(caiGouJiHuaMingXiBiao gong_cheng_bu.CaiGouJiHuaMingXiBiao) (err error) {
	err = global.GVA_DB.Delete(&caiGouJiHuaMingXiBiao).Error
	return err
}

// DeleteCaiGouJiHuaMingXiBiaoByIds 批量删除CaiGouJiHuaMingXiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (caiGouJiHuaMingXiBiaoService *CaiGouJiHuaMingXiBiaoService) DeleteCaiGouJiHuaMingXiBiaoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]gong_cheng_bu.CaiGouJiHuaMingXiBiao{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCaiGouJiHuaMingXiBiao 更新CaiGouJiHuaMingXiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (caiGouJiHuaMingXiBiaoService *CaiGouJiHuaMingXiBiaoService) UpdateCaiGouJiHuaMingXiBiao(caiGouJiHuaMingXiBiao gong_cheng_bu.CaiGouJiHuaMingXiBiao) (err error) {
	err = global.GVA_DB.Save(&caiGouJiHuaMingXiBiao).Error
	return err
}

// GetCaiGouJiHuaMingXiBiao 根据id获取CaiGouJiHuaMingXiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (caiGouJiHuaMingXiBiaoService *CaiGouJiHuaMingXiBiaoService) GetCaiGouJiHuaMingXiBiao(id uint) (caiGouJiHuaMingXiBiao gong_cheng_bu.CaiGouJiHuaMingXiBiao, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&caiGouJiHuaMingXiBiao).Error
	return
}

// GetCaiGouJiHuaMingXiBiaoInfoList 分页获取CaiGouJiHuaMingXiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (caiGouJiHuaMingXiBiaoService *CaiGouJiHuaMingXiBiaoService) GetCaiGouJiHuaMingXiBiaoInfoList(info gong_cheng_buReq.CaiGouJiHuaMingXiBiaoSearch) (list []gong_cheng_bu.CaiGouJiHuaMingXiBiao, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&gong_cheng_bu.CaiGouJiHuaMingXiBiao{})
	var caiGouJiHuaMingXiBiaos []gong_cheng_bu.CaiGouJiHuaMingXiBiao
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CaiGouJiHuaId != nil {
		db = db.Where("cai_gou_ji_hua_id = ?", info.CaiGouJiHuaId)
	}
	if info.CaiLiaoId != nil {
		db = db.Where("cai_liao_id = ?", info.CaiLiaoId)
	}
	if info.GongYingShangId != nil {
		db = db.Where("gong_ying_shang_id = ?", info.GongYingShangId)
	}
	if info.ZhuangTai != nil {
		db = db.Where("zhuang_tai = ?", info.ZhuangTai)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&caiGouJiHuaMingXiBiaos).Error
	return caiGouJiHuaMingXiBiaos, total, err
}
