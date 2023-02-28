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

type KaoQinTongJiApi struct {
}

var kaoQinTongJiService = service.ServiceGroupApp.HrServiceGroup.KaoQinTongJiService

// CreateKaoQinTongJi 创建KaoQinTongJi
// @Tags KaoQinTongJi
// @Summary 创建KaoQinTongJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.KaoQinTongJi true "创建KaoQinTongJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /kaoQinTongJi/createKaoQinTongJi [post]
func (kaoQinTongJiApi *KaoQinTongJiApi) CreateKaoQinTongJi(c *gin.Context) {
	var kaoQinTongJi hr.KaoQinTongJi
	err := c.ShouldBindJSON(&kaoQinTongJi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := kaoQinTongJiService.CreateKaoQinTongJi(kaoQinTongJi); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteKaoQinTongJi 删除KaoQinTongJi
// @Tags KaoQinTongJi
// @Summary 删除KaoQinTongJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.KaoQinTongJi true "删除KaoQinTongJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /kaoQinTongJi/deleteKaoQinTongJi [delete]
func (kaoQinTongJiApi *KaoQinTongJiApi) DeleteKaoQinTongJi(c *gin.Context) {
	var kaoQinTongJi hr.KaoQinTongJi
	err := c.ShouldBindJSON(&kaoQinTongJi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := kaoQinTongJiService.DeleteKaoQinTongJi(kaoQinTongJi); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteKaoQinTongJiByIds 批量删除KaoQinTongJi
// @Tags KaoQinTongJi
// @Summary 批量删除KaoQinTongJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除KaoQinTongJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /kaoQinTongJi/deleteKaoQinTongJiByIds [delete]
func (kaoQinTongJiApi *KaoQinTongJiApi) DeleteKaoQinTongJiByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := kaoQinTongJiService.DeleteKaoQinTongJiByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateKaoQinTongJi 更新KaoQinTongJi
// @Tags KaoQinTongJi
// @Summary 更新KaoQinTongJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.KaoQinTongJi true "更新KaoQinTongJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /kaoQinTongJi/updateKaoQinTongJi [put]
func (kaoQinTongJiApi *KaoQinTongJiApi) UpdateKaoQinTongJi(c *gin.Context) {
	var kaoQinTongJi hr.KaoQinTongJi
	err := c.ShouldBindJSON(&kaoQinTongJi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := kaoQinTongJiService.UpdateKaoQinTongJi(kaoQinTongJi); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindKaoQinTongJi 用id查询KaoQinTongJi
// @Tags KaoQinTongJi
// @Summary 用id查询KaoQinTongJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hr.KaoQinTongJi true "用id查询KaoQinTongJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /kaoQinTongJi/findKaoQinTongJi [get]
func (kaoQinTongJiApi *KaoQinTongJiApi) FindKaoQinTongJi(c *gin.Context) {
	var kaoQinTongJi hr.KaoQinTongJi
	err := c.ShouldBindQuery(&kaoQinTongJi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rekaoQinTongJi, err := kaoQinTongJiService.GetKaoQinTongJi(kaoQinTongJi.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rekaoQinTongJi": rekaoQinTongJi}, c)
	}
}

// GetKaoQinTongJiList 分页获取KaoQinTongJi列表
// @Tags KaoQinTongJi
// @Summary 分页获取KaoQinTongJi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hrReq.KaoQinTongJiSearch true "分页获取KaoQinTongJi列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /kaoQinTongJi/getKaoQinTongJiList [get]
func (kaoQinTongJiApi *KaoQinTongJiApi) GetKaoQinTongJiList(c *gin.Context) {
	var pageInfo hrReq.KaoQinTongJiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := kaoQinTongJiService.GetKaoQinTongJiInfoList(pageInfo); err != nil {
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
