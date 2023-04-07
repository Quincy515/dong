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

type XiangMuZhuCaiBiaoApi struct {
}

var xiangMuZhuCaiBiaoService = service.ServiceGroupApp.Gong_cheng_buServiceGroup.XiangMuZhuCaiBiaoService

// CreateXiangMuZhuCaiBiao 创建XiangMuZhuCaiBiao
// @Tags XiangMuZhuCaiBiao
// @Summary 创建XiangMuZhuCaiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.XiangMuZhuCaiBiao true "创建XiangMuZhuCaiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xiangMuZhuCaiBiao/createXiangMuZhuCaiBiao [post]
func (xiangMuZhuCaiBiaoApi *XiangMuZhuCaiBiaoApi) CreateXiangMuZhuCaiBiao(c *gin.Context) {
	var xiangMuZhuCaiBiao gong_cheng_bu.XiangMuZhuCaiBiao
	err := c.ShouldBindJSON(&xiangMuZhuCaiBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := xiangMuZhuCaiBiaoService.CreateXiangMuZhuCaiBiao(xiangMuZhuCaiBiao); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteXiangMuZhuCaiBiao 删除XiangMuZhuCaiBiao
// @Tags XiangMuZhuCaiBiao
// @Summary 删除XiangMuZhuCaiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.XiangMuZhuCaiBiao true "删除XiangMuZhuCaiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xiangMuZhuCaiBiao/deleteXiangMuZhuCaiBiao [delete]
func (xiangMuZhuCaiBiaoApi *XiangMuZhuCaiBiaoApi) DeleteXiangMuZhuCaiBiao(c *gin.Context) {
	var xiangMuZhuCaiBiao gong_cheng_bu.XiangMuZhuCaiBiao
	err := c.ShouldBindJSON(&xiangMuZhuCaiBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := xiangMuZhuCaiBiaoService.DeleteXiangMuZhuCaiBiao(xiangMuZhuCaiBiao); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteXiangMuZhuCaiBiaoByIds 批量删除XiangMuZhuCaiBiao
// @Tags XiangMuZhuCaiBiao
// @Summary 批量删除XiangMuZhuCaiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除XiangMuZhuCaiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /xiangMuZhuCaiBiao/deleteXiangMuZhuCaiBiaoByIds [delete]
func (xiangMuZhuCaiBiaoApi *XiangMuZhuCaiBiaoApi) DeleteXiangMuZhuCaiBiaoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := xiangMuZhuCaiBiaoService.DeleteXiangMuZhuCaiBiaoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateXiangMuZhuCaiBiao 更新XiangMuZhuCaiBiao
// @Tags XiangMuZhuCaiBiao
// @Summary 更新XiangMuZhuCaiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.XiangMuZhuCaiBiao true "更新XiangMuZhuCaiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /xiangMuZhuCaiBiao/updateXiangMuZhuCaiBiao [put]
func (xiangMuZhuCaiBiaoApi *XiangMuZhuCaiBiaoApi) UpdateXiangMuZhuCaiBiao(c *gin.Context) {
	var xiangMuZhuCaiBiao gong_cheng_bu.XiangMuZhuCaiBiao
	err := c.ShouldBindJSON(&xiangMuZhuCaiBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := xiangMuZhuCaiBiaoService.UpdateXiangMuZhuCaiBiao(xiangMuZhuCaiBiao); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindXiangMuZhuCaiBiao 用id查询XiangMuZhuCaiBiao
// @Tags XiangMuZhuCaiBiao
// @Summary 用id查询XiangMuZhuCaiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_bu.XiangMuZhuCaiBiao true "用id查询XiangMuZhuCaiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /xiangMuZhuCaiBiao/findXiangMuZhuCaiBiao [get]
func (xiangMuZhuCaiBiaoApi *XiangMuZhuCaiBiaoApi) FindXiangMuZhuCaiBiao(c *gin.Context) {
	var xiangMuZhuCaiBiao gong_cheng_bu.XiangMuZhuCaiBiao
	err := c.ShouldBindQuery(&xiangMuZhuCaiBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rexiangMuZhuCaiBiao, err := xiangMuZhuCaiBiaoService.GetXiangMuZhuCaiBiao(xiangMuZhuCaiBiao.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rexiangMuZhuCaiBiao": rexiangMuZhuCaiBiao}, c)
	}
}

// GetXiangMuZhuCaiBiaoList 分页获取XiangMuZhuCaiBiao列表
// @Tags XiangMuZhuCaiBiao
// @Summary 分页获取XiangMuZhuCaiBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_buReq.XiangMuZhuCaiBiaoSearch true "分页获取XiangMuZhuCaiBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xiangMuZhuCaiBiao/getXiangMuZhuCaiBiaoList [get]
func (xiangMuZhuCaiBiaoApi *XiangMuZhuCaiBiaoApi) GetXiangMuZhuCaiBiaoList(c *gin.Context) {
	var pageInfo gong_cheng_buReq.XiangMuZhuCaiBiaoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := xiangMuZhuCaiBiaoService.GetXiangMuZhuCaiBiaoInfoList(pageInfo); err != nil {
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
