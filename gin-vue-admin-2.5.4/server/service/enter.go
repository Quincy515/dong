package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/hr"
	"github.com/flipped-aurora/gin-vue-admin/server/service/jing_ying_bu"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup       system.ServiceGroup
	ExampleServiceGroup      example.ServiceGroup
	HrServiceGroup           hr.ServiceGroup
	Jing_ying_buServiceGroup jing_ying_bu.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
