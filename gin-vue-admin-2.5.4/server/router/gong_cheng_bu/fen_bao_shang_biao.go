package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FenBaoShangBiaoRouter struct {
}

// InitFenBaoShangBiaoRouter 初始化 FenBaoShangBiao 路由信息
func (s *FenBaoShangBiaoRouter) InitFenBaoShangBiaoRouter(Router *gin.RouterGroup) {
	fenBaoShangBiaoRouter := Router.Group("fenBaoShangBiao").Use(middleware.OperationRecord())
	fenBaoShangBiaoRouterWithoutRecord := Router.Group("fenBaoShangBiao")
	var fenBaoShangBiaoApi = v1.ApiGroupApp.Gong_cheng_buApiGroup.FenBaoShangBiaoApi
	{
		fenBaoShangBiaoRouter.POST("createFenBaoShangBiao", fenBaoShangBiaoApi.CreateFenBaoShangBiao)   // 新建FenBaoShangBiao
		fenBaoShangBiaoRouter.DELETE("deleteFenBaoShangBiao", fenBaoShangBiaoApi.DeleteFenBaoShangBiao) // 删除FenBaoShangBiao
		fenBaoShangBiaoRouter.DELETE("deleteFenBaoShangBiaoByIds", fenBaoShangBiaoApi.DeleteFenBaoShangBiaoByIds) // 批量删除FenBaoShangBiao
		fenBaoShangBiaoRouter.PUT("updateFenBaoShangBiao", fenBaoShangBiaoApi.UpdateFenBaoShangBiao)    // 更新FenBaoShangBiao
	}
	{
		fenBaoShangBiaoRouterWithoutRecord.GET("findFenBaoShangBiao", fenBaoShangBiaoApi.FindFenBaoShangBiao)        // 根据ID获取FenBaoShangBiao
		fenBaoShangBiaoRouterWithoutRecord.GET("getFenBaoShangBiaoList", fenBaoShangBiaoApi.GetFenBaoShangBiaoList)  // 获取FenBaoShangBiao列表
	}
}
