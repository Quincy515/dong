package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type QingJiaGuanLiRouter struct {
}

// InitQingJiaGuanLiRouter 初始化 QingJiaGuanLi 路由信息
func (s *QingJiaGuanLiRouter) InitQingJiaGuanLiRouter(Router *gin.RouterGroup) {
	qingJiaGuanLiRouter := Router.Group("qingJiaGuanLi").Use(middleware.OperationRecord())
	qingJiaGuanLiRouterWithoutRecord := Router.Group("qingJiaGuanLi")
	var qingJiaGuanLiApi = v1.ApiGroupApp.HrApiGroup.QingJiaGuanLiApi
	{
		qingJiaGuanLiRouter.POST("createQingJiaGuanLi", qingJiaGuanLiApi.CreateQingJiaGuanLi)   // 新建QingJiaGuanLi
		qingJiaGuanLiRouter.DELETE("deleteQingJiaGuanLi", qingJiaGuanLiApi.DeleteQingJiaGuanLi) // 删除QingJiaGuanLi
		qingJiaGuanLiRouter.DELETE("deleteQingJiaGuanLiByIds", qingJiaGuanLiApi.DeleteQingJiaGuanLiByIds) // 批量删除QingJiaGuanLi
		qingJiaGuanLiRouter.PUT("updateQingJiaGuanLi", qingJiaGuanLiApi.UpdateQingJiaGuanLi)    // 更新QingJiaGuanLi
	}
	{
		qingJiaGuanLiRouterWithoutRecord.GET("findQingJiaGuanLi", qingJiaGuanLiApi.FindQingJiaGuanLi)        // 根据ID获取QingJiaGuanLi
		qingJiaGuanLiRouterWithoutRecord.GET("getQingJiaGuanLiList", qingJiaGuanLiApi.GetQingJiaGuanLiList)  // 获取QingJiaGuanLi列表
	}
}
