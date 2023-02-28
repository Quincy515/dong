package jing_ying_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/jing_ying_bu"
	jing_ying_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/jing_ying_bu/request"
)

type XiangMuDengJiBiaoService struct {
}

// CreateXiangMuDengJiBiao 创建XiangMuDengJiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiangMuDengJiBiaoService *XiangMuDengJiBiaoService) CreateXiangMuDengJiBiao(xiangMuDengJiBiao jing_ying_bu.XiangMuDengJiBiao) (err error) {
	err = global.GVA_DB.Create(&xiangMuDengJiBiao).Error
	return err
}

// DeleteXiangMuDengJiBiao 删除XiangMuDengJiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiangMuDengJiBiaoService *XiangMuDengJiBiaoService) DeleteXiangMuDengJiBiao(xiangMuDengJiBiao jing_ying_bu.XiangMuDengJiBiao) (err error) {
	err = global.GVA_DB.Delete(&xiangMuDengJiBiao).Error
	return err
}

// DeleteXiangMuDengJiBiaoByIds 批量删除XiangMuDengJiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiangMuDengJiBiaoService *XiangMuDengJiBiaoService) DeleteXiangMuDengJiBiaoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]jing_ying_bu.XiangMuDengJiBiao{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateXiangMuDengJiBiao 更新XiangMuDengJiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiangMuDengJiBiaoService *XiangMuDengJiBiaoService) UpdateXiangMuDengJiBiao(xiangMuDengJiBiao jing_ying_bu.XiangMuDengJiBiao) (err error) {
	err = global.GVA_DB.Save(&xiangMuDengJiBiao).Error
	return err
}

// GetXiangMuDengJiBiao 根据id获取XiangMuDengJiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiangMuDengJiBiaoService *XiangMuDengJiBiaoService) GetXiangMuDengJiBiao(id uint) (xiangMuDengJiBiao jing_ying_bu.XiangMuDengJiBiao, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&xiangMuDengJiBiao).Error
	return
}

// GetXiangMuDengJiBiaoInfoList 分页获取XiangMuDengJiBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiangMuDengJiBiaoService *XiangMuDengJiBiaoService) GetXiangMuDengJiBiaoInfoList(info jing_ying_buReq.XiangMuDengJiBiaoSearch) (list []jing_ying_bu.XiangMuDengJiBiao, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&jing_ying_bu.XiangMuDengJiBiao{})
	var xiangMuDengJiBiaos []jing_ying_bu.XiangMuDengJiBiao
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Group != "" {
		db = db.Where("group = ?", info.Group)
	}
	if info.XiangMuMingCheng != "" {
		db = db.Where("xiang_mu_ming_cheng LIKE ?", "%"+info.XiangMuMingCheng+"%")
	}
	if info.XiangMuJianJie != "" {
		db = db.Where("xiang_mu_jian_jie LIKE ?", "%"+info.XiangMuJianJie+"%")
	}
	if info.XiangMuDengJiJinDu != "" {
		db = db.Where("xiang_mu_deng_ji_jin_du = ?", info.XiangMuDengJiJinDu)
	}
	if info.TouBiaoDanWeiMingCheng != "" {
		db = db.Where("tou_biao_dan_wei_ming_cheng LIKE ?", "%"+info.TouBiaoDanWeiMingCheng+"%")
	}
	if info.ZiZhiYaoQiu != nil {
		db = db.Where("zi_zhi_yao_qiu = ?", info.ZiZhiYaoQiu)
	}
	if info.DengJiYaoQiu != nil {
		db = db.Where("deng_ji_yao_qiu = ?", info.DengJiYaoQiu)
	}
	if info.SuoShuDiQu != nil {
		db = db.Where("suo_shu_di_qu = ?", info.SuoShuDiQu)
	}
	if info.YeWuFuZeRen != "" {
		db = db.Where("ye_wu_fu_ze_ren = ?", info.YeWuFuZeRen)
	}
	if info.BaoZhengJinXingShi != nil {
		db = db.Where("bao_zheng_jin_xing_shi = ?", info.BaoZhengJinXingShi)
	}
	if info.PingBiaoBanFa != nil {
		db = db.Where("ping_biao_ban_fa = ?", info.PingBiaoBanFa)
	}
	if info.BiaoShuLeiXing != nil {
		db = db.Where("biao_shu_lei_xing = ?", info.BiaoShuLeiXing)
	}
	if info.XiangMuZhuangTai != "" {
		db = db.Where("xiang_mu_zhuang_tai = ?", info.XiangMuZhuangTai)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("ID DESC").Limit(limit).Offset(offset).Find(&xiangMuDengJiBiaos).Error
	return xiangMuDengJiBiaos, total, err
}
