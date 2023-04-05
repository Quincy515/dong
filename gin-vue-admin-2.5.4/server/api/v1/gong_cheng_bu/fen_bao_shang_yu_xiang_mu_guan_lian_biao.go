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

type FenBaoShangYuXiangMuGuanLianBiaoApi struct {
}

var fenBaoShangYuXiangMuGuanLianBiaoService = service.ServiceGroupApp.Gong_cheng_buServiceGroup.FenBaoShangYuXiangMuGuanLianBiaoService

// CreateFenBaoShangYuXiangMuGuanLianBiao 创建FenBaoShangYuXiangMuGuanLianBiao
// @Tags FenBaoShangYuXiangMuGuanLianBiao
// @Summary 创建FenBaoShangYuXiangMuGuanLianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.FenBaoShangYuXiangMuGuanLianBiao true "创建FenBaoShangYuXiangMuGuanLianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fenBaoShangYuXiangMuGuanLianBiao/createFenBaoShangYuXiangMuGuanLianBiao [post]
func (fenBaoShangYuXiangMuGuanLianBiaoApi *FenBaoShangYuXiangMuGuanLianBiaoApi) CreateFenBaoShangYuXiangMuGuanLianBiao(c *gin.Context) {
	var fenBaoShangYuXiangMuGuanLianBiao gong_cheng_bu.FenBaoShangYuXiangMuGuanLianBiao
	err := c.ShouldBindJSON(&fenBaoShangYuXiangMuGuanLianBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fenBaoShangYuXiangMuGuanLianBiaoService.CreateFenBaoShangYuXiangMuGuanLianBiao(fenBaoShangYuXiangMuGuanLianBiao); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFenBaoShangYuXiangMuGuanLianBiao 删除FenBaoShangYuXiangMuGuanLianBiao
// @Tags FenBaoShangYuXiangMuGuanLianBiao
// @Summary 删除FenBaoShangYuXiangMuGuanLianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.FenBaoShangYuXiangMuGuanLianBiao true "删除FenBaoShangYuXiangMuGuanLianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fenBaoShangYuXiangMuGuanLianBiao/deleteFenBaoShangYuXiangMuGuanLianBiao [delete]
func (fenBaoShangYuXiangMuGuanLianBiaoApi *FenBaoShangYuXiangMuGuanLianBiaoApi) DeleteFenBaoShangYuXiangMuGuanLianBiao(c *gin.Context) {
	var fenBaoShangYuXiangMuGuanLianBiao gong_cheng_bu.FenBaoShangYuXiangMuGuanLianBiao
	err := c.ShouldBindJSON(&fenBaoShangYuXiangMuGuanLianBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fenBaoShangYuXiangMuGuanLianBiaoService.DeleteFenBaoShangYuXiangMuGuanLianBiao(fenBaoShangYuXiangMuGuanLianBiao); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFenBaoShangYuXiangMuGuanLianBiaoByIds 批量删除FenBaoShangYuXiangMuGuanLianBiao
// @Tags FenBaoShangYuXiangMuGuanLianBiao
// @Summary 批量删除FenBaoShangYuXiangMuGuanLianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FenBaoShangYuXiangMuGuanLianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fenBaoShangYuXiangMuGuanLianBiao/deleteFenBaoShangYuXiangMuGuanLianBiaoByIds [delete]
func (fenBaoShangYuXiangMuGuanLianBiaoApi *FenBaoShangYuXiangMuGuanLianBiaoApi) DeleteFenBaoShangYuXiangMuGuanLianBiaoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fenBaoShangYuXiangMuGuanLianBiaoService.DeleteFenBaoShangYuXiangMuGuanLianBiaoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFenBaoShangYuXiangMuGuanLianBiao 更新FenBaoShangYuXiangMuGuanLianBiao
// @Tags FenBaoShangYuXiangMuGuanLianBiao
// @Summary 更新FenBaoShangYuXiangMuGuanLianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.FenBaoShangYuXiangMuGuanLianBiao true "更新FenBaoShangYuXiangMuGuanLianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fenBaoShangYuXiangMuGuanLianBiao/updateFenBaoShangYuXiangMuGuanLianBiao [put]
func (fenBaoShangYuXiangMuGuanLianBiaoApi *FenBaoShangYuXiangMuGuanLianBiaoApi) UpdateFenBaoShangYuXiangMuGuanLianBiao(c *gin.Context) {
	var fenBaoShangYuXiangMuGuanLianBiao gong_cheng_bu.FenBaoShangYuXiangMuGuanLianBiao
	err := c.ShouldBindJSON(&fenBaoShangYuXiangMuGuanLianBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fenBaoShangYuXiangMuGuanLianBiaoService.UpdateFenBaoShangYuXiangMuGuanLianBiao(fenBaoShangYuXiangMuGuanLianBiao); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFenBaoShangYuXiangMuGuanLianBiao 用id查询FenBaoShangYuXiangMuGuanLianBiao
// @Tags FenBaoShangYuXiangMuGuanLianBiao
// @Summary 用id查询FenBaoShangYuXiangMuGuanLianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_bu.FenBaoShangYuXiangMuGuanLianBiao true "用id查询FenBaoShangYuXiangMuGuanLianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fenBaoShangYuXiangMuGuanLianBiao/findFenBaoShangYuXiangMuGuanLianBiao [get]
func (fenBaoShangYuXiangMuGuanLianBiaoApi *FenBaoShangYuXiangMuGuanLianBiaoApi) FindFenBaoShangYuXiangMuGuanLianBiao(c *gin.Context) {
	var fenBaoShangYuXiangMuGuanLianBiao gong_cheng_bu.FenBaoShangYuXiangMuGuanLianBiao
	err := c.ShouldBindQuery(&fenBaoShangYuXiangMuGuanLianBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refenBaoShangYuXiangMuGuanLianBiao, err := fenBaoShangYuXiangMuGuanLianBiaoService.GetFenBaoShangYuXiangMuGuanLianBiao(fenBaoShangYuXiangMuGuanLianBiao.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refenBaoShangYuXiangMuGuanLianBiao": refenBaoShangYuXiangMuGuanLianBiao}, c)
	}
}

// GetFenBaoShangYuXiangMuGuanLianBiaoList 分页获取FenBaoShangYuXiangMuGuanLianBiao列表
// @Tags FenBaoShangYuXiangMuGuanLianBiao
// @Summary 分页获取FenBaoShangYuXiangMuGuanLianBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_buReq.FenBaoShangYuXiangMuGuanLianBiaoSearch true "分页获取FenBaoShangYuXiangMuGuanLianBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fenBaoShangYuXiangMuGuanLianBiao/getFenBaoShangYuXiangMuGuanLianBiaoList [get]
func (fenBaoShangYuXiangMuGuanLianBiaoApi *FenBaoShangYuXiangMuGuanLianBiaoApi) GetFenBaoShangYuXiangMuGuanLianBiaoList(c *gin.Context) {
	var pageInfo gong_cheng_buReq.FenBaoShangYuXiangMuGuanLianBiaoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := fenBaoShangYuXiangMuGuanLianBiaoService.GetFenBaoShangYuXiangMuGuanLianBiaoInfoList(pageInfo); err != nil {
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
