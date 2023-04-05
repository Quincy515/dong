package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FenBaoShangYuXiangMuGuanLianBiaoRouter struct {
}

// InitFenBaoShangYuXiangMuGuanLianBiaoRouter 初始化 FenBaoShangYuXiangMuGuanLianBiao 路由信息
func (s *FenBaoShangYuXiangMuGuanLianBiaoRouter) InitFenBaoShangYuXiangMuGuanLianBiaoRouter(Router *gin.RouterGroup) {
	fenBaoShangYuXiangMuGuanLianBiaoRouter := Router.Group("fenBaoShangYuXiangMuGuanLianBiao").Use(middleware.OperationRecord())
	fenBaoShangYuXiangMuGuanLianBiaoRouterWithoutRecord := Router.Group("fenBaoShangYuXiangMuGuanLianBiao")
	var fenBaoShangYuXiangMuGuanLianBiaoApi = v1.ApiGroupApp.Gong_cheng_buApiGroup.FenBaoShangYuXiangMuGuanLianBiaoApi
	{
		fenBaoShangYuXiangMuGuanLianBiaoRouter.POST("createFenBaoShangYuXiangMuGuanLianBiao", fenBaoShangYuXiangMuGuanLianBiaoApi.CreateFenBaoShangYuXiangMuGuanLianBiao)   // 新建FenBaoShangYuXiangMuGuanLianBiao
		fenBaoShangYuXiangMuGuanLianBiaoRouter.DELETE("deleteFenBaoShangYuXiangMuGuanLianBiao", fenBaoShangYuXiangMuGuanLianBiaoApi.DeleteFenBaoShangYuXiangMuGuanLianBiao) // 删除FenBaoShangYuXiangMuGuanLianBiao
		fenBaoShangYuXiangMuGuanLianBiaoRouter.DELETE("deleteFenBaoShangYuXiangMuGuanLianBiaoByIds", fenBaoShangYuXiangMuGuanLianBiaoApi.DeleteFenBaoShangYuXiangMuGuanLianBiaoByIds) // 批量删除FenBaoShangYuXiangMuGuanLianBiao
		fenBaoShangYuXiangMuGuanLianBiaoRouter.PUT("updateFenBaoShangYuXiangMuGuanLianBiao", fenBaoShangYuXiangMuGuanLianBiaoApi.UpdateFenBaoShangYuXiangMuGuanLianBiao)    // 更新FenBaoShangYuXiangMuGuanLianBiao
	}
	{
		fenBaoShangYuXiangMuGuanLianBiaoRouterWithoutRecord.GET("findFenBaoShangYuXiangMuGuanLianBiao", fenBaoShangYuXiangMuGuanLianBiaoApi.FindFenBaoShangYuXiangMuGuanLianBiao)        // 根据ID获取FenBaoShangYuXiangMuGuanLianBiao
		fenBaoShangYuXiangMuGuanLianBiaoRouterWithoutRecord.GET("getFenBaoShangYuXiangMuGuanLianBiaoList", fenBaoShangYuXiangMuGuanLianBiaoApi.GetFenBaoShangYuXiangMuGuanLianBiaoList)  // 获取FenBaoShangYuXiangMuGuanLianBiao列表
	}
}
