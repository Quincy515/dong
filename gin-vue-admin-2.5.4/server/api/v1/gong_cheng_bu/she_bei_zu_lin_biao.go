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

type SheBeiZuLinBiaoApi struct {
}

var sheBeiZuLinBiaoService = service.ServiceGroupApp.Gong_cheng_buServiceGroup.SheBeiZuLinBiaoService

// CreateSheBeiZuLinBiao 创建SheBeiZuLinBiao
// @Tags SheBeiZuLinBiao
// @Summary 创建SheBeiZuLinBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.SheBeiZuLinBiao true "创建SheBeiZuLinBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sheBeiZuLinBiao/createSheBeiZuLinBiao [post]
func (sheBeiZuLinBiaoApi *SheBeiZuLinBiaoApi) CreateSheBeiZuLinBiao(c *gin.Context) {
	var sheBeiZuLinBiao gong_cheng_bu.SheBeiZuLinBiao
	err := c.ShouldBindJSON(&sheBeiZuLinBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sheBeiZuLinBiaoService.CreateSheBeiZuLinBiao(sheBeiZuLinBiao); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSheBeiZuLinBiao 删除SheBeiZuLinBiao
// @Tags SheBeiZuLinBiao
// @Summary 删除SheBeiZuLinBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.SheBeiZuLinBiao true "删除SheBeiZuLinBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sheBeiZuLinBiao/deleteSheBeiZuLinBiao [delete]
func (sheBeiZuLinBiaoApi *SheBeiZuLinBiaoApi) DeleteSheBeiZuLinBiao(c *gin.Context) {
	var sheBeiZuLinBiao gong_cheng_bu.SheBeiZuLinBiao
	err := c.ShouldBindJSON(&sheBeiZuLinBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sheBeiZuLinBiaoService.DeleteSheBeiZuLinBiao(sheBeiZuLinBiao); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSheBeiZuLinBiaoByIds 批量删除SheBeiZuLinBiao
// @Tags SheBeiZuLinBiao
// @Summary 批量删除SheBeiZuLinBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SheBeiZuLinBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sheBeiZuLinBiao/deleteSheBeiZuLinBiaoByIds [delete]
func (sheBeiZuLinBiaoApi *SheBeiZuLinBiaoApi) DeleteSheBeiZuLinBiaoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sheBeiZuLinBiaoService.DeleteSheBeiZuLinBiaoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSheBeiZuLinBiao 更新SheBeiZuLinBiao
// @Tags SheBeiZuLinBiao
// @Summary 更新SheBeiZuLinBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.SheBeiZuLinBiao true "更新SheBeiZuLinBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sheBeiZuLinBiao/updateSheBeiZuLinBiao [put]
func (sheBeiZuLinBiaoApi *SheBeiZuLinBiaoApi) UpdateSheBeiZuLinBiao(c *gin.Context) {
	var sheBeiZuLinBiao gong_cheng_bu.SheBeiZuLinBiao
	err := c.ShouldBindJSON(&sheBeiZuLinBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sheBeiZuLinBiaoService.UpdateSheBeiZuLinBiao(sheBeiZuLinBiao); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSheBeiZuLinBiao 用id查询SheBeiZuLinBiao
// @Tags SheBeiZuLinBiao
// @Summary 用id查询SheBeiZuLinBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_bu.SheBeiZuLinBiao true "用id查询SheBeiZuLinBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sheBeiZuLinBiao/findSheBeiZuLinBiao [get]
func (sheBeiZuLinBiaoApi *SheBeiZuLinBiaoApi) FindSheBeiZuLinBiao(c *gin.Context) {
	var sheBeiZuLinBiao gong_cheng_bu.SheBeiZuLinBiao
	err := c.ShouldBindQuery(&sheBeiZuLinBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resheBeiZuLinBiao, err := sheBeiZuLinBiaoService.GetSheBeiZuLinBiao(sheBeiZuLinBiao.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resheBeiZuLinBiao": resheBeiZuLinBiao}, c)
	}
}

// GetSheBeiZuLinBiaoList 分页获取SheBeiZuLinBiao列表
// @Tags SheBeiZuLinBiao
// @Summary 分页获取SheBeiZuLinBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_buReq.SheBeiZuLinBiaoSearch true "分页获取SheBeiZuLinBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sheBeiZuLinBiao/getSheBeiZuLinBiaoList [get]
func (sheBeiZuLinBiaoApi *SheBeiZuLinBiaoApi) GetSheBeiZuLinBiaoList(c *gin.Context) {
	var pageInfo gong_cheng_buReq.SheBeiZuLinBiaoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sheBeiZuLinBiaoService.GetSheBeiZuLinBiaoInfoList(pageInfo); err != nil {
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
