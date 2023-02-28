package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GongSiZhangHaoMiMaRouter struct {
}

// InitGongSiZhangHaoMiMaRouter 初始化 GongSiZhangHaoMiMa 路由信息
func (s *GongSiZhangHaoMiMaRouter) InitGongSiZhangHaoMiMaRouter(Router *gin.RouterGroup) {
	gongSiZhangHaoMiMaRouter := Router.Group("gongSiZhangHaoMiMa").Use(middleware.OperationRecord())
	gongSiZhangHaoMiMaRouterWithoutRecord := Router.Group("gongSiZhangHaoMiMa")
	var gongSiZhangHaoMiMaApi = v1.ApiGroupApp.HrApiGroup.GongSiZhangHaoMiMaApi
	{
		gongSiZhangHaoMiMaRouter.POST("createGongSiZhangHaoMiMa", gongSiZhangHaoMiMaApi.CreateGongSiZhangHaoMiMa)   // 新建GongSiZhangHaoMiMa
		gongSiZhangHaoMiMaRouter.DELETE("deleteGongSiZhangHaoMiMa", gongSiZhangHaoMiMaApi.DeleteGongSiZhangHaoMiMa) // 删除GongSiZhangHaoMiMa
		gongSiZhangHaoMiMaRouter.DELETE("deleteGongSiZhangHaoMiMaByIds", gongSiZhangHaoMiMaApi.DeleteGongSiZhangHaoMiMaByIds) // 批量删除GongSiZhangHaoMiMa
		gongSiZhangHaoMiMaRouter.PUT("updateGongSiZhangHaoMiMa", gongSiZhangHaoMiMaApi.UpdateGongSiZhangHaoMiMa)    // 更新GongSiZhangHaoMiMa
	}
	{
		gongSiZhangHaoMiMaRouterWithoutRecord.GET("findGongSiZhangHaoMiMa", gongSiZhangHaoMiMaApi.FindGongSiZhangHaoMiMa)        // 根据ID获取GongSiZhangHaoMiMa
		gongSiZhangHaoMiMaRouterWithoutRecord.GET("getGongSiZhangHaoMiMaList", gongSiZhangHaoMiMaApi.GetGongSiZhangHaoMiMaList)  // 获取GongSiZhangHaoMiMa列表
	}
}
