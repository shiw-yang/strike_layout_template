package service

import (
	"context"

	pb "strike_layout_template/api/helloworld/v1"
)

type MenuService struct {
	pb.UnimplementedMenuServer
}

func NewMenuService() *MenuService {
	return &MenuService{}
}

func (s *MenuService) CreateMenu(ctx context.Context, req *pb.CreateMenuRequest) (*pb.CreateMenuReply, error) {
	return &pb.CreateMenuReply{}, nil
}
func (s *MenuService) UpdateMenu(ctx context.Context, req *pb.UpdateMenuRequest) (*pb.UpdateMenuReply, error) {
	return &pb.UpdateMenuReply{}, nil
}
func (s *MenuService) DeleteMenu(ctx context.Context, req *pb.DeleteMenuRequest) (*pb.DeleteMenuReply, error) {
	return &pb.DeleteMenuReply{}, nil
}
func (s *MenuService) GetMenu(ctx context.Context, req *pb.GetMenuRequest) (*pb.GetMenuReply, error) {
	return &pb.GetMenuReply{}, nil
}
func (s *MenuService) ListMenu(ctx context.Context, req *pb.ListMenuRequest) (*pb.ListMenuReply, error) {
	return &pb.ListMenuReply{}, nil
}