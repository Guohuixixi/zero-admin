// Code generated by goctl. DO NOT EDIT.
// Source: cms.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/logic/topiccommentservice"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
)

type TopicCommentServiceServer struct {
	svcCtx *svc.ServiceContext
	cmsclient.UnimplementedTopicCommentServiceServer
}

func NewTopicCommentServiceServer(svcCtx *svc.ServiceContext) *TopicCommentServiceServer {
	return &TopicCommentServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加专题评论表
func (s *TopicCommentServiceServer) AddTopicComment(ctx context.Context, in *cmsclient.AddTopicCommentReq) (*cmsclient.AddTopicCommentResp, error) {
	l := topiccommentservicelogic.NewAddTopicCommentLogic(ctx, s.svcCtx)
	return l.AddTopicComment(in)
}

// 删除专题评论表
func (s *TopicCommentServiceServer) DeleteTopicComment(ctx context.Context, in *cmsclient.DeleteTopicCommentReq) (*cmsclient.DeleteTopicCommentResp, error) {
	l := topiccommentservicelogic.NewDeleteTopicCommentLogic(ctx, s.svcCtx)
	return l.DeleteTopicComment(in)
}

// 更新专题评论表
func (s *TopicCommentServiceServer) UpdateTopicComment(ctx context.Context, in *cmsclient.UpdateTopicCommentReq) (*cmsclient.UpdateTopicCommentResp, error) {
	l := topiccommentservicelogic.NewUpdateTopicCommentLogic(ctx, s.svcCtx)
	return l.UpdateTopicComment(in)
}

// 更新专题评论表状态
func (s *TopicCommentServiceServer) UpdateTopicCommentStatus(ctx context.Context, in *cmsclient.UpdateTopicCommentStatusReq) (*cmsclient.UpdateTopicCommentStatusResp, error) {
	l := topiccommentservicelogic.NewUpdateTopicCommentStatusLogic(ctx, s.svcCtx)
	return l.UpdateTopicCommentStatus(in)
}

// 查询专题评论表详情
func (s *TopicCommentServiceServer) QueryTopicCommentDetail(ctx context.Context, in *cmsclient.QueryTopicCommentDetailReq) (*cmsclient.QueryTopicCommentDetailResp, error) {
	l := topiccommentservicelogic.NewQueryTopicCommentDetailLogic(ctx, s.svcCtx)
	return l.QueryTopicCommentDetail(in)
}

// 查询专题评论表列表
func (s *TopicCommentServiceServer) QueryTopicCommentList(ctx context.Context, in *cmsclient.QueryTopicCommentListReq) (*cmsclient.QueryTopicCommentListResp, error) {
	l := topiccommentservicelogic.NewQueryTopicCommentListLogic(ctx, s.svcCtx)
	return l.QueryTopicCommentList(in)
}