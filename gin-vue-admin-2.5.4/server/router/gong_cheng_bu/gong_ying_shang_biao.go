package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GongYingShangBiaoRouter struct {
}

// InitGongYingShangBiaoRouter 初始化 GongYingShangBiao 路由信息
func (s *GongYingShangBiaoRouter) InitGongYingShangBiaoRouter(Router *gin.RouterGroup) {
	gongYingShangBiaoRouter := Router.Group("gongYingShangBiao").Use(middleware.OperationRecord())
	gongYingShangBiaoRouterWithoutRecord := Router.Group("gongYingShangBiao")
	var gongYingShangBiaoApi = v1.ApiGroupApp.Gong_cheng_buApiGroup.GongYingShangBiaoApi
	{
		gongYingShangBiaoRouter.POST("createGongYingShangBiao", gongYingShangBiaoApi.CreateGongYingShangBiao)   // 新建GongYingShangBiao
		gongYingShangBiaoRouter.DELETE("deleteGongYingShangBiao", gongYingShangBiaoApi.DeleteGongYingShangBiao) // 删除GongYingShangBiao
		gongYingShangBiaoRouter.DELETE("deleteGongYingShangBiaoByIds", gongYingShangBiaoApi.DeleteGongYingShangBiaoByIds) // 批量删除GongYingShangBiao
		gongYingShangBiaoRouter.PUT("updateGongYingShangBiao", gongYingShangBiaoApi.UpdateGongYingShangBiao)    // 更新GongYingShangBiao
	}
	{
		gongYingShangBiaoRouterWithoutRecord.GET("findGongYingShangBiao", gongYingShangBiaoApi.FindGongYingShangBiao)        // 根据ID获取GongYingShangBiao
		gongYingShangBiaoRouterWithoutRecord.GET("getGongYingShangBiaoList", gongYingShangBiaoApi.GetGongYingShangBiaoList)  // 获取GongYingShangBiao列表
	}
}
