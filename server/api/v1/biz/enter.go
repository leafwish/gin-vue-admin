package biz

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	CollaboratorApi
}

var (
	collaboratorService = service.ServiceGroupApp.BizServiceGroup.CollaboratorService
)
