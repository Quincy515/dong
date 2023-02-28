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

type MeiRiKaoQinApi struct {
}

var meiRiKaoQinService = service.ServiceGroupApp.HrServiceGroup.MeiRiKaoQinService

// CreateMeiRiKaoQin 创建MeiRiKaoQin
// @Tags MeiRiKaoQin
// @Summary 创建MeiRiKaoQin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.MeiRiKaoQin true "创建MeiRiKaoQin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /meiRiKaoQin/createMeiRiKaoQin [post]
func (meiRiKaoQinApi *MeiRiKaoQinApi) CreateMeiRiKaoQin(c *gin.Context) {
	var meiRiKaoQin hr.MeiRiKaoQin
	err := c.ShouldBindJSON(&meiRiKaoQin)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := meiRiKaoQinService.CreateMeiRiKaoQin(meiRiKaoQin); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMeiRiKaoQin 删除MeiRiKaoQin
// @Tags MeiRiKaoQin
// @Summary 删除MeiRiKaoQin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.MeiRiKaoQin true "删除MeiRiKaoQin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /meiRiKaoQin/deleteMeiRiKaoQin [delete]
func (meiRiKaoQinApi *MeiRiKaoQinApi) DeleteMeiRiKaoQin(c *gin.Context) {
	var meiRiKaoQin hr.MeiRiKaoQin
	err := c.ShouldBindJSON(&meiRiKaoQin)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := meiRiKaoQinService.DeleteMeiRiKaoQin(meiRiKaoQin); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMeiRiKaoQinByIds 批量删除MeiRiKaoQin
// @Tags MeiRiKaoQin
// @Summary 批量删除MeiRiKaoQin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MeiRiKaoQin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /meiRiKaoQin/deleteMeiRiKaoQinByIds [delete]
func (meiRiKaoQinApi *MeiRiKaoQinApi) DeleteMeiRiKaoQinByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := meiRiKaoQinService.DeleteMeiRiKaoQinByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMeiRiKaoQin 更新MeiRiKaoQin
// @Tags MeiRiKaoQin
// @Summary 更新MeiRiKaoQin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hr.MeiRiKaoQin true "更新MeiRiKaoQin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /meiRiKaoQin/updateMeiRiKaoQin [put]
func (meiRiKaoQinApi *MeiRiKaoQinApi) UpdateMeiRiKaoQin(c *gin.Context) {
	var meiRiKaoQin hr.MeiRiKaoQin
	err := c.ShouldBindJSON(&meiRiKaoQin)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := meiRiKaoQinService.UpdateMeiRiKaoQin(meiRiKaoQin); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMeiRiKaoQin 用id查询MeiRiKaoQin
// @Tags MeiRiKaoQin
// @Summary 用id查询MeiRiKaoQin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hr.MeiRiKaoQin true "用id查询MeiRiKaoQin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /meiRiKaoQin/findMeiRiKaoQin [get]
func (meiRiKaoQinApi *MeiRiKaoQinApi) FindMeiRiKaoQin(c *gin.Context) {
	var meiRiKaoQin hr.MeiRiKaoQin
	err := c.ShouldBindQuery(&meiRiKaoQin)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if remeiRiKaoQin, err := meiRiKaoQinService.GetMeiRiKaoQin(meiRiKaoQin.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remeiRiKaoQin": remeiRiKaoQin}, c)
	}
}

// GetMeiRiKaoQinList 分页获取MeiRiKaoQin列表
// @Tags MeiRiKaoQin
// @Summary 分页获取MeiRiKaoQin列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hrReq.MeiRiKaoQinSearch true "分页获取MeiRiKaoQin列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /meiRiKaoQin/getMeiRiKaoQinList [get]
func (meiRiKaoQinApi *MeiRiKaoQinApi) GetMeiRiKaoQinList(c *gin.Context) {
	var pageInfo hrReq.MeiRiKaoQinSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := meiRiKaoQinService.GetMeiRiKaoQinInfoList(pageInfo); err != nil {
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
