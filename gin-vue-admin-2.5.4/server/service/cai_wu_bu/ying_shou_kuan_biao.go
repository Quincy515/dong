package cai_wu_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cai_wu_bu"
	cai_wu_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/cai_wu_bu/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type YingShouKuanBiaoService struct {
}

// CreateYingShouKuanBiao 创建YingShouKuanBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (yingShouKuanBiaoService *YingShouKuanBiaoService) CreateYingShouKuanBiao(yingShouKuanBiao cai_wu_bu.YingShouKuanBiao) (err error) {
	err = global.GVA_DB.Create(&yingShouKuanBiao).Error
	return err
}

// DeleteYingShouKuanBiao 删除YingShouKuanBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (yingShouKuanBiaoService *YingShouKuanBiaoService) DeleteYingShouKuanBiao(yingShouKuanBiao cai_wu_bu.YingShouKuanBiao) (err error) {
	err = global.GVA_DB.Delete(&yingShouKuanBiao).Error
	return err
}

// DeleteYingShouKuanBiaoByIds 批量删除YingShouKuanBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (yingShouKuanBiaoService *YingShouKuanBiaoService) DeleteYingShouKuanBiaoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]cai_wu_bu.YingShouKuanBiao{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateYingShouKuanBiao 更新YingShouKuanBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (yingShouKuanBiaoService *YingShouKuanBiaoService) UpdateYingShouKuanBiao(yingShouKuanBiao cai_wu_bu.YingShouKuanBiao) (err error) {
	err = global.GVA_DB.Save(&yingShouKuanBiao).Error
	return err
}

// GetYingShouKuanBiao 根据id获取YingShouKuanBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (yingShouKuanBiaoService *YingShouKuanBiaoService) GetYingShouKuanBiao(id uint) (yingShouKuanBiao cai_wu_bu.YingShouKuanBiao, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&yingShouKuanBiao).Error
	return
}

// GetYingShouKuanBiaoInfoList 分页获取YingShouKuanBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (yingShouKuanBiaoService *YingShouKuanBiaoService) GetYingShouKuanBiaoInfoList(info cai_wu_buReq.YingShouKuanBiaoSearch) (list []cai_wu_bu.YingShouKuanBiao, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&cai_wu_bu.YingShouKuanBiao{})
	var yingShouKuanBiaos []cai_wu_bu.YingShouKuanBiao
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.FaPiaoLeiXing != nil {
		db = db.Where("fa_piao_lei_xing = ?", info.FaPiaoLeiXing)
	}
	if info.FaQiRenYuan != nil {
		db = db.Where("fa_qi_ren_yuan = ?", info.FaQiRenYuan)
	}
	if info.StartFeiYongFaShengRiQi != nil && info.EndFeiYongFaShengRiQi != nil {
		db = db.Where("fei_yong_fa_sheng_ri_qi BETWEEN ? AND ? ", info.StartFeiYongFaShengRiQi, info.EndFeiYongFaShengRiQi)
	}
	if info.Group != "" {
		db = db.Where("group = ?", info.Group)
	}
	if info.HeTongId != nil {
		db = db.Where("he_tong_id = ?", info.HeTongId)
	}
	if info.StartShenQingRiQi != nil && info.EndShenQingRiQi != nil {
		db = db.Where("shen_qing_ri_qi BETWEEN ? AND ? ", info.StartShenQingRiQi, info.EndShenQingRiQi)
	}
	if info.ShouCiShouKuanId != nil {
		db = db.Where("shou_ci_shou_kuan_id = ?", info.ShouCiShouKuanId)
	}
	if info.ShouKuanMingCheng != "" {
		db = db.Where("shou_kuan_ming_cheng LIKE ?", "%"+info.ShouKuanMingCheng+"%")
	}
	if info.XiangMuMingCheng != nil {
		db = db.Where("xiang_mu_ming_cheng = ?", info.XiangMuMingCheng)
	}
	if info.YingShouLeiXing != nil {
		db = db.Where("ying_shou_lei_xing = ?", info.YingShouLeiXing)
	}
	if info.YingShouShiXiang != nil {
		db = db.Where("ying_shou_shi_xiang = ?", info.YingShouShiXiang)
	}
	if info.ZhuangTai != nil {
		db = db.Where("zhuang_tai = ?", info.ZhuangTai)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&yingShouKuanBiaos).Error
	return yingShouKuanBiaos, total, err
}
