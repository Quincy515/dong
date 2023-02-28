package jing_ying_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type XiangMuDengJiBiaoRouter struct {
}

// InitXiangMuDengJiBiaoRouter 初始化 XiangMuDengJiBiao 路由信息
func (s *XiangMuDengJiBiaoRouter) InitXiangMuDengJiBiaoRouter(Router *gin.RouterGroup) {
	xiangMuDengJiBiaoRouter := Router.Group("xiangMuDengJiBiao").Use(middleware.OperationRecord())
	xiangMuDengJiBiaoRouterWithoutRecord := Router.Group("xiangMuDengJiBiao")
	var xiangMuDengJiBiaoApi = v1.ApiGroupApp.Jing_ying_buApiGroup.XiangMuDengJiBiaoApi
	{
		xiangMuDengJiBiaoRouter.POST("createXiangMuDengJiBiao", xiangMuDengJiBiaoApi.CreateXiangMuDengJiBiao)   // 新建XiangMuDengJiBiao
		xiangMuDengJiBiaoRouter.DELETE("deleteXiangMuDengJiBiao", xiangMuDengJiBiaoApi.DeleteXiangMuDengJiBiao) // 删除XiangMuDengJiBiao
		xiangMuDengJiBiaoRouter.DELETE("deleteXiangMuDengJiBiaoByIds", xiangMuDengJiBiaoApi.DeleteXiangMuDengJiBiaoByIds) // 批量删除XiangMuDengJiBiao
		xiangMuDengJiBiaoRouter.PUT("updateXiangMuDengJiBiao", xiangMuDengJiBiaoApi.UpdateXiangMuDengJiBiao)    // 更新XiangMuDengJiBiao
	}
	{
		xiangMuDengJiBiaoRouterWithoutRecord.GET("findXiangMuDengJiBiao", xiangMuDengJiBiaoApi.FindXiangMuDengJiBiao)        // 根据ID获取XiangMuDengJiBiao
		xiangMuDengJiBiaoRouterWithoutRecord.GET("getXiangMuDengJiBiaoList", xiangMuDengJiBiaoApi.GetXiangMuDengJiBiaoList)  // 获取XiangMuDengJiBiao列表
	}
}
