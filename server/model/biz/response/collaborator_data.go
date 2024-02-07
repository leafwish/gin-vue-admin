package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/biz"

type CollaboratorData struct {
	Collaborator biz.Collaborator `json:"collaborator"`
}
