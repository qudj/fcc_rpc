package handler

import (
	"context"
	"errors"
	"github.com/qudj/fcc_rpc/config"
	"github.com/qudj/fcc_rpc/models"
	"github.com/qudj/fcc_rpc/models/fcc_serv"
	"github.com/qudj/fcc_rpc/service"
	"gorm.io/gorm"
)

func SaveProject(ctx context.Context, req *fcc_serv.SaveProjectRequest) error {
	if req.Project == nil {
		return errors.New("param error")
	}
	var cur, pre models.FccProject
	objectType := "add"
	if err := config.FccWriteDB.WithContext(ctx).Where("project_key = ?", req.Project.ProjectKey).Last(&pre).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
		cur.Id = pre.Id
		cur.ProjectKey = pre.ProjectKey
		cur.ProjectName = req.Project.ProjectName
		cur.Description = req.Project.Description
		cur.Status = req.Project.Status
		objectType = "change"
	}
	if err := models.SaveProject(&cur); err != nil {
		return err
	}

	_ = service.SaveHistory(pre, cur, cur.TableName(), cur.ProjectKey, objectType, req.OpId)
	return nil
}

func SaveGroup(ctx context.Context, req *fcc_serv.SaveGroupRequest) error {
	if req.Group == nil {
		return errors.New("param error")
	}
	var pro models.FccProject
	if err := config.FccWriteDB.WithContext(ctx).Where("project_key = ?", req.Group.ProjectKey).Last(&pro).Error; err != nil {
		return err
	}

	var cur, pre models.FccGroup
	objectType := "add"
	if err := config.FccWriteDB.WithContext(ctx).Where("project_key = ? and group_key = ?", req.Group.ProjectKey, req.Group.GroupKey).Last(&pre).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
		cur.Id = pre.Id
		cur.ProjectKey = pre.ProjectKey
		cur.GroupKey = req.Group.GroupKey
		cur.GroupName = req.Group.GroupName
		cur.Description = req.Group.Description
		cur.Status = req.Group.Status
		objectType = "change"
	}
	if err := models.SaveGroup(&cur); err != nil {
		return err
	}
	_ = service.SaveHistory(pre, cur, cur.TableName(), cur.ProjectKey, objectType, req.OpId)
	return nil
}

func SaveConfig(ctx context.Context, req *fcc_serv.SaveConfigRequest) error {
	if req.Config == nil {
		return errors.New("param error")
	}
	var gro models.FccGroup
	if err := config.FccWriteDB.WithContext(ctx).Where("project_key = ? and group_key = ?", req.Config.ProjectKey, req.Config.GroupKey).Last(&gro).Error; err != nil {
		return err
	}

	var cur, pre models.FccConf
	objectType := "add"
	if err := config.FccWriteDB.WithContext(ctx).Where("project_key = ? and group_key = ? and conf_key = ?", req.Config.ProjectKey, req.Config.GroupKey, req.Config.ConfKey).Last(&pre).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
		cur.Id = pre.Id
		cur.ProjectKey = pre.ProjectKey
		cur.GroupKey = req.Config.GroupKey
		cur.ConfKey = req.Config.ConfKey
		cur.Description = req.Config.Description
		cur.Status = req.Config.Status
		objectType = "change"
	}
	if err := models.SaveConf(&cur); err != nil {
		return err
	}
	_ = service.SaveHistory(pre, cur, cur.TableName(), cur.ProjectKey, objectType, req.OpId)
	return nil
}
