package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CollaboratorService struct{}

func (cs *CollaboratorService) GetCollaborator(id uint) (collaborator biz.Collaborator, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&collaborator).Error
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
