package gong_cheng_bu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShiGongJinDuBiaoRouter struct {
}

// InitShiGongJinDuBiaoRouter 初始化 ShiGongJinDuBiao 路由信息
func (s *ShiGongJinDuBiaoRouter) InitShiGongJinDuBiaoRouter(Router *gin.RouterGroup) {
	shiGongJinDuBiaoRouter := Router.Group("shiGongJinDuBiao").Use(middleware.OperationRecord())
	shiGongJinDuBiaoRouterWithoutRecord := Router.Group("shiGongJinDuBiao")
	var shiGongJinDuBiaoApi = v1.ApiGroupApp.Gong_cheng_buApiGroup.ShiGongJinDuBiaoApi
	{
		shiGongJinDuBiaoRouter.POST("createShiGongJinDuBiao", shiGongJinDuBiaoApi.CreateShiGongJinDuBiao)   // 新建ShiGongJinDuBiao
		shiGongJinDuBiaoRouter.DELETE("deleteShiGongJinDuBiao", shiGongJinDuBiaoApi.DeleteShiGongJinDuBiao) // 删除ShiGongJinDuBiao
		shiGongJinDuBiaoRouter.DELETE("deleteShiGongJinDuBiaoByIds", shiGongJinDuBiaoApi.DeleteShiGongJinDuBiaoByIds) // 批量删除ShiGongJinDuBiao
		shiGongJinDuBiaoRouter.PUT("updateShiGongJinDuBiao", shiGongJinDuBiaoApi.UpdateShiGongJinDuBiao)    // 更新ShiGongJinDuBiao
	}
	{
		shiGongJinDuBiaoRouterWithoutRecord.GET("findShiGongJinDuBiao", shiGongJinDuBiaoApi.FindShiGongJinDuBiao)        // 根据ID获取ShiGongJinDuBiao
		shiGongJinDuBiaoRouterWithoutRecord.GET("getShiGongJinDuBiaoList", shiGongJinDuBiaoApi.GetShiGongJinDuBiaoList)  // 获取ShiGongJinDuBiao列表
	}
}
