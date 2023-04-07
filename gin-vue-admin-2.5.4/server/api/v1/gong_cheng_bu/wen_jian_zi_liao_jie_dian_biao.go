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

type WenJianZiLiaoJieDianBiaoApi struct {
}

var wenJianZiLiaoJieDianBiaoService = service.ServiceGroupApp.Gong_cheng_buServiceGroup.WenJianZiLiaoJieDianBiaoService

// CreateWenJianZiLiaoJieDianBiao 创建WenJianZiLiaoJieDianBiao
// @Tags WenJianZiLiaoJieDianBiao
// @Summary 创建WenJianZiLiaoJieDianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.WenJianZiLiaoJieDianBiao true "创建WenJianZiLiaoJieDianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wenJianZiLiaoJieDianBiao/createWenJianZiLiaoJieDianBiao [post]
func (wenJianZiLiaoJieDianBiaoApi *WenJianZiLiaoJieDianBiaoApi) CreateWenJianZiLiaoJieDianBiao(c *gin.Context) {
	var wenJianZiLiaoJieDianBiao gong_cheng_bu.WenJianZiLiaoJieDianBiao
	err := c.ShouldBindJSON(&wenJianZiLiaoJieDianBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wenJianZiLiaoJieDianBiaoService.CreateWenJianZiLiaoJieDianBiao(wenJianZiLiaoJieDianBiao); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWenJianZiLiaoJieDianBiao 删除WenJianZiLiaoJieDianBiao
// @Tags WenJianZiLiaoJieDianBiao
// @Summary 删除WenJianZiLiaoJieDianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.WenJianZiLiaoJieDianBiao true "删除WenJianZiLiaoJieDianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wenJianZiLiaoJieDianBiao/deleteWenJianZiLiaoJieDianBiao [delete]
func (wenJianZiLiaoJieDianBiaoApi *WenJianZiLiaoJieDianBiaoApi) DeleteWenJianZiLiaoJieDianBiao(c *gin.Context) {
	var wenJianZiLiaoJieDianBiao gong_cheng_bu.WenJianZiLiaoJieDianBiao
	err := c.ShouldBindJSON(&wenJianZiLiaoJieDianBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wenJianZiLiaoJieDianBiaoService.DeleteWenJianZiLiaoJieDianBiao(wenJianZiLiaoJieDianBiao); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWenJianZiLiaoJieDianBiaoByIds 批量删除WenJianZiLiaoJieDianBiao
// @Tags WenJianZiLiaoJieDianBiao
// @Summary 批量删除WenJianZiLiaoJieDianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WenJianZiLiaoJieDianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wenJianZiLiaoJieDianBiao/deleteWenJianZiLiaoJieDianBiaoByIds [delete]
func (wenJianZiLiaoJieDianBiaoApi *WenJianZiLiaoJieDianBiaoApi) DeleteWenJianZiLiaoJieDianBiaoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wenJianZiLiaoJieDianBiaoService.DeleteWenJianZiLiaoJieDianBiaoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWenJianZiLiaoJieDianBiao 更新WenJianZiLiaoJieDianBiao
// @Tags WenJianZiLiaoJieDianBiao
// @Summary 更新WenJianZiLiaoJieDianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.WenJianZiLiaoJieDianBiao true "更新WenJianZiLiaoJieDianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wenJianZiLiaoJieDianBiao/updateWenJianZiLiaoJieDianBiao [put]
func (wenJianZiLiaoJieDianBiaoApi *WenJianZiLiaoJieDianBiaoApi) UpdateWenJianZiLiaoJieDianBiao(c *gin.Context) {
	var wenJianZiLiaoJieDianBiao gong_cheng_bu.WenJianZiLiaoJieDianBiao
	err := c.ShouldBindJSON(&wenJianZiLiaoJieDianBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wenJianZiLiaoJieDianBiaoService.UpdateWenJianZiLiaoJieDianBiao(wenJianZiLiaoJieDianBiao); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWenJianZiLiaoJieDianBiao 用id查询WenJianZiLiaoJieDianBiao
// @Tags WenJianZiLiaoJieDianBiao
// @Summary 用id查询WenJianZiLiaoJieDianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_bu.WenJianZiLiaoJieDianBiao true "用id查询WenJianZiLiaoJieDianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wenJianZiLiaoJieDianBiao/findWenJianZiLiaoJieDianBiao [get]
func (wenJianZiLiaoJieDianBiaoApi *WenJianZiLiaoJieDianBiaoApi) FindWenJianZiLiaoJieDianBiao(c *gin.Context) {
	var wenJianZiLiaoJieDianBiao gong_cheng_bu.WenJianZiLiaoJieDianBiao
	err := c.ShouldBindQuery(&wenJianZiLiaoJieDianBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rewenJianZiLiaoJieDianBiao, err := wenJianZiLiaoJieDianBiaoService.GetWenJianZiLiaoJieDianBiao(wenJianZiLiaoJieDianBiao.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewenJianZiLiaoJieDianBiao": rewenJianZiLiaoJieDianBiao}, c)
	}
}

// GetWenJianZiLiaoJieDianBiaoList 分页获取WenJianZiLiaoJieDianBiao列表
// @Tags WenJianZiLiaoJieDianBiao
// @Summary 分页获取WenJianZiLiaoJieDianBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_buReq.WenJianZiLiaoJieDianBiaoSearch true "分页获取WenJianZiLiaoJieDianBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wenJianZiLiaoJieDianBiao/getWenJianZiLiaoJieDianBiaoList [get]
func (wenJianZiLiaoJieDianBiaoApi *WenJianZiLiaoJieDianBiaoApi) GetWenJianZiLiaoJieDianBiaoList(c *gin.Context) {
	var pageInfo gong_cheng_buReq.WenJianZiLiaoJieDianBiaoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wenJianZiLiaoJieDianBiaoService.GetWenJianZiLiaoJieDianBiaoInfoList(pageInfo); err != nil {
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
