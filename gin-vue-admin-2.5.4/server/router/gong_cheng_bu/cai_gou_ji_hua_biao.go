package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CaiGouJiHuaBiaoRouter struct {
}

// InitCaiGouJiHuaBiaoRouter 初始化 CaiGouJiHuaBiao 路由信息
func (s *CaiGouJiHuaBiaoRouter) InitCaiGouJiHuaBiaoRouter(Router *gin.RouterGroup) {
	caiGouJiHuaBiaoRouter := Router.Group("caiGouJiHuaBiao").Use(middleware.OperationRecord())
	caiGouJiHuaBiaoRouterWithoutRecord := Router.Group("caiGouJiHuaBiao")
	var caiGouJiHuaBiaoApi = v1.ApiGroupApp.Gong_cheng_buApiGroup.CaiGouJiHuaBiaoApi
	{
		caiGouJiHuaBiaoRouter.POST("createCaiGouJiHuaBiao", caiGouJiHuaBiaoApi.CreateCaiGouJiHuaBiao)   // 新建CaiGouJiHuaBiao
		caiGouJiHuaBiaoRouter.DELETE("deleteCaiGouJiHuaBiao", caiGouJiHuaBiaoApi.DeleteCaiGouJiHuaBiao) // 删除CaiGouJiHuaBiao
		caiGouJiHuaBiaoRouter.DELETE("deleteCaiGouJiHuaBiaoByIds", caiGouJiHuaBiaoApi.DeleteCaiGouJiHuaBiaoByIds) // 批量删除CaiGouJiHuaBiao
		caiGouJiHuaBiaoRouter.PUT("updateCaiGouJiHuaBiao", caiGouJiHuaBiaoApi.UpdateCaiGouJiHuaBiao)    // 更新CaiGouJiHuaBiao
	}
	{
		caiGouJiHuaBiaoRouterWithoutRecord.GET("findCaiGouJiHuaBiao", caiGouJiHuaBiaoApi.FindCaiGouJiHuaBiao)        // 根据ID获取CaiGouJiHuaBiao
		caiGouJiHuaBiaoRouterWithoutRecord.GET("getCaiGouJiHuaBiaoList", caiGouJiHuaBiaoApi.GetCaiGouJiHuaBiaoList)  // 获取CaiGouJiHuaBiao列表
	}
}
