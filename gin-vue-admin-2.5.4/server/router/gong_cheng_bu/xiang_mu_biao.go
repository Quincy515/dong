package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type XiangMuBiaoRouter struct {
}

// InitXiangMuBiaoRouter 初始化 XiangMuBiao 路由信息
func (s *XiangMuBiaoRouter) InitXiangMuBiaoRouter(Router *gin.RouterGroup) {
	xiangMuBiaoRouter := Router.Group("xiangMuBiao").Use(middleware.OperationRecord())
	xiangMuBiaoRouterWithoutRecord := Router.Group("xiangMuBiao")
	var xiangMuBiaoApi = v1.ApiGroupApp.Gong_cheng_buApiGroup.XiangMuBiaoApi
	{
		xiangMuBiaoRouter.POST("createXiangMuBiao", xiangMuBiaoApi.CreateXiangMuBiao)   // 新建XiangMuBiao
		xiangMuBiaoRouter.DELETE("deleteXiangMuBiao", xiangMuBiaoApi.DeleteXiangMuBiao) // 删除XiangMuBiao
		xiangMuBiaoRouter.DELETE("deleteXiangMuBiaoByIds", xiangMuBiaoApi.DeleteXiangMuBiaoByIds) // 批量删除XiangMuBiao
		xiangMuBiaoRouter.PUT("updateXiangMuBiao", xiangMuBiaoApi.UpdateXiangMuBiao)    // 更新XiangMuBiao
	}
	{
		xiangMuBiaoRouterWithoutRecord.GET("findXiangMuBiao", xiangMuBiaoApi.FindXiangMuBiao)        // 根据ID获取XiangMuBiao
		xiangMuBiaoRouterWithoutRecord.GET("getXiangMuBiaoList", xiangMuBiaoApi.GetXiangMuBiaoList)  // 获取XiangMuBiao列表
	}
}
