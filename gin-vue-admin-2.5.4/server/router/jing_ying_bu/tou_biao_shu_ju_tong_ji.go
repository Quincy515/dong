package jing_ying_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TouBiaoShuJuTongJiRouter struct {
}

// InitTouBiaoShuJuTongJiRouter 初始化 TouBiaoShuJuTongJi 路由信息
func (s *TouBiaoShuJuTongJiRouter) InitTouBiaoShuJuTongJiRouter(Router *gin.RouterGroup) {
	touBiaoShuJuTongJiRouter := Router.Group("touBiaoShuJuTongJi").Use(middleware.OperationRecord())
	touBiaoShuJuTongJiRouterWithoutRecord := Router.Group("touBiaoShuJuTongJi")
	var touBiaoShuJuTongJiApi = v1.ApiGroupApp.Jing_ying_buApiGroup.TouBiaoShuJuTongJiApi
	{
		touBiaoShuJuTongJiRouter.POST("createTouBiaoShuJuTongJi", touBiaoShuJuTongJiApi.CreateTouBiaoShuJuTongJi)   // 新建TouBiaoShuJuTongJi
		touBiaoShuJuTongJiRouter.DELETE("deleteTouBiaoShuJuTongJi", touBiaoShuJuTongJiApi.DeleteTouBiaoShuJuTongJi) // 删除TouBiaoShuJuTongJi
		touBiaoShuJuTongJiRouter.DELETE("deleteTouBiaoShuJuTongJiByIds", touBiaoShuJuTongJiApi.DeleteTouBiaoShuJuTongJiByIds) // 批量删除TouBiaoShuJuTongJi
		touBiaoShuJuTongJiRouter.PUT("updateTouBiaoShuJuTongJi", touBiaoShuJuTongJiApi.UpdateTouBiaoShuJuTongJi)    // 更新TouBiaoShuJuTongJi
	}
	{
		touBiaoShuJuTongJiRouterWithoutRecord.GET("findTouBiaoShuJuTongJi", touBiaoShuJuTongJiApi.FindTouBiaoShuJuTongJi)        // 根据ID获取TouBiaoShuJuTongJi
		touBiaoShuJuTongJiRouterWithoutRecord.GET("getTouBiaoShuJuTongJiList", touBiaoShuJuTongJiApi.GetTouBiaoShuJuTongJiList)  // 获取TouBiaoShuJuTongJi列表
	}
}
