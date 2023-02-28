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

type GongZhangDengJiApi struct {
}

var gongZhangDengJiService = service.ServiceGroupApp.HrServiceGroup.GongZhangDengJiService

// CreateGongZhangDengJi 创建GongZhangDengJi
// @Tags GongZhangDengJi
// @Summary 创建GongZhangDengJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.GongZhangDengJi true "创建GongZhangDengJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gongZhangDengJi/createGongZhangDengJi [post]
func (gongZhangDengJiApi *GongZhangDengJiApi) CreateGongZhangDengJi(c *gin.Context) {
	var gongZhangDengJi hr.GongZhangDengJi
	err := c.ShouldBindJSON(&gongZhangDengJi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := gongZhangDengJiService.CreateGongZhangDengJi(gongZhangDengJi); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteGongZhangDengJi 删除GongZhangDengJi
// @Tags GongZhangDengJi
// @Summary 删除GongZhangDengJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.GongZhangDengJi true "删除GongZhangDengJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gongZhangDengJi/deleteGongZhangDengJi [delete]
func (gongZhangDengJiApi *GongZhangDengJiApi) DeleteGongZhangDengJi(c *gin.Context) {
	var gongZhangDengJi hr.GongZhangDengJi
	err := c.ShouldBindJSON(&gongZhangDengJi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := gongZhangDengJiService.DeleteGongZhangDengJi(gongZhangDengJi); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteGongZhangDengJiByIds 批量删除GongZhangDengJi
// @Tags GongZhangDengJi
// @Summary 批量删除GongZhangDengJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GongZhangDengJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /gongZhangDengJi/deleteGongZhangDengJiByIds [delete]
func (gongZhangDengJiApi *GongZhangDengJiApi) DeleteGongZhangDengJiByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := gongZhangDengJiService.DeleteGongZhangDengJiByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateGongZhangDengJi 更新GongZhangDengJi
// @Tags GongZhangDengJi
// @Summary 更新GongZhangDengJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.GongZhangDengJi true "更新GongZhangDengJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /gongZhangDengJi/updateGongZhangDengJi [put]
func (gongZhangDengJiApi *GongZhangDengJiApi) UpdateGongZhangDengJi(c *gin.Context) {
	var gongZhangDengJi hr.GongZhangDengJi
	err := c.ShouldBindJSON(&gongZhangDengJi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := gongZhangDengJiService.UpdateGongZhangDengJi(gongZhangDengJi); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindGongZhangDengJi 用id查询GongZhangDengJi
// @Tags GongZhangDengJi
// @Summary 用id查询GongZhangDengJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hr.GongZhangDengJi true "用id查询GongZhangDengJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /gongZhangDengJi/findGongZhangDengJi [get]
func (gongZhangDengJiApi *GongZhangDengJiApi) FindGongZhangDengJi(c *gin.Context) {
	var gongZhangDengJi hr.GongZhangDengJi
	err := c.ShouldBindQuery(&gongZhangDengJi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if regongZhangDengJi, err := gongZhangDengJiService.GetGongZhangDengJi(gongZhangDengJi.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regongZhangDengJi": regongZhangDengJi}, c)
	}
}

// GetGongZhangDengJiList 分页获取GongZhangDengJi列表
// @Tags GongZhangDengJi
// @Summary 分页获取GongZhangDengJi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hrReq.GongZhangDengJiSearch true "分页获取GongZhangDengJi列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gongZhangDengJi/getGongZhangDengJiList [get]
func (gongZhangDengJiApi *GongZhangDengJiApi) GetGongZhangDengJiList(c *gin.Context) {
	var pageInfo hrReq.GongZhangDengJiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := gongZhangDengJiService.GetGongZhangDengJiInfoList(pageInfo); err != nil {
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
