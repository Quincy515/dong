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

type QingJiaGuanLiApi struct {
}

var qingJiaGuanLiService = service.ServiceGroupApp.HrServiceGroup.QingJiaGuanLiService

// CreateQingJiaGuanLi 创建QingJiaGuanLi
// @Tags QingJiaGuanLi
// @Summary 创建QingJiaGuanLi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.QingJiaGuanLi true "创建QingJiaGuanLi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /qingJiaGuanLi/createQingJiaGuanLi [post]
func (qingJiaGuanLiApi *QingJiaGuanLiApi) CreateQingJiaGuanLi(c *gin.Context) {
	var qingJiaGuanLi hr.QingJiaGuanLi
	err := c.ShouldBindJSON(&qingJiaGuanLi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := qingJiaGuanLiService.CreateQingJiaGuanLi(qingJiaGuanLi); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteQingJiaGuanLi 删除QingJiaGuanLi
// @Tags QingJiaGuanLi
// @Summary 删除QingJiaGuanLi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.QingJiaGuanLi true "删除QingJiaGuanLi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /qingJiaGuanLi/deleteQingJiaGuanLi [delete]
func (qingJiaGuanLiApi *QingJiaGuanLiApi) DeleteQingJiaGuanLi(c *gin.Context) {
	var qingJiaGuanLi hr.QingJiaGuanLi
	err := c.ShouldBindJSON(&qingJiaGuanLi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := qingJiaGuanLiService.DeleteQingJiaGuanLi(qingJiaGuanLi); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteQingJiaGuanLiByIds 批量删除QingJiaGuanLi
// @Tags QingJiaGuanLi
// @Summary 批量删除QingJiaGuanLi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除QingJiaGuanLi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /qingJiaGuanLi/deleteQingJiaGuanLiByIds [delete]
func (qingJiaGuanLiApi *QingJiaGuanLiApi) DeleteQingJiaGuanLiByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := qingJiaGuanLiService.DeleteQingJiaGuanLiByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateQingJiaGuanLi 更新QingJiaGuanLi
// @Tags QingJiaGuanLi
// @Summary 更新QingJiaGuanLi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.QingJiaGuanLi true "更新QingJiaGuanLi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /qingJiaGuanLi/updateQingJiaGuanLi [put]
func (qingJiaGuanLiApi *QingJiaGuanLiApi) UpdateQingJiaGuanLi(c *gin.Context) {
	var qingJiaGuanLi hr.QingJiaGuanLi
	err := c.ShouldBindJSON(&qingJiaGuanLi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := qingJiaGuanLiService.UpdateQingJiaGuanLi(qingJiaGuanLi); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindQingJiaGuanLi 用id查询QingJiaGuanLi
// @Tags QingJiaGuanLi
// @Summary 用id查询QingJiaGuanLi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hr.QingJiaGuanLi true "用id查询QingJiaGuanLi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /qingJiaGuanLi/findQingJiaGuanLi [get]
func (qingJiaGuanLiApi *QingJiaGuanLiApi) FindQingJiaGuanLi(c *gin.Context) {
	var qingJiaGuanLi hr.QingJiaGuanLi
	err := c.ShouldBindQuery(&qingJiaGuanLi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reqingJiaGuanLi, err := qingJiaGuanLiService.GetQingJiaGuanLi(qingJiaGuanLi.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reqingJiaGuanLi": reqingJiaGuanLi}, c)
	}
}

// GetQingJiaGuanLiList 分页获取QingJiaGuanLi列表
// @Tags QingJiaGuanLi
// @Summary 分页获取QingJiaGuanLi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hrReq.QingJiaGuanLiSearch true "分页获取QingJiaGuanLi列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /qingJiaGuanLi/getQingJiaGuanLiList [get]
func (qingJiaGuanLiApi *QingJiaGuanLiApi) GetQingJiaGuanLiList(c *gin.Context) {
	var pageInfo hrReq.QingJiaGuanLiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := qingJiaGuanLiService.GetQingJiaGuanLiInfoList(pageInfo); err != nil {
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
