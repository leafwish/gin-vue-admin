package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CollaboratorTeam 合作团队
type CollaboratorTeam struct {
	global.GVA_MODEL
	TeamName          string `json:"teamName" form:"teamName" gorm:"comment:合作团队名称"`
	LeaderName        string `json:"leaderName" form:"leaderName" gorm:"comment:合作团队负责人名称"`
	LeaderPhoneNumber string `json:"leaderPhoneNumber" form:"leaderPhoneNumber" gorm:"comment:合作团队负责人手机号"`
}
