package main

import (
	"context"
	"github.com/qudj/fcc_rpc/models"
	"github.com/qudj/fcc_rpc/models/fcc_serv"
)

type FccService struct {
}

func NewFccServiceServer() fcc_serv.FccServiceServer {
	return FccService{}
}

func (f FccService) FetchProjects(ctx context.Context, req *fcc_serv.FetchProjectsRequest) (*fcc_serv.FetchProjectsResponse, error) {
	ret := &fcc_serv.FetchProjectsResponse{
		BaseRet: &fcc_serv.BaseRet{},
	}
	filter := make(map[string]interface{})
	if req.ProjectKey != "" {
		filter["project_key"] = req.ProjectKey
	}
	if req.ProjectName != "" {
		filter["project_name"] = req.ProjectName
	}
	res, err := models.GetProjects(ctx, filter, int(req.Offset), int(req.Limit), "id")
	if err != nil {
		ret.BaseRet.Code = 400
		ret.BaseRet.Msg = err.Error()
		return ret, nil
	}
	ret.Data = FormatProjectRet(res)
	return ret, nil
}

func FormatProjectRet(res []*models.FccProject) []*fcc_serv.Project {
	ret := make([]*fcc_serv.Project, 0, len(res))
	for _, v := range res {
		one := &fcc_serv.Project{
			ProjectKey: v.ProjectKey,
			ProjectName: v.ProjectName,
			Description: v.Description,
			Status: v.Status,
		}
		ret = append(ret, one)
	}
	return ret
}


func (f FccService) FetchGroups(ctx context.Context, req *fcc_serv.FetchGroupsRequest) (*fcc_serv.FetchGroupsResponse, error) {

	return &fcc_serv.FetchGroupsResponse{}, nil
}

func (f FccService) FetchConfigs(ctx context.Context, req *fcc_serv.FetchConfigsRequest) (*fcc_serv.FetchConfigsResponse, error) {

	return &fcc_serv.FetchConfigsResponse{}, nil
}

func (f FccService) FetchMiniConfig(ctx context.Context, req *fcc_serv.FetchMiniConfigRequest) (*fcc_serv.FetchMiniConfigResponse, error) {

	return &fcc_serv.FetchMiniConfigResponse{}, nil
}