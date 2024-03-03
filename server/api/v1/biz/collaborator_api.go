package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	bizResp "github.com/flipped-aurora/gin-vue-admin/server/model/biz/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CollaboratorApi struct{}

// CreateCollaboratorTeam 新增合作团队
func (e *CollaboratorApi) CreateCollaboratorTeam(c *gin.Context) {
	var collaboratorTeam biz.CollaboratorTeam
	err := c.ShouldBindJSON(&collaboratorTeam)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(collaboratorTeam, utils.CollaboratorTeam)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = collaboratorService.CreateCollaboratorTeam(collaboratorTeam)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// UpdateCollaboratorTeam 更新合作团队
func (e *CollaboratorApi) UpdateCollaboratorTeam(c *gin.Context) {
	var collaboratorTeam biz.CollaboratorTeam
	err := c.ShouldBindJSON(&collaboratorTeam)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(collaboratorTeam.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(collaboratorTeam, utils.CollaboratorTeam)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = collaboratorService.UpdateCollaboratorTeam(&collaboratorTeam)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// DeleteCollaboratorTeam 删除合作团队
func (e *CollaboratorApi) DeleteCollaboratorTeam(c *gin.Context) {
	var collaboratorTeam biz.CollaboratorTeam
	err := c.ShouldBindJSON(&collaboratorTeam)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(collaboratorTeam.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = collaboratorService.DeleteCollaboratorTeam(collaboratorTeam)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetCollaboratorTeam 获取合作团队
func (e *CollaboratorApi) GetCollaboratorTeam(c *gin.Context) {
	var collaboratorTeam biz.CollaboratorTeam
	err := c.ShouldBindQuery(&collaboratorTeam)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(collaboratorTeam.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := collaboratorService.GetCollaboratorTeam(collaboratorTeam.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(bizResp.CollaboratorData{CollaboratorTeam: data}, "获取成功", c)
}

// GetCollaboratorTeamPage 获取合作团队分页
func (e *CollaboratorApi) GetCollaboratorTeamPage(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	collaboratorTeamList, total, err := collaboratorService.GetCollaboratorTeamPage(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     collaboratorTeamList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetCollaborator 获取合作人
func (e *CollaboratorApi) GetCollaborator(c *gin.Context) {
	var collaborator biz.Collaborator
	err := c.ShouldBindQuery(&collaborator)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(collaborator.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := collaboratorService.GetCollaborator(collaborator.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(bizResp.CollaboratorData{Collaborator: data}, "获取成功", c)
}

func (e *CollaboratorApi) GetCollaboratorPage(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	collaboratorList, total, err := collaboratorService.GetCollaboratorPage(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     collaboratorList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
