package cai_wu_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cai_wu_bu"
	cai_wu_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/cai_wu_bu/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type YingShouKuanBiaoApi struct {
}

var yingShouKuanBiaoService = service.ServiceGroupApp.Cai_wu_buServiceGroup.YingShouKuanBiaoService

// CreateYingShouKuanBiao 创建YingShouKuanBiao
// @Tags YingShouKuanBiao
// @Summary 创建YingShouKuanBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cai_wu_bu.YingShouKuanBiao true "创建YingShouKuanBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yingShouKuanBiao/createYingShouKuanBiao [post]
func (yingShouKuanBiaoApi *YingShouKuanBiaoApi) CreateYingShouKuanBiao(c *gin.Context) {
	var yingShouKuanBiao cai_wu_bu.YingShouKuanBiao
	err := c.ShouldBindJSON(&yingShouKuanBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := yingShouKuanBiaoService.CreateYingShouKuanBiao(yingShouKuanBiao); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteYingShouKuanBiao 删除YingShouKuanBiao
// @Tags YingShouKuanBiao
// @Summary 删除YingShouKuanBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cai_wu_bu.YingShouKuanBiao true "删除YingShouKuanBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /yingShouKuanBiao/deleteYingShouKuanBiao [delete]
func (yingShouKuanBiaoApi *YingShouKuanBiaoApi) DeleteYingShouKuanBiao(c *gin.Context) {
	var yingShouKuanBiao cai_wu_bu.YingShouKuanBiao
	err := c.ShouldBindJSON(&yingShouKuanBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := yingShouKuanBiaoService.DeleteYingShouKuanBiao(yingShouKuanBiao); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteYingShouKuanBiaoByIds 批量删除YingShouKuanBiao
// @Tags YingShouKuanBiao
// @Summary 批量删除YingShouKuanBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除YingShouKuanBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /yingShouKuanBiao/deleteYingShouKuanBiaoByIds [delete]
func (yingShouKuanBiaoApi *YingShouKuanBiaoApi) DeleteYingShouKuanBiaoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := yingShouKuanBiaoService.DeleteYingShouKuanBiaoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateYingShouKuanBiao 更新YingShouKuanBiao
// @Tags YingShouKuanBiao
// @Summary 更新YingShouKuanBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cai_wu_bu.YingShouKuanBiao true "更新YingShouKuanBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /yingShouKuanBiao/updateYingShouKuanBiao [put]
func (yingShouKuanBiaoApi *YingShouKuanBiaoApi) UpdateYingShouKuanBiao(c *gin.Context) {
	var yingShouKuanBiao cai_wu_bu.YingShouKuanBiao
	err := c.ShouldBindJSON(&yingShouKuanBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := yingShouKuanBiaoService.UpdateYingShouKuanBiao(yingShouKuanBiao); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindYingShouKuanBiao 用id查询YingShouKuanBiao
// @Tags YingShouKuanBiao
// @Summary 用id查询YingShouKuanBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cai_wu_bu.YingShouKuanBiao true "用id查询YingShouKuanBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /yingShouKuanBiao/findYingShouKuanBiao [get]
func (yingShouKuanBiaoApi *YingShouKuanBiaoApi) FindYingShouKuanBiao(c *gin.Context) {
	var yingShouKuanBiao cai_wu_bu.YingShouKuanBiao
	err := c.ShouldBindQuery(&yingShouKuanBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reyingShouKuanBiao, err := yingShouKuanBiaoService.GetYingShouKuanBiao(yingShouKuanBiao.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reyingShouKuanBiao": reyingShouKuanBiao}, c)
	}
}

// GetYingShouKuanBiaoList 分页获取YingShouKuanBiao列表
// @Tags YingShouKuanBiao
// @Summary 分页获取YingShouKuanBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cai_wu_buReq.YingShouKuanBiaoSearch true "分页获取YingShouKuanBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yingShouKuanBiao/getYingShouKuanBiaoList [get]
func (yingShouKuanBiaoApi *YingShouKuanBiaoApi) GetYingShouKuanBiaoList(c *gin.Context) {
	var pageInfo cai_wu_buReq.YingShouKuanBiaoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := yingShouKuanBiaoService.GetYingShouKuanBiaoInfoList(pageInfo); err != nil {
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
