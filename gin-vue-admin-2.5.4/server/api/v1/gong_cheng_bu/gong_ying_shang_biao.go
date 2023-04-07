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

type GongYingShangBiaoApi struct {
}

var gongYingShangBiaoService = service.ServiceGroupApp.Gong_cheng_buServiceGroup.GongYingShangBiaoService

// CreateGongYingShangBiao 创建GongYingShangBiao
// @Tags GongYingShangBiao
// @Summary 创建GongYingShangBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.GongYingShangBiao true "创建GongYingShangBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gongYingShangBiao/createGongYingShangBiao [post]
func (gongYingShangBiaoApi *GongYingShangBiaoApi) CreateGongYingShangBiao(c *gin.Context) {
	var gongYingShangBiao gong_cheng_bu.GongYingShangBiao
	err := c.ShouldBindJSON(&gongYingShangBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := gongYingShangBiaoService.CreateGongYingShangBiao(gongYingShangBiao); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteGongYingShangBiao 删除GongYingShangBiao
// @Tags GongYingShangBiao
// @Summary 删除GongYingShangBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.GongYingShangBiao true "删除GongYingShangBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gongYingShangBiao/deleteGongYingShangBiao [delete]
func (gongYingShangBiaoApi *GongYingShangBiaoApi) DeleteGongYingShangBiao(c *gin.Context) {
	var gongYingShangBiao gong_cheng_bu.GongYingShangBiao
	err := c.ShouldBindJSON(&gongYingShangBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := gongYingShangBiaoService.DeleteGongYingShangBiao(gongYingShangBiao); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteGongYingShangBiaoByIds 批量删除GongYingShangBiao
// @Tags GongYingShangBiao
// @Summary 批量删除GongYingShangBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GongYingShangBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /gongYingShangBiao/deleteGongYingShangBiaoByIds [delete]
func (gongYingShangBiaoApi *GongYingShangBiaoApi) DeleteGongYingShangBiaoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := gongYingShangBiaoService.DeleteGongYingShangBiaoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateGongYingShangBiao 更新GongYingShangBiao
// @Tags GongYingShangBiao
// @Summary 更新GongYingShangBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.GongYingShangBiao true "更新GongYingShangBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /gongYingShangBiao/updateGongYingShangBiao [put]
func (gongYingShangBiaoApi *GongYingShangBiaoApi) UpdateGongYingShangBiao(c *gin.Context) {
	var gongYingShangBiao gong_cheng_bu.GongYingShangBiao
	err := c.ShouldBindJSON(&gongYingShangBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := gongYingShangBiaoService.UpdateGongYingShangBiao(gongYingShangBiao); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindGongYingShangBiao 用id查询GongYingShangBiao
// @Tags GongYingShangBiao
// @Summary 用id查询GongYingShangBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_bu.GongYingShangBiao true "用id查询GongYingShangBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /gongYingShangBiao/findGongYingShangBiao [get]
func (gongYingShangBiaoApi *GongYingShangBiaoApi) FindGongYingShangBiao(c *gin.Context) {
	var gongYingShangBiao gong_cheng_bu.GongYingShangBiao
	err := c.ShouldBindQuery(&gongYingShangBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if regongYingShangBiao, err := gongYingShangBiaoService.GetGongYingShangBiao(gongYingShangBiao.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regongYingShangBiao": regongYingShangBiao}, c)
	}
}

// GetGongYingShangBiaoList 分页获取GongYingShangBiao列表
// @Tags GongYingShangBiao
// @Summary 分页获取GongYingShangBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_buReq.GongYingShangBiaoSearch true "分页获取GongYingShangBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gongYingShangBiao/getGongYingShangBiaoList [get]
func (gongYingShangBiaoApi *GongYingShangBiaoApi) GetGongYingShangBiaoList(c *gin.Context) {
	var pageInfo gong_cheng_buReq.GongYingShangBiaoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := gongYingShangBiaoService.GetGongYingShangBiaoInfoList(pageInfo); err != nil {
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
