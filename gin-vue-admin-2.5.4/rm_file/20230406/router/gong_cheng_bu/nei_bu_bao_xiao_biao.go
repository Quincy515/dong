package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type NeiBuBaoXiaoBiaoRouter struct {
}

// InitNeiBuBaoXiaoBiaoRouter 初始化 NeiBuBaoXiaoBiao 路由信息
func (s *NeiBuBaoXiaoBiaoRouter) InitNeiBuBaoXiaoBiaoRouter(Router *gin.RouterGroup) {
	neiBuBaoXiaoBiaoRouter := Router.Group("neiBuBaoXiaoBiao").Use(middleware.OperationRecord())
	neiBuBaoXiaoBiaoRouterWithoutRecord := Router.Group("neiBuBaoXiaoBiao")
	var neiBuBaoXiaoBiaoApi = v1.ApiGroupApp.Gong_cheng_buApiGroup.NeiBuBaoXiaoBiaoApi
	{
		neiBuBaoXiaoBiaoRouter.POST("createNeiBuBaoXiaoBiao", neiBuBaoXiaoBiaoApi.CreateNeiBuBaoXiaoBiao)   // 新建NeiBuBaoXiaoBiao
		neiBuBaoXiaoBiaoRouter.DELETE("deleteNeiBuBaoXiaoBiao", neiBuBaoXiaoBiaoApi.DeleteNeiBuBaoXiaoBiao) // 删除NeiBuBaoXiaoBiao
		neiBuBaoXiaoBiaoRouter.DELETE("deleteNeiBuBaoXiaoBiaoByIds", neiBuBaoXiaoBiaoApi.DeleteNeiBuBaoXiaoBiaoByIds) // 批量删除NeiBuBaoXiaoBiao
		neiBuBaoXiaoBiaoRouter.PUT("updateNeiBuBaoXiaoBiao", neiBuBaoXiaoBiaoApi.UpdateNeiBuBaoXiaoBiao)    // 更新NeiBuBaoXiaoBiao
	}
	{
		neiBuBaoXiaoBiaoRouterWithoutRecord.GET("findNeiBuBaoXiaoBiao", neiBuBaoXiaoBiaoApi.FindNeiBuBaoXiaoBiao)        // 根据ID获取NeiBuBaoXiaoBiao
		neiBuBaoXiaoBiaoRouterWithoutRecord.GET("getNeiBuBaoXiaoBiaoList", neiBuBaoXiaoBiaoApi.GetNeiBuBaoXiaoBiaoList)  // 获取NeiBuBaoXiaoBiao列表
	}
}
