package jing_ying_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/jing_ying_bu"
	jing_ying_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/jing_ying_bu/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type XiangMuDengJiBiaoApi struct {
}

var xiangMuDengJiBiaoService = service.ServiceGroupApp.Jing_ying_buServiceGroup.XiangMuDengJiBiaoService

// CreateXiangMuDengJiBiao 创建XiangMuDengJiBiao
// @Tags XiangMuDengJiBiao
// @Summary 创建XiangMuDengJiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body jing_ying_bu.XiangMuDengJiBiao true "创建XiangMuDengJiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xiangMuDengJiBiao/createXiangMuDengJiBiao [post]
func (xiangMuDengJiBiaoApi *XiangMuDengJiBiaoApi) CreateXiangMuDengJiBiao(c *gin.Context) {
	var xiangMuDengJiBiao jing_ying_bu.XiangMuDengJiBiao
	err := c.ShouldBindJSON(&xiangMuDengJiBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := xiangMuDengJiBiaoService.CreateXiangMuDengJiBiao(xiangMuDengJiBiao); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteXiangMuDengJiBiao 删除XiangMuDengJiBiao
// @Tags XiangMuDengJiBiao
// @Summary 删除XiangMuDengJiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body jing_ying_bu.XiangMuDengJiBiao true "删除XiangMuDengJiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xiangMuDengJiBiao/deleteXiangMuDengJiBiao [delete]
func (xiangMuDengJiBiaoApi *XiangMuDengJiBiaoApi) DeleteXiangMuDengJiBiao(c *gin.Context) {
	var xiangMuDengJiBiao jing_ying_bu.XiangMuDengJiBiao
	err := c.ShouldBindJSON(&xiangMuDengJiBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := xiangMuDengJiBiaoService.DeleteXiangMuDengJiBiao(xiangMuDengJiBiao); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteXiangMuDengJiBiaoByIds 批量删除XiangMuDengJiBiao
// @Tags XiangMuDengJiBiao
// @Summary 批量删除XiangMuDengJiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除XiangMuDengJiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /xiangMuDengJiBiao/deleteXiangMuDengJiBiaoByIds [delete]
func (xiangMuDengJiBiaoApi *XiangMuDengJiBiaoApi) DeleteXiangMuDengJiBiaoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := xiangMuDengJiBiaoService.DeleteXiangMuDengJiBiaoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateXiangMuDengJiBiao 更新XiangMuDengJiBiao
// @Tags XiangMuDengJiBiao
// @Summary 更新XiangMuDengJiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body jing_ying_bu.XiangMuDengJiBiao true "更新XiangMuDengJiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /xiangMuDengJiBiao/updateXiangMuDengJiBiao [put]
func (xiangMuDengJiBiaoApi *XiangMuDengJiBiaoApi) UpdateXiangMuDengJiBiao(c *gin.Context) {
	var xiangMuDengJiBiao jing_ying_bu.XiangMuDengJiBiao
	err := c.ShouldBindJSON(&xiangMuDengJiBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := xiangMuDengJiBiaoService.UpdateXiangMuDengJiBiao(xiangMuDengJiBiao); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindXiangMuDengJiBiao 用id查询XiangMuDengJiBiao
// @Tags XiangMuDengJiBiao
// @Summary 用id查询XiangMuDengJiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query jing_ying_bu.XiangMuDengJiBiao true "用id查询XiangMuDengJiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /xiangMuDengJiBiao/findXiangMuDengJiBiao [get]
func (xiangMuDengJiBiaoApi *XiangMuDengJiBiaoApi) FindXiangMuDengJiBiao(c *gin.Context) {
	var xiangMuDengJiBiao jing_ying_bu.XiangMuDengJiBiao
	err := c.ShouldBindQuery(&xiangMuDengJiBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rexiangMuDengJiBiao, err := xiangMuDengJiBiaoService.GetXiangMuDengJiBiao(xiangMuDengJiBiao.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rexiangMuDengJiBiao": rexiangMuDengJiBiao}, c)
	}
}

// GetXiangMuDengJiBiaoList 分页获取XiangMuDengJiBiao列表
// @Tags XiangMuDengJiBiao
// @Summary 分页获取XiangMuDengJiBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query jing_ying_buReq.XiangMuDengJiBiaoSearch true "分页获取XiangMuDengJiBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xiangMuDengJiBiao/getXiangMuDengJiBiaoList [get]
func (xiangMuDengJiBiaoApi *XiangMuDengJiBiaoApi) GetXiangMuDengJiBiaoList(c *gin.Context) {
	var pageInfo jing_ying_buReq.XiangMuDengJiBiaoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := xiangMuDengJiBiaoService.GetXiangMuDengJiBiaoInfoList(pageInfo); err != nil {
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
