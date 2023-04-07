package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CaiGouJiHuaMingXiBiaoRouter struct {
}

// InitCaiGouJiHuaMingXiBiaoRouter 初始化 CaiGouJiHuaMingXiBiao 路由信息
func (s *CaiGouJiHuaMingXiBiaoRouter) InitCaiGouJiHuaMingXiBiaoRouter(Router *gin.RouterGroup) {
	caiGouJiHuaMingXiBiaoRouter := Router.Group("caiGouJiHuaMingXiBiao").Use(middleware.OperationRecord())
	caiGouJiHuaMingXiBiaoRouterWithoutRecord := Router.Group("caiGouJiHuaMingXiBiao")
	var caiGouJiHuaMingXiBiaoApi = v1.ApiGroupApp.Gong_cheng_buApiGroup.CaiGouJiHuaMingXiBiaoApi
	{
		caiGouJiHuaMingXiBiaoRouter.POST("createCaiGouJiHuaMingXiBiao", caiGouJiHuaMingXiBiaoApi.CreateCaiGouJiHuaMingXiBiao)   // 新建CaiGouJiHuaMingXiBiao
		caiGouJiHuaMingXiBiaoRouter.DELETE("deleteCaiGouJiHuaMingXiBiao", caiGouJiHuaMingXiBiaoApi.DeleteCaiGouJiHuaMingXiBiao) // 删除CaiGouJiHuaMingXiBiao
		caiGouJiHuaMingXiBiaoRouter.DELETE("deleteCaiGouJiHuaMingXiBiaoByIds", caiGouJiHuaMingXiBiaoApi.DeleteCaiGouJiHuaMingXiBiaoByIds) // 批量删除CaiGouJiHuaMingXiBiao
		caiGouJiHuaMingXiBiaoRouter.PUT("updateCaiGouJiHuaMingXiBiao", caiGouJiHuaMingXiBiaoApi.UpdateCaiGouJiHuaMingXiBiao)    // 更新CaiGouJiHuaMingXiBiao
	}
	{
		caiGouJiHuaMingXiBiaoRouterWithoutRecord.GET("findCaiGouJiHuaMingXiBiao", caiGouJiHuaMingXiBiaoApi.FindCaiGouJiHuaMingXiBiao)        // 根据ID获取CaiGouJiHuaMingXiBiao
		caiGouJiHuaMingXiBiaoRouterWithoutRecord.GET("getCaiGouJiHuaMingXiBiaoList", caiGouJiHuaMingXiBiaoApi.GetCaiGouJiHuaMingXiBiaoList)  // 获取CaiGouJiHuaMingXiBiao列表
	}
}
