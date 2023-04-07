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

type ChuRuKuMingXiBiaoApi struct {
}

var chuRuKuMingXiBiaoService = service.ServiceGroupApp.Gong_cheng_buServiceGroup.ChuRuKuMingXiBiaoService

// CreateChuRuKuMingXiBiao 创建ChuRuKuMingXiBiao
// @Tags ChuRuKuMingXiBiao
// @Summary 创建ChuRuKuMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.ChuRuKuMingXiBiao true "创建ChuRuKuMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /chuRuKuMingXiBiao/createChuRuKuMingXiBiao [post]
func (chuRuKuMingXiBiaoApi *ChuRuKuMingXiBiaoApi) CreateChuRuKuMingXiBiao(c *gin.Context) {
	var chuRuKuMingXiBiao gong_cheng_bu.ChuRuKuMingXiBiao
	err := c.ShouldBindJSON(&chuRuKuMingXiBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := chuRuKuMingXiBiaoService.CreateChuRuKuMingXiBiao(chuRuKuMingXiBiao); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteChuRuKuMingXiBiao 删除ChuRuKuMingXiBiao
// @Tags ChuRuKuMingXiBiao
// @Summary 删除ChuRuKuMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.ChuRuKuMingXiBiao true "删除ChuRuKuMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /chuRuKuMingXiBiao/deleteChuRuKuMingXiBiao [delete]
func (chuRuKuMingXiBiaoApi *ChuRuKuMingXiBiaoApi) DeleteChuRuKuMingXiBiao(c *gin.Context) {
	var chuRuKuMingXiBiao gong_cheng_bu.ChuRuKuMingXiBiao
	err := c.ShouldBindJSON(&chuRuKuMingXiBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := chuRuKuMingXiBiaoService.DeleteChuRuKuMingXiBiao(chuRuKuMingXiBiao); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteChuRuKuMingXiBiaoByIds 批量删除ChuRuKuMingXiBiao
// @Tags ChuRuKuMingXiBiao
// @Summary 批量删除ChuRuKuMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ChuRuKuMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /chuRuKuMingXiBiao/deleteChuRuKuMingXiBiaoByIds [delete]
func (chuRuKuMingXiBiaoApi *ChuRuKuMingXiBiaoApi) DeleteChuRuKuMingXiBiaoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := chuRuKuMingXiBiaoService.DeleteChuRuKuMingXiBiaoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateChuRuKuMingXiBiao 更新ChuRuKuMingXiBiao
// @Tags ChuRuKuMingXiBiao
// @Summary 更新ChuRuKuMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.ChuRuKuMingXiBiao true "更新ChuRuKuMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /chuRuKuMingXiBiao/updateChuRuKuMingXiBiao [put]
func (chuRuKuMingXiBiaoApi *ChuRuKuMingXiBiaoApi) UpdateChuRuKuMingXiBiao(c *gin.Context) {
	var chuRuKuMingXiBiao gong_cheng_bu.ChuRuKuMingXiBiao
	err := c.ShouldBindJSON(&chuRuKuMingXiBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := chuRuKuMingXiBiaoService.UpdateChuRuKuMingXiBiao(chuRuKuMingXiBiao); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindChuRuKuMingXiBiao 用id查询ChuRuKuMingXiBiao
// @Tags ChuRuKuMingXiBiao
// @Summary 用id查询ChuRuKuMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_bu.ChuRuKuMingXiBiao true "用id查询ChuRuKuMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /chuRuKuMingXiBiao/findChuRuKuMingXiBiao [get]
func (chuRuKuMingXiBiaoApi *ChuRuKuMingXiBiaoApi) FindChuRuKuMingXiBiao(c *gin.Context) {
	var chuRuKuMingXiBiao gong_cheng_bu.ChuRuKuMingXiBiao
	err := c.ShouldBindQuery(&chuRuKuMingXiBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rechuRuKuMingXiBiao, err := chuRuKuMingXiBiaoService.GetChuRuKuMingXiBiao(chuRuKuMingXiBiao.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rechuRuKuMingXiBiao": rechuRuKuMingXiBiao}, c)
	}
}

// GetChuRuKuMingXiBiaoList 分页获取ChuRuKuMingXiBiao列表
// @Tags ChuRuKuMingXiBiao
// @Summary 分页获取ChuRuKuMingXiBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_buReq.ChuRuKuMingXiBiaoSearch true "分页获取ChuRuKuMingXiBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /chuRuKuMingXiBiao/getChuRuKuMingXiBiaoList [get]
func (chuRuKuMingXiBiaoApi *ChuRuKuMingXiBiaoApi) GetChuRuKuMingXiBiaoList(c *gin.Context) {
	var pageInfo gong_cheng_buReq.ChuRuKuMingXiBiaoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := chuRuKuMingXiBiaoService.GetChuRuKuMingXiBiaoInfoList(pageInfo); err != nil {
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
