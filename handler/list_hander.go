package handler

import (
	"context"
	"github.com/qudj/fcc_rpc/models"
	"github.com/qudj/fcc_rpc/models/fcc_serv"
)

func FetchProjects(ctx context.Context, req *fcc_serv.FetchProjectsRequest) (*fcc_serv.FetchProjectsRet, error) {
	filter := make(map[string]interface{})
	if req.ProjectKey != "" {
		filter["project_key"] = req.ProjectKey
	}
	if req.ProjectName != "" {
		filter["project_name"] = req.ProjectName
	}
	list, count, err := models.GetProjects(ctx, filter, int(req.Offset), int(req.Limit), "id")
	if err != nil {
		return nil, err
	}
	ret := &fcc_serv.FetchProjectsRet{
		Total: count,
		List:  FormatProjectRetList(list),
	}
	return ret, nil
}

func FormatProjectRetList(res []*models.FccProject) []*fcc_serv.Project {
	ret := make([]*fcc_serv.Project, 0, len(res))
	for _, v := range res {
		one := &fcc_serv.Project{
			ProjectKey:  v.ProjectKey,
			ProjectName: v.ProjectName,
			Description: v.Description,
			Status:      v.Status,
		}
		ret = append(ret, one)
	}
	return ret
}

func FetchGroups(ctx context.Context, req *fcc_serv.FetchGroupsRequest) (*fcc_serv.FetchGroupsRet, error) {
	filter := make(map[string]interface{})
	if req.GroupKey != "" {
		filter["group_key"] = req.GroupKey
	}
	if req.GroupName != "" {
		filter["group_name"] = req.GroupName
	}
	list, count, err := models.GetGroups(ctx, req.ProjectKey, filter, int(req.Offset), int(req.Limit), "id")
	if err != nil {
		return nil, err
	}
	ret := &fcc_serv.FetchGroupsRet{
		Total: count,
		List:  FormatGroupRetList(list),
	}
	return ret, nil
}

func FormatGroupRetList(res []*models.FccGroup) []*fcc_serv.Group {
	ret := make([]*fcc_serv.Group, 0, len(res))
	for _, v := range res {
		one := &fcc_serv.Group{
			ProjectKey:  v.ProjectKey,
			GroupKey:    v.GroupKey,
			GroupName:   v.GroupName,
			Description: v.Description,
			Status:      v.Status,
		}
		ret = append(ret, one)
	}
	return ret
}

func FetchConfigs(ctx context.Context, req *fcc_serv.FetchConfigsRequest) (*fcc_serv.FetchConfigsRet, error) {
	filter := make(map[string]interface{})
	if req.ConfKey != "" {
		filter["conf_key"] = req.ConfKey
	}
	list, count, err := models.GetConfigs(ctx, req.ProjectKey, req.GroupKey, filter, int(req.Offset), int(req.Limit), "id")
	if err != nil {
		return nil, err
	}
	ret := &fcc_serv.FetchConfigsRet{
		Total: count,
		List:  FormatConfigRetList(list),
	}
	return ret, nil
}

func FormatConfigRetList(res []*models.FccConf) []*fcc_serv.Config {
	ret := make([]*fcc_serv.Config, 0, len(res))
	for _, v := range res {
		one := &fcc_serv.Config{
			ProjectKey:  v.ProjectKey,
			GroupKey:    v.GroupKey,
			ConfKey:     v.ConfKey,
			Description: v.Description,
			Value:       v.Value,
			PreValue:    v.Value,
			Status:      v.Status,
		}
		ret = append(ret, one)
	}
	return ret
}
