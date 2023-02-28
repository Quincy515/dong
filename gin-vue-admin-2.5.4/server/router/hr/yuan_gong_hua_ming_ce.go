package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type YuanGongHuaMingCeRouter struct {
}

// InitYuanGongHuaMingCeRouter 初始化 YuanGongHuaMingCe 路由信息
func (s *YuanGongHuaMingCeRouter) InitYuanGongHuaMingCeRouter(Router *gin.RouterGroup) {
	yuanGongHuaMingCeRouter := Router.Group("yuanGongHuaMingCe").Use(middleware.OperationRecord())
	yuanGongHuaMingCeRouterWithoutRecord := Router.Group("yuanGongHuaMingCe")
	var yuanGongHuaMingCeApi = v1.ApiGroupApp.HrApiGroup.YuanGongHuaMingCeApi
	{
		yuanGongHuaMingCeRouter.POST("createYuanGongHuaMingCe", yuanGongHuaMingCeApi.CreateYuanGongHuaMingCe)   // 新建YuanGongHuaMingCe
		yuanGongHuaMingCeRouter.DELETE("deleteYuanGongHuaMingCe", yuanGongHuaMingCeApi.DeleteYuanGongHuaMingCe) // 删除YuanGongHuaMingCe
		yuanGongHuaMingCeRouter.DELETE("deleteYuanGongHuaMingCeByIds", yuanGongHuaMingCeApi.DeleteYuanGongHuaMingCeByIds) // 批量删除YuanGongHuaMingCe
		yuanGongHuaMingCeRouter.PUT("updateYuanGongHuaMingCe", yuanGongHuaMingCeApi.UpdateYuanGongHuaMingCe)    // 更新YuanGongHuaMingCe
	}
	{
		yuanGongHuaMingCeRouterWithoutRecord.GET("findYuanGongHuaMingCe", yuanGongHuaMingCeApi.FindYuanGongHuaMingCe)        // 根据ID获取YuanGongHuaMingCe
		yuanGongHuaMingCeRouterWithoutRecord.GET("getYuanGongHuaMingCeList", yuanGongHuaMingCeApi.GetYuanGongHuaMingCeList)  // 获取YuanGongHuaMingCe列表
	}
}
