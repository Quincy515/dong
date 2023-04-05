package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu"
	gong_cheng_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu/request"
)

type FenBaoShangYuXiangMuGuanLianBiaoService struct {
}

// CreateFenBaoShangYuXiangMuGuanLianBiao 创建FenBaoShangYuXiangMuGuanLianBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (fenBaoShangYuXiangMuGuanLianBiaoService *FenBaoShangYuXiangMuGuanLianBiaoService) CreateFenBaoShangYuXiangMuGuanLianBiao(fenBaoShangYuXiangMuGuanLianBiao gong_cheng_bu.FenBaoShangYuXiangMuGuanLianBiao) (err error) {
	err = global.GVA_DB.Create(&fenBaoShangYuXiangMuGuanLianBiao).Error
	return err
}

// DeleteFenBaoShangYuXiangMuGuanLianBiao 删除FenBaoShangYuXiangMuGuanLianBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (fenBaoShangYuXiangMuGuanLianBiaoService *FenBaoShangYuXiangMuGuanLianBiaoService) DeleteFenBaoShangYuXiangMuGuanLianBiao(fenBaoShangYuXiangMuGuanLianBiao gong_cheng_bu.FenBaoShangYuXiangMuGuanLianBiao) (err error) {
	err = global.GVA_DB.Delete(&fenBaoShangYuXiangMuGuanLianBiao).Error
	return err
}

// DeleteFenBaoShangYuXiangMuGuanLianBiaoByIds 批量删除FenBaoShangYuXiangMuGuanLianBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (fenBaoShangYuXiangMuGuanLianBiaoService *FenBaoShangYuXiangMuGuanLianBiaoService) DeleteFenBaoShangYuXiangMuGuanLianBiaoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]gong_cheng_bu.FenBaoShangYuXiangMuGuanLianBiao{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateFenBaoShangYuXiangMuGuanLianBiao 更新FenBaoShangYuXiangMuGuanLianBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (fenBaoShangYuXiangMuGuanLianBiaoService *FenBaoShangYuXiangMuGuanLianBiaoService) UpdateFenBaoShangYuXiangMuGuanLianBiao(fenBaoShangYuXiangMuGuanLianBiao gong_cheng_bu.FenBaoShangYuXiangMuGuanLianBiao) (err error) {
	err = global.GVA_DB.Save(&fenBaoShangYuXiangMuGuanLianBiao).Error
	return err
}

// GetFenBaoShangYuXiangMuGuanLianBiao 根据id获取FenBaoShangYuXiangMuGuanLianBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (fenBaoShangYuXiangMuGuanLianBiaoService *FenBaoShangYuXiangMuGuanLianBiaoService) GetFenBaoShangYuXiangMuGuanLianBiao(id uint) (fenBaoShangYuXiangMuGuanLianBiao gong_cheng_bu.FenBaoShangYuXiangMuGuanLianBiao, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&fenBaoShangYuXiangMuGuanLianBiao).Error
	return
}

// GetFenBaoShangYuXiangMuGuanLianBiaoInfoList 分页获取FenBaoShangYuXiangMuGuanLianBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (fenBaoShangYuXiangMuGuanLianBiaoService *FenBaoShangYuXiangMuGuanLianBiaoService) GetFenBaoShangYuXiangMuGuanLianBiaoInfoList(info gong_cheng_buReq.FenBaoShangYuXiangMuGuanLianBiaoSearch) (list []gong_cheng_bu.FenBaoShangYuXiangMuGuanLianBiao, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&gong_cheng_bu.FenBaoShangYuXiangMuGuanLianBiao{})
	var fenBaoShangYuXiangMuGuanLianBiaos []gong_cheng_bu.FenBaoShangYuXiangMuGuanLianBiao
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.FenBaoShangId != nil {
		db = db.Where("fen_bao_shang_id = ?", info.FenBaoShangId)
	}
	if info.Group != "" {
		db = db.Where("group = ?", info.Group)
	}
	if info.HeTongBianHao != "" {
		db = db.Where("he_tong_bian_hao LIKE ?", "%"+info.HeTongBianHao+"%")
	}
	if info.HeTongMingCheng != "" {
		db = db.Where("he_tong_ming_cheng LIKE ?", "%"+info.HeTongMingCheng+"%")
	}
	if info.WeiTuoRen != "" {
		db = db.Where("wei_tuo_ren LIKE ?", "%"+info.WeiTuoRen+"%")
	}
	if info.XiangMuId != nil {
		db = db.Where("xiang_mu_id = ?", info.XiangMuId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&fenBaoShangYuXiangMuGuanLianBiaos).Error
	return fenBaoShangYuXiangMuGuanLianBiaos, total, err
}
