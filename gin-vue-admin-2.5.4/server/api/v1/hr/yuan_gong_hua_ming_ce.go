package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hr"
	hrReq "github.com/flipped-aurora/gin-vue-admin/server/model/hr/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type YuanGongHuaMingCeApi struct {
}

var yuanGongHuaMingCeService = service.ServiceGroupApp.HrServiceGroup.YuanGongHuaMingCeService

// CreateYuanGongHuaMingCe 创建YuanGongHuaMingCe
// @Tags YuanGongHuaMingCe
// @Summary 创建YuanGongHuaMingCe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.YuanGongHuaMingCe true "创建YuanGongHuaMingCe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yuanGongHuaMingCe/createYuanGongHuaMingCe [post]
func (yuanGongHuaMingCeApi *YuanGongHuaMingCeApi) CreateYuanGongHuaMingCe(c *gin.Context) {
	var yuanGongHuaMingCe hr.YuanGongHuaMingCe
	err := c.ShouldBindJSON(&yuanGongHuaMingCe)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := yuanGongHuaMingCeService.CreateYuanGongHuaMingCe(yuanGongHuaMingCe); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteYuanGongHuaMingCe 删除YuanGongHuaMingCe
// @Tags YuanGongHuaMingCe
// @Summary 删除YuanGongHuaMingCe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.YuanGongHuaMingCe true "删除YuanGongHuaMingCe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /yuanGongHuaMingCe/deleteYuanGongHuaMingCe [delete]
func (yuanGongHuaMingCeApi *YuanGongHuaMingCeApi) DeleteYuanGongHuaMingCe(c *gin.Context) {
	var yuanGongHuaMingCe hr.YuanGongHuaMingCe
	err := c.ShouldBindJSON(&yuanGongHuaMingCe)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := yuanGongHuaMingCeService.DeleteYuanGongHuaMingCe(yuanGongHuaMingCe); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteYuanGongHuaMingCeByIds 批量删除YuanGongHuaMingCe
// @Tags YuanGongHuaMingCe
// @Summary 批量删除YuanGongHuaMingCe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除YuanGongHuaMingCe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /yuanGongHuaMingCe/deleteYuanGongHuaMingCeByIds [delete]
func (yuanGongHuaMingCeApi *YuanGongHuaMingCeApi) DeleteYuanGongHuaMingCeByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := yuanGongHuaMingCeService.DeleteYuanGongHuaMingCeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateYuanGongHuaMingCe 更新YuanGongHuaMingCe
// @Tags YuanGongHuaMingCe
// @Summary 更新YuanGongHuaMingCe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.YuanGongHuaMingCe true "更新YuanGongHuaMingCe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /yuanGongHuaMingCe/updateYuanGongHuaMingCe [put]
func (yuanGongHuaMingCeApi *YuanGongHuaMingCeApi) UpdateYuanGongHuaMingCe(c *gin.Context) {
	var yuanGongHuaMingCe hr.YuanGongHuaMingCe
	err := c.ShouldBindJSON(&yuanGongHuaMingCe)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := yuanGongHuaMingCeService.UpdateYuanGongHuaMingCe(yuanGongHuaMingCe); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindYuanGongHuaMingCe 用id查询YuanGongHuaMingCe
// @Tags YuanGongHuaMingCe
// @Summary 用id查询YuanGongHuaMingCe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hr.YuanGongHuaMingCe true "用id查询YuanGongHuaMingCe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /yuanGongHuaMingCe/findYuanGongHuaMingCe [get]
func (yuanGongHuaMingCeApi *YuanGongHuaMingCeApi) FindYuanGongHuaMingCe(c *gin.Context) {
	var yuanGongHuaMingCe hr.YuanGongHuaMingCe
	err := c.ShouldBindQuery(&yuanGongHuaMingCe)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reyuanGongHuaMingCe, err := yuanGongHuaMingCeService.GetYuanGongHuaMingCe(yuanGongHuaMingCe.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reyuanGongHuaMingCe": reyuanGongHuaMingCe}, c)
	}
}

// GetYuanGongHuaMingCeList 分页获取YuanGongHuaMingCe列表
// @Tags YuanGongHuaMingCe
// @Summary 分页获取YuanGongHuaMingCe列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hrReq.YuanGongHuaMingCeSearch true "分页获取YuanGongHuaMingCe列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yuanGongHuaMingCe/getYuanGongHuaMingCeList [get]
func (yuanGongHuaMingCeApi *YuanGongHuaMingCeApi) GetYuanGongHuaMingCeList(c *gin.Context) {
	var pageInfo hrReq.YuanGongHuaMingCeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := yuanGongHuaMingCeService.GetYuanGongHuaMingCeInfoList(pageInfo); err != nil {
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
