package service

import (
	"context"

	"acd/dao"
	pb "acd/proto"
)

type AcdService struct {
	*pb.UnimplementedACDServiceServer
}

func (svc AcdService) AgentState(ctx context.Context, in *pb.RequestAgentState) (*pb.ResponseAgentState, error) {
	if len(in.AppId) == 0 || len(in.AgentId) == 0 {
		return nil, ErrInvalidParam
	}

	st, err := dao.GetAgentState(in.AppId, in.AgentId)
	if nil != err {
		return nil, ErrDBOperationFailed
	}

	rst := &pb.ResponseAgentState{
		AppId:   in.AppId,
		AgentId: in.AgentId,
		State:   st,
	}

	return rst, nil
}
