package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    gong_cheng_buReq "github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type XiangMuBiaoApi struct {
}

var xiangMuBiaoService = service.ServiceGroupApp.Gong_cheng_buServiceGroup.XiangMuBiaoService


// CreateXiangMuBiao 创建XiangMuBiao
// @Tags XiangMuBiao
// @Summary 创建XiangMuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.XiangMuBiao true "创建XiangMuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xiangMuBiao/createXiangMuBiao [post]
func (xiangMuBiaoApi *XiangMuBiaoApi) CreateXiangMuBiao(c *gin.Context) {
	var xiangMuBiao gong_cheng_bu.XiangMuBiao
	err := c.ShouldBindJSON(&xiangMuBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    xiangMuBiao.CreatedBy = utils.GetUserID(c)
	if err := xiangMuBiaoService.CreateXiangMuBiao(xiangMuBiao); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteXiangMuBiao 删除XiangMuBiao
// @Tags XiangMuBiao
// @Summary 删除XiangMuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.XiangMuBiao true "删除XiangMuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xiangMuBiao/deleteXiangMuBiao [delete]
func (xiangMuBiaoApi *XiangMuBiaoApi) DeleteXiangMuBiao(c *gin.Context) {
	var xiangMuBiao gong_cheng_bu.XiangMuBiao
	err := c.ShouldBindJSON(&xiangMuBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    xiangMuBiao.DeletedBy = utils.GetUserID(c)
	if err := xiangMuBiaoService.DeleteXiangMuBiao(xiangMuBiao); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteXiangMuBiaoByIds 批量删除XiangMuBiao
// @Tags XiangMuBiao
// @Summary 批量删除XiangMuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除XiangMuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /xiangMuBiao/deleteXiangMuBiaoByIds [delete]
func (xiangMuBiaoApi *XiangMuBiaoApi) DeleteXiangMuBiaoByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := xiangMuBiaoService.DeleteXiangMuBiaoByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateXiangMuBiao 更新XiangMuBiao
// @Tags XiangMuBiao
// @Summary 更新XiangMuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gong_cheng_bu.XiangMuBiao true "更新XiangMuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /xiangMuBiao/updateXiangMuBiao [put]
func (xiangMuBiaoApi *XiangMuBiaoApi) UpdateXiangMuBiao(c *gin.Context) {
	var xiangMuBiao gong_cheng_bu.XiangMuBiao
	err := c.ShouldBindJSON(&xiangMuBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    xiangMuBiao.UpdatedBy = utils.GetUserID(c)
	if err := xiangMuBiaoService.UpdateXiangMuBiao(xiangMuBiao); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindXiangMuBiao 用id查询XiangMuBiao
// @Tags XiangMuBiao
// @Summary 用id查询XiangMuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_bu.XiangMuBiao true "用id查询XiangMuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /xiangMuBiao/findXiangMuBiao [get]
func (xiangMuBiaoApi *XiangMuBiaoApi) FindXiangMuBiao(c *gin.Context) {
	var xiangMuBiao gong_cheng_bu.XiangMuBiao
	err := c.ShouldBindQuery(&xiangMuBiao)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rexiangMuBiao, err := xiangMuBiaoService.GetXiangMuBiao(xiangMuBiao.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rexiangMuBiao": rexiangMuBiao}, c)
	}
}

// GetXiangMuBiaoList 分页获取XiangMuBiao列表
// @Tags XiangMuBiao
// @Summary 分页获取XiangMuBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gong_cheng_buReq.XiangMuBiaoSearch true "分页获取XiangMuBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xiangMuBiao/getXiangMuBiaoList [get]
func (xiangMuBiaoApi *XiangMuBiaoApi) GetXiangMuBiaoList(c *gin.Context) {
	var pageInfo gong_cheng_buReq.XiangMuBiaoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := xiangMuBiaoService.GetXiangMuBiaoInfoList(pageInfo); err != nil {
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
