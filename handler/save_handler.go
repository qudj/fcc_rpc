package handler

import (
	"context"
	"errors"
	"github.com/qudj/fcc_rpc/config"
	"github.com/qudj/fcc_rpc/models"
	"github.com/qudj/fcc_rpc/models/fcc_serv"
	"github.com/qudj/fcc_rpc/service"
	"gorm.io/gorm"
	"time"
)

func SaveProject(ctx context.Context, req *fcc_serv.SaveProjectRequest) error {
	if req.Project == nil {
		return errors.New("param error")
	}
	pre := &models.FccProject{}
	objectType := "update"
	if err := config.FccWriteDB.WithContext(ctx).Debug().Where("project_key = ?", req.Project.ProjectKey).Last(pre).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
		objectType = "add"
	}
	cur, err := GetCurProject(pre, req)
	if err != nil {
		return err
	}
	if err := models.SaveProject(cur); err != nil {
		return err
	}
	_ = service.SaveHistory(pre, cur, cur.TableName(), cur.ProjectKey, objectType, req.OpId)
	return nil
}

func GetCurProject(pre *models.FccProject, req *fcc_serv.SaveProjectRequest) (*models.FccProject, error) {
	if req.Project.ProjectKey == "" {
		return nil, errors.New("project_key need")
	}
	cur := &models.FccProject{}
	curTime := time.Now().Unix()
	if pre == nil || pre.Id == 0 {
		if req.Project.ProjectName == "" {
			return nil, errors.New("add project need project_key")
		}
		if req.Project.Description == "" {
			return nil, errors.New("add project need description")
		}
		if req.Project.Status == 0 {
			return nil, errors.New("add project need status")
		}
		cur.ProjectKey = req.Project.ProjectKey
		cur.CreateTime = curTime
	} else {
		cur = pre
	}
	if req.Project.ProjectName != "" {
		cur.ProjectName = req.Project.ProjectName
	}
	if req.Project.Description != "" {
		cur.Description = req.Project.Description
	}
	if req.Project.Status != 0 {
		cur.Status = req.Project.Status
	}
	cur.UpdateTime = curTime
	return cur, nil
}

func SaveGroup(ctx context.Context, req *fcc_serv.SaveGroupRequest) error {
	if req.Group == nil {
		return errors.New("param error")
	}
	pro := &models.FccProject{}
	if err := config.FccWriteDB.WithContext(ctx).Where("project_key = ?", req.Group.ProjectKey).Last(pro).Error; err != nil {
		return err
	}
	pre := &models.FccGroup{}
	objectType := "update"
	if err := config.FccWriteDB.WithContext(ctx).Where("project_key = ? and group_key = ?", req.Group.ProjectKey, req.Group.GroupKey).Last(pre).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
		objectType = "add"
	}
	cur, err := GetCurGroup(pre, req)
	if err != nil {
		return err
	}
	if err := models.SaveGroup(cur); err != nil {
		return err
	}
	_ = service.SaveHistory(pre, cur, cur.TableName(), cur.ProjectKey, objectType, req.OpId)
	return nil
}

func GetCurGroup(pre *models.FccGroup, req *fcc_serv.SaveGroupRequest) (*models.FccGroup, error) {
	if req.Group.ProjectKey == "" || req.Group.GroupKey == "" {
		return nil, errors.New("project_key and group_key need")
	}
	cur := &models.FccGroup{}
	curTime := time.Now().Unix()
	if pre == nil || pre.Id == 0 {
		if req.Group.GroupName == "" {
			return nil, errors.New("add project need project_name")
		}
		if req.Group.Description == "" {
			return nil, errors.New("add project need description")
		}
		if req.Group.Status == 0 {
			return nil, errors.New("add project need status")
		}
		cur.ProjectKey = req.Group.ProjectKey
		cur.GroupKey = req.Group.GroupKey
		cur.CreateTime = curTime
	} else {
		cur = pre
	}
	if req.Group.GroupName != "" {
		cur.GroupName = req.Group.GroupName
	}
	if req.Group.Description != "" {
		cur.Description = req.Group.Description
	}
	if req.Group.Status != 0 {
		cur.Status = req.Group.Status
	}
	cur.UpdateTime = curTime
	return cur, nil
}

func SaveConfig(ctx context.Context, req *fcc_serv.SaveConfigRequest) error {
	if req.Config == nil {
		return errors.New("param error")
	}
	gro := &models.FccGroup{}
	if err := config.FccWriteDB.WithContext(ctx).Where("project_key = ? and group_key = ?", req.Config.ProjectKey, req.Config.GroupKey).Last(gro).Error; err != nil {
		return err
	}

	pre := &models.FccConf{}
	objectType := "update"
	if err := config.FccWriteDB.WithContext(ctx).Where("project_key = ? and group_key = ? and conf_key = ?", req.Config.ProjectKey, req.Config.GroupKey, req.Config.ConfKey).Last(pre).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
		objectType = "add"
	}
	cur, err := GetCurConfig(pre, req)
	if err != nil {
		return err
	}
	if err := models.SaveConf(cur); err != nil {
		return err
	}
	_ = service.SaveHistory(pre, cur, cur.TableName(), cur.ProjectKey, objectType, req.OpId)
	return nil
}

func GetCurConfig(pre *models.FccConf, req *fcc_serv.SaveConfigRequest) (*models.FccConf, error) {
	if req.Config.ProjectKey == "" || req.Config.GroupKey == "" || req.Config.ConfKey == "" {
		return nil, errors.New("project_key and group_key and conf_key need")
	}
	cur := &models.FccConf{}
	curTime := time.Now().Unix()
	if pre == nil || pre.Id == 0 {
		if req.Config.Description == "" {
			return nil, errors.New("add project need description")
		}
		if req.Config.Status == 0 {
			return nil, errors.New("add project need status")
		}
		cur.ProjectKey = req.Config.ProjectKey
		cur.GroupKey = req.Config.GroupKey
		cur.ConfKey = req.Config.ConfKey
		cur.CreateTime = curTime
	} else {
		cur = pre
	}
	if req.Config.Description != "" {
		cur.Description = req.Config.Description
	}
	if req.Config.Status != 0 {
		cur.Status = req.Config.Status
	}
	if req.Config.PreValue != "" {
		cur.PreValue = req.Config.PreValue
	}
	if req.Config.Value != "" {
		cur.Value = req.Config.Value
	}
	cur.UpdateTime = curTime
	return cur, nil
}
