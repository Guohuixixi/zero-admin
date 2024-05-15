// Code generated by goctl. DO NOT EDIT.
// Source: pms.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/internal/logic/productcategoryservice"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
)

type ProductCategoryServiceServer struct {
	svcCtx *svc.ServiceContext
	pmsclient.UnimplementedProductCategoryServiceServer
}

func NewProductCategoryServiceServer(svcCtx *svc.ServiceContext) *ProductCategoryServiceServer {
	return &ProductCategoryServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加商品分类
func (s *ProductCategoryServiceServer) ProductCategoryAdd(ctx context.Context, in *pmsclient.ProductCategoryAddReq) (*pmsclient.ProductCategoryAddResp, error) {
	l := productcategoryservicelogic.NewProductCategoryAddLogic(ctx, s.svcCtx)
	return l.ProductCategoryAdd(in)
}

// 查询商品分类
func (s *ProductCategoryServiceServer) ProductCategoryList(ctx context.Context, in *pmsclient.ProductCategoryListReq) (*pmsclient.ProductCategoryListResp, error) {
	l := productcategoryservicelogic.NewProductCategoryListLogic(ctx, s.svcCtx)
	return l.ProductCategoryList(in)
}

// 更新商品分类
func (s *ProductCategoryServiceServer) ProductCategoryUpdate(ctx context.Context, in *pmsclient.ProductCategoryUpdateReq) (*pmsclient.ProductCategoryUpdateResp, error) {
	l := productcategoryservicelogic.NewProductCategoryUpdateLogic(ctx, s.svcCtx)
	return l.ProductCategoryUpdate(in)
}

// 更新商品分类导航显示状态
func (s *ProductCategoryServiceServer) UpdateCategoryNavStatus(ctx context.Context, in *pmsclient.UpdateProductCategoryStatusReq) (*pmsclient.ProductCategoryUpdateResp, error) {
	l := productcategoryservicelogic.NewUpdateCategoryNavStatusLogic(ctx, s.svcCtx)
	return l.UpdateCategoryNavStatus(in)
}

// 更新商品分类显示状态
func (s *ProductCategoryServiceServer) UpdateCategoryShowStatus(ctx context.Context, in *pmsclient.UpdateProductCategoryStatusReq) (*pmsclient.ProductCategoryUpdateResp, error) {
	l := productcategoryservicelogic.NewUpdateCategoryShowStatusLogic(ctx, s.svcCtx)
	return l.UpdateCategoryShowStatus(in)
}

// 查询商品分类
func (s *ProductCategoryServiceServer) ProductCategoryDelete(ctx context.Context, in *pmsclient.ProductCategoryDeleteReq) (*pmsclient.ProductCategoryDeleteResp, error) {
	l := productcategoryservicelogic.NewProductCategoryDeleteLogic(ctx, s.svcCtx)
	return l.ProductCategoryDelete(in)
}

// 查询商品分类
func (s *ProductCategoryServiceServer) QueryProductCategoryList(ctx context.Context, in *pmsclient.QueryProductCategoryListReq) (*pmsclient.QueryProductCategoryListResp, error) {
	l := productcategoryservicelogic.NewQueryProductCategoryListLogic(ctx, s.svcCtx)
	return l.QueryProductCategoryList(in)
}
