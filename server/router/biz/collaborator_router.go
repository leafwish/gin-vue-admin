package biz

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CollaboratorRouter struct{}

func (e *CollaboratorRouter) InitCollaboratorRouter(Router *gin.RouterGroup, ExternalRouter *gin.RouterGroup) {
	//collaboratorRouter := Router.Group("collaborator").Use(middleware.OperationRecord())
	routerWithoutRecord := Router.Group("collaborator")
	publicRouterWithoutRecord := ExternalRouter.Group("collaborator")

	collaboratorApi := v1.ApiGroupApp.BizApiGroup.CollaboratorApi
	{
		//collaboratorRouter.POST("collaborator", collaboratorApi.CreateExaCustomer)   // 创建客户
		//collaboratorRouter.PUT("collaborator", collaboratorApi.UpdateExaCustomer)    // 更新客户
		//collaboratorRouter.DELETE("collaborator", collaboratorApi.DeleteExaCustomer) // 删除客户
	}
	{
		routerWithoutRecord.GET("collaborator", collaboratorApi.GetCollaborator)          // 获取合作人详情
		routerWithoutRecord.GET("collaborator/page", collaboratorApi.GetCollaboratorPage) // 获取合作人分页
	}
	{
		//publicRouterWithoutRecord.GET("collaborator/page", collaboratorApi.GetCollaboratorPage) // 获取合作人分页
		publicRouterWithoutRecord.GET("team/page", collaboratorApi.GetCollaboratorTeamPage) // 获取合作团队分页
		publicRouterWithoutRecord.GET("team", collaboratorApi.GetCollaboratorTeam)          // 获取合作团队
		publicRouterWithoutRecord.POST("team", collaboratorApi.CreateCollaboratorTeam)      // 新增合作团队
		publicRouterWithoutRecord.PUT("team", collaboratorApi.UpdateCollaboratorTeam)       // 修改合作团队
		publicRouterWithoutRecord.DELETE("team", collaboratorApi.DeleteCollaboratorTeam)    // 删除合作团队
	}
}
