package main

import (
	"context"
	"github.com/qudj/fcc_rpc/handler"
	"github.com/qudj/fcc_rpc/models/fcc_serv"
)

type FccService struct{}

func NewFccServiceServer() fcc_serv.FccServiceServer {
	return FccService{}
}

func (f FccService) FetchProjects(ctx context.Context, req *fcc_serv.FetchProjectsRequest) (*fcc_serv.FetchProjectsResponse, error) {
	ret := &fcc_serv.FetchProjectsResponse{
		BaseRet: &fcc_serv.BaseRet{},
	}
	data, err := handler.FetchProjects(ctx, req)
	if err != nil {
		ret.BaseRet.Code = 400
		ret.BaseRet.Msg = err.Error()
		return ret, nil
	}
	ret.Data = data
	return ret, nil
}

func (f FccService) FetchGroups(ctx context.Context, req *fcc_serv.FetchGroupsRequest) (*fcc_serv.FetchGroupsResponse, error) {
	ret := &fcc_serv.FetchGroupsResponse{
		BaseRet: &fcc_serv.BaseRet{},
		Data:    &fcc_serv.FetchGroupsRet{},
	}
	data, err := handler.FetchGroups(ctx, req)
	if err != nil {
		ret.BaseRet.Code = 400
		ret.BaseRet.Msg = err.Error()
		return ret, nil
	}
	ret.Data = data
	return ret, nil
}

func (f FccService) FetchConfigs(ctx context.Context, req *fcc_serv.FetchConfigsRequest) (*fcc_serv.FetchConfigsResponse, error) {
	ret := &fcc_serv.FetchConfigsResponse{
		BaseRet: &fcc_serv.BaseRet{},
		Data:    &fcc_serv.FetchConfigsRet{},
	}
	data, err := handler.FetchConfigs(ctx, req)
	if err != nil {
		ret.BaseRet.Code = 400
		ret.BaseRet.Msg = err.Error()
		return ret, nil
	}
	ret.Data = data
	return ret, nil
}

func (f FccService) SaveProject(ctx context.Context, req *fcc_serv.SaveProjectRequest) (*fcc_serv.SaveProjectResponse, error) {
	ret := &fcc_serv.SaveProjectResponse{
		BaseRet: &fcc_serv.BaseRet{},
	}
	err := handler.SaveProject(ctx, req)
	if err != nil {
		ret.BaseRet.Code = 400
		ret.BaseRet.Msg = err.Error()
		return ret, nil
	}
	return ret, nil
}

func (f FccService) SaveGroup(ctx context.Context, req *fcc_serv.SaveGroupRequest) (*fcc_serv.SaveGroupResponse, error) {
	ret := &fcc_serv.SaveGroupResponse{
		BaseRet: &fcc_serv.BaseRet{},
	}
	err := handler.SaveGroup(ctx, req)
	if err != nil {
		ret.BaseRet.Code = 400
		ret.BaseRet.Msg = err.Error()
		return ret, nil
	}
	return ret, nil
}

func (f FccService) SaveConfig(ctx context.Context, req *fcc_serv.SaveConfigRequest) (*fcc_serv.SaveConfigResponse, error) {
	ret := &fcc_serv.SaveConfigResponse{
		BaseRet: &fcc_serv.BaseRet{},
	}
	err := handler.SaveConfig(ctx, req)
	if err != nil {
		ret.BaseRet.Code = 400
		ret.BaseRet.Msg = err.Error()
		return ret, nil
	}
	return ret, nil
}

func (f FccService) PrePublish(ctx context.Context, req *fcc_serv.PrePublishRequest) (*fcc_serv.PrePublishResponse, error) {
	ret := &fcc_serv.PrePublishResponse{
		BaseRet: &fcc_serv.BaseRet{},
	}
	err := handler.PrePublish(ctx, req)
	if err != nil {
		ret.BaseRet.Code = 400
		ret.BaseRet.Msg = err.Error()
		return ret, nil
	}
	return ret, nil
}

func (f FccService) Publish(ctx context.Context, req *fcc_serv.PublishRequest) (*fcc_serv.PublishResponse, error) {
	ret := &fcc_serv.PublishResponse{
		BaseRet: &fcc_serv.BaseRet{},
	}
	err := handler.Publish(ctx, req)
	if err != nil {
		ret.BaseRet.Code = 400
		ret.BaseRet.Msg = err.Error()
		return ret, nil
	}
	return ret, nil
}

func (f FccService) FetchConfig(ctx context.Context, req *fcc_serv.FetchConfigRequest) (*fcc_serv.FetchConfigResponse, error) {
	ret := &fcc_serv.FetchConfigResponse{
		BaseRet: &fcc_serv.BaseRet{},
		Data:    &fcc_serv.Config{},
	}
	value, err := handler.FetchConfig(ctx, req)
	if err != nil {
		ret.BaseRet.Code = 400
		ret.BaseRet.Msg = err.Error()
		return ret, nil
	}
	ret.Data = value
	return ret, nil
}
