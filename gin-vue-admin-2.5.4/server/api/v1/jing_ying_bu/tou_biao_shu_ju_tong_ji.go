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

type TouBiaoShuJuTongJiApi struct {
}

var touBiaoShuJuTongJiService = service.ServiceGroupApp.Jing_ying_buServiceGroup.TouBiaoShuJuTongJiService

// CreateTouBiaoShuJuTongJi 创建TouBiaoShuJuTongJi
// @Tags TouBiaoShuJuTongJi
// @Summary 创建TouBiaoShuJuTongJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body jing_ying_bu.TouBiaoShuJuTongJi true "创建TouBiaoShuJuTongJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /touBiaoShuJuTongJi/createTouBiaoShuJuTongJi [post]
func (touBiaoShuJuTongJiApi *TouBiaoShuJuTongJiApi) CreateTouBiaoShuJuTongJi(c *gin.Context) {
	var touBiaoShuJuTongJi jing_ying_bu.TouBiaoShuJuTongJi
	err := c.ShouldBindJSON(&touBiaoShuJuTongJi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := touBiaoShuJuTongJiService.CreateTouBiaoShuJuTongJi(touBiaoShuJuTongJi); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTouBiaoShuJuTongJi 删除TouBiaoShuJuTongJi
// @Tags TouBiaoShuJuTongJi
// @Summary 删除TouBiaoShuJuTongJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body jing_ying_bu.TouBiaoShuJuTongJi true "删除TouBiaoShuJuTongJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /touBiaoShuJuTongJi/deleteTouBiaoShuJuTongJi [delete]
func (touBiaoShuJuTongJiApi *TouBiaoShuJuTongJiApi) DeleteTouBiaoShuJuTongJi(c *gin.Context) {
	var touBiaoShuJuTongJi jing_ying_bu.TouBiaoShuJuTongJi
	err := c.ShouldBindJSON(&touBiaoShuJuTongJi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := touBiaoShuJuTongJiService.DeleteTouBiaoShuJuTongJi(touBiaoShuJuTongJi); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTouBiaoShuJuTongJiByIds 批量删除TouBiaoShuJuTongJi
// @Tags TouBiaoShuJuTongJi
// @Summary 批量删除TouBiaoShuJuTongJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TouBiaoShuJuTongJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /touBiaoShuJuTongJi/deleteTouBiaoShuJuTongJiByIds [delete]
func (touBiaoShuJuTongJiApi *TouBiaoShuJuTongJiApi) DeleteTouBiaoShuJuTongJiByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := touBiaoShuJuTongJiService.DeleteTouBiaoShuJuTongJiByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTouBiaoShuJuTongJi 更新TouBiaoShuJuTongJi
// @Tags TouBiaoShuJuTongJi
// @Summary 更新TouBiaoShuJuTongJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body jing_ying_bu.TouBiaoShuJuTongJi true "更新TouBiaoShuJuTongJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /touBiaoShuJuTongJi/updateTouBiaoShuJuTongJi [put]
func (touBiaoShuJuTongJiApi *TouBiaoShuJuTongJiApi) UpdateTouBiaoShuJuTongJi(c *gin.Context) {
	var touBiaoShuJuTongJi jing_ying_bu.TouBiaoShuJuTongJi
	err := c.ShouldBindJSON(&touBiaoShuJuTongJi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := touBiaoShuJuTongJiService.UpdateTouBiaoShuJuTongJi(touBiaoShuJuTongJi); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTouBiaoShuJuTongJi 用id查询TouBiaoShuJuTongJi
// @Tags TouBiaoShuJuTongJi
// @Summary 用id查询TouBiaoShuJuTongJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query jing_ying_bu.TouBiaoShuJuTongJi true "用id查询TouBiaoShuJuTongJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /touBiaoShuJuTongJi/findTouBiaoShuJuTongJi [get]
func (touBiaoShuJuTongJiApi *TouBiaoShuJuTongJiApi) FindTouBiaoShuJuTongJi(c *gin.Context) {
	var touBiaoShuJuTongJi jing_ying_bu.TouBiaoShuJuTongJi
	err := c.ShouldBindQuery(&touBiaoShuJuTongJi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retouBiaoShuJuTongJi, err := touBiaoShuJuTongJiService.GetTouBiaoShuJuTongJi(touBiaoShuJuTongJi.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retouBiaoShuJuTongJi": retouBiaoShuJuTongJi}, c)
	}
}

// GetTouBiaoShuJuTongJiList 分页获取TouBiaoShuJuTongJi列表
// @Tags TouBiaoShuJuTongJi
// @Summary 分页获取TouBiaoShuJuTongJi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query jing_ying_buReq.TouBiaoShuJuTongJiSearch true "分页获取TouBiaoShuJuTongJi列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /touBiaoShuJuTongJi/getTouBiaoShuJuTongJiList [get]
func (touBiaoShuJuTongJiApi *TouBiaoShuJuTongJiApi) GetTouBiaoShuJuTongJiList(c *gin.Context) {
	var pageInfo jing_ying_buReq.TouBiaoShuJuTongJiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := touBiaoShuJuTongJiService.GetTouBiaoShuJuTongJiInfoList(pageInfo); err != nil {
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
