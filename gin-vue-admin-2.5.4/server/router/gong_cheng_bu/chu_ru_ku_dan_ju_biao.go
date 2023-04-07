package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ChuRuKuDanJuBiaoRouter struct {
}

// InitChuRuKuDanJuBiaoRouter 初始化 ChuRuKuDanJuBiao 路由信息
func (s *ChuRuKuDanJuBiaoRouter) InitChuRuKuDanJuBiaoRouter(Router *gin.RouterGroup) {
	chuRuKuDanJuBiaoRouter := Router.Group("chuRuKuDanJuBiao").Use(middleware.OperationRecord())
	chuRuKuDanJuBiaoRouterWithoutRecord := Router.Group("chuRuKuDanJuBiao")
	var chuRuKuDanJuBiaoApi = v1.ApiGroupApp.Gong_cheng_buApiGroup.ChuRuKuDanJuBiaoApi
	{
		chuRuKuDanJuBiaoRouter.POST("createChuRuKuDanJuBiao", chuRuKuDanJuBiaoApi.CreateChuRuKuDanJuBiao)   // 新建ChuRuKuDanJuBiao
		chuRuKuDanJuBiaoRouter.DELETE("deleteChuRuKuDanJuBiao", chuRuKuDanJuBiaoApi.DeleteChuRuKuDanJuBiao) // 删除ChuRuKuDanJuBiao
		chuRuKuDanJuBiaoRouter.DELETE("deleteChuRuKuDanJuBiaoByIds", chuRuKuDanJuBiaoApi.DeleteChuRuKuDanJuBiaoByIds) // 批量删除ChuRuKuDanJuBiao
		chuRuKuDanJuBiaoRouter.PUT("updateChuRuKuDanJuBiao", chuRuKuDanJuBiaoApi.UpdateChuRuKuDanJuBiao)    // 更新ChuRuKuDanJuBiao
	}
	{
		chuRuKuDanJuBiaoRouterWithoutRecord.GET("findChuRuKuDanJuBiao", chuRuKuDanJuBiaoApi.FindChuRuKuDanJuBiao)        // 根据ID获取ChuRuKuDanJuBiao
		chuRuKuDanJuBiaoRouterWithoutRecord.GET("getChuRuKuDanJuBiaoList", chuRuKuDanJuBiaoApi.GetChuRuKuDanJuBiaoList)  // 获取ChuRuKuDanJuBiao列表
	}
}
