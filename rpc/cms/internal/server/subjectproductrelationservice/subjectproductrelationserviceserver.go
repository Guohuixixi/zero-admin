// Code generated by goctl. DO NOT EDIT.
// Source: cms.proto

package server

import (
	"context"

	"zero-admin/rpc/cms/cmsclient"
	"zero-admin/rpc/cms/internal/logic/subjectproductrelationservice"
	"zero-admin/rpc/cms/internal/svc"
)

type SubjectProductRelationServiceServer struct {
	svcCtx *svc.ServiceContext
	cmsclient.UnimplementedSubjectProductRelationServiceServer
}

func NewSubjectProductRelationServiceServer(svcCtx *svc.ServiceContext) *SubjectProductRelationServiceServer {
	return &SubjectProductRelationServiceServer{
		svcCtx: svcCtx,
	}
}

// 专题关联
func (s *SubjectProductRelationServiceServer) SubjectProductRelationAdd(ctx context.Context, in *cmsclient.SubjectProductRelationAddReq) (*cmsclient.SubjectProductRelationAddResp, error) {
	l := subjectproductrelationservicelogic.NewSubjectProductRelationAddLogic(ctx, s.svcCtx)
	return l.SubjectProductRelationAdd(in)
}

func (s *SubjectProductRelationServiceServer) SubjectProductRelationDelete(ctx context.Context, in *cmsclient.SubjectProductRelationDeleteReq) (*cmsclient.SubjectProductRelationDeleteResp, error) {
	l := subjectproductrelationservicelogic.NewSubjectProductRelationDeleteLogic(ctx, s.svcCtx)
	return l.SubjectProductRelationDelete(in)
}

func (s *SubjectProductRelationServiceServer) SubjectProductRelationUpdate(ctx context.Context, in *cmsclient.SubjectProductRelationUpdateReq) (*cmsclient.SubjectProductRelationUpdateResp, error) {
	l := subjectproductrelationservicelogic.NewSubjectProductRelationUpdateLogic(ctx, s.svcCtx)
	return l.SubjectProductRelationUpdate(in)
}

func (s *SubjectProductRelationServiceServer) SubjectProductRelationList(ctx context.Context, in *cmsclient.SubjectProductRelationListReq) (*cmsclient.SubjectProductRelationListResp, error) {
	l := subjectproductrelationservicelogic.NewSubjectProductRelationListLogic(ctx, s.svcCtx)
	return l.SubjectProductRelationList(in)
}
