// Code generated by goctl. DO NOT EDIT.
// Source: sms.proto

package couponproductrelationservice

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CouponAddOrUpdateReq                          = smsclient.CouponAddOrUpdateReq
	CouponAddOrUpdateResp                         = smsclient.CouponAddOrUpdateResp
	CouponCountReq                                = smsclient.CouponCountReq
	CouponCountResp                               = smsclient.CouponCountResp
	CouponDeleteReq                               = smsclient.CouponDeleteReq
	CouponDeleteResp                              = smsclient.CouponDeleteResp
	CouponFindByIdReq                             = smsclient.CouponFindByIdReq
	CouponFindByIdResp                            = smsclient.CouponFindByIdResp
	CouponFindByIdsReq                            = smsclient.CouponFindByIdsReq
	CouponFindByIdsResp                           = smsclient.CouponFindByIdsResp
	CouponFindByProductIdAndProductCategoryIdReq  = smsclient.CouponFindByProductIdAndProductCategoryIdReq
	CouponFindByProductIdAndProductCategoryIdResp = smsclient.CouponFindByProductIdAndProductCategoryIdResp
	CouponHistoryAddReq                           = smsclient.CouponHistoryAddReq
	CouponHistoryAddResp                          = smsclient.CouponHistoryAddResp
	CouponHistoryDeleteReq                        = smsclient.CouponHistoryDeleteReq
	CouponHistoryDeleteResp                       = smsclient.CouponHistoryDeleteResp
	CouponHistoryDetailListData                   = smsclient.CouponHistoryDetailListData
	CouponHistoryDetailListReq                    = smsclient.CouponHistoryDetailListReq
	CouponHistoryDetailListResp                   = smsclient.CouponHistoryDetailListResp
	CouponHistoryListData                         = smsclient.CouponHistoryListData
	CouponHistoryListReq                          = smsclient.CouponHistoryListReq
	CouponHistoryListResp                         = smsclient.CouponHistoryListResp
	CouponHistoryUpdateReq                        = smsclient.CouponHistoryUpdateReq
	CouponHistoryUpdateResp                       = smsclient.CouponHistoryUpdateResp
	CouponListData                                = smsclient.CouponListData
	CouponListReq                                 = smsclient.CouponListReq
	CouponListResp                                = smsclient.CouponListResp
	CouponProductCategoryRelationAddReq           = smsclient.CouponProductCategoryRelationAddReq
	CouponProductCategoryRelationAddResp          = smsclient.CouponProductCategoryRelationAddResp
	CouponProductCategoryRelationDeleteReq        = smsclient.CouponProductCategoryRelationDeleteReq
	CouponProductCategoryRelationDeleteResp       = smsclient.CouponProductCategoryRelationDeleteResp
	CouponProductCategoryRelationListData         = smsclient.CouponProductCategoryRelationListData
	CouponProductCategoryRelationUpdateReq        = smsclient.CouponProductCategoryRelationUpdateReq
	CouponProductCategoryRelationUpdateResp       = smsclient.CouponProductCategoryRelationUpdateResp
	CouponProductRelationAddReq                   = smsclient.CouponProductRelationAddReq
	CouponProductRelationAddResp                  = smsclient.CouponProductRelationAddResp
	CouponProductRelationDeleteReq                = smsclient.CouponProductRelationDeleteReq
	CouponProductRelationDeleteResp               = smsclient.CouponProductRelationDeleteResp
	CouponProductRelationListData                 = smsclient.CouponProductRelationListData
	CouponProductRelationUpdateReq                = smsclient.CouponProductRelationUpdateReq
	CouponProductRelationUpdateResp               = smsclient.CouponProductRelationUpdateResp
	FlashPromotionAddReq                          = smsclient.FlashPromotionAddReq
	FlashPromotionAddResp                         = smsclient.FlashPromotionAddResp
	FlashPromotionDeleteReq                       = smsclient.FlashPromotionDeleteReq
	FlashPromotionDeleteResp                      = smsclient.FlashPromotionDeleteResp
	FlashPromotionListByDateReq                   = smsclient.FlashPromotionListByDateReq
	FlashPromotionListByDateResp                  = smsclient.FlashPromotionListByDateResp
	FlashPromotionListData                        = smsclient.FlashPromotionListData
	FlashPromotionListReq                         = smsclient.FlashPromotionListReq
	FlashPromotionListResp                        = smsclient.FlashPromotionListResp
	FlashPromotionLogAddReq                       = smsclient.FlashPromotionLogAddReq
	FlashPromotionLogAddResp                      = smsclient.FlashPromotionLogAddResp
	FlashPromotionLogDeleteReq                    = smsclient.FlashPromotionLogDeleteReq
	FlashPromotionLogDeleteResp                   = smsclient.FlashPromotionLogDeleteResp
	FlashPromotionLogListData                     = smsclient.FlashPromotionLogListData
	FlashPromotionLogListReq                      = smsclient.FlashPromotionLogListReq
	FlashPromotionLogListResp                     = smsclient.FlashPromotionLogListResp
	FlashPromotionLogUpdateReq                    = smsclient.FlashPromotionLogUpdateReq
	FlashPromotionLogUpdateResp                   = smsclient.FlashPromotionLogUpdateResp
	FlashPromotionProductRelationAddReq           = smsclient.FlashPromotionProductRelationAddReq
	FlashPromotionProductRelationAddResp          = smsclient.FlashPromotionProductRelationAddResp
	FlashPromotionProductRelationDeleteReq        = smsclient.FlashPromotionProductRelationDeleteReq
	FlashPromotionProductRelationDeleteResp       = smsclient.FlashPromotionProductRelationDeleteResp
	FlashPromotionProductRelationListData         = smsclient.FlashPromotionProductRelationListData
	FlashPromotionProductRelationListReq          = smsclient.FlashPromotionProductRelationListReq
	FlashPromotionProductRelationListResp         = smsclient.FlashPromotionProductRelationListResp
	FlashPromotionProductRelationUpdateReq        = smsclient.FlashPromotionProductRelationUpdateReq
	FlashPromotionProductRelationUpdateResp       = smsclient.FlashPromotionProductRelationUpdateResp
	FlashPromotionSessionAddReq                   = smsclient.FlashPromotionSessionAddReq
	FlashPromotionSessionAddResp                  = smsclient.FlashPromotionSessionAddResp
	FlashPromotionSessionByTimeReq                = smsclient.FlashPromotionSessionByTimeReq
	FlashPromotionSessionByTimeResp               = smsclient.FlashPromotionSessionByTimeResp
	FlashPromotionSessionDeleteReq                = smsclient.FlashPromotionSessionDeleteReq
	FlashPromotionSessionDeleteResp               = smsclient.FlashPromotionSessionDeleteResp
	FlashPromotionSessionListData                 = smsclient.FlashPromotionSessionListData
	FlashPromotionSessionListReq                  = smsclient.FlashPromotionSessionListReq
	FlashPromotionSessionListResp                 = smsclient.FlashPromotionSessionListResp
	FlashPromotionSessionUpdateReq                = smsclient.FlashPromotionSessionUpdateReq
	FlashPromotionSessionUpdateResp               = smsclient.FlashPromotionSessionUpdateResp
	FlashPromotionUpdateReq                       = smsclient.FlashPromotionUpdateReq
	FlashPromotionUpdateResp                      = smsclient.FlashPromotionUpdateResp
	HomeAdvertiseAddReq                           = smsclient.HomeAdvertiseAddReq
	HomeAdvertiseAddResp                          = smsclient.HomeAdvertiseAddResp
	HomeAdvertiseDeleteReq                        = smsclient.HomeAdvertiseDeleteReq
	HomeAdvertiseDeleteResp                       = smsclient.HomeAdvertiseDeleteResp
	HomeAdvertiseListData                         = smsclient.HomeAdvertiseListData
	HomeAdvertiseListReq                          = smsclient.HomeAdvertiseListReq
	HomeAdvertiseListResp                         = smsclient.HomeAdvertiseListResp
	HomeAdvertiseUpdateReq                        = smsclient.HomeAdvertiseUpdateReq
	HomeAdvertiseUpdateResp                       = smsclient.HomeAdvertiseUpdateResp
	HomeBrandAddData                              = smsclient.HomeBrandAddData
	HomeBrandAddReq                               = smsclient.HomeBrandAddReq
	HomeBrandAddResp                              = smsclient.HomeBrandAddResp
	HomeBrandDeleteReq                            = smsclient.HomeBrandDeleteReq
	HomeBrandDeleteResp                           = smsclient.HomeBrandDeleteResp
	HomeBrandListData                             = smsclient.HomeBrandListData
	HomeBrandListReq                              = smsclient.HomeBrandListReq
	HomeBrandListResp                             = smsclient.HomeBrandListResp
	HomeNewProductAddData                         = smsclient.HomeNewProductAddData
	HomeNewProductAddReq                          = smsclient.HomeNewProductAddReq
	HomeNewProductAddResp                         = smsclient.HomeNewProductAddResp
	HomeNewProductDeleteReq                       = smsclient.HomeNewProductDeleteReq
	HomeNewProductDeleteResp                      = smsclient.HomeNewProductDeleteResp
	HomeNewProductListData                        = smsclient.HomeNewProductListData
	HomeNewProductListReq                         = smsclient.HomeNewProductListReq
	HomeNewProductListResp                        = smsclient.HomeNewProductListResp
	HomeRecommendProductAddData                   = smsclient.HomeRecommendProductAddData
	HomeRecommendProductAddReq                    = smsclient.HomeRecommendProductAddReq
	HomeRecommendProductAddResp                   = smsclient.HomeRecommendProductAddResp
	HomeRecommendProductDeleteReq                 = smsclient.HomeRecommendProductDeleteReq
	HomeRecommendProductDeleteResp                = smsclient.HomeRecommendProductDeleteResp
	HomeRecommendProductListData                  = smsclient.HomeRecommendProductListData
	HomeRecommendProductListReq                   = smsclient.HomeRecommendProductListReq
	HomeRecommendProductListResp                  = smsclient.HomeRecommendProductListResp
	HomeRecommendSubjectAddData                   = smsclient.HomeRecommendSubjectAddData
	HomeRecommendSubjectAddReq                    = smsclient.HomeRecommendSubjectAddReq
	HomeRecommendSubjectAddResp                   = smsclient.HomeRecommendSubjectAddResp
	HomeRecommendSubjectDeleteReq                 = smsclient.HomeRecommendSubjectDeleteReq
	HomeRecommendSubjectDeleteResp                = smsclient.HomeRecommendSubjectDeleteResp
	HomeRecommendSubjectListData                  = smsclient.HomeRecommendSubjectListData
	HomeRecommendSubjectListReq                   = smsclient.HomeRecommendSubjectListReq
	HomeRecommendSubjectListResp                  = smsclient.HomeRecommendSubjectListResp
	QueryFlashPromotionByProductReq               = smsclient.QueryFlashPromotionByProductReq
	QueryFlashPromotionByProductResp              = smsclient.QueryFlashPromotionByProductResp
	QueryMemberCouponListReq                      = smsclient.QueryMemberCouponListReq
	QueryMemberCouponListResp                     = smsclient.QueryMemberCouponListResp
	UpdateCouponStatusReq                         = smsclient.UpdateCouponStatusReq
	UpdateCouponStatusResp                        = smsclient.UpdateCouponStatusResp
	UpdateHomeAdvertiseStatusReq                  = smsclient.UpdateHomeAdvertiseStatusReq
	UpdateHomeBrandSortReq                        = smsclient.UpdateHomeBrandSortReq
	UpdateHomeBrandSortResp                       = smsclient.UpdateHomeBrandSortResp
	UpdateHomeBrandStatusReq                      = smsclient.UpdateHomeBrandStatusReq
	UpdateHomeBrandStatusResp                     = smsclient.UpdateHomeBrandStatusResp
	UpdateNewProductSortReq                       = smsclient.UpdateNewProductSortReq
	UpdateNewProductSortResp                      = smsclient.UpdateNewProductSortResp
	UpdateNewProductStatusReq                     = smsclient.UpdateNewProductStatusReq
	UpdateNewProductStatusResp                    = smsclient.UpdateNewProductStatusResp
	UpdateRecommendProductSortReq                 = smsclient.UpdateRecommendProductSortReq
	UpdateRecommendProductSortResp                = smsclient.UpdateRecommendProductSortResp
	UpdateRecommendProductStatusReq               = smsclient.UpdateRecommendProductStatusReq
	UpdateRecommendProductStatusResp              = smsclient.UpdateRecommendProductStatusResp
	UpdateRecommendSubjectSortReq                 = smsclient.UpdateRecommendSubjectSortReq
	UpdateRecommendSubjectSortResp                = smsclient.UpdateRecommendSubjectSortResp
	UpdateRecommendSubjectStatusReq               = smsclient.UpdateRecommendSubjectStatusReq
	UpdateRecommendSubjectStatusResp              = smsclient.UpdateRecommendSubjectStatusResp

	CouponProductRelationService interface {
		CouponProductRelationAdd(ctx context.Context, in *CouponProductRelationAddReq, opts ...grpc.CallOption) (*CouponProductRelationAddResp, error)
		CouponProductRelationUpdate(ctx context.Context, in *CouponProductRelationUpdateReq, opts ...grpc.CallOption) (*CouponProductRelationUpdateResp, error)
		CouponProductRelationDelete(ctx context.Context, in *CouponProductRelationDeleteReq, opts ...grpc.CallOption) (*CouponProductRelationDeleteResp, error)
	}

	defaultCouponProductRelationService struct {
		cli zrpc.Client
	}
)

func NewCouponProductRelationService(cli zrpc.Client) CouponProductRelationService {
	return &defaultCouponProductRelationService{
		cli: cli,
	}
}

func (m *defaultCouponProductRelationService) CouponProductRelationAdd(ctx context.Context, in *CouponProductRelationAddReq, opts ...grpc.CallOption) (*CouponProductRelationAddResp, error) {
	client := smsclient.NewCouponProductRelationServiceClient(m.cli.Conn())
	return client.CouponProductRelationAdd(ctx, in, opts...)
}

func (m *defaultCouponProductRelationService) CouponProductRelationUpdate(ctx context.Context, in *CouponProductRelationUpdateReq, opts ...grpc.CallOption) (*CouponProductRelationUpdateResp, error) {
	client := smsclient.NewCouponProductRelationServiceClient(m.cli.Conn())
	return client.CouponProductRelationUpdate(ctx, in, opts...)
}

func (m *defaultCouponProductRelationService) CouponProductRelationDelete(ctx context.Context, in *CouponProductRelationDeleteReq, opts ...grpc.CallOption) (*CouponProductRelationDeleteResp, error) {
	client := smsclient.NewCouponProductRelationServiceClient(m.cli.Conn())
	return client.CouponProductRelationDelete(ctx, in, opts...)
}
