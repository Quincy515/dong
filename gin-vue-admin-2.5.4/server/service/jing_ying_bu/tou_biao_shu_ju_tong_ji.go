package jing_ying_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/jing_ying_bu"
	jing_ying_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/jing_ying_bu/request"
)

type TouBiaoShuJuTongJiService struct {
}

// CreateTouBiaoShuJuTongJi 创建TouBiaoShuJuTongJi记录
// Author [piexlmax](https://github.com/piexlmax)
func (touBiaoShuJuTongJiService *TouBiaoShuJuTongJiService) CreateTouBiaoShuJuTongJi(touBiaoShuJuTongJi jing_ying_bu.TouBiaoShuJuTongJi) (err error) {
	err = global.GVA_DB.Create(&touBiaoShuJuTongJi).Error
	return err
}

// DeleteTouBiaoShuJuTongJi 删除TouBiaoShuJuTongJi记录
// Author [piexlmax](https://github.com/piexlmax)
func (touBiaoShuJuTongJiService *TouBiaoShuJuTongJiService) DeleteTouBiaoShuJuTongJi(touBiaoShuJuTongJi jing_ying_bu.TouBiaoShuJuTongJi) (err error) {
	err = global.GVA_DB.Delete(&touBiaoShuJuTongJi).Error
	return err
}

// DeleteTouBiaoShuJuTongJiByIds 批量删除TouBiaoShuJuTongJi记录
// Author [piexlmax](https://github.com/piexlmax)
func (touBiaoShuJuTongJiService *TouBiaoShuJuTongJiService) DeleteTouBiaoShuJuTongJiByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]jing_ying_bu.TouBiaoShuJuTongJi{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTouBiaoShuJuTongJi 更新TouBiaoShuJuTongJi记录
// Author [piexlmax](https://github.com/piexlmax)
func (touBiaoShuJuTongJiService *TouBiaoShuJuTongJiService) UpdateTouBiaoShuJuTongJi(touBiaoShuJuTongJi jing_ying_bu.TouBiaoShuJuTongJi) (err error) {
	if touBiaoShuJuTongJi.CreatedAt.IsZero() && touBiaoShuJuTongJi.ID != 0 {
		var temp jing_ying_bu.TouBiaoShuJuTongJi
		err = global.GVA_DB.Where("id = ?", touBiaoShuJuTongJi.ID).First(&temp).Error
		if err != nil {
			return err
		}
		touBiaoShuJuTongJi.CreatedAt = temp.CreatedAt
	}
	err = global.GVA_DB.Save(&touBiaoShuJuTongJi).Error
	return err
}

// GetTouBiaoShuJuTongJi 根据id获取TouBiaoShuJuTongJi记录
// Author [piexlmax](https://github.com/piexlmax)
func (touBiaoShuJuTongJiService *TouBiaoShuJuTongJiService) GetTouBiaoShuJuTongJi(id uint) (touBiaoShuJuTongJi jing_ying_bu.TouBiaoShuJuTongJi, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&touBiaoShuJuTongJi).Error
	return
}

// GetTouBiaoShuJuTongJiInfoList 分页获取TouBiaoShuJuTongJi记录
// Author [piexlmax](https://github.com/piexlmax)
func (touBiaoShuJuTongJiService *TouBiaoShuJuTongJiService) GetTouBiaoShuJuTongJiInfoList(info jing_ying_buReq.TouBiaoShuJuTongJiSearch) (list []jing_ying_bu.TouBiaoShuJuTongJi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&jing_ying_bu.TouBiaoShuJuTongJi{})
	var touBiaoShuJuTongJis []jing_ying_bu.TouBiaoShuJuTongJi
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.DengJiYaoQiu != nil {
		db = db.Where("deng_ji_yao_qiu = ?", info.DengJiYaoQiu)
	}
	if info.SuoShuDiQu != nil {
		db = db.Where("suo_shu_di_qu = ?", info.SuoShuDiQu)
	}
	if info.XiangMuMingCheng != "" {
		db = db.Where("xiang_mu_ming_cheng LIKE ?", "%"+info.XiangMuMingCheng+"%")
	}
	if info.ZiZhiYaoQiu != nil {
		db = db.Where("zi_zhi_yao_qiu = ?", info.ZiZhiYaoQiu)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("ID DESC").Limit(limit).Offset(offset).Find(&touBiaoShuJuTongJis).Error
	return touBiaoShuJuTongJis, total, err
}
