package cai_wu_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cai_wu_bu"
	cai_wu_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/cai_wu_bu/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WaiBuFeiYongQingKuanBiaoApi struct {
}

var waiBuFeiYongQingKuanBiaoService = service.ServiceGroupApp.Cai_wu_buServiceGroup.WaiBuFeiYongQingKuanBiaoService

// CreateWaiBuFeiYongQingKuanBiao 创建WaiBuFeiYongQingKuanBiao
// @Tags WaiBuFeiYongQingKuanBiao
// @Summary 创建WaiBuFeiYongQingKuanBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cai_wu_bu.WaiBuFeiYongQingKuanBiao true "创建WaiBuFeiYongQingKuanBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /waiBuFeiYongQingKuanBiao/createWaiBuFeiYongQingKuanBiao [post]
func (waiBuFeiYongQingKuanBiaoApi *WaiBuFeiYongQingKuanBiaoApi) CreateWaiBuFeiYongQingKuanBiao(c *gin.Context) {
	var waiBuFeiYongQingKuanBiao cai_wu_bu.WaiBuFeiYongQingKuanBiao
	err := c.ShouldBindJSON(&waiBuFeiYongQingKuanBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := waiBuFeiYongQingKuanBiaoService.CreateWaiBuFeiYongQingKuanBiao(waiBuFeiYongQingKuanBiao); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWaiBuFeiYongQingKuanBiao 删除WaiBuFeiYongQingKuanBiao
// @Tags WaiBuFeiYongQingKuanBiao
// @Summary 删除WaiBuFeiYongQingKuanBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cai_wu_bu.WaiBuFeiYongQingKuanBiao true "删除WaiBuFeiYongQingKuanBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /waiBuFeiYongQingKuanBiao/deleteWaiBuFeiYongQingKuanBiao [delete]
func (waiBuFeiYongQingKuanBiaoApi *WaiBuFeiYongQingKuanBiaoApi) DeleteWaiBuFeiYongQingKuanBiao(c *gin.Context) {
	var waiBuFeiYongQingKuanBiao cai_wu_bu.WaiBuFeiYongQingKuanBiao
	err := c.ShouldBindJSON(&waiBuFeiYongQingKuanBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := waiBuFeiYongQingKuanBiaoService.DeleteWaiBuFeiYongQingKuanBiao(waiBuFeiYongQingKuanBiao); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWaiBuFeiYongQingKuanBiaoByIds 批量删除WaiBuFeiYongQingKuanBiao
// @Tags WaiBuFeiYongQingKuanBiao
// @Summary 批量删除WaiBuFeiYongQingKuanBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WaiBuFeiYongQingKuanBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /waiBuFeiYongQingKuanBiao/deleteWaiBuFeiYongQingKuanBiaoByIds [delete]
func (waiBuFeiYongQingKuanBiaoApi *WaiBuFeiYongQingKuanBiaoApi) DeleteWaiBuFeiYongQingKuanBiaoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := waiBuFeiYongQingKuanBiaoService.DeleteWaiBuFeiYongQingKuanBiaoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWaiBuFeiYongQingKuanBiao 更新WaiBuFeiYongQingKuanBiao
// @Tags WaiBuFeiYongQingKuanBiao
// @Summary 更新WaiBuFeiYongQingKuanBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cai_wu_bu.WaiBuFeiYongQingKuanBiao true "更新WaiBuFeiYongQingKuanBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /waiBuFeiYongQingKuanBiao/updateWaiBuFeiYongQingKuanBiao [put]
func (waiBuFeiYongQingKuanBiaoApi *WaiBuFeiYongQingKuanBiaoApi) UpdateWaiBuFeiYongQingKuanBiao(c *gin.Context) {
	var waiBuFeiYongQingKuanBiao cai_wu_bu.WaiBuFeiYongQingKuanBiao
	err := c.ShouldBindJSON(&waiBuFeiYongQingKuanBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := waiBuFeiYongQingKuanBiaoService.UpdateWaiBuFeiYongQingKuanBiao(waiBuFeiYongQingKuanBiao); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWaiBuFeiYongQingKuanBiao 用id查询WaiBuFeiYongQingKuanBiao
// @Tags WaiBuFeiYongQingKuanBiao
// @Summary 用id查询WaiBuFeiYongQingKuanBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cai_wu_bu.WaiBuFeiYongQingKuanBiao true "用id查询WaiBuFeiYongQingKuanBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /waiBuFeiYongQingKuanBiao/findWaiBuFeiYongQingKuanBiao [get]
func (waiBuFeiYongQingKuanBiaoApi *WaiBuFeiYongQingKuanBiaoApi) FindWaiBuFeiYongQingKuanBiao(c *gin.Context) {
	var waiBuFeiYongQingKuanBiao cai_wu_bu.WaiBuFeiYongQingKuanBiao
	err := c.ShouldBindQuery(&waiBuFeiYongQingKuanBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rewaiBuFeiYongQingKuanBiao, err := waiBuFeiYongQingKuanBiaoService.GetWaiBuFeiYongQingKuanBiao(waiBuFeiYongQingKuanBiao.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewaiBuFeiYongQingKuanBiao": rewaiBuFeiYongQingKuanBiao}, c)
	}
}

// GetWaiBuFeiYongQingKuanBiaoList 分页获取WaiBuFeiYongQingKuanBiao列表
// @Tags WaiBuFeiYongQingKuanBiao
// @Summary 分页获取WaiBuFeiYongQingKuanBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cai_wu_buReq.WaiBuFeiYongQingKuanBiaoSearch true "分页获取WaiBuFeiYongQingKuanBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /waiBuFeiYongQingKuanBiao/getWaiBuFeiYongQingKuanBiaoList [get]
func (waiBuFeiYongQingKuanBiaoApi *WaiBuFeiYongQingKuanBiaoApi) GetWaiBuFeiYongQingKuanBiaoList(c *gin.Context) {
	var pageInfo cai_wu_buReq.WaiBuFeiYongQingKuanBiaoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := waiBuFeiYongQingKuanBiaoService.GetWaiBuFeiYongQingKuanBiaoInfoList(pageInfo); err != nil {
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
