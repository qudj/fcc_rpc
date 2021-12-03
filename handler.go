package main

import (
	"context"
	srvpb "github.com/qudj/fly_lib/models/fly_conf"
)

type FccService struct {
}

func NewFccServiceServer() srvpb.FccServiceServer {
	return FccService{}
}

func (f FccService) FetchProjects(ctx context.Context, req *srvpb.FetchProjectsRequest) (*srvpb.FetchProjectsResponse, error) {

	return &srvpb.FetchProjectsResponse{}, nil
}

func (f FccService) FetchGroups(ctx context.Context, req *srvpb.FetchGroupsRequest) (*srvpb.FetchGroupsResponse, error) {

	return &srvpb.FetchGroupsResponse{}, nil
}

func (f FccService) FetchConfigs(ctx context.Context, req *srvpb.FetchConfigsRequest) (*srvpb.FetchConfigsResponse, error) {

	return &srvpb.FetchConfigsResponse{}, nil
}

func (f FccService) FetchMiniConfig(ctx context.Context, req *srvpb.FetchMiniConfigRequest) (*srvpb.FetchMiniConfigResponse, error) {

	return &srvpb.FetchMiniConfigResponse{}, nil
}