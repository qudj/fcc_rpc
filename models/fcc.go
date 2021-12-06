package models

import (
	"context"
	"fmt"
	"github.com/qudj/fcc_rpc/config"
)

type FccProject struct {
	Id          int64  `json:"id"`
	ProjectKey  string `json:"project_key"`
	ProjectName string `json:"project_name"`
	Description string `json:"description"`
	Status      int64  `json:"status"`
}

type FccGroup struct {
	Id          int64  `json:"id"`
	ProjectKey  string `json:"project_key"`
	GroupKey    string `json:"group_key"`
	Description string `json:"description"`
	Status      int64  `json:"status"`
}

type FccConf struct {
	Id          int64  `json:"id"`
	ProjectKey  string `json:"project_key"`
	GroupKey    string `json:"group_key"`
	ConfKey     string `json:"conf_key"`
	Description string `json:"description"`
	Status      int64  `json:"status"`
}

type FccHistoryLog struct {
	Id          int64  `json:"id"`
	Table       string `json:"table"`
	ObjectId    string `json:"object_id"`
	ObjectType  string `json:"object_type"`
	OpId        string `json:"op_id"`
	ChangeData  string `json:"change_data"`
	HistoryData string `json:"history_data"`
}

type FccMiniConf struct {
	Id    int64  `json:"id" form:"id"`
	Value string `json:"value" form:"value"`
}

func (FccProject) TableName() string {
	return "fcc_project"
}

func (FccGroup) TableName() string {
	return "fcc_project"
}

func (FccConf) TableName() string {
	return "fcc_project"
}

func (FccHistoryLog) TableName() string {
	return "fcc_project"
}

func (FccMiniConf) TableName() string {
	return "fcc_conf"
}

func GetProjects(ctx context.Context, filter map[string]interface{}, offset, limit int, orderBy string) ([]*FccProject, error) {
	var ret []*FccProject
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
	if err := config.FccReadDB.WithContext(ctx).Where(whereStr, whereArgs...).Order(orderBy).Offset(offset).Limit(limit).Find(&ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}

func GetGroups(ctx context.Context, proKey string, filter map[string]interface{}, offset, limit int, orderBy string) ([]*FccGroup, error) {
	var ret []*FccGroup
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
	if err := config.FccReadDB.WithContext(ctx).Where(whereStr, whereArgs...).Order(orderBy).Offset(offset).Limit(limit).Find(&ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}

func GetConfigs(ctx context.Context, proKey, grKey string, filter map[string]interface{}, offset, limit int, orderBy string) ([]*FccConf, error) {
	var ret []*FccConf
	whereStr := "project_key = ? and group_key = ?"
	whereArgs := []interface{}{proKey, grKey}
	if v, ok := filter["conf_key"]; ok {
		whereStr += " and conf_key = ?"
		whereArgs = append(whereArgs, v)
	}
	if err := config.FccReadDB.WithContext(ctx).Where(whereStr, whereArgs...).Order(orderBy).Offset(offset).Limit(limit).Find(&ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}

func GetMiniConfig(ctx context.Context, proKey, grKey, confKey string) (*FccConf, error) {
	var ret *FccConf
	if err := config.FccReadDB.WithContext(ctx).Where("project_key = ? and group_key = ? and conf_key", proKey, grKey, confKey).First(ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
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

func GetFccConf(project, group, key string) *FccMiniConf {
	conf := &FccMiniConf{}
	if err := config.FccReadDB.Where("project_key = ? and group_key = ? and conf_key = ?", project, group, key).First(conf).Error; err != nil {
		fmt.Println(err)
	}
	return conf
}
