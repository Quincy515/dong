package hr

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type XiaoXiTongZhiRouter struct {
}

// InitXiaoXiTongZhiRouter 初始化 XiaoXiTongZhi 路由信息
func (s *XiaoXiTongZhiRouter) InitXiaoXiTongZhiRouter(Router *gin.RouterGroup) {
	xiaoXiTongZhiRouter := Router.Group("xiaoXiTongZhi").Use(middleware.OperationRecord())
	xiaoXiTongZhiRouterWithoutRecord := Router.Group("xiaoXiTongZhi")
	var xiaoXiTongZhiApi = v1.ApiGroupApp.HrApiGroup.XiaoXiTongZhiApi
	{
		xiaoXiTongZhiRouter.POST("createXiaoXiTongZhi", xiaoXiTongZhiApi.CreateXiaoXiTongZhi)             // 新建XiaoXiTongZhi
		xiaoXiTongZhiRouter.DELETE("deleteXiaoXiTongZhi", xiaoXiTongZhiApi.DeleteXiaoXiTongZhi)           // 删除XiaoXiTongZhi
		xiaoXiTongZhiRouter.DELETE("deleteXiaoXiTongZhiByIds", xiaoXiTongZhiApi.DeleteXiaoXiTongZhiByIds) // 批量删除XiaoXiTongZhi
		xiaoXiTongZhiRouter.PUT("updateXiaoXiTongZhi", xiaoXiTongZhiApi.UpdateXiaoXiTongZhi)              // 更新XiaoXiTongZhi
	}
	{
		xiaoXiTongZhiRouterWithoutRecord.GET("findXiaoXiTongZhi", xiaoXiTongZhiApi.FindXiaoXiTongZhi)       // 根据ID获取XiaoXiTongZhi
		xiaoXiTongZhiRouterWithoutRecord.GET("getXiaoXiTongZhiList", xiaoXiTongZhiApi.GetXiaoXiTongZhiList) // 获取XiaoXiTongZhi列表
	}
}
