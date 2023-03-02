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

type ZhengShuJieChuJiLuBiaoApi struct {
}

var zhengShuJieChuJiLuBiaoService = service.ServiceGroupApp.HrServiceGroup.ZhengShuJieChuJiLuBiaoService

// CreateZhengShuJieChuJiLuBiao 创建ZhengShuJieChuJiLuBiao
// @Tags ZhengShuJieChuJiLuBiao
// @Summary 创建ZhengShuJieChuJiLuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.ZhengShuJieChuJiLuBiao true "创建ZhengShuJieChuJiLuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /zhengShuJieChuJiLuBiao/createZhengShuJieChuJiLuBiao [post]
func (zhengShuJieChuJiLuBiaoApi *ZhengShuJieChuJiLuBiaoApi) CreateZhengShuJieChuJiLuBiao(c *gin.Context) {
	var zhengShuJieChuJiLuBiao hr.ZhengShuJieChuJiLuBiao
	err := c.ShouldBindJSON(&zhengShuJieChuJiLuBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := zhengShuJieChuJiLuBiaoService.CreateZhengShuJieChuJiLuBiao(zhengShuJieChuJiLuBiao); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteZhengShuJieChuJiLuBiao 删除ZhengShuJieChuJiLuBiao
// @Tags ZhengShuJieChuJiLuBiao
// @Summary 删除ZhengShuJieChuJiLuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.ZhengShuJieChuJiLuBiao true "删除ZhengShuJieChuJiLuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /zhengShuJieChuJiLuBiao/deleteZhengShuJieChuJiLuBiao [delete]
func (zhengShuJieChuJiLuBiaoApi *ZhengShuJieChuJiLuBiaoApi) DeleteZhengShuJieChuJiLuBiao(c *gin.Context) {
	var zhengShuJieChuJiLuBiao hr.ZhengShuJieChuJiLuBiao
	err := c.ShouldBindJSON(&zhengShuJieChuJiLuBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := zhengShuJieChuJiLuBiaoService.DeleteZhengShuJieChuJiLuBiao(zhengShuJieChuJiLuBiao); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteZhengShuJieChuJiLuBiaoByIds 批量删除ZhengShuJieChuJiLuBiao
// @Tags ZhengShuJieChuJiLuBiao
// @Summary 批量删除ZhengShuJieChuJiLuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ZhengShuJieChuJiLuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /zhengShuJieChuJiLuBiao/deleteZhengShuJieChuJiLuBiaoByIds [delete]
func (zhengShuJieChuJiLuBiaoApi *ZhengShuJieChuJiLuBiaoApi) DeleteZhengShuJieChuJiLuBiaoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := zhengShuJieChuJiLuBiaoService.DeleteZhengShuJieChuJiLuBiaoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateZhengShuJieChuJiLuBiao 更新ZhengShuJieChuJiLuBiao
// @Tags ZhengShuJieChuJiLuBiao
// @Summary 更新ZhengShuJieChuJiLuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.ZhengShuJieChuJiLuBiao true "更新ZhengShuJieChuJiLuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /zhengShuJieChuJiLuBiao/updateZhengShuJieChuJiLuBiao [put]
func (zhengShuJieChuJiLuBiaoApi *ZhengShuJieChuJiLuBiaoApi) UpdateZhengShuJieChuJiLuBiao(c *gin.Context) {
	var zhengShuJieChuJiLuBiao hr.ZhengShuJieChuJiLuBiao
	err := c.ShouldBindJSON(&zhengShuJieChuJiLuBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := zhengShuJieChuJiLuBiaoService.UpdateZhengShuJieChuJiLuBiao(zhengShuJieChuJiLuBiao); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindZhengShuJieChuJiLuBiao 用id查询ZhengShuJieChuJiLuBiao
// @Tags ZhengShuJieChuJiLuBiao
// @Summary 用id查询ZhengShuJieChuJiLuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hr.ZhengShuJieChuJiLuBiao true "用id查询ZhengShuJieChuJiLuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /zhengShuJieChuJiLuBiao/findZhengShuJieChuJiLuBiao [get]
func (zhengShuJieChuJiLuBiaoApi *ZhengShuJieChuJiLuBiaoApi) FindZhengShuJieChuJiLuBiao(c *gin.Context) {
	var zhengShuJieChuJiLuBiao hr.ZhengShuJieChuJiLuBiao
	err := c.ShouldBindQuery(&zhengShuJieChuJiLuBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rezhengShuJieChuJiLuBiao, err := zhengShuJieChuJiLuBiaoService.GetZhengShuJieChuJiLuBiao(zhengShuJieChuJiLuBiao.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rezhengShuJieChuJiLuBiao": rezhengShuJieChuJiLuBiao}, c)
	}
}

// GetZhengShuJieChuJiLuBiaoList 分页获取ZhengShuJieChuJiLuBiao列表
// @Tags ZhengShuJieChuJiLuBiao
// @Summary 分页获取ZhengShuJieChuJiLuBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hrReq.ZhengShuJieChuJiLuBiaoSearch true "分页获取ZhengShuJieChuJiLuBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /zhengShuJieChuJiLuBiao/getZhengShuJieChuJiLuBiaoList [get]
func (zhengShuJieChuJiLuBiaoApi *ZhengShuJieChuJiLuBiaoApi) GetZhengShuJieChuJiLuBiaoList(c *gin.Context) {
	var pageInfo hrReq.ZhengShuJieChuJiLuBiaoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := zhengShuJieChuJiLuBiaoService.GetZhengShuJieChuJiLuBiaoInfoList(pageInfo); err != nil {
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
