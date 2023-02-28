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

type HuiYuanGuanLiApi struct {
}

var huiYuanGuanLiService = service.ServiceGroupApp.HrServiceGroup.HuiYuanGuanLiService

// CreateHuiYuanGuanLi 创建HuiYuanGuanLi
// @Tags HuiYuanGuanLi
// @Summary 创建HuiYuanGuanLi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.HuiYuanGuanLi true "创建HuiYuanGuanLi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /huiYuanGuanLi/createHuiYuanGuanLi [post]
func (huiYuanGuanLiApi *HuiYuanGuanLiApi) CreateHuiYuanGuanLi(c *gin.Context) {
	var huiYuanGuanLi hr.HuiYuanGuanLi
	err := c.ShouldBindJSON(&huiYuanGuanLi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := huiYuanGuanLiService.CreateHuiYuanGuanLi(huiYuanGuanLi); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHuiYuanGuanLi 删除HuiYuanGuanLi
// @Tags HuiYuanGuanLi
// @Summary 删除HuiYuanGuanLi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.HuiYuanGuanLi true "删除HuiYuanGuanLi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /huiYuanGuanLi/deleteHuiYuanGuanLi [delete]
func (huiYuanGuanLiApi *HuiYuanGuanLiApi) DeleteHuiYuanGuanLi(c *gin.Context) {
	var huiYuanGuanLi hr.HuiYuanGuanLi
	err := c.ShouldBindJSON(&huiYuanGuanLi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := huiYuanGuanLiService.DeleteHuiYuanGuanLi(huiYuanGuanLi); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHuiYuanGuanLiByIds 批量删除HuiYuanGuanLi
// @Tags HuiYuanGuanLi
// @Summary 批量删除HuiYuanGuanLi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HuiYuanGuanLi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /huiYuanGuanLi/deleteHuiYuanGuanLiByIds [delete]
func (huiYuanGuanLiApi *HuiYuanGuanLiApi) DeleteHuiYuanGuanLiByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := huiYuanGuanLiService.DeleteHuiYuanGuanLiByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHuiYuanGuanLi 更新HuiYuanGuanLi
// @Tags HuiYuanGuanLi
// @Summary 更新HuiYuanGuanLi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.HuiYuanGuanLi true "更新HuiYuanGuanLi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /huiYuanGuanLi/updateHuiYuanGuanLi [put]
func (huiYuanGuanLiApi *HuiYuanGuanLiApi) UpdateHuiYuanGuanLi(c *gin.Context) {
	var huiYuanGuanLi hr.HuiYuanGuanLi
	err := c.ShouldBindJSON(&huiYuanGuanLi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := huiYuanGuanLiService.UpdateHuiYuanGuanLi(huiYuanGuanLi); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHuiYuanGuanLi 用id查询HuiYuanGuanLi
// @Tags HuiYuanGuanLi
// @Summary 用id查询HuiYuanGuanLi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hr.HuiYuanGuanLi true "用id查询HuiYuanGuanLi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /huiYuanGuanLi/findHuiYuanGuanLi [get]
func (huiYuanGuanLiApi *HuiYuanGuanLiApi) FindHuiYuanGuanLi(c *gin.Context) {
	var huiYuanGuanLi hr.HuiYuanGuanLi
	err := c.ShouldBindQuery(&huiYuanGuanLi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehuiYuanGuanLi, err := huiYuanGuanLiService.GetHuiYuanGuanLi(huiYuanGuanLi.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehuiYuanGuanLi": rehuiYuanGuanLi}, c)
	}
}

// GetHuiYuanGuanLiList 分页获取HuiYuanGuanLi列表
// @Tags HuiYuanGuanLi
// @Summary 分页获取HuiYuanGuanLi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hrReq.HuiYuanGuanLiSearch true "分页获取HuiYuanGuanLi列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /huiYuanGuanLi/getHuiYuanGuanLiList [get]
func (huiYuanGuanLiApi *HuiYuanGuanLiApi) GetHuiYuanGuanLiList(c *gin.Context) {
	var pageInfo hrReq.HuiYuanGuanLiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := huiYuanGuanLiService.GetHuiYuanGuanLiInfoList(pageInfo); err != nil {
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
