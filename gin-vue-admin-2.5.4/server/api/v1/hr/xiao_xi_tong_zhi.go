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

type XiaoXiTongZhiApi struct {
}

var xiaoXiTongZhiService = service.ServiceGroupApp.HrServiceGroup.XiaoXiTongZhiService

// CreateXiaoXiTongZhi 创建XiaoXiTongZhi
// @Tags XiaoXiTongZhi
// @Summary 创建XiaoXiTongZhi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.XiaoXiTongZhi true "创建XiaoXiTongZhi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xiaoXiTongZhi/createXiaoXiTongZhi [post]
func (xiaoXiTongZhiApi *XiaoXiTongZhiApi) CreateXiaoXiTongZhi(c *gin.Context) {
	var xiaoXiTongZhi hr.XiaoXiTongZhi
	err := c.ShouldBindJSON(&xiaoXiTongZhi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := xiaoXiTongZhiService.CreateXiaoXiTongZhi(xiaoXiTongZhi); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteXiaoXiTongZhi 删除XiaoXiTongZhi
// @Tags XiaoXiTongZhi
// @Summary 删除XiaoXiTongZhi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.XiaoXiTongZhi true "删除XiaoXiTongZhi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xiaoXiTongZhi/deleteXiaoXiTongZhi [delete]
func (xiaoXiTongZhiApi *XiaoXiTongZhiApi) DeleteXiaoXiTongZhi(c *gin.Context) {
	var xiaoXiTongZhi hr.XiaoXiTongZhi
	err := c.ShouldBindJSON(&xiaoXiTongZhi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := xiaoXiTongZhiService.DeleteXiaoXiTongZhi(xiaoXiTongZhi); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteXiaoXiTongZhiByIds 批量删除XiaoXiTongZhi
// @Tags XiaoXiTongZhi
// @Summary 批量删除XiaoXiTongZhi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除XiaoXiTongZhi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /xiaoXiTongZhi/deleteXiaoXiTongZhiByIds [delete]
func (xiaoXiTongZhiApi *XiaoXiTongZhiApi) DeleteXiaoXiTongZhiByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := xiaoXiTongZhiService.DeleteXiaoXiTongZhiByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateXiaoXiTongZhi 更新XiaoXiTongZhi
// @Tags XiaoXiTongZhi
// @Summary 更新XiaoXiTongZhi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.XiaoXiTongZhi true "更新XiaoXiTongZhi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /xiaoXiTongZhi/updateXiaoXiTongZhi [put]
func (xiaoXiTongZhiApi *XiaoXiTongZhiApi) UpdateXiaoXiTongZhi(c *gin.Context) {
	var xiaoXiTongZhi hr.XiaoXiTongZhi
	err := c.ShouldBindJSON(&xiaoXiTongZhi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := xiaoXiTongZhiService.UpdateXiaoXiTongZhi(xiaoXiTongZhi); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindXiaoXiTongZhi 用id查询XiaoXiTongZhi
// @Tags XiaoXiTongZhi
// @Summary 用id查询XiaoXiTongZhi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hr.XiaoXiTongZhi true "用id查询XiaoXiTongZhi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /xiaoXiTongZhi/findXiaoXiTongZhi [get]
func (xiaoXiTongZhiApi *XiaoXiTongZhiApi) FindXiaoXiTongZhi(c *gin.Context) {
	var xiaoXiTongZhi hr.XiaoXiTongZhi
	err := c.ShouldBindQuery(&xiaoXiTongZhi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rexiaoXiTongZhi, err := xiaoXiTongZhiService.GetXiaoXiTongZhi(xiaoXiTongZhi.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rexiaoXiTongZhi": rexiaoXiTongZhi}, c)
	}
}

// GetXiaoXiTongZhiList 分页获取XiaoXiTongZhi列表
// @Tags XiaoXiTongZhi
// @Summary 分页获取XiaoXiTongZhi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hrReq.XiaoXiTongZhiSearch true "分页获取XiaoXiTongZhi列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xiaoXiTongZhi/getXiaoXiTongZhiList [get]
func (xiaoXiTongZhiApi *XiaoXiTongZhiApi) GetXiaoXiTongZhiList(c *gin.Context) {
	var pageInfo hrReq.XiaoXiTongZhiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := xiaoXiTongZhiService.GetXiaoXiTongZhiInfoList(pageInfo); err != nil {
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
