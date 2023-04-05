package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu"
	gong_cheng_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu/request"
)

type HeTongWenJianBiaoService struct {
}

// CreateHeTongWenJianBiao 创建HeTongWenJianBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (heTongWenJianBiaoService *HeTongWenJianBiaoService) CreateHeTongWenJianBiao(heTongWenJianBiao gong_cheng_bu.HeTongWenJianBiao) (err error) {
	err = global.GVA_DB.Create(&heTongWenJianBiao).Error
	return err
}

// DeleteHeTongWenJianBiao 删除HeTongWenJianBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (heTongWenJianBiaoService *HeTongWenJianBiaoService) DeleteHeTongWenJianBiao(heTongWenJianBiao gong_cheng_bu.HeTongWenJianBiao) (err error) {
	err = global.GVA_DB.Delete(&heTongWenJianBiao).Error
	return err
}

// DeleteHeTongWenJianBiaoByIds 批量删除HeTongWenJianBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (heTongWenJianBiaoService *HeTongWenJianBiaoService) DeleteHeTongWenJianBiaoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]gong_cheng_bu.HeTongWenJianBiao{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHeTongWenJianBiao 更新HeTongWenJianBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (heTongWenJianBiaoService *HeTongWenJianBiaoService) UpdateHeTongWenJianBiao(heTongWenJianBiao gong_cheng_bu.HeTongWenJianBiao) (err error) {
	err = global.GVA_DB.Save(&heTongWenJianBiao).Error
	return err
}

// GetHeTongWenJianBiao 根据id获取HeTongWenJianBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (heTongWenJianBiaoService *HeTongWenJianBiaoService) GetHeTongWenJianBiao(id uint) (heTongWenJianBiao gong_cheng_bu.HeTongWenJianBiao, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&heTongWenJianBiao).Error
	return
}

// GetHeTongWenJianBiaoInfoList 分页获取HeTongWenJianBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (heTongWenJianBiaoService *HeTongWenJianBiaoService) GetHeTongWenJianBiaoInfoList(info gong_cheng_buReq.HeTongWenJianBiaoSearch) (list []gong_cheng_bu.HeTongWenJianBiao, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&gong_cheng_bu.HeTongWenJianBiao{})
	var heTongWenJianBiaos []gong_cheng_bu.HeTongWenJianBiao
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.FaPiaoLeiXing != nil {
		db = db.Where("fa_piao_lei_xing = ?", info.FaPiaoLeiXing)
	}
	if info.FuKuanFangShi != nil {
		db = db.Where("fu_kuan_fang_shi = ?", info.FuKuanFangShi)
	}
	if info.HeTongBianHao != "" {
		db = db.Where("he_tong_bian_hao LIKE ?", "%"+info.HeTongBianHao+"%")
	}
	if info.HeTongLeiBie != nil {
		db = db.Where("he_tong_lei_bie = ?", info.HeTongLeiBie)
	}
	if info.HeTongLeiXing != nil {
		db = db.Where("he_tong_lei_xing = ?", info.HeTongLeiXing)
	}
	if info.HeTongMingCheng != "" {
		db = db.Where("he_tong_ming_cheng LIKE ?", "%"+info.HeTongMingCheng+"%")
	}
	if info.HeTongZhuangTai != nil {
		db = db.Where("he_tong_zhuang_tai = ?", info.HeTongZhuangTai)
	}
	if info.JiaFangGongSi != nil {
		db = db.Where("jia_fang_gong_si = ?", info.JiaFangGongSi)
	}
	if info.JieSuanFangShi != nil {
		db = db.Where("jie_suan_fang_shi = ?", info.JieSuanFangShi)
	}
	if info.QianDingRen != nil {
		db = db.Where("qian_ding_ren = ?", info.QianDingRen)
	}
	if info.TianJiaRen != nil {
		db = db.Where("tian_jia_ren = ?", info.TianJiaRen)
	}
	if info.XiangMuId != nil {
		db = db.Where("xiang_mu_id = ?", info.XiangMuId)
	}
	if info.YiFangGongSi != nil {
		db = db.Where("yi_fang_gong_si = ?", info.YiFangGongSi)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&heTongWenJianBiaos).Error
	return heTongWenJianBiaos, total, err
}
