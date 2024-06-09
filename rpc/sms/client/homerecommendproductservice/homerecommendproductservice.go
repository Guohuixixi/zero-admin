// Code generated by goctl. DO NOT EDIT.
// Source: sms.proto

package homerecommendproductservice

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
	CouponProductCategoryRelationListData         = smsclient.CouponProductCategoryRelationListData
	CouponProductRelationListData                 = smsclient.CouponProductRelationListData
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
	UpdateFlashPromotionSessionStatusReq          = smsclient.UpdateFlashPromotionSessionStatusReq
	UpdateFlashPromotionSessionStatusResp         = smsclient.UpdateFlashPromotionSessionStatusResp
	UpdateFlashPromotionStatusReq                 = smsclient.UpdateFlashPromotionStatusReq
	UpdateFlashPromotionStatusResp                = smsclient.UpdateFlashPromotionStatusResp
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

	HomeRecommendProductService interface {
		// 添加人气推荐
		HomeRecommendProductAdd(ctx context.Context, in *HomeRecommendProductAddReq, opts ...grpc.CallOption) (*HomeRecommendProductAddResp, error)
		// 查询人气推荐
		HomeRecommendProductList(ctx context.Context, in *HomeRecommendProductListReq, opts ...grpc.CallOption) (*HomeRecommendProductListResp, error)
		// 修改推荐排序
		UpdateRecommendProductSort(ctx context.Context, in *UpdateRecommendProductSortReq, opts ...grpc.CallOption) (*UpdateRecommendProductSortResp, error)
		// 批量修改推荐状态
		UpdateRecommendProductStatus(ctx context.Context, in *UpdateRecommendProductStatusReq, opts ...grpc.CallOption) (*UpdateRecommendProductStatusResp, error)
		// 删除人气推荐
		HomeRecommendProductDelete(ctx context.Context, in *HomeRecommendProductDeleteReq, opts ...grpc.CallOption) (*HomeRecommendProductDeleteResp, error)
	}

	defaultHomeRecommendProductService struct {
		cli zrpc.Client
	}
)

func NewHomeRecommendProductService(cli zrpc.Client) HomeRecommendProductService {
	return &defaultHomeRecommendProductService{
		cli: cli,
	}
}

// 添加人气推荐
func (m *defaultHomeRecommendProductService) HomeRecommendProductAdd(ctx context.Context, in *HomeRecommendProductAddReq, opts ...grpc.CallOption) (*HomeRecommendProductAddResp, error) {
	client := smsclient.NewHomeRecommendProductServiceClient(m.cli.Conn())
	return client.HomeRecommendProductAdd(ctx, in, opts...)
}

// 查询人气推荐
func (m *defaultHomeRecommendProductService) HomeRecommendProductList(ctx context.Context, in *HomeRecommendProductListReq, opts ...grpc.CallOption) (*HomeRecommendProductListResp, error) {
	client := smsclient.NewHomeRecommendProductServiceClient(m.cli.Conn())
	return client.HomeRecommendProductList(ctx, in, opts...)
}

// 修改推荐排序
func (m *defaultHomeRecommendProductService) UpdateRecommendProductSort(ctx context.Context, in *UpdateRecommendProductSortReq, opts ...grpc.CallOption) (*UpdateRecommendProductSortResp, error) {
	client := smsclient.NewHomeRecommendProductServiceClient(m.cli.Conn())
	return client.UpdateRecommendProductSort(ctx, in, opts...)
}

// 批量修改推荐状态
func (m *defaultHomeRecommendProductService) UpdateRecommendProductStatus(ctx context.Context, in *UpdateRecommendProductStatusReq, opts ...grpc.CallOption) (*UpdateRecommendProductStatusResp, error) {
	client := smsclient.NewHomeRecommendProductServiceClient(m.cli.Conn())
	return client.UpdateRecommendProductStatus(ctx, in, opts...)
}

// 删除人气推荐
func (m *defaultHomeRecommendProductService) HomeRecommendProductDelete(ctx context.Context, in *HomeRecommendProductDeleteReq, opts ...grpc.CallOption) (*HomeRecommendProductDeleteResp, error) {
	client := smsclient.NewHomeRecommendProductServiceClient(m.cli.Conn())
	return client.HomeRecommendProductDelete(ctx, in, opts...)
}
