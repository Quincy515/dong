package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ZhengShuJieChuJiLuBiaoRouter struct {
}

// InitZhengShuJieChuJiLuBiaoRouter 初始化 ZhengShuJieChuJiLuBiao 路由信息
func (s *ZhengShuJieChuJiLuBiaoRouter) InitZhengShuJieChuJiLuBiaoRouter(Router *gin.RouterGroup) {
	zhengShuJieChuJiLuBiaoRouter := Router.Group("zhengShuJieChuJiLuBiao").Use(middleware.OperationRecord())
	zhengShuJieChuJiLuBiaoRouterWithoutRecord := Router.Group("zhengShuJieChuJiLuBiao")
	var zhengShuJieChuJiLuBiaoApi = v1.ApiGroupApp.HrApiGroup.ZhengShuJieChuJiLuBiaoApi
	{
		zhengShuJieChuJiLuBiaoRouter.POST("createZhengShuJieChuJiLuBiao", zhengShuJieChuJiLuBiaoApi.CreateZhengShuJieChuJiLuBiao)   // 新建ZhengShuJieChuJiLuBiao
		zhengShuJieChuJiLuBiaoRouter.DELETE("deleteZhengShuJieChuJiLuBiao", zhengShuJieChuJiLuBiaoApi.DeleteZhengShuJieChuJiLuBiao) // 删除ZhengShuJieChuJiLuBiao
		zhengShuJieChuJiLuBiaoRouter.DELETE("deleteZhengShuJieChuJiLuBiaoByIds", zhengShuJieChuJiLuBiaoApi.DeleteZhengShuJieChuJiLuBiaoByIds) // 批量删除ZhengShuJieChuJiLuBiao
		zhengShuJieChuJiLuBiaoRouter.PUT("updateZhengShuJieChuJiLuBiao", zhengShuJieChuJiLuBiaoApi.UpdateZhengShuJieChuJiLuBiao)    // 更新ZhengShuJieChuJiLuBiao
	}
	{
		zhengShuJieChuJiLuBiaoRouterWithoutRecord.GET("findZhengShuJieChuJiLuBiao", zhengShuJieChuJiLuBiaoApi.FindZhengShuJieChuJiLuBiao)        // 根据ID获取ZhengShuJieChuJiLuBiao
		zhengShuJieChuJiLuBiaoRouterWithoutRecord.GET("getZhengShuJieChuJiLuBiaoList", zhengShuJieChuJiLuBiaoApi.GetZhengShuJieChuJiLuBiaoList)  // 获取ZhengShuJieChuJiLuBiao列表
	}
}
