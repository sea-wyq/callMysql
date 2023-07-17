package service

import (
	"context"
	"strconv"

	pb "callMysql/api"
	"callMysql/internal/biz"
)

type CallMysqlService struct {
	pb.UnimplementedGreeterServer

	uc *biz.DataControllerUsecase
}

func NewCallMysqlService(uc *biz.DataControllerUsecase) *CallMysqlService {
	return &CallMysqlService{uc: uc}
}

func (s *CallMysqlService) GetMysqlMess(ctx context.Context, req *pb.CallRequest) (*pb.CallReply, error) {
	id, _ := strconv.ParseInt(req.Id, 10, 64)
	g, err := s.uc.CreateMess(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.CallReply{Message: g.Info}, err
}
