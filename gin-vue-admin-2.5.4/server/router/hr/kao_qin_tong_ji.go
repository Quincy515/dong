package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type KaoQinTongJiRouter struct {
}

// InitKaoQinTongJiRouter 初始化 KaoQinTongJi 路由信息
func (s *KaoQinTongJiRouter) InitKaoQinTongJiRouter(Router *gin.RouterGroup) {
	kaoQinTongJiRouter := Router.Group("kaoQinTongJi").Use(middleware.OperationRecord())
	kaoQinTongJiRouterWithoutRecord := Router.Group("kaoQinTongJi")
	var kaoQinTongJiApi = v1.ApiGroupApp.HrApiGroup.KaoQinTongJiApi
	{
		kaoQinTongJiRouter.POST("createKaoQinTongJi", kaoQinTongJiApi.CreateKaoQinTongJi)   // 新建KaoQinTongJi
		kaoQinTongJiRouter.DELETE("deleteKaoQinTongJi", kaoQinTongJiApi.DeleteKaoQinTongJi) // 删除KaoQinTongJi
		kaoQinTongJiRouter.DELETE("deleteKaoQinTongJiByIds", kaoQinTongJiApi.DeleteKaoQinTongJiByIds) // 批量删除KaoQinTongJi
		kaoQinTongJiRouter.PUT("updateKaoQinTongJi", kaoQinTongJiApi.UpdateKaoQinTongJi)    // 更新KaoQinTongJi
	}
	{
		kaoQinTongJiRouterWithoutRecord.GET("findKaoQinTongJi", kaoQinTongJiApi.FindKaoQinTongJi)        // 根据ID获取KaoQinTongJi
		kaoQinTongJiRouterWithoutRecord.GET("getKaoQinTongJiList", kaoQinTongJiApi.GetKaoQinTongJiList)  // 获取KaoQinTongJi列表
	}
}
