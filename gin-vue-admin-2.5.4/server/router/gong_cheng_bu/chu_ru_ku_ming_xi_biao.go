package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ChuRuKuMingXiBiaoRouter struct {
}

// InitChuRuKuMingXiBiaoRouter 初始化 ChuRuKuMingXiBiao 路由信息
func (s *ChuRuKuMingXiBiaoRouter) InitChuRuKuMingXiBiaoRouter(Router *gin.RouterGroup) {
	chuRuKuMingXiBiaoRouter := Router.Group("chuRuKuMingXiBiao").Use(middleware.OperationRecord())
	chuRuKuMingXiBiaoRouterWithoutRecord := Router.Group("chuRuKuMingXiBiao")
	var chuRuKuMingXiBiaoApi = v1.ApiGroupApp.Gong_cheng_buApiGroup.ChuRuKuMingXiBiaoApi
	{
		chuRuKuMingXiBiaoRouter.POST("createChuRuKuMingXiBiao", chuRuKuMingXiBiaoApi.CreateChuRuKuMingXiBiao)   // 新建ChuRuKuMingXiBiao
		chuRuKuMingXiBiaoRouter.DELETE("deleteChuRuKuMingXiBiao", chuRuKuMingXiBiaoApi.DeleteChuRuKuMingXiBiao) // 删除ChuRuKuMingXiBiao
		chuRuKuMingXiBiaoRouter.DELETE("deleteChuRuKuMingXiBiaoByIds", chuRuKuMingXiBiaoApi.DeleteChuRuKuMingXiBiaoByIds) // 批量删除ChuRuKuMingXiBiao
		chuRuKuMingXiBiaoRouter.PUT("updateChuRuKuMingXiBiao", chuRuKuMingXiBiaoApi.UpdateChuRuKuMingXiBiao)    // 更新ChuRuKuMingXiBiao
	}
	{
		chuRuKuMingXiBiaoRouterWithoutRecord.GET("findChuRuKuMingXiBiao", chuRuKuMingXiBiaoApi.FindChuRuKuMingXiBiao)        // 根据ID获取ChuRuKuMingXiBiao
		chuRuKuMingXiBiaoRouterWithoutRecord.GET("getChuRuKuMingXiBiaoList", chuRuKuMingXiBiaoApi.GetChuRuKuMingXiBiaoList)  // 获取ChuRuKuMingXiBiao列表
	}
}
