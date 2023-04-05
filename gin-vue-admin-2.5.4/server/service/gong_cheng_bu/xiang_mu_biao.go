package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    gong_cheng_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu/request"
    "gorm.io/gorm"
)

type XiangMuBiaoService struct {
}

// CreateXiangMuBiao 创建XiangMuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiangMuBiaoService *XiangMuBiaoService) CreateXiangMuBiao(xiangMuBiao gong_cheng_bu.XiangMuBiao) (err error) {
	err = global.GVA_DB.Create(&xiangMuBiao).Error
	return err
}

// DeleteXiangMuBiao 删除XiangMuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiangMuBiaoService *XiangMuBiaoService)DeleteXiangMuBiao(xiangMuBiao gong_cheng_bu.XiangMuBiao) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&gong_cheng_bu.XiangMuBiao{}).Where("id = ?", xiangMuBiao.ID).Update("deleted_by", xiangMuBiao.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&xiangMuBiao).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteXiangMuBiaoByIds 批量删除XiangMuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiangMuBiaoService *XiangMuBiaoService)DeleteXiangMuBiaoByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&gong_cheng_bu.XiangMuBiao{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&gong_cheng_bu.XiangMuBiao{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateXiangMuBiao 更新XiangMuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiangMuBiaoService *XiangMuBiaoService)UpdateXiangMuBiao(xiangMuBiao gong_cheng_bu.XiangMuBiao) (err error) {
	err = global.GVA_DB.Save(&xiangMuBiao).Error
	return err
}

// GetXiangMuBiao 根据id获取XiangMuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiangMuBiaoService *XiangMuBiaoService)GetXiangMuBiao(id uint) (xiangMuBiao gong_cheng_bu.XiangMuBiao, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&xiangMuBiao).Error
	return
}

// GetXiangMuBiaoInfoList 分页获取XiangMuBiao记录
// Author [piexlmax](https://github.com/piexlmax)
func (xiangMuBiaoService *XiangMuBiaoService)GetXiangMuBiaoInfoList(info gong_cheng_buReq.XiangMuBiaoSearch) (list []gong_cheng_bu.XiangMuBiao, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&gong_cheng_bu.XiangMuBiao{})
    var xiangMuBiaos []gong_cheng_bu.XiangMuBiao
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.XiangMuBianHao != nil {
        db = db.Where("xiang_mu_bian_hao = ?",info.XiangMuBianHao)
    }
    if info.BaoZhangJinZhuangTai != nil {
        db = db.Where("bao_zhang_jin_zhuang_tai = ?",info.BaoZhangJinZhuangTai)
    }
    if info.BaoZhengJinZhuangTai != nil {
        db = db.Where("bao_zheng_jin_zhuang_tai = ?",info.BaoZhengJinZhuangTai)
    }
    if info.FuKuanFangShi != nil {
        db = db.Where("fu_kuan_fang_shi = ?",info.FuKuanFangShi)
    }
    if info.JieSuanZhuangTai != nil {
        db = db.Where("jie_suan_zhuang_tai = ?",info.JieSuanZhuangTai)
    }
    if info.XiangMuZhuangTai != nil {
        db = db.Where("xiang_mu_zhuang_tai = ?",info.XiangMuZhuangTai)
    }
    if info.XiangMuJingLi != nil {
        db = db.Where("xiang_mu_jing_li = ?",info.XiangMuJingLi)
    }
    if info.XiangMuZhuGuan != nil {
        db = db.Where("xiang_mu_zhu_guan = ?",info.XiangMuZhuGuan)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&xiangMuBiaos).Error
	return  xiangMuBiaos, total, err
}
