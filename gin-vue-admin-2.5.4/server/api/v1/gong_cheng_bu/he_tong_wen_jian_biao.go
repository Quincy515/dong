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

type HeTongWenJianBiaoApi struct {
}

var heTongWenJianBiaoService = service.ServiceGroupApp.Gong_cheng_buServiceGroup.HeTongWenJianBiaoService

// CreateHeTongWenJianBiao 创建HeTongWenJianBiao
// @Tags HeTongWenJianBiao
// @Summary 创建HeTongWenJianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.HeTongWenJianBiao true "创建HeTongWenJianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /heTongWenJianBiao/createHeTongWenJianBiao [post]
func (heTongWenJianBiaoApi *HeTongWenJianBiaoApi) CreateHeTongWenJianBiao(c *gin.Context) {
	var heTongWenJianBiao gong_cheng_bu.HeTongWenJianBiao
	err := c.ShouldBindJSON(&heTongWenJianBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := heTongWenJianBiaoService.CreateHeTongWenJianBiao(heTongWenJianBiao); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHeTongWenJianBiao 删除HeTongWenJianBiao
// @Tags HeTongWenJianBiao
// @Summary 删除HeTongWenJianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.HeTongWenJianBiao true "删除HeTongWenJianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /heTongWenJianBiao/deleteHeTongWenJianBiao [delete]
func (heTongWenJianBiaoApi *HeTongWenJianBiaoApi) DeleteHeTongWenJianBiao(c *gin.Context) {
	var heTongWenJianBiao gong_cheng_bu.HeTongWenJianBiao
	err := c.ShouldBindJSON(&heTongWenJianBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := heTongWenJianBiaoService.DeleteHeTongWenJianBiao(heTongWenJianBiao); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHeTongWenJianBiaoByIds 批量删除HeTongWenJianBiao
// @Tags HeTongWenJianBiao
// @Summary 批量删除HeTongWenJianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HeTongWenJianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /heTongWenJianBiao/deleteHeTongWenJianBiaoByIds [delete]
func (heTongWenJianBiaoApi *HeTongWenJianBiaoApi) DeleteHeTongWenJianBiaoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := heTongWenJianBiaoService.DeleteHeTongWenJianBiaoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHeTongWenJianBiao 更新HeTongWenJianBiao
// @Tags HeTongWenJianBiao
// @Summary 更新HeTongWenJianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.HeTongWenJianBiao true "更新HeTongWenJianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /heTongWenJianBiao/updateHeTongWenJianBiao [put]
func (heTongWenJianBiaoApi *HeTongWenJianBiaoApi) UpdateHeTongWenJianBiao(c *gin.Context) {
	var heTongWenJianBiao gong_cheng_bu.HeTongWenJianBiao
	err := c.ShouldBindJSON(&heTongWenJianBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := heTongWenJianBiaoService.UpdateHeTongWenJianBiao(heTongWenJianBiao); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHeTongWenJianBiao 用id查询HeTongWenJianBiao
// @Tags HeTongWenJianBiao
// @Summary 用id查询HeTongWenJianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_bu.HeTongWenJianBiao true "用id查询HeTongWenJianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /heTongWenJianBiao/findHeTongWenJianBiao [get]
func (heTongWenJianBiaoApi *HeTongWenJianBiaoApi) FindHeTongWenJianBiao(c *gin.Context) {
	var heTongWenJianBiao gong_cheng_bu.HeTongWenJianBiao
	err := c.ShouldBindQuery(&heTongWenJianBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reheTongWenJianBiao, err := heTongWenJianBiaoService.GetHeTongWenJianBiao(heTongWenJianBiao.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reheTongWenJianBiao": reheTongWenJianBiao}, c)
	}
}

// GetHeTongWenJianBiaoList 分页获取HeTongWenJianBiao列表
// @Tags HeTongWenJianBiao
// @Summary 分页获取HeTongWenJianBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_buReq.HeTongWenJianBiaoSearch true "分页获取HeTongWenJianBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /heTongWenJianBiao/getHeTongWenJianBiaoList [get]
func (heTongWenJianBiaoApi *HeTongWenJianBiaoApi) GetHeTongWenJianBiaoList(c *gin.Context) {
	var pageInfo gong_cheng_buReq.HeTongWenJianBiaoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := heTongWenJianBiaoService.GetHeTongWenJianBiaoInfoList(pageInfo); err != nil {
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
