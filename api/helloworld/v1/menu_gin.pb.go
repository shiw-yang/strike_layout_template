// Code generated by github.com/shiw-yang/strike/cmd/protoc-gen-gin. DO NOT EDIT.
// versions:
// - protoc-gen-gin v1.0.0
// - protoc         v3.20.0
// source: helloworld/v1/menu.proto
package v1

import (
	context "context"
	errors "errors"
	gin "github.com/gin-gonic/gin"
	metadata "google.golang.org/grpc/metadata"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the github.com/shiw-yang/strike/cmd/protoc-gen-gin package it is being compiled against.
// context.
// metadata.
// gin.errors.

type MenuHTTPServer interface {
	CreateMenu(context.Context, *CreateMenuRequest) (*CreateMenuReply, error)

	DeleteMenu(context.Context, *DeleteMenuRequest) (*DeleteMenuReply, error)

	GetMenu(context.Context, *GetMenuRequest) (*GetMenuReply, error)

	ListMenu(context.Context, *ListMenuRequest) (*ListMenuReply, error)

	UpdateMenu(context.Context, *UpdateMenuRequest) (*UpdateMenuReply, error)
}

func RegisterMenuHTTPServer(r gin.IRouter, srv MenuHTTPServer) {
	s := Menu{
		server: srv,
		router: r,
		resp:   defaultMenuResp{},
	}
	s.RegisterService()
}

type Menu struct {
	server MenuHTTPServer
	router gin.IRouter
	resp   interface {
		Error(ctx *gin.Context, err error)
		ParamsError(ctx *gin.Context, err error)
		Success(ctx *gin.Context, data interface{})
	}
}

// Resp 返回值
type defaultMenuResp struct{}

func (resp defaultMenuResp) response(ctx *gin.Context, status, code int, msg string, data interface{}) {
	ctx.JSON(status, map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// Error 返回错误信息
func (resp defaultMenuResp) Error(ctx *gin.Context, err error) {
	code := -1
	status := 500
	msg := err.Error()

	if err == nil {
		msg += ", err is nil"
		resp.response(ctx, status, code, msg, nil)
		return
	}

	type iCode interface {
		HTTPCode() int
		Message() string
		Code() int
	}

	var c iCode
	if errors.As(err, &c) {
		status = c.HTTPCode()
		code = c.Code()
		msg = c.Message()
	}

	_ = ctx.Error(err)

	resp.response(ctx, status, code, msg, nil)
}

// ParamsError 参数错误
func (resp defaultMenuResp) ParamsError(ctx *gin.Context, err error) {
	_ = ctx.Error(err)
	resp.response(ctx, 400, 400, err.Error(), nil)
}

// Success 返回成功信息
func (resp defaultMenuResp) Success(ctx *gin.Context, data interface{}) {
	resp.response(ctx, 200, 0, "success", data)
}

func (s *Menu) CreateMenu_0(ctx *gin.Context) {
	var in CreateMenuRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		s.resp.ParamsError(ctx, err)
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx.Request.Context(), md)
	out, err := s.server.(MenuHTTPServer).CreateMenu(newCtx, &in)
	if err != nil {
		s.resp.Error(ctx, err)
		return
	}

	s.resp.Success(ctx, out)
}

func (s *Menu) UpdateMenu_0(ctx *gin.Context) {
	var in UpdateMenuRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		s.resp.ParamsError(ctx, err)
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx.Request.Context(), md)
	out, err := s.server.(MenuHTTPServer).UpdateMenu(newCtx, &in)
	if err != nil {
		s.resp.Error(ctx, err)
		return
	}

	s.resp.Success(ctx, out)
}

func (s *Menu) DeleteMenu_0(ctx *gin.Context) {
	var in DeleteMenuRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		s.resp.ParamsError(ctx, err)
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx.Request.Context(), md)
	out, err := s.server.(MenuHTTPServer).DeleteMenu(newCtx, &in)
	if err != nil {
		s.resp.Error(ctx, err)
		return
	}

	s.resp.Success(ctx, out)
}

func (s *Menu) GetMenu_0(ctx *gin.Context) {
	var in GetMenuRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		s.resp.ParamsError(ctx, err)
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx.Request.Context(), md)
	out, err := s.server.(MenuHTTPServer).GetMenu(newCtx, &in)
	if err != nil {
		s.resp.Error(ctx, err)
		return
	}

	s.resp.Success(ctx, out)
}

func (s *Menu) ListMenu_0(ctx *gin.Context) {
	var in ListMenuRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		s.resp.ParamsError(ctx, err)
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx.Request.Context(), md)
	out, err := s.server.(MenuHTTPServer).ListMenu(newCtx, &in)
	if err != nil {
		s.resp.Error(ctx, err)
		return
	}

	s.resp.Success(ctx, out)
}

func (s *Menu) RegisterService() {

	s.router.Handle("POST", "", s.CreateMenu_0)

	s.router.Handle("POST", "", s.UpdateMenu_0)

	s.router.Handle("POST", "", s.DeleteMenu_0)

	s.router.Handle("POST", "", s.GetMenu_0)

	s.router.Handle("POST", "", s.ListMenu_0)

}
