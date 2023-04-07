package cai_wu_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WaiBuFeiYongQingKuanBiaoRouter struct {
}

// InitWaiBuFeiYongQingKuanBiaoRouter 初始化 WaiBuFeiYongQingKuanBiao 路由信息
func (s *WaiBuFeiYongQingKuanBiaoRouter) InitWaiBuFeiYongQingKuanBiaoRouter(Router *gin.RouterGroup) {
	waiBuFeiYongQingKuanBiaoRouter := Router.Group("waiBuFeiYongQingKuanBiao").Use(middleware.OperationRecord())
	waiBuFeiYongQingKuanBiaoRouterWithoutRecord := Router.Group("waiBuFeiYongQingKuanBiao")
	var waiBuFeiYongQingKuanBiaoApi = v1.ApiGroupApp.Cai_wu_buApiGroup.WaiBuFeiYongQingKuanBiaoApi
	{
		waiBuFeiYongQingKuanBiaoRouter.POST("createWaiBuFeiYongQingKuanBiao", waiBuFeiYongQingKuanBiaoApi.CreateWaiBuFeiYongQingKuanBiao)   // 新建WaiBuFeiYongQingKuanBiao
		waiBuFeiYongQingKuanBiaoRouter.DELETE("deleteWaiBuFeiYongQingKuanBiao", waiBuFeiYongQingKuanBiaoApi.DeleteWaiBuFeiYongQingKuanBiao) // 删除WaiBuFeiYongQingKuanBiao
		waiBuFeiYongQingKuanBiaoRouter.DELETE("deleteWaiBuFeiYongQingKuanBiaoByIds", waiBuFeiYongQingKuanBiaoApi.DeleteWaiBuFeiYongQingKuanBiaoByIds) // 批量删除WaiBuFeiYongQingKuanBiao
		waiBuFeiYongQingKuanBiaoRouter.PUT("updateWaiBuFeiYongQingKuanBiao", waiBuFeiYongQingKuanBiaoApi.UpdateWaiBuFeiYongQingKuanBiao)    // 更新WaiBuFeiYongQingKuanBiao
	}
	{
		waiBuFeiYongQingKuanBiaoRouterWithoutRecord.GET("findWaiBuFeiYongQingKuanBiao", waiBuFeiYongQingKuanBiaoApi.FindWaiBuFeiYongQingKuanBiao)        // 根据ID获取WaiBuFeiYongQingKuanBiao
		waiBuFeiYongQingKuanBiaoRouterWithoutRecord.GET("getWaiBuFeiYongQingKuanBiaoList", waiBuFeiYongQingKuanBiaoApi.GetWaiBuFeiYongQingKuanBiaoList)  // 获取WaiBuFeiYongQingKuanBiao列表
	}
}
