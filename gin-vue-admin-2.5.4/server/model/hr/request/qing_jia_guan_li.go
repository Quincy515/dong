package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/hr"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type QingJiaGuanLiSearch struct{
    hr.QingJiaGuanLi
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
