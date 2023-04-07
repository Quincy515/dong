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

type ShiGongJinDuBiaoApi struct {
}

var shiGongJinDuBiaoService = service.ServiceGroupApp.Gong_cheng_buServiceGroup.ShiGongJinDuBiaoService

// CreateShiGongJinDuBiao 创建ShiGongJinDuBiao
// @Tags ShiGongJinDuBiao
// @Summary 创建ShiGongJinDuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.ShiGongJinDuBiao true "创建ShiGongJinDuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shiGongJinDuBiao/createShiGongJinDuBiao [post]
func (shiGongJinDuBiaoApi *ShiGongJinDuBiaoApi) CreateShiGongJinDuBiao(c *gin.Context) {
	var shiGongJinDuBiao gong_cheng_bu.ShiGongJinDuBiao
	err := c.ShouldBindJSON(&shiGongJinDuBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shiGongJinDuBiaoService.CreateShiGongJinDuBiao(shiGongJinDuBiao); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteShiGongJinDuBiao 删除ShiGongJinDuBiao
// @Tags ShiGongJinDuBiao
// @Summary 删除ShiGongJinDuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.ShiGongJinDuBiao true "删除ShiGongJinDuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shiGongJinDuBiao/deleteShiGongJinDuBiao [delete]
func (shiGongJinDuBiaoApi *ShiGongJinDuBiaoApi) DeleteShiGongJinDuBiao(c *gin.Context) {
	var shiGongJinDuBiao gong_cheng_bu.ShiGongJinDuBiao
	err := c.ShouldBindJSON(&shiGongJinDuBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shiGongJinDuBiaoService.DeleteShiGongJinDuBiao(shiGongJinDuBiao); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteShiGongJinDuBiaoByIds 批量删除ShiGongJinDuBiao
// @Tags ShiGongJinDuBiao
// @Summary 批量删除ShiGongJinDuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ShiGongJinDuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /shiGongJinDuBiao/deleteShiGongJinDuBiaoByIds [delete]
func (shiGongJinDuBiaoApi *ShiGongJinDuBiaoApi) DeleteShiGongJinDuBiaoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shiGongJinDuBiaoService.DeleteShiGongJinDuBiaoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateShiGongJinDuBiao 更新ShiGongJinDuBiao
// @Tags ShiGongJinDuBiao
// @Summary 更新ShiGongJinDuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.ShiGongJinDuBiao true "更新ShiGongJinDuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shiGongJinDuBiao/updateShiGongJinDuBiao [put]
func (shiGongJinDuBiaoApi *ShiGongJinDuBiaoApi) UpdateShiGongJinDuBiao(c *gin.Context) {
	var shiGongJinDuBiao gong_cheng_bu.ShiGongJinDuBiao
	err := c.ShouldBindJSON(&shiGongJinDuBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shiGongJinDuBiaoService.UpdateShiGongJinDuBiao(shiGongJinDuBiao); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindShiGongJinDuBiao 用id查询ShiGongJinDuBiao
// @Tags ShiGongJinDuBiao
// @Summary 用id查询ShiGongJinDuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_bu.ShiGongJinDuBiao true "用id查询ShiGongJinDuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shiGongJinDuBiao/findShiGongJinDuBiao [get]
func (shiGongJinDuBiaoApi *ShiGongJinDuBiaoApi) FindShiGongJinDuBiao(c *gin.Context) {
	var shiGongJinDuBiao gong_cheng_bu.ShiGongJinDuBiao
	err := c.ShouldBindQuery(&shiGongJinDuBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reshiGongJinDuBiao, err := shiGongJinDuBiaoService.GetShiGongJinDuBiao(shiGongJinDuBiao.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reshiGongJinDuBiao": reshiGongJinDuBiao}, c)
	}
}

// GetShiGongJinDuBiaoList 分页获取ShiGongJinDuBiao列表
// @Tags ShiGongJinDuBiao
// @Summary 分页获取ShiGongJinDuBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_buReq.ShiGongJinDuBiaoSearch true "分页获取ShiGongJinDuBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shiGongJinDuBiao/getShiGongJinDuBiaoList [get]
func (shiGongJinDuBiaoApi *ShiGongJinDuBiaoApi) GetShiGongJinDuBiaoList(c *gin.Context) {
	var pageInfo gong_cheng_buReq.ShiGongJinDuBiaoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := shiGongJinDuBiaoService.GetShiGongJinDuBiaoInfoList(pageInfo); err != nil {
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
