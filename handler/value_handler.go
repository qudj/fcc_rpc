package handler

import (
	"context"
	"fmt"
	"github.com/qudj/fcc_rpc/config"
	"github.com/qudj/fcc_rpc/models"
	"github.com/qudj/fcc_rpc/models/fcc_serv"
	"github.com/qudj/fcc_rpc/service"
	"golang.org/x/sync/singleflight"
	"time"
)

var gsf singleflight.Group

func PrePublish(ctx context.Context, req *fcc_serv.PrePublishRequest) error {
	var pre, cur models.FccConf
	if err := config.FccWriteDB.WithContext(ctx).Where("project_key = ? and group_key = ? and conf_key = ?", req.ProjectKey, req.GroupKey, req.ConfKey).Last(&pre).Error; err != nil {
		return err
	}
	cur = pre
	cur.PreValue = req.PreValue
	cur.PublishStatus = config.UnPublishStatus
	cur.OpId = req.OpId
	cur.OpName = req.OpName
	cur.UpdateTime = time.Now().Unix()
	if err := models.SaveConf(&cur); err != nil {
		return err
	}
	_ = service.SaveHistory("", req.PreValue, cur.TableName(), cur.ProjectKey, "pre_publish", req.OpId)
	return nil
}

func Publish(ctx context.Context, req *fcc_serv.PublishRequest) error {
	var cur, pre models.FccConf
	if err := config.FccWriteDB.WithContext(ctx).Where("project_key = ? and group_key = ? and conf_key = ?", req.ProjectKey, req.GroupKey, req.ConfKey).Last(&pre).Error; err != nil {
		return err
	}
	cur = pre
	cur.PreValue = ""
	cur.Value = pre.PreValue
	cur.PublishStatus = config.PublishedStatus
	cur.OpId = req.OpId
	cur.OpName = req.OpName
	cur.UpdateTime = time.Now().Unix()
	if err := models.SaveConf(&cur); err != nil {
		return err
	}
	_ = service.SaveHistory(pre.Value, cur.Value, cur.TableName(), cur.ProjectKey, "publish", req.OpId)
	return nil
}

func FetchConfig(ctx context.Context, req *fcc_serv.FetchConfigRequest) (*fcc_serv.Config, error) {
	key := fmt.Sprintf("MC:%s_%s_%s", req.ProjectKey, req.GroupKey, req.ConfKey)
	gRes, err, _ := gsf.Do(key, func() (interface{}, error) {
		return models.GetFccConf(ctx, req.ProjectKey, req.GroupKey, req.ConfKey)
	})
	if err != nil {
		return nil, err
	}
	res := gRes.(*models.FccConf)
	return FormatConfigRet(res), nil
}

func FormatConfigRet(conf *models.FccConf) *fcc_serv.Config {
	ret := &fcc_serv.Config{
		ProjectKey:    conf.ProjectKey,
		GroupKey:      conf.GroupKey,
		ConfKey:       conf.ConfKey,
		Description:   conf.Description,
		Value:         conf.Value,
		PreValue:      conf.PreValue,
		Status:        conf.Status,
		PublishStatus: conf.PublishStatus,
		OpId:          conf.OpId,
		OpName:        conf.OpName,
		CreateTime:    conf.CreateTime,
		UpdateTime:    conf.UpdateTime,
	}
	return ret
}
