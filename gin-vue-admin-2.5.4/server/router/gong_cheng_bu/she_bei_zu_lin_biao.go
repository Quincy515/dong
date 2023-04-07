package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SheBeiZuLinBiaoRouter struct {
}

// InitSheBeiZuLinBiaoRouter 初始化 SheBeiZuLinBiao 路由信息
func (s *SheBeiZuLinBiaoRouter) InitSheBeiZuLinBiaoRouter(Router *gin.RouterGroup) {
	sheBeiZuLinBiaoRouter := Router.Group("sheBeiZuLinBiao").Use(middleware.OperationRecord())
	sheBeiZuLinBiaoRouterWithoutRecord := Router.Group("sheBeiZuLinBiao")
	var sheBeiZuLinBiaoApi = v1.ApiGroupApp.Gong_cheng_buApiGroup.SheBeiZuLinBiaoApi
	{
		sheBeiZuLinBiaoRouter.POST("createSheBeiZuLinBiao", sheBeiZuLinBiaoApi.CreateSheBeiZuLinBiao)   // 新建SheBeiZuLinBiao
		sheBeiZuLinBiaoRouter.DELETE("deleteSheBeiZuLinBiao", sheBeiZuLinBiaoApi.DeleteSheBeiZuLinBiao) // 删除SheBeiZuLinBiao
		sheBeiZuLinBiaoRouter.DELETE("deleteSheBeiZuLinBiaoByIds", sheBeiZuLinBiaoApi.DeleteSheBeiZuLinBiaoByIds) // 批量删除SheBeiZuLinBiao
		sheBeiZuLinBiaoRouter.PUT("updateSheBeiZuLinBiao", sheBeiZuLinBiaoApi.UpdateSheBeiZuLinBiao)    // 更新SheBeiZuLinBiao
	}
	{
		sheBeiZuLinBiaoRouterWithoutRecord.GET("findSheBeiZuLinBiao", sheBeiZuLinBiaoApi.FindSheBeiZuLinBiao)        // 根据ID获取SheBeiZuLinBiao
		sheBeiZuLinBiaoRouterWithoutRecord.GET("getSheBeiZuLinBiaoList", sheBeiZuLinBiaoApi.GetSheBeiZuLinBiaoList)  // 获取SheBeiZuLinBiao列表
	}
}
