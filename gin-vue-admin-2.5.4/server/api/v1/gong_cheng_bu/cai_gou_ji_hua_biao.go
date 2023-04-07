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

type CaiGouJiHuaBiaoApi struct {
}

var caiGouJiHuaBiaoService = service.ServiceGroupApp.Gong_cheng_buServiceGroup.CaiGouJiHuaBiaoService

// CreateCaiGouJiHuaBiao 创建CaiGouJiHuaBiao
// @Tags CaiGouJiHuaBiao
// @Summary 创建CaiGouJiHuaBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.CaiGouJiHuaBiao true "创建CaiGouJiHuaBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /caiGouJiHuaBiao/createCaiGouJiHuaBiao [post]
func (caiGouJiHuaBiaoApi *CaiGouJiHuaBiaoApi) CreateCaiGouJiHuaBiao(c *gin.Context) {
	var caiGouJiHuaBiao gong_cheng_bu.CaiGouJiHuaBiao
	err := c.ShouldBindJSON(&caiGouJiHuaBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := caiGouJiHuaBiaoService.CreateCaiGouJiHuaBiao(caiGouJiHuaBiao); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCaiGouJiHuaBiao 删除CaiGouJiHuaBiao
// @Tags CaiGouJiHuaBiao
// @Summary 删除CaiGouJiHuaBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.CaiGouJiHuaBiao true "删除CaiGouJiHuaBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /caiGouJiHuaBiao/deleteCaiGouJiHuaBiao [delete]
func (caiGouJiHuaBiaoApi *CaiGouJiHuaBiaoApi) DeleteCaiGouJiHuaBiao(c *gin.Context) {
	var caiGouJiHuaBiao gong_cheng_bu.CaiGouJiHuaBiao
	err := c.ShouldBindJSON(&caiGouJiHuaBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := caiGouJiHuaBiaoService.DeleteCaiGouJiHuaBiao(caiGouJiHuaBiao); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCaiGouJiHuaBiaoByIds 批量删除CaiGouJiHuaBiao
// @Tags CaiGouJiHuaBiao
// @Summary 批量删除CaiGouJiHuaBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CaiGouJiHuaBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /caiGouJiHuaBiao/deleteCaiGouJiHuaBiaoByIds [delete]
func (caiGouJiHuaBiaoApi *CaiGouJiHuaBiaoApi) DeleteCaiGouJiHuaBiaoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := caiGouJiHuaBiaoService.DeleteCaiGouJiHuaBiaoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCaiGouJiHuaBiao 更新CaiGouJiHuaBiao
// @Tags CaiGouJiHuaBiao
// @Summary 更新CaiGouJiHuaBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.CaiGouJiHuaBiao true "更新CaiGouJiHuaBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /caiGouJiHuaBiao/updateCaiGouJiHuaBiao [put]
func (caiGouJiHuaBiaoApi *CaiGouJiHuaBiaoApi) UpdateCaiGouJiHuaBiao(c *gin.Context) {
	var caiGouJiHuaBiao gong_cheng_bu.CaiGouJiHuaBiao
	err := c.ShouldBindJSON(&caiGouJiHuaBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := caiGouJiHuaBiaoService.UpdateCaiGouJiHuaBiao(caiGouJiHuaBiao); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCaiGouJiHuaBiao 用id查询CaiGouJiHuaBiao
// @Tags CaiGouJiHuaBiao
// @Summary 用id查询CaiGouJiHuaBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_bu.CaiGouJiHuaBiao true "用id查询CaiGouJiHuaBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /caiGouJiHuaBiao/findCaiGouJiHuaBiao [get]
func (caiGouJiHuaBiaoApi *CaiGouJiHuaBiaoApi) FindCaiGouJiHuaBiao(c *gin.Context) {
	var caiGouJiHuaBiao gong_cheng_bu.CaiGouJiHuaBiao
	err := c.ShouldBindQuery(&caiGouJiHuaBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recaiGouJiHuaBiao, err := caiGouJiHuaBiaoService.GetCaiGouJiHuaBiao(caiGouJiHuaBiao.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recaiGouJiHuaBiao": recaiGouJiHuaBiao}, c)
	}
}

// GetCaiGouJiHuaBiaoList 分页获取CaiGouJiHuaBiao列表
// @Tags CaiGouJiHuaBiao
// @Summary 分页获取CaiGouJiHuaBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_buReq.CaiGouJiHuaBiaoSearch true "分页获取CaiGouJiHuaBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /caiGouJiHuaBiao/getCaiGouJiHuaBiaoList [get]
func (caiGouJiHuaBiaoApi *CaiGouJiHuaBiaoApi) GetCaiGouJiHuaBiaoList(c *gin.Context) {
	var pageInfo gong_cheng_buReq.CaiGouJiHuaBiaoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := caiGouJiHuaBiaoService.GetCaiGouJiHuaBiaoInfoList(pageInfo); err != nil {
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
