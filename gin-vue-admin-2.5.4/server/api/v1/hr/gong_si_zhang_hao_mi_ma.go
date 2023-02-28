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

type GongSiZhangHaoMiMaApi struct {
}

var gongSiZhangHaoMiMaService = service.ServiceGroupApp.HrServiceGroup.GongSiZhangHaoMiMaService

// CreateGongSiZhangHaoMiMa 创建GongSiZhangHaoMiMa
// @Tags GongSiZhangHaoMiMa
// @Summary 创建GongSiZhangHaoMiMa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.GongSiZhangHaoMiMa true "创建GongSiZhangHaoMiMa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gongSiZhangHaoMiMa/createGongSiZhangHaoMiMa [post]
func (gongSiZhangHaoMiMaApi *GongSiZhangHaoMiMaApi) CreateGongSiZhangHaoMiMa(c *gin.Context) {
	var gongSiZhangHaoMiMa hr.GongSiZhangHaoMiMa
	err := c.ShouldBindJSON(&gongSiZhangHaoMiMa)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := gongSiZhangHaoMiMaService.CreateGongSiZhangHaoMiMa(gongSiZhangHaoMiMa); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteGongSiZhangHaoMiMa 删除GongSiZhangHaoMiMa
// @Tags GongSiZhangHaoMiMa
// @Summary 删除GongSiZhangHaoMiMa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.GongSiZhangHaoMiMa true "删除GongSiZhangHaoMiMa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gongSiZhangHaoMiMa/deleteGongSiZhangHaoMiMa [delete]
func (gongSiZhangHaoMiMaApi *GongSiZhangHaoMiMaApi) DeleteGongSiZhangHaoMiMa(c *gin.Context) {
	var gongSiZhangHaoMiMa hr.GongSiZhangHaoMiMa
	err := c.ShouldBindJSON(&gongSiZhangHaoMiMa)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := gongSiZhangHaoMiMaService.DeleteGongSiZhangHaoMiMa(gongSiZhangHaoMiMa); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteGongSiZhangHaoMiMaByIds 批量删除GongSiZhangHaoMiMa
// @Tags GongSiZhangHaoMiMa
// @Summary 批量删除GongSiZhangHaoMiMa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GongSiZhangHaoMiMa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /gongSiZhangHaoMiMa/deleteGongSiZhangHaoMiMaByIds [delete]
func (gongSiZhangHaoMiMaApi *GongSiZhangHaoMiMaApi) DeleteGongSiZhangHaoMiMaByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := gongSiZhangHaoMiMaService.DeleteGongSiZhangHaoMiMaByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateGongSiZhangHaoMiMa 更新GongSiZhangHaoMiMa
// @Tags GongSiZhangHaoMiMa
// @Summary 更新GongSiZhangHaoMiMa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.GongSiZhangHaoMiMa true "更新GongSiZhangHaoMiMa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /gongSiZhangHaoMiMa/updateGongSiZhangHaoMiMa [put]
func (gongSiZhangHaoMiMaApi *GongSiZhangHaoMiMaApi) UpdateGongSiZhangHaoMiMa(c *gin.Context) {
	var gongSiZhangHaoMiMa hr.GongSiZhangHaoMiMa
	err := c.ShouldBindJSON(&gongSiZhangHaoMiMa)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := gongSiZhangHaoMiMaService.UpdateGongSiZhangHaoMiMa(gongSiZhangHaoMiMa); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindGongSiZhangHaoMiMa 用id查询GongSiZhangHaoMiMa
// @Tags GongSiZhangHaoMiMa
// @Summary 用id查询GongSiZhangHaoMiMa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hr.GongSiZhangHaoMiMa true "用id查询GongSiZhangHaoMiMa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /gongSiZhangHaoMiMa/findGongSiZhangHaoMiMa [get]
func (gongSiZhangHaoMiMaApi *GongSiZhangHaoMiMaApi) FindGongSiZhangHaoMiMa(c *gin.Context) {
	var gongSiZhangHaoMiMa hr.GongSiZhangHaoMiMa
	err := c.ShouldBindQuery(&gongSiZhangHaoMiMa)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if regongSiZhangHaoMiMa, err := gongSiZhangHaoMiMaService.GetGongSiZhangHaoMiMa(gongSiZhangHaoMiMa.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regongSiZhangHaoMiMa": regongSiZhangHaoMiMa}, c)
	}
}

// GetGongSiZhangHaoMiMaList 分页获取GongSiZhangHaoMiMa列表
// @Tags GongSiZhangHaoMiMa
// @Summary 分页获取GongSiZhangHaoMiMa列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hrReq.GongSiZhangHaoMiMaSearch true "分页获取GongSiZhangHaoMiMa列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gongSiZhangHaoMiMa/getGongSiZhangHaoMiMaList [get]
func (gongSiZhangHaoMiMaApi *GongSiZhangHaoMiMaApi) GetGongSiZhangHaoMiMaList(c *gin.Context) {
	var pageInfo hrReq.GongSiZhangHaoMiMaSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := gongSiZhangHaoMiMaService.GetGongSiZhangHaoMiMaInfoList(pageInfo); err != nil {
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
