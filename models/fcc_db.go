package models

import (
	"context"
	"github.com/qudj/fcc_rpc/config"
)

type HistoryChange interface {
	TableName() string
}

type FccProject struct {
	Id          int64  `json:"id"`
	ProjectKey  string `json:"project_key"`
	ProjectName string `json:"project_name"`
	Description string `json:"description"`
	Status      int64  `json:"status"`
	UpdateTime  int64  `json:"update_time"`
	CreateTime  int64  `json:"create_time"`
}

type FccGroup struct {
	Id          int64  `json:"id"`
	ProjectKey  string `json:"project_key"`
	GroupKey    string `json:"group_key"`
	GroupName   string `json:"group_name"`
	Description string `json:"description"`
	Status      int64  `json:"status"`
	UpdateTime  int64  `json:"update_time"`
	CreateTime  int64  `json:"create_time"`
}

type FccConf struct {
	Id          int64  `json:"id"`
	ProjectKey  string `json:"project_key"`
	GroupKey    string `json:"group_key"`
	ConfKey     string `json:"conf_key"`
	Description string `json:"description"`
	Value       string `json:"value"`
	PreValue    string `json:"pre_value"`
	Status      int64  `json:"status"`
	UpdateTime  int64  `json:"update_time"`
	CreateTime  int64  `json:"create_time"`
}

type FccHistoryLog struct {
	Id          int64  `json:"id"`
	Table       string `json:"table"`
	ObjectKey   string `json:"object_key"`
	ObjectType  string `json:"object_type"`
	OpId        string `json:"op_id"`
	ChangeData  string `json:"change_data"`
	HistoryData string `json:"history_data"`
	CreateTime  int64  `json:"create_time"`
}

type FccMiniConf struct {
	Id    int64  `json:"id" form:"id"`
	Value string `json:"value" form:"value"`
}

func (FccProject) TableName() string {
	return "fcc_project"
}

func (FccGroup) TableName() string {
	return "fcc_group"
}

func (FccConf) TableName() string {
	return "fcc_conf"
}

func (FccHistoryLog) TableName() string {
	return "fcc_history_log"
}

func (FccMiniConf) TableName() string {
	return "fcc_conf"
}

func GetProjects(ctx context.Context, filter map[string]interface{}, offset, limit int, orderBy string) ([]*FccProject, int64, error) {
	var ret []*FccProject
	var count int64
	whereStr := "id > 0"
	whereArgs := make([]interface{}, 0)
	if v, ok := filter["project_key"]; ok {
		whereStr += " and project_key = ?"
		whereArgs = append(whereArgs, v)
	}
	if v, ok := filter["project_name"]; ok {
		whereStr += " and project_name = ?"
		whereArgs = append(whereArgs, v)
	}
	if err := config.FccReadDB.Table("fcc_project").WithContext(ctx).Where(whereStr, whereArgs...).Debug().Count(&count).
		Order(orderBy).Offset(offset).Limit(limit).Find(&ret).Error; err != nil {
		return nil, 0, err
	}
	return ret, count, nil
}

func GetGroups(ctx context.Context, proKey string, filter map[string]interface{}, offset, limit int, orderBy string) ([]*FccGroup, int64, error) {
	var ret []*FccGroup
	var count int64
	whereStr := "project_key = ?"
	whereArgs := []interface{}{proKey}
	if v, ok := filter["group_key"]; ok {
		whereStr += " and group_key = ?"
		whereArgs = append(whereArgs, v)
	}
	if v, ok := filter["group_name"]; ok {
		whereStr += " and group_name = ?"
		whereArgs = append(whereArgs, v)
	}
	if err := config.FccReadDB.Table("fcc_group").WithContext(ctx).Where(whereStr, whereArgs...).Debug().Count(&count).
		Order(orderBy).Offset(offset).Limit(limit).Find(&ret).Error; err != nil {
		return nil, 0, err
	}
	return ret, count, nil
}

func GetConfigs(ctx context.Context, proKey, grKey string, filter map[string]interface{}, offset, limit int, orderBy string) ([]*FccConf, int64, error) {
	var ret []*FccConf
	var count int64
	whereStr := "project_key = ? and group_key = ?"
	whereArgs := []interface{}{proKey, grKey}
	if v, ok := filter["conf_key"]; ok {
		whereStr += " and conf_key = ?"
		whereArgs = append(whereArgs, v)
	}
	if err := config.FccReadDB.Table("fcc_conf").WithContext(ctx).Where(whereStr, whereArgs...).Debug().Count(&count).
		Order(orderBy).Offset(offset).Limit(limit).Find(&ret).Error; err != nil {
		return nil, 0, err
	}
	return ret, count, nil
}

func SaveProject(project *FccProject) error {
	if err := config.FccReadDB.Save(project).Error; err != nil {
		return err
	}
	return nil
}

func SaveGroup(group *FccGroup) error {
	if err := config.FccReadDB.Save(group).Error; err != nil {
		return err
	}
	return nil
}

func SaveConf(conf *FccConf) error {
	if err := config.FccReadDB.Save(conf).Error; err != nil {
		return err
	}
	return nil
}

func SaveHistory(history *FccHistoryLog) error {
	if err := config.FccReadDB.Save(history).Error; err != nil {
		return err
	}
	return nil
}

func GetFccConf(ctx context.Context, project, group, key string) (*FccConf, error) {
	conf := &FccConf{}
	if err := config.FccReadDB.WithContext(ctx).Where("project_key = ? and group_key = ? and conf_key = ?", project, group, key).First(conf).Error; err != nil {
		return nil, err
	}
	return conf, nil
}
