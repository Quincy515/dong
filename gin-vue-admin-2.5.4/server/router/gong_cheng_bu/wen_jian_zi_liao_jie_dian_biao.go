package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WenJianZiLiaoJieDianBiaoRouter struct {
}

// InitWenJianZiLiaoJieDianBiaoRouter 初始化 WenJianZiLiaoJieDianBiao 路由信息
func (s *WenJianZiLiaoJieDianBiaoRouter) InitWenJianZiLiaoJieDianBiaoRouter(Router *gin.RouterGroup) {
	wenJianZiLiaoJieDianBiaoRouter := Router.Group("wenJianZiLiaoJieDianBiao").Use(middleware.OperationRecord())
	wenJianZiLiaoJieDianBiaoRouterWithoutRecord := Router.Group("wenJianZiLiaoJieDianBiao")
	var wenJianZiLiaoJieDianBiaoApi = v1.ApiGroupApp.Gong_cheng_buApiGroup.WenJianZiLiaoJieDianBiaoApi
	{
		wenJianZiLiaoJieDianBiaoRouter.POST("createWenJianZiLiaoJieDianBiao", wenJianZiLiaoJieDianBiaoApi.CreateWenJianZiLiaoJieDianBiao)   // 新建WenJianZiLiaoJieDianBiao
		wenJianZiLiaoJieDianBiaoRouter.DELETE("deleteWenJianZiLiaoJieDianBiao", wenJianZiLiaoJieDianBiaoApi.DeleteWenJianZiLiaoJieDianBiao) // 删除WenJianZiLiaoJieDianBiao
		wenJianZiLiaoJieDianBiaoRouter.DELETE("deleteWenJianZiLiaoJieDianBiaoByIds", wenJianZiLiaoJieDianBiaoApi.DeleteWenJianZiLiaoJieDianBiaoByIds) // 批量删除WenJianZiLiaoJieDianBiao
		wenJianZiLiaoJieDianBiaoRouter.PUT("updateWenJianZiLiaoJieDianBiao", wenJianZiLiaoJieDianBiaoApi.UpdateWenJianZiLiaoJieDianBiao)    // 更新WenJianZiLiaoJieDianBiao
	}
	{
		wenJianZiLiaoJieDianBiaoRouterWithoutRecord.GET("findWenJianZiLiaoJieDianBiao", wenJianZiLiaoJieDianBiaoApi.FindWenJianZiLiaoJieDianBiao)        // 根据ID获取WenJianZiLiaoJieDianBiao
		wenJianZiLiaoJieDianBiaoRouterWithoutRecord.GET("getWenJianZiLiaoJieDianBiaoList", wenJianZiLiaoJieDianBiaoApi.GetWenJianZiLiaoJieDianBiaoList)  // 获取WenJianZiLiaoJieDianBiao列表
	}
}
