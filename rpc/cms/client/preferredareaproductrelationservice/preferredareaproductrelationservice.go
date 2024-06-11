// Code generated by goctl. DO NOT EDIT.
// Source: cms.proto

package preferredareaproductrelationservice

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddHelpCategoryReq                        = cmsclient.AddHelpCategoryReq
	AddHelpCategoryResp                       = cmsclient.AddHelpCategoryResp
	AddHelpReq                                = cmsclient.AddHelpReq
	AddHelpResp                               = cmsclient.AddHelpResp
	AddMemberReportReq                        = cmsclient.AddMemberReportReq
	AddMemberReportResp                       = cmsclient.AddMemberReportResp
	AddPreferredAreaProductRelationReq        = cmsclient.AddPreferredAreaProductRelationReq
	AddPreferredAreaProductRelationResp       = cmsclient.AddPreferredAreaProductRelationResp
	AddPreferredAreaReq                       = cmsclient.AddPreferredAreaReq
	AddPreferredAreaResp                      = cmsclient.AddPreferredAreaResp
	AddSubjectCategoryReq                     = cmsclient.AddSubjectCategoryReq
	AddSubjectCategoryResp                    = cmsclient.AddSubjectCategoryResp
	AddSubjectCommentReq                      = cmsclient.AddSubjectCommentReq
	AddSubjectCommentResp                     = cmsclient.AddSubjectCommentResp
	AddSubjectProductRelationReq              = cmsclient.AddSubjectProductRelationReq
	AddSubjectProductRelationResp             = cmsclient.AddSubjectProductRelationResp
	AddSubjectReq                             = cmsclient.AddSubjectReq
	AddSubjectResp                            = cmsclient.AddSubjectResp
	AddTopicCategoryReq                       = cmsclient.AddTopicCategoryReq
	AddTopicCategoryResp                      = cmsclient.AddTopicCategoryResp
	AddTopicCommentReq                        = cmsclient.AddTopicCommentReq
	AddTopicCommentResp                       = cmsclient.AddTopicCommentResp
	AddTopicReq                               = cmsclient.AddTopicReq
	AddTopicResp                              = cmsclient.AddTopicResp
	DeleteHelpCategoryReq                     = cmsclient.DeleteHelpCategoryReq
	DeleteHelpCategoryResp                    = cmsclient.DeleteHelpCategoryResp
	DeleteHelpReq                             = cmsclient.DeleteHelpReq
	DeleteHelpResp                            = cmsclient.DeleteHelpResp
	DeleteMemberReportReq                     = cmsclient.DeleteMemberReportReq
	DeleteMemberReportResp                    = cmsclient.DeleteMemberReportResp
	DeletePreferredAreaReq                    = cmsclient.DeletePreferredAreaReq
	DeletePreferredAreaResp                   = cmsclient.DeletePreferredAreaResp
	DeleteSubjectCategoryReq                  = cmsclient.DeleteSubjectCategoryReq
	DeleteSubjectCategoryResp                 = cmsclient.DeleteSubjectCategoryResp
	DeleteSubjectCommentReq                   = cmsclient.DeleteSubjectCommentReq
	DeleteSubjectCommentResp                  = cmsclient.DeleteSubjectCommentResp
	DeleteSubjectReq                          = cmsclient.DeleteSubjectReq
	DeleteSubjectResp                         = cmsclient.DeleteSubjectResp
	DeleteTopicCategoryReq                    = cmsclient.DeleteTopicCategoryReq
	DeleteTopicCategoryResp                   = cmsclient.DeleteTopicCategoryResp
	DeleteTopicCommentReq                     = cmsclient.DeleteTopicCommentReq
	DeleteTopicCommentResp                    = cmsclient.DeleteTopicCommentResp
	DeleteTopicReq                            = cmsclient.DeleteTopicReq
	DeleteTopicResp                           = cmsclient.DeleteTopicResp
	HelpCategoryListData                      = cmsclient.HelpCategoryListData
	HelpListData                              = cmsclient.HelpListData
	MemberReportListData                      = cmsclient.MemberReportListData
	PreferredAreaListData                     = cmsclient.PreferredAreaListData
	QueryHelpCategoryDetailReq                = cmsclient.QueryHelpCategoryDetailReq
	QueryHelpCategoryDetailResp               = cmsclient.QueryHelpCategoryDetailResp
	QueryHelpCategoryListReq                  = cmsclient.QueryHelpCategoryListReq
	QueryHelpCategoryListResp                 = cmsclient.QueryHelpCategoryListResp
	QueryHelpDetailReq                        = cmsclient.QueryHelpDetailReq
	QueryHelpDetailResp                       = cmsclient.QueryHelpDetailResp
	QueryHelpListReq                          = cmsclient.QueryHelpListReq
	QueryHelpListResp                         = cmsclient.QueryHelpListResp
	QueryMemberReportDetailReq                = cmsclient.QueryMemberReportDetailReq
	QueryMemberReportDetailResp               = cmsclient.QueryMemberReportDetailResp
	QueryMemberReportListReq                  = cmsclient.QueryMemberReportListReq
	QueryMemberReportListResp                 = cmsclient.QueryMemberReportListResp
	QueryPreferredAreaDetailReq               = cmsclient.QueryPreferredAreaDetailReq
	QueryPreferredAreaDetailResp              = cmsclient.QueryPreferredAreaDetailResp
	QueryPreferredAreaListReq                 = cmsclient.QueryPreferredAreaListReq
	QueryPreferredAreaListResp                = cmsclient.QueryPreferredAreaListResp
	QueryPreferredAreaProductRelationListReq  = cmsclient.QueryPreferredAreaProductRelationListReq
	QueryPreferredAreaProductRelationListResp = cmsclient.QueryPreferredAreaProductRelationListResp
	QuerySubjectCategoryDetailReq             = cmsclient.QuerySubjectCategoryDetailReq
	QuerySubjectCategoryDetailResp            = cmsclient.QuerySubjectCategoryDetailResp
	QuerySubjectCategoryListReq               = cmsclient.QuerySubjectCategoryListReq
	QuerySubjectCategoryListResp              = cmsclient.QuerySubjectCategoryListResp
	QuerySubjectCommentDetailReq              = cmsclient.QuerySubjectCommentDetailReq
	QuerySubjectCommentDetailResp             = cmsclient.QuerySubjectCommentDetailResp
	QuerySubjectCommentListReq                = cmsclient.QuerySubjectCommentListReq
	QuerySubjectCommentListResp               = cmsclient.QuerySubjectCommentListResp
	QuerySubjectDetailReq                     = cmsclient.QuerySubjectDetailReq
	QuerySubjectDetailResp                    = cmsclient.QuerySubjectDetailResp
	QuerySubjectListReq                       = cmsclient.QuerySubjectListReq
	QuerySubjectListResp                      = cmsclient.QuerySubjectListResp
	QuerySubjectProductRelationListReq        = cmsclient.QuerySubjectProductRelationListReq
	QuerySubjectProductRelationListResp       = cmsclient.QuerySubjectProductRelationListResp
	QueryTopicCategoryDetailReq               = cmsclient.QueryTopicCategoryDetailReq
	QueryTopicCategoryDetailResp              = cmsclient.QueryTopicCategoryDetailResp
	QueryTopicCategoryListReq                 = cmsclient.QueryTopicCategoryListReq
	QueryTopicCategoryListResp                = cmsclient.QueryTopicCategoryListResp
	QueryTopicCommentDetailReq                = cmsclient.QueryTopicCommentDetailReq
	QueryTopicCommentDetailResp               = cmsclient.QueryTopicCommentDetailResp
	QueryTopicCommentListReq                  = cmsclient.QueryTopicCommentListReq
	QueryTopicCommentListResp                 = cmsclient.QueryTopicCommentListResp
	QueryTopicDetailReq                       = cmsclient.QueryTopicDetailReq
	QueryTopicDetailResp                      = cmsclient.QueryTopicDetailResp
	QueryTopicListReq                         = cmsclient.QueryTopicListReq
	QueryTopicListResp                        = cmsclient.QueryTopicListResp
	SubjectCategoryListData                   = cmsclient.SubjectCategoryListData
	SubjectCommentListData                    = cmsclient.SubjectCommentListData
	SubjectListByIdsReq                       = cmsclient.SubjectListByIdsReq
	SubjectListData                           = cmsclient.SubjectListData
	TopicCategoryListData                     = cmsclient.TopicCategoryListData
	TopicCommentListData                      = cmsclient.TopicCommentListData
	TopicListData                             = cmsclient.TopicListData
	UpdateHelpCategoryReq                     = cmsclient.UpdateHelpCategoryReq
	UpdateHelpCategoryResp                    = cmsclient.UpdateHelpCategoryResp
	UpdateHelpCategoryStatusReq               = cmsclient.UpdateHelpCategoryStatusReq
	UpdateHelpCategoryStatusResp              = cmsclient.UpdateHelpCategoryStatusResp
	UpdateHelpReq                             = cmsclient.UpdateHelpReq
	UpdateHelpResp                            = cmsclient.UpdateHelpResp
	UpdateHelpStatusReq                       = cmsclient.UpdateHelpStatusReq
	UpdateHelpStatusResp                      = cmsclient.UpdateHelpStatusResp
	UpdateMemberReportReq                     = cmsclient.UpdateMemberReportReq
	UpdateMemberReportResp                    = cmsclient.UpdateMemberReportResp
	UpdateMemberReportStatusReq               = cmsclient.UpdateMemberReportStatusReq
	UpdateMemberReportStatusResp              = cmsclient.UpdateMemberReportStatusResp
	UpdatePreferredAreaReq                    = cmsclient.UpdatePreferredAreaReq
	UpdatePreferredAreaResp                   = cmsclient.UpdatePreferredAreaResp
	UpdatePreferredAreaStatusReq              = cmsclient.UpdatePreferredAreaStatusReq
	UpdatePreferredAreaStatusResp             = cmsclient.UpdatePreferredAreaStatusResp
	UpdateSubjectCategoryReq                  = cmsclient.UpdateSubjectCategoryReq
	UpdateSubjectCategoryResp                 = cmsclient.UpdateSubjectCategoryResp
	UpdateSubjectCategoryStatusReq            = cmsclient.UpdateSubjectCategoryStatusReq
	UpdateSubjectCategoryStatusResp           = cmsclient.UpdateSubjectCategoryStatusResp
	UpdateSubjectCommentReq                   = cmsclient.UpdateSubjectCommentReq
	UpdateSubjectCommentResp                  = cmsclient.UpdateSubjectCommentResp
	UpdateSubjectCommentStatusReq             = cmsclient.UpdateSubjectCommentStatusReq
	UpdateSubjectCommentStatusResp            = cmsclient.UpdateSubjectCommentStatusResp
	UpdateSubjectRecommendStatusReq           = cmsclient.UpdateSubjectRecommendStatusReq
	UpdateSubjectRecommendStatusResp          = cmsclient.UpdateSubjectRecommendStatusResp
	UpdateSubjectReq                          = cmsclient.UpdateSubjectReq
	UpdateSubjectResp                         = cmsclient.UpdateSubjectResp
	UpdateSubjectStatusReq                    = cmsclient.UpdateSubjectStatusReq
	UpdateSubjectStatusResp                   = cmsclient.UpdateSubjectStatusResp
	UpdateTopicCategoryReq                    = cmsclient.UpdateTopicCategoryReq
	UpdateTopicCategoryResp                   = cmsclient.UpdateTopicCategoryResp
	UpdateTopicCategoryStatusReq              = cmsclient.UpdateTopicCategoryStatusReq
	UpdateTopicCategoryStatusResp             = cmsclient.UpdateTopicCategoryStatusResp
	UpdateTopicCommentReq                     = cmsclient.UpdateTopicCommentReq
	UpdateTopicCommentResp                    = cmsclient.UpdateTopicCommentResp
	UpdateTopicCommentStatusReq               = cmsclient.UpdateTopicCommentStatusReq
	UpdateTopicCommentStatusResp              = cmsclient.UpdateTopicCommentStatusResp
	UpdateTopicReq                            = cmsclient.UpdateTopicReq
	UpdateTopicResp                           = cmsclient.UpdateTopicResp

	PreferredAreaProductRelationService interface {
		// 添加优选专区和产品关系表
		AddPreferredAreaProductRelation(ctx context.Context, in *AddPreferredAreaProductRelationReq, opts ...grpc.CallOption) (*AddPreferredAreaProductRelationResp, error)
		// 查询优选专区和产品关系表列表
		QueryPreferredAreaProductRelationList(ctx context.Context, in *QueryPreferredAreaProductRelationListReq, opts ...grpc.CallOption) (*QueryPreferredAreaProductRelationListResp, error)
	}

	defaultPreferredAreaProductRelationService struct {
		cli zrpc.Client
	}
)

func NewPreferredAreaProductRelationService(cli zrpc.Client) PreferredAreaProductRelationService {
	return &defaultPreferredAreaProductRelationService{
		cli: cli,
	}
}

// 添加优选专区和产品关系表
func (m *defaultPreferredAreaProductRelationService) AddPreferredAreaProductRelation(ctx context.Context, in *AddPreferredAreaProductRelationReq, opts ...grpc.CallOption) (*AddPreferredAreaProductRelationResp, error) {
	client := cmsclient.NewPreferredAreaProductRelationServiceClient(m.cli.Conn())
	return client.AddPreferredAreaProductRelation(ctx, in, opts...)
}

// 查询优选专区和产品关系表列表
func (m *defaultPreferredAreaProductRelationService) QueryPreferredAreaProductRelationList(ctx context.Context, in *QueryPreferredAreaProductRelationListReq, opts ...grpc.CallOption) (*QueryPreferredAreaProductRelationListResp, error) {
	client := cmsclient.NewPreferredAreaProductRelationServiceClient(m.cli.Conn())
	return client.QueryPreferredAreaProductRelationList(ctx, in, opts...)
}
