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

type FenBaoShangBiaoApi struct {
}

var fenBaoShangBiaoService = service.ServiceGroupApp.Gong_cheng_buServiceGroup.FenBaoShangBiaoService

// CreateFenBaoShangBiao 创建FenBaoShangBiao
// @Tags FenBaoShangBiao
// @Summary 创建FenBaoShangBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.FenBaoShangBiao true "创建FenBaoShangBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fenBaoShangBiao/createFenBaoShangBiao [post]
func (fenBaoShangBiaoApi *FenBaoShangBiaoApi) CreateFenBaoShangBiao(c *gin.Context) {
	var fenBaoShangBiao gong_cheng_bu.FenBaoShangBiao
	err := c.ShouldBindJSON(&fenBaoShangBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fenBaoShangBiaoService.CreateFenBaoShangBiao(fenBaoShangBiao); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFenBaoShangBiao 删除FenBaoShangBiao
// @Tags FenBaoShangBiao
// @Summary 删除FenBaoShangBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.FenBaoShangBiao true "删除FenBaoShangBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fenBaoShangBiao/deleteFenBaoShangBiao [delete]
func (fenBaoShangBiaoApi *FenBaoShangBiaoApi) DeleteFenBaoShangBiao(c *gin.Context) {
	var fenBaoShangBiao gong_cheng_bu.FenBaoShangBiao
	err := c.ShouldBindJSON(&fenBaoShangBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fenBaoShangBiaoService.DeleteFenBaoShangBiao(fenBaoShangBiao); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFenBaoShangBiaoByIds 批量删除FenBaoShangBiao
// @Tags FenBaoShangBiao
// @Summary 批量删除FenBaoShangBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FenBaoShangBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fenBaoShangBiao/deleteFenBaoShangBiaoByIds [delete]
func (fenBaoShangBiaoApi *FenBaoShangBiaoApi) DeleteFenBaoShangBiaoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fenBaoShangBiaoService.DeleteFenBaoShangBiaoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFenBaoShangBiao 更新FenBaoShangBiao
// @Tags FenBaoShangBiao
// @Summary 更新FenBaoShangBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.FenBaoShangBiao true "更新FenBaoShangBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fenBaoShangBiao/updateFenBaoShangBiao [put]
func (fenBaoShangBiaoApi *FenBaoShangBiaoApi) UpdateFenBaoShangBiao(c *gin.Context) {
	var fenBaoShangBiao gong_cheng_bu.FenBaoShangBiao
	err := c.ShouldBindJSON(&fenBaoShangBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fenBaoShangBiaoService.UpdateFenBaoShangBiao(fenBaoShangBiao); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFenBaoShangBiao 用id查询FenBaoShangBiao
// @Tags FenBaoShangBiao
// @Summary 用id查询FenBaoShangBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_bu.FenBaoShangBiao true "用id查询FenBaoShangBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fenBaoShangBiao/findFenBaoShangBiao [get]
func (fenBaoShangBiaoApi *FenBaoShangBiaoApi) FindFenBaoShangBiao(c *gin.Context) {
	var fenBaoShangBiao gong_cheng_bu.FenBaoShangBiao
	err := c.ShouldBindQuery(&fenBaoShangBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refenBaoShangBiao, err := fenBaoShangBiaoService.GetFenBaoShangBiao(fenBaoShangBiao.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refenBaoShangBiao": refenBaoShangBiao}, c)
	}
}

// GetFenBaoShangBiaoList 分页获取FenBaoShangBiao列表
// @Tags FenBaoShangBiao
// @Summary 分页获取FenBaoShangBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_buReq.FenBaoShangBiaoSearch true "分页获取FenBaoShangBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fenBaoShangBiao/getFenBaoShangBiaoList [get]
func (fenBaoShangBiaoApi *FenBaoShangBiaoApi) GetFenBaoShangBiaoList(c *gin.Context) {
	var pageInfo gong_cheng_buReq.FenBaoShangBiaoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := fenBaoShangBiaoService.GetFenBaoShangBiaoInfoList(pageInfo); err != nil {
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
