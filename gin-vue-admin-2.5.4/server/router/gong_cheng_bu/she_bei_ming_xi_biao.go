package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SheBeiMingXiBiaoRouter struct {
}

// InitSheBeiMingXiBiaoRouter 初始化 SheBeiMingXiBiao 路由信息
func (s *SheBeiMingXiBiaoRouter) InitSheBeiMingXiBiaoRouter(Router *gin.RouterGroup) {
	sheBeiMingXiBiaoRouter := Router.Group("sheBeiMingXiBiao").Use(middleware.OperationRecord())
	sheBeiMingXiBiaoRouterWithoutRecord := Router.Group("sheBeiMingXiBiao")
	var sheBeiMingXiBiaoApi = v1.ApiGroupApp.Gong_cheng_buApiGroup.SheBeiMingXiBiaoApi
	{
		sheBeiMingXiBiaoRouter.POST("createSheBeiMingXiBiao", sheBeiMingXiBiaoApi.CreateSheBeiMingXiBiao)   // 新建SheBeiMingXiBiao
		sheBeiMingXiBiaoRouter.DELETE("deleteSheBeiMingXiBiao", sheBeiMingXiBiaoApi.DeleteSheBeiMingXiBiao) // 删除SheBeiMingXiBiao
		sheBeiMingXiBiaoRouter.DELETE("deleteSheBeiMingXiBiaoByIds", sheBeiMingXiBiaoApi.DeleteSheBeiMingXiBiaoByIds) // 批量删除SheBeiMingXiBiao
		sheBeiMingXiBiaoRouter.PUT("updateSheBeiMingXiBiao", sheBeiMingXiBiaoApi.UpdateSheBeiMingXiBiao)    // 更新SheBeiMingXiBiao
	}
	{
		sheBeiMingXiBiaoRouterWithoutRecord.GET("findSheBeiMingXiBiao", sheBeiMingXiBiaoApi.FindSheBeiMingXiBiao)        // 根据ID获取SheBeiMingXiBiao
		sheBeiMingXiBiaoRouterWithoutRecord.GET("getSheBeiMingXiBiaoList", sheBeiMingXiBiaoApi.GetSheBeiMingXiBiaoList)  // 获取SheBeiMingXiBiao列表
	}
}
