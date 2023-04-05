package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/gong_cheng_bu"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type XiangMuBiaoSearch struct{
    gong_cheng_bu.XiangMuBiao
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
