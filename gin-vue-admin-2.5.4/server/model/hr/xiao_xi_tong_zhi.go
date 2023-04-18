// 自动生成模板XiaoXiTongZhi
package hr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// XiaoXiTongZhi 结构体
type XiaoXiTongZhi struct {
      global.GVA_MODEL
      BeiZhu  string `json:"beiZhu" form:"beiZhu" gorm:"column:bei_zhu;comment:备注;"`
      Content  string `json:"content" form:"content" gorm:"column:content;comment:内容;"`
      Group  string `json:"group" form:"group" gorm:"column:group;comment:所属;size:100;"`
      ReceiverId  *int `json:"receiverId" form:"receiverId" gorm:"column:receiver_id;comment:接收方 id;size:10;"`
      RedirectUrl  string `json:"redirectUrl" form:"redirectUrl" gorm:"column:redirect_url;comment:跳转地址;size:100;"`
      SenderId  *int `json:"senderId" form:"senderId" gorm:"column:sender_id;comment:发送方 id;size:10;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态 0.未读 1.已读;size:10;"`
      Title  string `json:"title" form:"title" gorm:"column:title;comment:标题;size:100;"`
      Type  string `json:"type" form:"type" gorm:"column:type;comment:类型;size:100;"`
}


// TableName XiaoXiTongZhi 表名
func (XiaoXiTongZhi) TableName() string {
  return "xiao_xi_tong_zhi"
}

