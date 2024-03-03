package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/biz"

type CollaboratorData struct {
	CollaboratorTeam biz.CollaboratorTeam `json:"collaboratorTeam"`
	Collaborator     biz.Collaborator     `json:"collaborator"`
}
