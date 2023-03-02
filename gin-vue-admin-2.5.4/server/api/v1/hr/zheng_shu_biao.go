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

type ZhengShuBiaoApi struct {
}

var zhengShuBiaoService = service.ServiceGroupApp.HrServiceGroup.ZhengShuBiaoService

// CreateZhengShuBiao 创建ZhengShuBiao
// @Tags ZhengShuBiao
// @Summary 创建ZhengShuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.ZhengShuBiao true "创建ZhengShuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /zhengShuBiao/createZhengShuBiao [post]
func (zhengShuBiaoApi *ZhengShuBiaoApi) CreateZhengShuBiao(c *gin.Context) {
	var zhengShuBiao hr.ZhengShuBiao
	err := c.ShouldBindJSON(&zhengShuBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := zhengShuBiaoService.CreateZhengShuBiao(zhengShuBiao); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteZhengShuBiao 删除ZhengShuBiao
// @Tags ZhengShuBiao
// @Summary 删除ZhengShuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.ZhengShuBiao true "删除ZhengShuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /zhengShuBiao/deleteZhengShuBiao [delete]
func (zhengShuBiaoApi *ZhengShuBiaoApi) DeleteZhengShuBiao(c *gin.Context) {
	var zhengShuBiao hr.ZhengShuBiao
	err := c.ShouldBindJSON(&zhengShuBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := zhengShuBiaoService.DeleteZhengShuBiao(zhengShuBiao); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteZhengShuBiaoByIds 批量删除ZhengShuBiao
// @Tags ZhengShuBiao
// @Summary 批量删除ZhengShuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ZhengShuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /zhengShuBiao/deleteZhengShuBiaoByIds [delete]
func (zhengShuBiaoApi *ZhengShuBiaoApi) DeleteZhengShuBiaoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := zhengShuBiaoService.DeleteZhengShuBiaoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateZhengShuBiao 更新ZhengShuBiao
// @Tags ZhengShuBiao
// @Summary 更新ZhengShuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.ZhengShuBiao true "更新ZhengShuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /zhengShuBiao/updateZhengShuBiao [put]
func (zhengShuBiaoApi *ZhengShuBiaoApi) UpdateZhengShuBiao(c *gin.Context) {
	var zhengShuBiao hr.ZhengShuBiao
	err := c.ShouldBindJSON(&zhengShuBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := zhengShuBiaoService.UpdateZhengShuBiao(zhengShuBiao); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindZhengShuBiao 用id查询ZhengShuBiao
// @Tags ZhengShuBiao
// @Summary 用id查询ZhengShuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hr.ZhengShuBiao true "用id查询ZhengShuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /zhengShuBiao/findZhengShuBiao [get]
func (zhengShuBiaoApi *ZhengShuBiaoApi) FindZhengShuBiao(c *gin.Context) {
	var zhengShuBiao hr.ZhengShuBiao
	err := c.ShouldBindQuery(&zhengShuBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rezhengShuBiao, err := zhengShuBiaoService.GetZhengShuBiao(zhengShuBiao.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rezhengShuBiao": rezhengShuBiao}, c)
	}
}

// GetZhengShuBiaoList 分页获取ZhengShuBiao列表
// @Tags ZhengShuBiao
// @Summary 分页获取ZhengShuBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hrReq.ZhengShuBiaoSearch true "分页获取ZhengShuBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /zhengShuBiao/getZhengShuBiaoList [get]
func (zhengShuBiaoApi *ZhengShuBiaoApi) GetZhengShuBiaoList(c *gin.Context) {
	var pageInfo hrReq.ZhengShuBiaoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := zhengShuBiaoService.GetZhengShuBiaoInfoList(pageInfo); err != nil {
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
