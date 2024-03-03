package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CollaboratorService struct{}

func (cs *CollaboratorService) CreateCollaboratorTeam(collaboratorTeam biz.CollaboratorTeam) (err error) {
	err = global.GVA_DB.Create(&collaboratorTeam).Error
	return err
}

func (cs *CollaboratorService) UpdateCollaboratorTeam(collaboratorTeam *biz.CollaboratorTeam) (err error) {
	err = global.GVA_DB.Save(collaboratorTeam).Error
	return err
}

func (cs *CollaboratorService) DeleteCollaboratorTeam(collaboratorTeam biz.CollaboratorTeam) (err error) {
	err = global.GVA_DB.Delete(&collaboratorTeam).Error
	if err != nil {
		return err
	}
	return global.GVA_DB.Delete(&biz.Collaborator{}, "collaborator_team_id = ?", collaboratorTeam.ID).Error
}

func (cs *CollaboratorService) GetCollaboratorTeam(id uint) (collaboratorTeam biz.CollaboratorTeam, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&collaboratorTeam).Error
	return
}

func (cs *CollaboratorService) GetCollaboratorTeamPage(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&biz.CollaboratorTeam{})
	var CollaboratorTeamList []biz.CollaboratorTeam
	err = db.Count(&total).Error
	if err != nil || total == 0 {
		CollaboratorTeamList = make([]biz.CollaboratorTeam, 0)
		return CollaboratorTeamList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&CollaboratorTeamList).Error
	return CollaboratorTeamList, total, err
}

func (cs *CollaboratorService) GetCollaborator(id uint) (collaborator biz.Collaborator, err error) {
	err = global.GVA_DB.Where("id = ?", id).Preload("CollaboratorTeam").First(&collaborator).Error
	return
}

// GetCollaboratorPage
// @description: 分页获取合作人列表
func (cs *CollaboratorService) GetCollaboratorPage(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&biz.Collaborator{})
	var CollaboratorList []biz.Collaborator
	err = db.Limit(limit).Offset(offset).Find(&CollaboratorList).Error
	return CollaboratorList, total, err
}
