package handler

import (
	"context"
	"fmt"
	"github.com/qudj/fcc_rpc/config"
	"github.com/qudj/fcc_rpc/models"
	"github.com/qudj/fcc_rpc/models/fcc_serv"
	"github.com/qudj/fcc_rpc/service"
	"golang.org/x/sync/singleflight"
)

var gsf singleflight.Group

func PrePublish(ctx context.Context, req *fcc_serv.PrePublishRequest) error {
	var pre, cur models.FccConf
	if err := config.FccWriteDB.WithContext(ctx).Where("project_key = ? and group_key = ? and conf_key = ?", req.ProjectKey, req.GroupKey, req.ConfKey).Last(&pre).Error; err != nil {
		return err
	}
	cur = pre
	cur.PreValue = req.PreValue
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
	if err := models.SaveConf(&cur); err != nil {
		return err
	}
	_ = service.SaveHistory(pre, cur, cur.TableName(), cur.ProjectKey, "publish", req.OpId)
	return nil
}

func FetchMiniConfig(ctx context.Context, req *fcc_serv.FetchMiniConfigRequest) (string, error) {
	key := fmt.Sprintf("MC:%s_%s_%s", req.ProjectKey, req.GroupKey, req.ConfKey)
	res, err, _ := gsf.Do(key, func() (interface{}, error) {
		return models.GetMiniConfig(ctx, req.ProjectKey, req.GroupKey, req.ConfKey)
	})
	if err != nil {
		return "", err
	}
	return res.(*models.FccMiniConf).Value, nil
}
