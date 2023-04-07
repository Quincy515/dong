package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/cai_wu_bu"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type NeiBuBaoXiaoBiaoSearch struct{
    cai_wu_bu.NeiBuBaoXiaoBiao
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
