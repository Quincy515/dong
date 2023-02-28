package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GongZhangDengJiRouter struct {
}

// InitGongZhangDengJiRouter 初始化 GongZhangDengJi 路由信息
func (s *GongZhangDengJiRouter) InitGongZhangDengJiRouter(Router *gin.RouterGroup) {
	gongZhangDengJiRouter := Router.Group("gongZhangDengJi").Use(middleware.OperationRecord())
	gongZhangDengJiRouterWithoutRecord := Router.Group("gongZhangDengJi")
	var gongZhangDengJiApi = v1.ApiGroupApp.HrApiGroup.GongZhangDengJiApi
	{
		gongZhangDengJiRouter.POST("createGongZhangDengJi", gongZhangDengJiApi.CreateGongZhangDengJi)   // 新建GongZhangDengJi
		gongZhangDengJiRouter.DELETE("deleteGongZhangDengJi", gongZhangDengJiApi.DeleteGongZhangDengJi) // 删除GongZhangDengJi
		gongZhangDengJiRouter.DELETE("deleteGongZhangDengJiByIds", gongZhangDengJiApi.DeleteGongZhangDengJiByIds) // 批量删除GongZhangDengJi
		gongZhangDengJiRouter.PUT("updateGongZhangDengJi", gongZhangDengJiApi.UpdateGongZhangDengJi)    // 更新GongZhangDengJi
	}
	{
		gongZhangDengJiRouterWithoutRecord.GET("findGongZhangDengJi", gongZhangDengJiApi.FindGongZhangDengJi)        // 根据ID获取GongZhangDengJi
		gongZhangDengJiRouterWithoutRecord.GET("getGongZhangDengJiList", gongZhangDengJiApi.GetGongZhangDengJiList)  // 获取GongZhangDengJi列表
	}
}
