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

type CaiGouJiHuaMingXiBiaoApi struct {
}

var caiGouJiHuaMingXiBiaoService = service.ServiceGroupApp.Gong_cheng_buServiceGroup.CaiGouJiHuaMingXiBiaoService

// CreateCaiGouJiHuaMingXiBiao 创建CaiGouJiHuaMingXiBiao
// @Tags CaiGouJiHuaMingXiBiao
// @Summary 创建CaiGouJiHuaMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.CaiGouJiHuaMingXiBiao true "创建CaiGouJiHuaMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /caiGouJiHuaMingXiBiao/createCaiGouJiHuaMingXiBiao [post]
func (caiGouJiHuaMingXiBiaoApi *CaiGouJiHuaMingXiBiaoApi) CreateCaiGouJiHuaMingXiBiao(c *gin.Context) {
	var caiGouJiHuaMingXiBiao gong_cheng_bu.CaiGouJiHuaMingXiBiao
	err := c.ShouldBindJSON(&caiGouJiHuaMingXiBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := caiGouJiHuaMingXiBiaoService.CreateCaiGouJiHuaMingXiBiao(caiGouJiHuaMingXiBiao); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCaiGouJiHuaMingXiBiao 删除CaiGouJiHuaMingXiBiao
// @Tags CaiGouJiHuaMingXiBiao
// @Summary 删除CaiGouJiHuaMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.CaiGouJiHuaMingXiBiao true "删除CaiGouJiHuaMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /caiGouJiHuaMingXiBiao/deleteCaiGouJiHuaMingXiBiao [delete]
func (caiGouJiHuaMingXiBiaoApi *CaiGouJiHuaMingXiBiaoApi) DeleteCaiGouJiHuaMingXiBiao(c *gin.Context) {
	var caiGouJiHuaMingXiBiao gong_cheng_bu.CaiGouJiHuaMingXiBiao
	err := c.ShouldBindJSON(&caiGouJiHuaMingXiBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := caiGouJiHuaMingXiBiaoService.DeleteCaiGouJiHuaMingXiBiao(caiGouJiHuaMingXiBiao); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCaiGouJiHuaMingXiBiaoByIds 批量删除CaiGouJiHuaMingXiBiao
// @Tags CaiGouJiHuaMingXiBiao
// @Summary 批量删除CaiGouJiHuaMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CaiGouJiHuaMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /caiGouJiHuaMingXiBiao/deleteCaiGouJiHuaMingXiBiaoByIds [delete]
func (caiGouJiHuaMingXiBiaoApi *CaiGouJiHuaMingXiBiaoApi) DeleteCaiGouJiHuaMingXiBiaoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := caiGouJiHuaMingXiBiaoService.DeleteCaiGouJiHuaMingXiBiaoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCaiGouJiHuaMingXiBiao 更新CaiGouJiHuaMingXiBiao
// @Tags CaiGouJiHuaMingXiBiao
// @Summary 更新CaiGouJiHuaMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.CaiGouJiHuaMingXiBiao true "更新CaiGouJiHuaMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /caiGouJiHuaMingXiBiao/updateCaiGouJiHuaMingXiBiao [put]
func (caiGouJiHuaMingXiBiaoApi *CaiGouJiHuaMingXiBiaoApi) UpdateCaiGouJiHuaMingXiBiao(c *gin.Context) {
	var caiGouJiHuaMingXiBiao gong_cheng_bu.CaiGouJiHuaMingXiBiao
	err := c.ShouldBindJSON(&caiGouJiHuaMingXiBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := caiGouJiHuaMingXiBiaoService.UpdateCaiGouJiHuaMingXiBiao(caiGouJiHuaMingXiBiao); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCaiGouJiHuaMingXiBiao 用id查询CaiGouJiHuaMingXiBiao
// @Tags CaiGouJiHuaMingXiBiao
// @Summary 用id查询CaiGouJiHuaMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_bu.CaiGouJiHuaMingXiBiao true "用id查询CaiGouJiHuaMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /caiGouJiHuaMingXiBiao/findCaiGouJiHuaMingXiBiao [get]
func (caiGouJiHuaMingXiBiaoApi *CaiGouJiHuaMingXiBiaoApi) FindCaiGouJiHuaMingXiBiao(c *gin.Context) {
	var caiGouJiHuaMingXiBiao gong_cheng_bu.CaiGouJiHuaMingXiBiao
	err := c.ShouldBindQuery(&caiGouJiHuaMingXiBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recaiGouJiHuaMingXiBiao, err := caiGouJiHuaMingXiBiaoService.GetCaiGouJiHuaMingXiBiao(caiGouJiHuaMingXiBiao.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recaiGouJiHuaMingXiBiao": recaiGouJiHuaMingXiBiao}, c)
	}
}

// GetCaiGouJiHuaMingXiBiaoList 分页获取CaiGouJiHuaMingXiBiao列表
// @Tags CaiGouJiHuaMingXiBiao
// @Summary 分页获取CaiGouJiHuaMingXiBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_buReq.CaiGouJiHuaMingXiBiaoSearch true "分页获取CaiGouJiHuaMingXiBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /caiGouJiHuaMingXiBiao/getCaiGouJiHuaMingXiBiaoList [get]
func (caiGouJiHuaMingXiBiaoApi *CaiGouJiHuaMingXiBiaoApi) GetCaiGouJiHuaMingXiBiaoList(c *gin.Context) {
	var pageInfo gong_cheng_buReq.CaiGouJiHuaMingXiBiaoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := caiGouJiHuaMingXiBiaoService.GetCaiGouJiHuaMingXiBiaoInfoList(pageInfo); err != nil {
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
