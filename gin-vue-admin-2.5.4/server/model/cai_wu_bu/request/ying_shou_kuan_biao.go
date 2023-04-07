package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/cai_wu_bu"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type YingShouKuanBiaoSearch struct{
    cai_wu_bu.YingShouKuanBiao
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    StartFeiYongFaShengRiQi  *time.Time  `json:"startFeiYongFaShengRiQi" form:"startFeiYongFaShengRiQi"`
    EndFeiYongFaShengRiQi  *time.Time  `json:"endFeiYongFaShengRiQi" form:"endFeiYongFaShengRiQi"`
    StartShenQingRiQi  *time.Time  `json:"startShenQingRiQi" form:"startShenQingRiQi"`
    EndShenQingRiQi  *time.Time  `json:"endShenQingRiQi" form:"endShenQingRiQi"`
    request.PageInfo
}
