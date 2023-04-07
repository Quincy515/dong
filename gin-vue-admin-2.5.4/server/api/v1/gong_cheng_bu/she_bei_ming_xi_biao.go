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

type SheBeiMingXiBiaoApi struct {
}

var sheBeiMingXiBiaoService = service.ServiceGroupApp.Gong_cheng_buServiceGroup.SheBeiMingXiBiaoService

// CreateSheBeiMingXiBiao 创建SheBeiMingXiBiao
// @Tags SheBeiMingXiBiao
// @Summary 创建SheBeiMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.SheBeiMingXiBiao true "创建SheBeiMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sheBeiMingXiBiao/createSheBeiMingXiBiao [post]
func (sheBeiMingXiBiaoApi *SheBeiMingXiBiaoApi) CreateSheBeiMingXiBiao(c *gin.Context) {
	var sheBeiMingXiBiao gong_cheng_bu.SheBeiMingXiBiao
	err := c.ShouldBindJSON(&sheBeiMingXiBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sheBeiMingXiBiaoService.CreateSheBeiMingXiBiao(sheBeiMingXiBiao); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSheBeiMingXiBiao 删除SheBeiMingXiBiao
// @Tags SheBeiMingXiBiao
// @Summary 删除SheBeiMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.SheBeiMingXiBiao true "删除SheBeiMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sheBeiMingXiBiao/deleteSheBeiMingXiBiao [delete]
func (sheBeiMingXiBiaoApi *SheBeiMingXiBiaoApi) DeleteSheBeiMingXiBiao(c *gin.Context) {
	var sheBeiMingXiBiao gong_cheng_bu.SheBeiMingXiBiao
	err := c.ShouldBindJSON(&sheBeiMingXiBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sheBeiMingXiBiaoService.DeleteSheBeiMingXiBiao(sheBeiMingXiBiao); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSheBeiMingXiBiaoByIds 批量删除SheBeiMingXiBiao
// @Tags SheBeiMingXiBiao
// @Summary 批量删除SheBeiMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SheBeiMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sheBeiMingXiBiao/deleteSheBeiMingXiBiaoByIds [delete]
func (sheBeiMingXiBiaoApi *SheBeiMingXiBiaoApi) DeleteSheBeiMingXiBiaoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sheBeiMingXiBiaoService.DeleteSheBeiMingXiBiaoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSheBeiMingXiBiao 更新SheBeiMingXiBiao
// @Tags SheBeiMingXiBiao
// @Summary 更新SheBeiMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.SheBeiMingXiBiao true "更新SheBeiMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sheBeiMingXiBiao/updateSheBeiMingXiBiao [put]
func (sheBeiMingXiBiaoApi *SheBeiMingXiBiaoApi) UpdateSheBeiMingXiBiao(c *gin.Context) {
	var sheBeiMingXiBiao gong_cheng_bu.SheBeiMingXiBiao
	err := c.ShouldBindJSON(&sheBeiMingXiBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sheBeiMingXiBiaoService.UpdateSheBeiMingXiBiao(sheBeiMingXiBiao); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSheBeiMingXiBiao 用id查询SheBeiMingXiBiao
// @Tags SheBeiMingXiBiao
// @Summary 用id查询SheBeiMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_bu.SheBeiMingXiBiao true "用id查询SheBeiMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sheBeiMingXiBiao/findSheBeiMingXiBiao [get]
func (sheBeiMingXiBiaoApi *SheBeiMingXiBiaoApi) FindSheBeiMingXiBiao(c *gin.Context) {
	var sheBeiMingXiBiao gong_cheng_bu.SheBeiMingXiBiao
	err := c.ShouldBindQuery(&sheBeiMingXiBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resheBeiMingXiBiao, err := sheBeiMingXiBiaoService.GetSheBeiMingXiBiao(sheBeiMingXiBiao.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resheBeiMingXiBiao": resheBeiMingXiBiao}, c)
	}
}

// GetSheBeiMingXiBiaoList 分页获取SheBeiMingXiBiao列表
// @Tags SheBeiMingXiBiao
// @Summary 分页获取SheBeiMingXiBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_buReq.SheBeiMingXiBiaoSearch true "分页获取SheBeiMingXiBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sheBeiMingXiBiao/getSheBeiMingXiBiaoList [get]
func (sheBeiMingXiBiaoApi *SheBeiMingXiBiaoApi) GetSheBeiMingXiBiaoList(c *gin.Context) {
	var pageInfo gong_cheng_buReq.SheBeiMingXiBiaoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sheBeiMingXiBiaoService.GetSheBeiMingXiBiaoInfoList(pageInfo); err != nil {
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
