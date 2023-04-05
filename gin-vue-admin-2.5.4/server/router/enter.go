package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/gong_cheng_bu"
	"github.com/flipped-aurora/gin-vue-admin/server/router/hr"
	"github.com/flipped-aurora/gin-vue-admin/server/router/jing_ying_bu"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System        system.RouterGroup
	Example       example.RouterGroup
	Hr            hr.RouterGroup
	Jing_ying_bu  jing_ying_bu.RouterGroup
	Gong_cheng_bu gong_cheng_bu.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
