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

type XiangMuShiGongBiaoApi struct {
}

var xiangMuShiGongBiaoService = service.ServiceGroupApp.Gong_cheng_buServiceGroup.XiangMuShiGongBiaoService

// CreateXiangMuShiGongBiao 创建XiangMuShiGongBiao
// @Tags XiangMuShiGongBiao
// @Summary 创建XiangMuShiGongBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.XiangMuShiGongBiao true "创建XiangMuShiGongBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xiangMuShiGongBiao/createXiangMuShiGongBiao [post]
func (xiangMuShiGongBiaoApi *XiangMuShiGongBiaoApi) CreateXiangMuShiGongBiao(c *gin.Context) {
	var xiangMuShiGongBiao gong_cheng_bu.XiangMuShiGongBiao
	err := c.ShouldBindJSON(&xiangMuShiGongBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := xiangMuShiGongBiaoService.CreateXiangMuShiGongBiao(xiangMuShiGongBiao); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteXiangMuShiGongBiao 删除XiangMuShiGongBiao
// @Tags XiangMuShiGongBiao
// @Summary 删除XiangMuShiGongBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.XiangMuShiGongBiao true "删除XiangMuShiGongBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xiangMuShiGongBiao/deleteXiangMuShiGongBiao [delete]
func (xiangMuShiGongBiaoApi *XiangMuShiGongBiaoApi) DeleteXiangMuShiGongBiao(c *gin.Context) {
	var xiangMuShiGongBiao gong_cheng_bu.XiangMuShiGongBiao
	err := c.ShouldBindJSON(&xiangMuShiGongBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := xiangMuShiGongBiaoService.DeleteXiangMuShiGongBiao(xiangMuShiGongBiao); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteXiangMuShiGongBiaoByIds 批量删除XiangMuShiGongBiao
// @Tags XiangMuShiGongBiao
// @Summary 批量删除XiangMuShiGongBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除XiangMuShiGongBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /xiangMuShiGongBiao/deleteXiangMuShiGongBiaoByIds [delete]
func (xiangMuShiGongBiaoApi *XiangMuShiGongBiaoApi) DeleteXiangMuShiGongBiaoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := xiangMuShiGongBiaoService.DeleteXiangMuShiGongBiaoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateXiangMuShiGongBiao 更新XiangMuShiGongBiao
// @Tags XiangMuShiGongBiao
// @Summary 更新XiangMuShiGongBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.XiangMuShiGongBiao true "更新XiangMuShiGongBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /xiangMuShiGongBiao/updateXiangMuShiGongBiao [put]
func (xiangMuShiGongBiaoApi *XiangMuShiGongBiaoApi) UpdateXiangMuShiGongBiao(c *gin.Context) {
	var xiangMuShiGongBiao gong_cheng_bu.XiangMuShiGongBiao
	err := c.ShouldBindJSON(&xiangMuShiGongBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := xiangMuShiGongBiaoService.UpdateXiangMuShiGongBiao(xiangMuShiGongBiao); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindXiangMuShiGongBiao 用id查询XiangMuShiGongBiao
// @Tags XiangMuShiGongBiao
// @Summary 用id查询XiangMuShiGongBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_bu.XiangMuShiGongBiao true "用id查询XiangMuShiGongBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /xiangMuShiGongBiao/findXiangMuShiGongBiao [get]
func (xiangMuShiGongBiaoApi *XiangMuShiGongBiaoApi) FindXiangMuShiGongBiao(c *gin.Context) {
	var xiangMuShiGongBiao gong_cheng_bu.XiangMuShiGongBiao
	err := c.ShouldBindQuery(&xiangMuShiGongBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rexiangMuShiGongBiao, err := xiangMuShiGongBiaoService.GetXiangMuShiGongBiao(xiangMuShiGongBiao.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rexiangMuShiGongBiao": rexiangMuShiGongBiao}, c)
	}
}

// GetXiangMuShiGongBiaoList 分页获取XiangMuShiGongBiao列表
// @Tags XiangMuShiGongBiao
// @Summary 分页获取XiangMuShiGongBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_buReq.XiangMuShiGongBiaoSearch true "分页获取XiangMuShiGongBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xiangMuShiGongBiao/getXiangMuShiGongBiaoList [get]
func (xiangMuShiGongBiaoApi *XiangMuShiGongBiaoApi) GetXiangMuShiGongBiaoList(c *gin.Context) {
	var pageInfo gong_cheng_buReq.XiangMuShiGongBiaoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := xiangMuShiGongBiaoService.GetXiangMuShiGongBiaoInfoList(pageInfo); err != nil {
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
