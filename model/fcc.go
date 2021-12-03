package model

import (
	"fmt"
	"github.com/qudj/fcc_rpc/config"
)

type FccProject struct {
	Id          int64  `json:"id"`
	ProjectKey  string `json:"project_key""`
	Description string `json:"description"`
	Status      int64  `json:"status"`
}

type FccGroup struct {
	Id          int64  `json:"id"`
	ProjectKey  string `json:"project_key""`
	GroupKey    string `json:"group_key""`
	Description string `json:"description"`
	Status      int64  `json:"status"`
}

type FccConf struct {
	Id          int64  `json:"id"`
	ProjectKey  string `json:"project_key""`
	GroupKey    string `json:"group_key""`
	ConfKey     string `json:"conf_key""`
	Description string `json:"description"`
	Status      int64  `json:"status"`
}

type FccHistoryLog struct {
	Id          int64  `json:"id"`
	Table       string `json:"table""`
	ObjectId    string `json:"object_id""`
	ObjectType  string `json:"object_type""`
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

func GetProject() *FccProject {
	project := &FccProject{}
	if err := config.FccReadDB.Where("id = 1").First(project).Error; err != nil {
		fmt.Println(err)
	}
	return project
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