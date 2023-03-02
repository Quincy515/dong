package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/hr"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ZhengShuJieChuJiLuBiaoSearch struct{
    hr.ZhengShuJieChuJiLuBiao
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
