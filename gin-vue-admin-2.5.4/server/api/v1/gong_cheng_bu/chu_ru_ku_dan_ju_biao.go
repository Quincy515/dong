package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu"
	gong_cheng_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ChuRuKuDanJuBiaoApi struct {
}

var chuRuKuDanJuBiaoService = service.ServiceGroupApp.Gong_cheng_buServiceGroup.ChuRuKuDanJuBiaoService

// CreateChuRuKuDanJuBiao 创建ChuRuKuDanJuBiao
// @Tags ChuRuKuDanJuBiao
// @Summary 创建ChuRuKuDanJuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.ChuRuKuDanJuBiao true "创建ChuRuKuDanJuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /chuRuKuDanJuBiao/createChuRuKuDanJuBiao [post]
func (chuRuKuDanJuBiaoApi *ChuRuKuDanJuBiaoApi) CreateChuRuKuDanJuBiao(c *gin.Context) {
	var chuRuKuDanJuBiao gong_cheng_bu.ChuRuKuDanJuBiao
	err := c.ShouldBindJSON(&chuRuKuDanJuBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := chuRuKuDanJuBiaoService.CreateChuRuKuDanJuBiao(chuRuKuDanJuBiao); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteChuRuKuDanJuBiao 删除ChuRuKuDanJuBiao
// @Tags ChuRuKuDanJuBiao
// @Summary 删除ChuRuKuDanJuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.ChuRuKuDanJuBiao true "删除ChuRuKuDanJuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /chuRuKuDanJuBiao/deleteChuRuKuDanJuBiao [delete]
func (chuRuKuDanJuBiaoApi *ChuRuKuDanJuBiaoApi) DeleteChuRuKuDanJuBiao(c *gin.Context) {
	var chuRuKuDanJuBiao gong_cheng_bu.ChuRuKuDanJuBiao
	err := c.ShouldBindJSON(&chuRuKuDanJuBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := chuRuKuDanJuBiaoService.DeleteChuRuKuDanJuBiao(chuRuKuDanJuBiao); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteChuRuKuDanJuBiaoByIds 批量删除ChuRuKuDanJuBiao
// @Tags ChuRuKuDanJuBiao
// @Summary 批量删除ChuRuKuDanJuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ChuRuKuDanJuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /chuRuKuDanJuBiao/deleteChuRuKuDanJuBiaoByIds [delete]
func (chuRuKuDanJuBiaoApi *ChuRuKuDanJuBiaoApi) DeleteChuRuKuDanJuBiaoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := chuRuKuDanJuBiaoService.DeleteChuRuKuDanJuBiaoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateChuRuKuDanJuBiao 更新ChuRuKuDanJuBiao
// @Tags ChuRuKuDanJuBiao
// @Summary 更新ChuRuKuDanJuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.ChuRuKuDanJuBiao true "更新ChuRuKuDanJuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /chuRuKuDanJuBiao/updateChuRuKuDanJuBiao [put]
func (chuRuKuDanJuBiaoApi *ChuRuKuDanJuBiaoApi) UpdateChuRuKuDanJuBiao(c *gin.Context) {
	var chuRuKuDanJuBiao gong_cheng_bu.ChuRuKuDanJuBiao
	err := c.ShouldBindJSON(&chuRuKuDanJuBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := chuRuKuDanJuBiaoService.UpdateChuRuKuDanJuBiao(chuRuKuDanJuBiao); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindChuRuKuDanJuBiao 用id查询ChuRuKuDanJuBiao
// @Tags ChuRuKuDanJuBiao
// @Summary 用id查询ChuRuKuDanJuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_bu.ChuRuKuDanJuBiao true "用id查询ChuRuKuDanJuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /chuRuKuDanJuBiao/findChuRuKuDanJuBiao [get]
func (chuRuKuDanJuBiaoApi *ChuRuKuDanJuBiaoApi) FindChuRuKuDanJuBiao(c *gin.Context) {
	var chuRuKuDanJuBiao gong_cheng_bu.ChuRuKuDanJuBiao
	err := c.ShouldBindQuery(&chuRuKuDanJuBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rechuRuKuDanJuBiao, err := chuRuKuDanJuBiaoService.GetChuRuKuDanJuBiao(chuRuKuDanJuBiao.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rechuRuKuDanJuBiao": rechuRuKuDanJuBiao}, c)
	}
}

// GetChuRuKuDanJuBiaoList 分页获取ChuRuKuDanJuBiao列表
// @Tags ChuRuKuDanJuBiao
// @Summary 分页获取ChuRuKuDanJuBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_buReq.ChuRuKuDanJuBiaoSearch true "分页获取ChuRuKuDanJuBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /chuRuKuDanJuBiao/getChuRuKuDanJuBiaoList [get]
func (chuRuKuDanJuBiaoApi *ChuRuKuDanJuBiaoApi) GetChuRuKuDanJuBiaoList(c *gin.Context) {
	var pageInfo gong_cheng_buReq.ChuRuKuDanJuBiaoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := chuRuKuDanJuBiaoService.GetChuRuKuDanJuBiaoInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
