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

type NeiBuBaoXiaoBiaoApi struct {
}

var neiBuBaoXiaoBiaoService = service.ServiceGroupApp.Cai_wu_buServiceGroup.NeiBuBaoXiaoBiaoService

// CreateNeiBuBaoXiaoBiao 创建NeiBuBaoXiaoBiao
// @Tags NeiBuBaoXiaoBiao
// @Summary 创建NeiBuBaoXiaoBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cai_wu_bu.NeiBuBaoXiaoBiao true "创建NeiBuBaoXiaoBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /neiBuBaoXiaoBiao/createNeiBuBaoXiaoBiao [post]
func (neiBuBaoXiaoBiaoApi *NeiBuBaoXiaoBiaoApi) CreateNeiBuBaoXiaoBiao(c *gin.Context) {
	var neiBuBaoXiaoBiao cai_wu_bu.NeiBuBaoXiaoBiao
	err := c.ShouldBindJSON(&neiBuBaoXiaoBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := neiBuBaoXiaoBiaoService.CreateNeiBuBaoXiaoBiao(neiBuBaoXiaoBiao); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteNeiBuBaoXiaoBiao 删除NeiBuBaoXiaoBiao
// @Tags NeiBuBaoXiaoBiao
// @Summary 删除NeiBuBaoXiaoBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cai_wu_bu.NeiBuBaoXiaoBiao true "删除NeiBuBaoXiaoBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /neiBuBaoXiaoBiao/deleteNeiBuBaoXiaoBiao [delete]
func (neiBuBaoXiaoBiaoApi *NeiBuBaoXiaoBiaoApi) DeleteNeiBuBaoXiaoBiao(c *gin.Context) {
	var neiBuBaoXiaoBiao cai_wu_bu.NeiBuBaoXiaoBiao
	err := c.ShouldBindJSON(&neiBuBaoXiaoBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := neiBuBaoXiaoBiaoService.DeleteNeiBuBaoXiaoBiao(neiBuBaoXiaoBiao); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteNeiBuBaoXiaoBiaoByIds 批量删除NeiBuBaoXiaoBiao
// @Tags NeiBuBaoXiaoBiao
// @Summary 批量删除NeiBuBaoXiaoBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除NeiBuBaoXiaoBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /neiBuBaoXiaoBiao/deleteNeiBuBaoXiaoBiaoByIds [delete]
func (neiBuBaoXiaoBiaoApi *NeiBuBaoXiaoBiaoApi) DeleteNeiBuBaoXiaoBiaoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := neiBuBaoXiaoBiaoService.DeleteNeiBuBaoXiaoBiaoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateNeiBuBaoXiaoBiao 更新NeiBuBaoXiaoBiao
// @Tags NeiBuBaoXiaoBiao
// @Summary 更新NeiBuBaoXiaoBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cai_wu_bu.NeiBuBaoXiaoBiao true "更新NeiBuBaoXiaoBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /neiBuBaoXiaoBiao/updateNeiBuBaoXiaoBiao [put]
func (neiBuBaoXiaoBiaoApi *NeiBuBaoXiaoBiaoApi) UpdateNeiBuBaoXiaoBiao(c *gin.Context) {
	var neiBuBaoXiaoBiao cai_wu_bu.NeiBuBaoXiaoBiao
	err := c.ShouldBindJSON(&neiBuBaoXiaoBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := neiBuBaoXiaoBiaoService.UpdateNeiBuBaoXiaoBiao(neiBuBaoXiaoBiao); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindNeiBuBaoXiaoBiao 用id查询NeiBuBaoXiaoBiao
// @Tags NeiBuBaoXiaoBiao
// @Summary 用id查询NeiBuBaoXiaoBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cai_wu_bu.NeiBuBaoXiaoBiao true "用id查询NeiBuBaoXiaoBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /neiBuBaoXiaoBiao/findNeiBuBaoXiaoBiao [get]
func (neiBuBaoXiaoBiaoApi *NeiBuBaoXiaoBiaoApi) FindNeiBuBaoXiaoBiao(c *gin.Context) {
	var neiBuBaoXiaoBiao cai_wu_bu.NeiBuBaoXiaoBiao
	err := c.ShouldBindQuery(&neiBuBaoXiaoBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reneiBuBaoXiaoBiao, err := neiBuBaoXiaoBiaoService.GetNeiBuBaoXiaoBiao(neiBuBaoXiaoBiao.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reneiBuBaoXiaoBiao": reneiBuBaoXiaoBiao}, c)
	}
}

// GetNeiBuBaoXiaoBiaoList 分页获取NeiBuBaoXiaoBiao列表
// @Tags NeiBuBaoXiaoBiao
// @Summary 分页获取NeiBuBaoXiaoBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cai_wu_buReq.NeiBuBaoXiaoBiaoSearch true "分页获取NeiBuBaoXiaoBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /neiBuBaoXiaoBiao/getNeiBuBaoXiaoBiaoList [get]
func (neiBuBaoXiaoBiaoApi *NeiBuBaoXiaoBiaoApi) GetNeiBuBaoXiaoBiaoList(c *gin.Context) {
	var pageInfo cai_wu_buReq.NeiBuBaoXiaoBiaoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := neiBuBaoXiaoBiaoService.GetNeiBuBaoXiaoBiaoInfoList(pageInfo); err != nil {
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
