package service

import (
	"context"

	pb "strike_layout_template/api/helloworld/v1"
	"strike_layout_template/internal/biz"
)

type GreeterService struct {
	pb.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

func (s *GreeterService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
