package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Collaborator 合作人
type Collaborator struct {
	global.GVA_MODEL
	CollaboratorName string `json:"name" form:"name" gorm:"comment:合作人名称"`
	PhoneNumber      string `json:"phoneNumber" form:"phoneNumber" gorm:"comment:合作人手机号"`
	//CollaboratorTeam CollaboratorTeam `json:"collaboratorTeam"`
}
