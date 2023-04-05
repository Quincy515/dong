package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/gong_cheng_bu"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/hr"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/jing_ying_bu"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup        system.ApiGroup
	ExampleApiGroup       example.ApiGroup
	HrApiGroup            hr.ApiGroup
	Jing_ying_buApiGroup  jing_ying_bu.ApiGroup
	Gong_cheng_buApiGroup gong_cheng_bu.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
