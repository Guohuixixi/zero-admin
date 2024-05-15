// Code generated by goctl. DO NOT EDIT.
// Source: pms.proto

package skustockservice

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AlbumAddReq                                = pmsclient.AlbumAddReq
	AlbumAddResp                               = pmsclient.AlbumAddResp
	AlbumDeleteReq                             = pmsclient.AlbumDeleteReq
	AlbumDeleteResp                            = pmsclient.AlbumDeleteResp
	AlbumListData                              = pmsclient.AlbumListData
	AlbumListReq                               = pmsclient.AlbumListReq
	AlbumListResp                              = pmsclient.AlbumListResp
	AlbumPicAddReq                             = pmsclient.AlbumPicAddReq
	AlbumPicAddResp                            = pmsclient.AlbumPicAddResp
	AlbumPicDeleteReq                          = pmsclient.AlbumPicDeleteReq
	AlbumPicDeleteResp                         = pmsclient.AlbumPicDeleteResp
	AlbumPicListData                           = pmsclient.AlbumPicListData
	AlbumPicListReq                            = pmsclient.AlbumPicListReq
	AlbumPicListResp                           = pmsclient.AlbumPicListResp
	AlbumPicUpdateReq                          = pmsclient.AlbumPicUpdateReq
	AlbumPicUpdateResp                         = pmsclient.AlbumPicUpdateResp
	AlbumUpdateReq                             = pmsclient.AlbumUpdateReq
	AlbumUpdateResp                            = pmsclient.AlbumUpdateResp
	BrandAddReq                                = pmsclient.BrandAddReq
	BrandAddResp                               = pmsclient.BrandAddResp
	BrandDeleteReq                             = pmsclient.BrandDeleteReq
	BrandDeleteResp                            = pmsclient.BrandDeleteResp
	BrandListByIdsReq                          = pmsclient.BrandListByIdsReq
	BrandListData                              = pmsclient.BrandListData
	BrandListReq                               = pmsclient.BrandListReq
	BrandListResp                              = pmsclient.BrandListResp
	BrandUpdateReq                             = pmsclient.BrandUpdateReq
	BrandUpdateResp                            = pmsclient.BrandUpdateResp
	CommentAddReq                              = pmsclient.CommentAddReq
	CommentAddResp                             = pmsclient.CommentAddResp
	CommentDeleteReq                           = pmsclient.CommentDeleteReq
	CommentDeleteResp                          = pmsclient.CommentDeleteResp
	CommentListData                            = pmsclient.CommentListData
	CommentListReq                             = pmsclient.CommentListReq
	CommentListResp                            = pmsclient.CommentListResp
	CommentReplayAddReq                        = pmsclient.CommentReplayAddReq
	CommentReplayAddResp                       = pmsclient.CommentReplayAddResp
	CommentReplayDeleteReq                     = pmsclient.CommentReplayDeleteReq
	CommentReplayDeleteResp                    = pmsclient.CommentReplayDeleteResp
	CommentReplayListData                      = pmsclient.CommentReplayListData
	CommentReplayListReq                       = pmsclient.CommentReplayListReq
	CommentReplayListResp                      = pmsclient.CommentReplayListResp
	CommentReplayUpdateReq                     = pmsclient.CommentReplayUpdateReq
	CommentReplayUpdateResp                    = pmsclient.CommentReplayUpdateResp
	CommentUpdateReq                           = pmsclient.CommentUpdateReq
	CommentUpdateResp                          = pmsclient.CommentUpdateResp
	FeightTemplateAddReq                       = pmsclient.FeightTemplateAddReq
	FeightTemplateAddResp                      = pmsclient.FeightTemplateAddResp
	FeightTemplateDeleteReq                    = pmsclient.FeightTemplateDeleteReq
	FeightTemplateDeleteResp                   = pmsclient.FeightTemplateDeleteResp
	FeightTemplateListData                     = pmsclient.FeightTemplateListData
	FeightTemplateListReq                      = pmsclient.FeightTemplateListReq
	FeightTemplateListResp                     = pmsclient.FeightTemplateListResp
	FeightTemplateUpdateReq                    = pmsclient.FeightTemplateUpdateReq
	FeightTemplateUpdateResp                   = pmsclient.FeightTemplateUpdateResp
	LockSkuStockLockData                       = pmsclient.LockSkuStockLockData
	LockSkuStockLockReq                        = pmsclient.LockSkuStockLockReq
	LockSkuStockLockResp                       = pmsclient.LockSkuStockLockResp
	MemberPriceAddReq                          = pmsclient.MemberPriceAddReq
	MemberPriceAddResp                         = pmsclient.MemberPriceAddResp
	MemberPriceDeleteReq                       = pmsclient.MemberPriceDeleteReq
	MemberPriceDeleteResp                      = pmsclient.MemberPriceDeleteResp
	MemberPriceList                            = pmsclient.MemberPriceList
	MemberPriceListData                        = pmsclient.MemberPriceListData
	MemberPriceListReq                         = pmsclient.MemberPriceListReq
	MemberPriceListResp                        = pmsclient.MemberPriceListResp
	MemberPriceUpdateReq                       = pmsclient.MemberPriceUpdateReq
	MemberPriceUpdateResp                      = pmsclient.MemberPriceUpdateResp
	ProductAddReq                              = pmsclient.ProductAddReq
	ProductAddResp                             = pmsclient.ProductAddResp
	ProductAttributeAddReq                     = pmsclient.ProductAttributeAddReq
	ProductAttributeAddResp                    = pmsclient.ProductAttributeAddResp
	ProductAttributeCategoryAddReq             = pmsclient.ProductAttributeCategoryAddReq
	ProductAttributeCategoryAddResp            = pmsclient.ProductAttributeCategoryAddResp
	ProductAttributeCategoryDeleteReq          = pmsclient.ProductAttributeCategoryDeleteReq
	ProductAttributeCategoryDeleteResp         = pmsclient.ProductAttributeCategoryDeleteResp
	ProductAttributeCategoryListData           = pmsclient.ProductAttributeCategoryListData
	ProductAttributeCategoryListReq            = pmsclient.ProductAttributeCategoryListReq
	ProductAttributeCategoryListResp           = pmsclient.ProductAttributeCategoryListResp
	ProductAttributeCategoryUpdateReq          = pmsclient.ProductAttributeCategoryUpdateReq
	ProductAttributeCategoryUpdateResp         = pmsclient.ProductAttributeCategoryUpdateResp
	ProductAttributeDeleteReq                  = pmsclient.ProductAttributeDeleteReq
	ProductAttributeDeleteResp                 = pmsclient.ProductAttributeDeleteResp
	ProductAttributeListData                   = pmsclient.ProductAttributeListData
	ProductAttributeListReq                    = pmsclient.ProductAttributeListReq
	ProductAttributeListResp                   = pmsclient.ProductAttributeListResp
	ProductAttributeUpdateReq                  = pmsclient.ProductAttributeUpdateReq
	ProductAttributeUpdateResp                 = pmsclient.ProductAttributeUpdateResp
	ProductAttributeValueAddReq                = pmsclient.ProductAttributeValueAddReq
	ProductAttributeValueAddResp               = pmsclient.ProductAttributeValueAddResp
	ProductAttributeValueDeleteReq             = pmsclient.ProductAttributeValueDeleteReq
	ProductAttributeValueDeleteResp            = pmsclient.ProductAttributeValueDeleteResp
	ProductAttributeValueList                  = pmsclient.ProductAttributeValueList
	ProductAttributeValueListData              = pmsclient.ProductAttributeValueListData
	ProductAttributeValueListReq               = pmsclient.ProductAttributeValueListReq
	ProductAttributeValueListResp              = pmsclient.ProductAttributeValueListResp
	ProductAttributeValueUpdateReq             = pmsclient.ProductAttributeValueUpdateReq
	ProductAttributeValueUpdateResp            = pmsclient.ProductAttributeValueUpdateResp
	ProductByIdsReq                            = pmsclient.ProductByIdsReq
	ProductCategoryAddReq                      = pmsclient.ProductCategoryAddReq
	ProductCategoryAddResp                     = pmsclient.ProductCategoryAddResp
	ProductCategoryAttributeRelationAddReq     = pmsclient.ProductCategoryAttributeRelationAddReq
	ProductCategoryAttributeRelationAddResp    = pmsclient.ProductCategoryAttributeRelationAddResp
	ProductCategoryAttributeRelationDeleteReq  = pmsclient.ProductCategoryAttributeRelationDeleteReq
	ProductCategoryAttributeRelationDeleteResp = pmsclient.ProductCategoryAttributeRelationDeleteResp
	ProductCategoryAttributeRelationListData   = pmsclient.ProductCategoryAttributeRelationListData
	ProductCategoryAttributeRelationListReq    = pmsclient.ProductCategoryAttributeRelationListReq
	ProductCategoryAttributeRelationListResp   = pmsclient.ProductCategoryAttributeRelationListResp
	ProductCategoryAttributeRelationUpdateReq  = pmsclient.ProductCategoryAttributeRelationUpdateReq
	ProductCategoryAttributeRelationUpdateResp = pmsclient.ProductCategoryAttributeRelationUpdateResp
	ProductCategoryDeleteReq                   = pmsclient.ProductCategoryDeleteReq
	ProductCategoryDeleteResp                  = pmsclient.ProductCategoryDeleteResp
	ProductCategoryListData                    = pmsclient.ProductCategoryListData
	ProductCategoryListReq                     = pmsclient.ProductCategoryListReq
	ProductCategoryListResp                    = pmsclient.ProductCategoryListResp
	ProductCategoryUpdateReq                   = pmsclient.ProductCategoryUpdateReq
	ProductCategoryUpdateResp                  = pmsclient.ProductCategoryUpdateResp
	ProductDeleteReq                           = pmsclient.ProductDeleteReq
	ProductDeleteResp                          = pmsclient.ProductDeleteResp
	ProductDetailByIdReq                       = pmsclient.ProductDetailByIdReq
	ProductDetailByIdResp                      = pmsclient.ProductDetailByIdResp
	ProductFullReductionAddReq                 = pmsclient.ProductFullReductionAddReq
	ProductFullReductionAddResp                = pmsclient.ProductFullReductionAddResp
	ProductFullReductionDeleteReq              = pmsclient.ProductFullReductionDeleteReq
	ProductFullReductionDeleteResp             = pmsclient.ProductFullReductionDeleteResp
	ProductFullReductionList                   = pmsclient.ProductFullReductionList
	ProductFullReductionListData               = pmsclient.ProductFullReductionListData
	ProductFullReductionListReq                = pmsclient.ProductFullReductionListReq
	ProductFullReductionListResp               = pmsclient.ProductFullReductionListResp
	ProductFullReductionUpdateReq              = pmsclient.ProductFullReductionUpdateReq
	ProductFullReductionUpdateResp             = pmsclient.ProductFullReductionUpdateResp
	ProductLadderAddReq                        = pmsclient.ProductLadderAddReq
	ProductLadderAddResp                       = pmsclient.ProductLadderAddResp
	ProductLadderDeleteReq                     = pmsclient.ProductLadderDeleteReq
	ProductLadderDeleteResp                    = pmsclient.ProductLadderDeleteResp
	ProductLadderList                          = pmsclient.ProductLadderList
	ProductLadderListData                      = pmsclient.ProductLadderListData
	ProductLadderListReq                       = pmsclient.ProductLadderListReq
	ProductLadderListResp                      = pmsclient.ProductLadderListResp
	ProductLadderUpdateReq                     = pmsclient.ProductLadderUpdateReq
	ProductLadderUpdateResp                    = pmsclient.ProductLadderUpdateResp
	ProductListData                            = pmsclient.ProductListData
	ProductListReq                             = pmsclient.ProductListReq
	ProductListResp                            = pmsclient.ProductListResp
	ProductOperateLogAddReq                    = pmsclient.ProductOperateLogAddReq
	ProductOperateLogAddResp                   = pmsclient.ProductOperateLogAddResp
	ProductOperateLogDeleteReq                 = pmsclient.ProductOperateLogDeleteReq
	ProductOperateLogDeleteResp                = pmsclient.ProductOperateLogDeleteResp
	ProductOperateLogListData                  = pmsclient.ProductOperateLogListData
	ProductOperateLogListReq                   = pmsclient.ProductOperateLogListReq
	ProductOperateLogListResp                  = pmsclient.ProductOperateLogListResp
	ProductOperateLogUpdateReq                 = pmsclient.ProductOperateLogUpdateReq
	ProductOperateLogUpdateResp                = pmsclient.ProductOperateLogUpdateResp
	ProductUpdateReq                           = pmsclient.ProductUpdateReq
	ProductUpdateResp                          = pmsclient.ProductUpdateResp
	ProductVertifyRecordAddReq                 = pmsclient.ProductVertifyRecordAddReq
	ProductVertifyRecordAddResp                = pmsclient.ProductVertifyRecordAddResp
	ProductVertifyRecordDeleteReq              = pmsclient.ProductVertifyRecordDeleteReq
	ProductVertifyRecordDeleteResp             = pmsclient.ProductVertifyRecordDeleteResp
	ProductVertifyRecordListData               = pmsclient.ProductVertifyRecordListData
	ProductVertifyRecordListReq                = pmsclient.ProductVertifyRecordListReq
	ProductVertifyRecordListResp               = pmsclient.ProductVertifyRecordListResp
	ProductVertifyRecordUpdateReq              = pmsclient.ProductVertifyRecordUpdateReq
	ProductVertifyRecordUpdateResp             = pmsclient.ProductVertifyRecordUpdateResp
	QueryProductCategoryListData               = pmsclient.QueryProductCategoryListData
	QueryProductCategoryListReq                = pmsclient.QueryProductCategoryListReq
	QueryProductCategoryListResp               = pmsclient.QueryProductCategoryListResp
	QuerySkuStockByProductSkuIdReq             = pmsclient.QuerySkuStockByProductSkuIdReq
	ReleaseSkuStockLockData                    = pmsclient.ReleaseSkuStockLockData
	ReleaseSkuStockLockReq                     = pmsclient.ReleaseSkuStockLockReq
	ReleaseSkuStockLockResp                    = pmsclient.ReleaseSkuStockLockResp
	SkuStockAddReq                             = pmsclient.SkuStockAddReq
	SkuStockAddResp                            = pmsclient.SkuStockAddResp
	SkuStockDeleteReq                          = pmsclient.SkuStockDeleteReq
	SkuStockDeleteResp                         = pmsclient.SkuStockDeleteResp
	SkuStockList                               = pmsclient.SkuStockList
	SkuStockListData                           = pmsclient.SkuStockListData
	SkuStockListReq                            = pmsclient.SkuStockListReq
	SkuStockListResp                           = pmsclient.SkuStockListResp
	SkuStockUpdateReq                          = pmsclient.SkuStockUpdateReq
	SkuStockUpdateResp                         = pmsclient.SkuStockUpdateResp
	UpdateBrandFactoryStatusReq                = pmsclient.UpdateBrandFactoryStatusReq
	UpdateBrandFactoryStatusResp               = pmsclient.UpdateBrandFactoryStatusResp
	UpdateBrandShowStatusReq                   = pmsclient.UpdateBrandShowStatusReq
	UpdateBrandShowStatusResp                  = pmsclient.UpdateBrandShowStatusResp
	UpdateProductCategoryStatusReq             = pmsclient.UpdateProductCategoryStatusReq

	SkuStockService interface {
		SkuStockAdd(ctx context.Context, in *SkuStockAddReq, opts ...grpc.CallOption) (*SkuStockAddResp, error)
		SkuStockList(ctx context.Context, in *SkuStockListReq, opts ...grpc.CallOption) (*SkuStockListResp, error)
		SkuStockUpdate(ctx context.Context, in *SkuStockUpdateReq, opts ...grpc.CallOption) (*SkuStockUpdateResp, error)
		SkuStockDelete(ctx context.Context, in *SkuStockDeleteReq, opts ...grpc.CallOption) (*SkuStockDeleteResp, error)
		// 取消订单的时候,释放库存
		ReleaseSkuStockLock(ctx context.Context, in *ReleaseSkuStockLockReq, opts ...grpc.CallOption) (*ReleaseSkuStockLockResp, error)
		// 下单的时候,锁定库存
		LockSkuStockLock(ctx context.Context, in *LockSkuStockLockReq, opts ...grpc.CallOption) (*LockSkuStockLockResp, error)
		// 根据ProductSkuId查询sku
		QuerySkuStockByProductSkuId(ctx context.Context, in *QuerySkuStockByProductSkuIdReq, opts ...grpc.CallOption) (*SkuStockListData, error)
	}

	defaultSkuStockService struct {
		cli zrpc.Client
	}
)

func NewSkuStockService(cli zrpc.Client) SkuStockService {
	return &defaultSkuStockService{
		cli: cli,
	}
}

func (m *defaultSkuStockService) SkuStockAdd(ctx context.Context, in *SkuStockAddReq, opts ...grpc.CallOption) (*SkuStockAddResp, error) {
	client := pmsclient.NewSkuStockServiceClient(m.cli.Conn())
	return client.SkuStockAdd(ctx, in, opts...)
}

func (m *defaultSkuStockService) SkuStockList(ctx context.Context, in *SkuStockListReq, opts ...grpc.CallOption) (*SkuStockListResp, error) {
	client := pmsclient.NewSkuStockServiceClient(m.cli.Conn())
	return client.SkuStockList(ctx, in, opts...)
}

func (m *defaultSkuStockService) SkuStockUpdate(ctx context.Context, in *SkuStockUpdateReq, opts ...grpc.CallOption) (*SkuStockUpdateResp, error) {
	client := pmsclient.NewSkuStockServiceClient(m.cli.Conn())
	return client.SkuStockUpdate(ctx, in, opts...)
}

func (m *defaultSkuStockService) SkuStockDelete(ctx context.Context, in *SkuStockDeleteReq, opts ...grpc.CallOption) (*SkuStockDeleteResp, error) {
	client := pmsclient.NewSkuStockServiceClient(m.cli.Conn())
	return client.SkuStockDelete(ctx, in, opts...)
}

// 取消订单的时候,释放库存
func (m *defaultSkuStockService) ReleaseSkuStockLock(ctx context.Context, in *ReleaseSkuStockLockReq, opts ...grpc.CallOption) (*ReleaseSkuStockLockResp, error) {
	client := pmsclient.NewSkuStockServiceClient(m.cli.Conn())
	return client.ReleaseSkuStockLock(ctx, in, opts...)
}

// 下单的时候,锁定库存
func (m *defaultSkuStockService) LockSkuStockLock(ctx context.Context, in *LockSkuStockLockReq, opts ...grpc.CallOption) (*LockSkuStockLockResp, error) {
	client := pmsclient.NewSkuStockServiceClient(m.cli.Conn())
	return client.LockSkuStockLock(ctx, in, opts...)
}

// 根据ProductSkuId查询sku
func (m *defaultSkuStockService) QuerySkuStockByProductSkuId(ctx context.Context, in *QuerySkuStockByProductSkuIdReq, opts ...grpc.CallOption) (*SkuStockListData, error) {
	client := pmsclient.NewSkuStockServiceClient(m.cli.Conn())
	return client.QuerySkuStockByProductSkuId(ctx, in, opts...)
}
