package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WenJianZiLiaoBiaoRouter struct {
}

// InitWenJianZiLiaoBiaoRouter 初始化 WenJianZiLiaoBiao 路由信息
func (s *WenJianZiLiaoBiaoRouter) InitWenJianZiLiaoBiaoRouter(Router *gin.RouterGroup) {
	wenJianZiLiaoBiaoRouter := Router.Group("wenJianZiLiaoBiao").Use(middleware.OperationRecord())
	wenJianZiLiaoBiaoRouterWithoutRecord := Router.Group("wenJianZiLiaoBiao")
	var wenJianZiLiaoBiaoApi = v1.ApiGroupApp.Gong_cheng_buApiGroup.WenJianZiLiaoBiaoApi
	{
		wenJianZiLiaoBiaoRouter.POST("createWenJianZiLiaoBiao", wenJianZiLiaoBiaoApi.CreateWenJianZiLiaoBiao)   // 新建WenJianZiLiaoBiao
		wenJianZiLiaoBiaoRouter.DELETE("deleteWenJianZiLiaoBiao", wenJianZiLiaoBiaoApi.DeleteWenJianZiLiaoBiao) // 删除WenJianZiLiaoBiao
		wenJianZiLiaoBiaoRouter.DELETE("deleteWenJianZiLiaoBiaoByIds", wenJianZiLiaoBiaoApi.DeleteWenJianZiLiaoBiaoByIds) // 批量删除WenJianZiLiaoBiao
		wenJianZiLiaoBiaoRouter.PUT("updateWenJianZiLiaoBiao", wenJianZiLiaoBiaoApi.UpdateWenJianZiLiaoBiao)    // 更新WenJianZiLiaoBiao
	}
	{
		wenJianZiLiaoBiaoRouterWithoutRecord.GET("findWenJianZiLiaoBiao", wenJianZiLiaoBiaoApi.FindWenJianZiLiaoBiao)        // 根据ID获取WenJianZiLiaoBiao
		wenJianZiLiaoBiaoRouterWithoutRecord.GET("getWenJianZiLiaoBiaoList", wenJianZiLiaoBiaoApi.GetWenJianZiLiaoBiaoList)  // 获取WenJianZiLiaoBiao列表
	}
}
