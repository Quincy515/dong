package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HeTongWenJianBiaoRouter struct {
}

// InitHeTongWenJianBiaoRouter 初始化 HeTongWenJianBiao 路由信息
func (s *HeTongWenJianBiaoRouter) InitHeTongWenJianBiaoRouter(Router *gin.RouterGroup) {
	heTongWenJianBiaoRouter := Router.Group("heTongWenJianBiao").Use(middleware.OperationRecord())
	heTongWenJianBiaoRouterWithoutRecord := Router.Group("heTongWenJianBiao")
	var heTongWenJianBiaoApi = v1.ApiGroupApp.Gong_cheng_buApiGroup.HeTongWenJianBiaoApi
	{
		heTongWenJianBiaoRouter.POST("createHeTongWenJianBiao", heTongWenJianBiaoApi.CreateHeTongWenJianBiao)   // 新建HeTongWenJianBiao
		heTongWenJianBiaoRouter.DELETE("deleteHeTongWenJianBiao", heTongWenJianBiaoApi.DeleteHeTongWenJianBiao) // 删除HeTongWenJianBiao
		heTongWenJianBiaoRouter.DELETE("deleteHeTongWenJianBiaoByIds", heTongWenJianBiaoApi.DeleteHeTongWenJianBiaoByIds) // 批量删除HeTongWenJianBiao
		heTongWenJianBiaoRouter.PUT("updateHeTongWenJianBiao", heTongWenJianBiaoApi.UpdateHeTongWenJianBiao)    // 更新HeTongWenJianBiao
	}
	{
		heTongWenJianBiaoRouterWithoutRecord.GET("findHeTongWenJianBiao", heTongWenJianBiaoApi.FindHeTongWenJianBiao)        // 根据ID获取HeTongWenJianBiao
		heTongWenJianBiaoRouterWithoutRecord.GET("getHeTongWenJianBiaoList", heTongWenJianBiaoApi.GetHeTongWenJianBiaoList)  // 获取HeTongWenJianBiao列表
	}
}
