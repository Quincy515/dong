package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MeiRiKaoQinRouter struct {
}

// InitMeiRiKaoQinRouter 初始化 MeiRiKaoQin 路由信息
func (s *MeiRiKaoQinRouter) InitMeiRiKaoQinRouter(Router *gin.RouterGroup) {
	meiRiKaoQinRouter := Router.Group("meiRiKaoQin").Use(middleware.OperationRecord())
	meiRiKaoQinRouterWithoutRecord := Router.Group("meiRiKaoQin")
	var meiRiKaoQinApi = v1.ApiGroupApp.HrApiGroup.MeiRiKaoQinApi
	{
		meiRiKaoQinRouter.POST("createMeiRiKaoQin", meiRiKaoQinApi.CreateMeiRiKaoQin)   // 新建MeiRiKaoQin
		meiRiKaoQinRouter.DELETE("deleteMeiRiKaoQin", meiRiKaoQinApi.DeleteMeiRiKaoQin) // 删除MeiRiKaoQin
		meiRiKaoQinRouter.DELETE("deleteMeiRiKaoQinByIds", meiRiKaoQinApi.DeleteMeiRiKaoQinByIds) // 批量删除MeiRiKaoQin
		meiRiKaoQinRouter.PUT("updateMeiRiKaoQin", meiRiKaoQinApi.UpdateMeiRiKaoQin)    // 更新MeiRiKaoQin
	}
	{
		meiRiKaoQinRouterWithoutRecord.GET("findMeiRiKaoQin", meiRiKaoQinApi.FindMeiRiKaoQin)        // 根据ID获取MeiRiKaoQin
		meiRiKaoQinRouterWithoutRecord.GET("getMeiRiKaoQinList", meiRiKaoQinApi.GetMeiRiKaoQinList)  // 获取MeiRiKaoQin列表
	}
}
