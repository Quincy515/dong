package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/jing_ying_bu"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type TouBiaoShuJuTongJiSearch struct{
    jing_ying_bu.TouBiaoShuJuTongJi
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
