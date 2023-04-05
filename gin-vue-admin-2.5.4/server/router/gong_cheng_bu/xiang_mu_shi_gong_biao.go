package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type XiangMuShiGongBiaoRouter struct {
}

// InitXiangMuShiGongBiaoRouter 初始化 XiangMuShiGongBiao 路由信息
func (s *XiangMuShiGongBiaoRouter) InitXiangMuShiGongBiaoRouter(Router *gin.RouterGroup) {
	xiangMuShiGongBiaoRouter := Router.Group("xiangMuShiGongBiao").Use(middleware.OperationRecord())
	xiangMuShiGongBiaoRouterWithoutRecord := Router.Group("xiangMuShiGongBiao")
	var xiangMuShiGongBiaoApi = v1.ApiGroupApp.Gong_cheng_buApiGroup.XiangMuShiGongBiaoApi
	{
		xiangMuShiGongBiaoRouter.POST("createXiangMuShiGongBiao", xiangMuShiGongBiaoApi.CreateXiangMuShiGongBiao)   // 新建XiangMuShiGongBiao
		xiangMuShiGongBiaoRouter.DELETE("deleteXiangMuShiGongBiao", xiangMuShiGongBiaoApi.DeleteXiangMuShiGongBiao) // 删除XiangMuShiGongBiao
		xiangMuShiGongBiaoRouter.DELETE("deleteXiangMuShiGongBiaoByIds", xiangMuShiGongBiaoApi.DeleteXiangMuShiGongBiaoByIds) // 批量删除XiangMuShiGongBiao
		xiangMuShiGongBiaoRouter.PUT("updateXiangMuShiGongBiao", xiangMuShiGongBiaoApi.UpdateXiangMuShiGongBiao)    // 更新XiangMuShiGongBiao
	}
	{
		xiangMuShiGongBiaoRouterWithoutRecord.GET("findXiangMuShiGongBiao", xiangMuShiGongBiaoApi.FindXiangMuShiGongBiao)        // 根据ID获取XiangMuShiGongBiao
		xiangMuShiGongBiaoRouterWithoutRecord.GET("getXiangMuShiGongBiaoList", xiangMuShiGongBiaoApi.GetXiangMuShiGongBiaoList)  // 获取XiangMuShiGongBiao列表
	}
}
