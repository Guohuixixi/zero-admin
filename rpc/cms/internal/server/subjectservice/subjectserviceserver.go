// Code generated by goctl. DO NOT EDIT.
// Source: cms.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/logic/subjectservice"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
)

type SubjectServiceServer struct {
	svcCtx *svc.ServiceContext
	cmsclient.UnimplementedSubjectServiceServer
}

func NewSubjectServiceServer(svcCtx *svc.ServiceContext) *SubjectServiceServer {
	return &SubjectServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加专题表
func (s *SubjectServiceServer) AddSubject(ctx context.Context, in *cmsclient.AddSubjectReq) (*cmsclient.AddSubjectResp, error) {
	l := subjectservicelogic.NewAddSubjectLogic(ctx, s.svcCtx)
	return l.AddSubject(in)
}

// 删除专题表
func (s *SubjectServiceServer) DeleteSubject(ctx context.Context, in *cmsclient.DeleteSubjectReq) (*cmsclient.DeleteSubjectResp, error) {
	l := subjectservicelogic.NewDeleteSubjectLogic(ctx, s.svcCtx)
	return l.DeleteSubject(in)
}

// 更新专题表
func (s *SubjectServiceServer) UpdateSubject(ctx context.Context, in *cmsclient.UpdateSubjectReq) (*cmsclient.UpdateSubjectResp, error) {
	l := subjectservicelogic.NewUpdateSubjectLogic(ctx, s.svcCtx)
	return l.UpdateSubject(in)
}

// 更新专题表状态
func (s *SubjectServiceServer) UpdateSubjectStatus(ctx context.Context, in *cmsclient.UpdateSubjectStatusReq) (*cmsclient.UpdateSubjectStatusResp, error) {
	l := subjectservicelogic.NewUpdateSubjectStatusLogic(ctx, s.svcCtx)
	return l.UpdateSubjectStatus(in)
}

// 查询专题表详情
func (s *SubjectServiceServer) QuerySubjectDetail(ctx context.Context, in *cmsclient.QuerySubjectDetailReq) (*cmsclient.QuerySubjectDetailResp, error) {
	l := subjectservicelogic.NewQuerySubjectDetailLogic(ctx, s.svcCtx)
	return l.QuerySubjectDetail(in)
}

// 查询专题表列表
func (s *SubjectServiceServer) QuerySubjectList(ctx context.Context, in *cmsclient.QuerySubjectListReq) (*cmsclient.QuerySubjectListResp, error) {
	l := subjectservicelogic.NewQuerySubjectListLogic(ctx, s.svcCtx)
	return l.QuerySubjectList(in)
}

func (s *SubjectServiceServer) SubjectListByIds(ctx context.Context, in *cmsclient.SubjectListByIdsReq) (*cmsclient.QuerySubjectListResp, error) {
	l := subjectservicelogic.NewSubjectListByIdsLogic(ctx, s.svcCtx)
	return l.SubjectListByIds(in)
}

// 批量更新状态
func (s *SubjectServiceServer) UpdateSubjectRecommendStatus(ctx context.Context, in *cmsclient.UpdateSubjectRecommendStatusReq) (*cmsclient.UpdateSubjectRecommendStatusResp, error) {
	l := subjectservicelogic.NewUpdateSubjectRecommendStatusLogic(ctx, s.svcCtx)
	return l.UpdateSubjectRecommendStatus(in)
}
