package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HuiYuanGuanLiRouter struct {
}

// InitHuiYuanGuanLiRouter 初始化 HuiYuanGuanLi 路由信息
func (s *HuiYuanGuanLiRouter) InitHuiYuanGuanLiRouter(Router *gin.RouterGroup) {
	huiYuanGuanLiRouter := Router.Group("huiYuanGuanLi").Use(middleware.OperationRecord())
	huiYuanGuanLiRouterWithoutRecord := Router.Group("huiYuanGuanLi")
	var huiYuanGuanLiApi = v1.ApiGroupApp.HrApiGroup.HuiYuanGuanLiApi
	{
		huiYuanGuanLiRouter.POST("createHuiYuanGuanLi", huiYuanGuanLiApi.CreateHuiYuanGuanLi)   // 新建HuiYuanGuanLi
		huiYuanGuanLiRouter.DELETE("deleteHuiYuanGuanLi", huiYuanGuanLiApi.DeleteHuiYuanGuanLi) // 删除HuiYuanGuanLi
		huiYuanGuanLiRouter.DELETE("deleteHuiYuanGuanLiByIds", huiYuanGuanLiApi.DeleteHuiYuanGuanLiByIds) // 批量删除HuiYuanGuanLi
		huiYuanGuanLiRouter.PUT("updateHuiYuanGuanLi", huiYuanGuanLiApi.UpdateHuiYuanGuanLi)    // 更新HuiYuanGuanLi
	}
	{
		huiYuanGuanLiRouterWithoutRecord.GET("findHuiYuanGuanLi", huiYuanGuanLiApi.FindHuiYuanGuanLi)        // 根据ID获取HuiYuanGuanLi
		huiYuanGuanLiRouterWithoutRecord.GET("getHuiYuanGuanLiList", huiYuanGuanLiApi.GetHuiYuanGuanLiList)  // 获取HuiYuanGuanLi列表
	}
}
