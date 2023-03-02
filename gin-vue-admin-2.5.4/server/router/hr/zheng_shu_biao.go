package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ZhengShuBiaoRouter struct {
}

// InitZhengShuBiaoRouter 初始化 ZhengShuBiao 路由信息
func (s *ZhengShuBiaoRouter) InitZhengShuBiaoRouter(Router *gin.RouterGroup) {
	zhengShuBiaoRouter := Router.Group("zhengShuBiao").Use(middleware.OperationRecord())
	zhengShuBiaoRouterWithoutRecord := Router.Group("zhengShuBiao")
	var zhengShuBiaoApi = v1.ApiGroupApp.HrApiGroup.ZhengShuBiaoApi
	{
		zhengShuBiaoRouter.POST("createZhengShuBiao", zhengShuBiaoApi.CreateZhengShuBiao)   // 新建ZhengShuBiao
		zhengShuBiaoRouter.DELETE("deleteZhengShuBiao", zhengShuBiaoApi.DeleteZhengShuBiao) // 删除ZhengShuBiao
		zhengShuBiaoRouter.DELETE("deleteZhengShuBiaoByIds", zhengShuBiaoApi.DeleteZhengShuBiaoByIds) // 批量删除ZhengShuBiao
		zhengShuBiaoRouter.PUT("updateZhengShuBiao", zhengShuBiaoApi.UpdateZhengShuBiao)    // 更新ZhengShuBiao
	}
	{
		zhengShuBiaoRouterWithoutRecord.GET("findZhengShuBiao", zhengShuBiaoApi.FindZhengShuBiao)        // 根据ID获取ZhengShuBiao
		zhengShuBiaoRouterWithoutRecord.GET("getZhengShuBiaoList", zhengShuBiaoApi.GetZhengShuBiaoList)  // 获取ZhengShuBiao列表
	}
}
