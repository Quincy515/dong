package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type XiangMuZhuCaiBiaoRouter struct {
}

// InitXiangMuZhuCaiBiaoRouter 初始化 XiangMuZhuCaiBiao 路由信息
func (s *XiangMuZhuCaiBiaoRouter) InitXiangMuZhuCaiBiaoRouter(Router *gin.RouterGroup) {
	xiangMuZhuCaiBiaoRouter := Router.Group("xiangMuZhuCaiBiao").Use(middleware.OperationRecord())
	xiangMuZhuCaiBiaoRouterWithoutRecord := Router.Group("xiangMuZhuCaiBiao")
	var xiangMuZhuCaiBiaoApi = v1.ApiGroupApp.Gong_cheng_buApiGroup.XiangMuZhuCaiBiaoApi
	{
		xiangMuZhuCaiBiaoRouter.POST("createXiangMuZhuCaiBiao", xiangMuZhuCaiBiaoApi.CreateXiangMuZhuCaiBiao)   // 新建XiangMuZhuCaiBiao
		xiangMuZhuCaiBiaoRouter.DELETE("deleteXiangMuZhuCaiBiao", xiangMuZhuCaiBiaoApi.DeleteXiangMuZhuCaiBiao) // 删除XiangMuZhuCaiBiao
		xiangMuZhuCaiBiaoRouter.DELETE("deleteXiangMuZhuCaiBiaoByIds", xiangMuZhuCaiBiaoApi.DeleteXiangMuZhuCaiBiaoByIds) // 批量删除XiangMuZhuCaiBiao
		xiangMuZhuCaiBiaoRouter.PUT("updateXiangMuZhuCaiBiao", xiangMuZhuCaiBiaoApi.UpdateXiangMuZhuCaiBiao)    // 更新XiangMuZhuCaiBiao
	}
	{
		xiangMuZhuCaiBiaoRouterWithoutRecord.GET("findXiangMuZhuCaiBiao", xiangMuZhuCaiBiaoApi.FindXiangMuZhuCaiBiao)        // 根据ID获取XiangMuZhuCaiBiao
		xiangMuZhuCaiBiaoRouterWithoutRecord.GET("getXiangMuZhuCaiBiaoList", xiangMuZhuCaiBiaoApi.GetXiangMuZhuCaiBiaoList)  // 获取XiangMuZhuCaiBiao列表
	}
}
