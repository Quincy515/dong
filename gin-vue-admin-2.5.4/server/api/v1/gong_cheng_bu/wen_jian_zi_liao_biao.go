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

type WenJianZiLiaoBiaoApi struct {
}

var wenJianZiLiaoBiaoService = service.ServiceGroupApp.Gong_cheng_buServiceGroup.WenJianZiLiaoBiaoService

// CreateWenJianZiLiaoBiao 创建WenJianZiLiaoBiao
// @Tags WenJianZiLiaoBiao
// @Summary 创建WenJianZiLiaoBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.WenJianZiLiaoBiao true "创建WenJianZiLiaoBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wenJianZiLiaoBiao/createWenJianZiLiaoBiao [post]
func (wenJianZiLiaoBiaoApi *WenJianZiLiaoBiaoApi) CreateWenJianZiLiaoBiao(c *gin.Context) {
	var wenJianZiLiaoBiao gong_cheng_bu.WenJianZiLiaoBiao
	err := c.ShouldBindJSON(&wenJianZiLiaoBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wenJianZiLiaoBiaoService.CreateWenJianZiLiaoBiao(wenJianZiLiaoBiao); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWenJianZiLiaoBiao 删除WenJianZiLiaoBiao
// @Tags WenJianZiLiaoBiao
// @Summary 删除WenJianZiLiaoBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.WenJianZiLiaoBiao true "删除WenJianZiLiaoBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wenJianZiLiaoBiao/deleteWenJianZiLiaoBiao [delete]
func (wenJianZiLiaoBiaoApi *WenJianZiLiaoBiaoApi) DeleteWenJianZiLiaoBiao(c *gin.Context) {
	var wenJianZiLiaoBiao gong_cheng_bu.WenJianZiLiaoBiao
	err := c.ShouldBindJSON(&wenJianZiLiaoBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wenJianZiLiaoBiaoService.DeleteWenJianZiLiaoBiao(wenJianZiLiaoBiao); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWenJianZiLiaoBiaoByIds 批量删除WenJianZiLiaoBiao
// @Tags WenJianZiLiaoBiao
// @Summary 批量删除WenJianZiLiaoBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WenJianZiLiaoBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wenJianZiLiaoBiao/deleteWenJianZiLiaoBiaoByIds [delete]
func (wenJianZiLiaoBiaoApi *WenJianZiLiaoBiaoApi) DeleteWenJianZiLiaoBiaoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wenJianZiLiaoBiaoService.DeleteWenJianZiLiaoBiaoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWenJianZiLiaoBiao 更新WenJianZiLiaoBiao
// @Tags WenJianZiLiaoBiao
// @Summary 更新WenJianZiLiaoBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.WenJianZiLiaoBiao true "更新WenJianZiLiaoBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wenJianZiLiaoBiao/updateWenJianZiLiaoBiao [put]
func (wenJianZiLiaoBiaoApi *WenJianZiLiaoBiaoApi) UpdateWenJianZiLiaoBiao(c *gin.Context) {
	var wenJianZiLiaoBiao gong_cheng_bu.WenJianZiLiaoBiao
	err := c.ShouldBindJSON(&wenJianZiLiaoBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wenJianZiLiaoBiaoService.UpdateWenJianZiLiaoBiao(wenJianZiLiaoBiao); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWenJianZiLiaoBiao 用id查询WenJianZiLiaoBiao
// @Tags WenJianZiLiaoBiao
// @Summary 用id查询WenJianZiLiaoBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_bu.WenJianZiLiaoBiao true "用id查询WenJianZiLiaoBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wenJianZiLiaoBiao/findWenJianZiLiaoBiao [get]
func (wenJianZiLiaoBiaoApi *WenJianZiLiaoBiaoApi) FindWenJianZiLiaoBiao(c *gin.Context) {
	var wenJianZiLiaoBiao gong_cheng_bu.WenJianZiLiaoBiao
	err := c.ShouldBindQuery(&wenJianZiLiaoBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rewenJianZiLiaoBiao, err := wenJianZiLiaoBiaoService.GetWenJianZiLiaoBiao(wenJianZiLiaoBiao.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewenJianZiLiaoBiao": rewenJianZiLiaoBiao}, c)
	}
}

// GetWenJianZiLiaoBiaoList 分页获取WenJianZiLiaoBiao列表
// @Tags WenJianZiLiaoBiao
// @Summary 分页获取WenJianZiLiaoBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_buReq.WenJianZiLiaoBiaoSearch true "分页获取WenJianZiLiaoBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wenJianZiLiaoBiao/getWenJianZiLiaoBiaoList [get]
func (wenJianZiLiaoBiaoApi *WenJianZiLiaoBiaoApi) GetWenJianZiLiaoBiaoList(c *gin.Context) {
	var pageInfo gong_cheng_buReq.WenJianZiLiaoBiaoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wenJianZiLiaoBiaoService.GetWenJianZiLiaoBiaoInfoList(pageInfo); err != nil {
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
