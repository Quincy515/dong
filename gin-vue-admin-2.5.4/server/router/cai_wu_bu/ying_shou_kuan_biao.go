package cai_wu_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type YingShouKuanBiaoRouter struct {
}

// InitYingShouKuanBiaoRouter 初始化 YingShouKuanBiao 路由信息
func (s *YingShouKuanBiaoRouter) InitYingShouKuanBiaoRouter(Router *gin.RouterGroup) {
	yingShouKuanBiaoRouter := Router.Group("yingShouKuanBiao").Use(middleware.OperationRecord())
	yingShouKuanBiaoRouterWithoutRecord := Router.Group("yingShouKuanBiao")
	var yingShouKuanBiaoApi = v1.ApiGroupApp.Cai_wu_buApiGroup.YingShouKuanBiaoApi
	{
		yingShouKuanBiaoRouter.POST("createYingShouKuanBiao", yingShouKuanBiaoApi.CreateYingShouKuanBiao)   // 新建YingShouKuanBiao
		yingShouKuanBiaoRouter.DELETE("deleteYingShouKuanBiao", yingShouKuanBiaoApi.DeleteYingShouKuanBiao) // 删除YingShouKuanBiao
		yingShouKuanBiaoRouter.DELETE("deleteYingShouKuanBiaoByIds", yingShouKuanBiaoApi.DeleteYingShouKuanBiaoByIds) // 批量删除YingShouKuanBiao
		yingShouKuanBiaoRouter.PUT("updateYingShouKuanBiao", yingShouKuanBiaoApi.UpdateYingShouKuanBiao)    // 更新YingShouKuanBiao
	}
	{
		yingShouKuanBiaoRouterWithoutRecord.GET("findYingShouKuanBiao", yingShouKuanBiaoApi.FindYingShouKuanBiao)        // 根据ID获取YingShouKuanBiao
		yingShouKuanBiaoRouterWithoutRecord.GET("getYingShouKuanBiaoList", yingShouKuanBiaoApi.GetYingShouKuanBiaoList)  // 获取YingShouKuanBiao列表
	}
}
