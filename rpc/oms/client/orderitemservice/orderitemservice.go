// Code generated by goctl. DO NOT EDIT.
// Source: oms.proto

package orderitemservice

import (
	"context"

	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CartItemAddReq                = omsclient.CartItemAddReq
	CartItemAddResp               = omsclient.CartItemAddResp
	CartItemClearReq              = omsclient.CartItemClearReq
	CartItemClearResp             = omsclient.CartItemClearResp
	CartItemDeleteReq             = omsclient.CartItemDeleteReq
	CartItemDeleteResp            = omsclient.CartItemDeleteResp
	CartItemListData              = omsclient.CartItemListData
	CartItemListReq               = omsclient.CartItemListReq
	CartItemListResp              = omsclient.CartItemListResp
	CartItemUpdateReq             = omsclient.CartItemUpdateReq
	CartItemUpdateResp            = omsclient.CartItemUpdateResp
	CompanyAddressAddReq          = omsclient.CompanyAddressAddReq
	CompanyAddressAddResp         = omsclient.CompanyAddressAddResp
	CompanyAddressDeleteReq       = omsclient.CompanyAddressDeleteReq
	CompanyAddressDeleteResp      = omsclient.CompanyAddressDeleteResp
	CompanyAddressListData        = omsclient.CompanyAddressListData
	CompanyAddressListReq         = omsclient.CompanyAddressListReq
	CompanyAddressListResp        = omsclient.CompanyAddressListResp
	CompanyAddressUpdateReq       = omsclient.CompanyAddressUpdateReq
	CompanyAddressUpdateResp      = omsclient.CompanyAddressUpdateResp
	GoodsListByMemberIdData       = omsclient.GoodsListByMemberIdData
	OrderAddReq                   = omsclient.OrderAddReq
	OrderAddResp                  = omsclient.OrderAddResp
	OrderCancelReq                = omsclient.OrderCancelReq
	OrderCancelResp               = omsclient.OrderCancelResp
	OrderConfirmReq               = omsclient.OrderConfirmReq
	OrderConfirmResp              = omsclient.OrderConfirmResp
	OrderDeleteByIdReq            = omsclient.OrderDeleteByIdReq
	OrderDeleteReq                = omsclient.OrderDeleteReq
	OrderDeleteResp               = omsclient.OrderDeleteResp
	OrderItemAddReq               = omsclient.OrderItemAddReq
	OrderItemAddResp              = omsclient.OrderItemAddResp
	OrderItemDeleteReq            = omsclient.OrderItemDeleteReq
	OrderItemDeleteResp           = omsclient.OrderItemDeleteResp
	OrderItemListData             = omsclient.OrderItemListData
	OrderItemListReq              = omsclient.OrderItemListReq
	OrderItemListResp             = omsclient.OrderItemListResp
	OrderItemUpdateReq            = omsclient.OrderItemUpdateReq
	OrderItemUpdateResp           = omsclient.OrderItemUpdateResp
	OrderListByMemberIdData       = omsclient.OrderListByMemberIdData
	OrderListByMemberIdReq        = omsclient.OrderListByMemberIdReq
	OrderListByMemberIdResp       = omsclient.OrderListByMemberIdResp
	OrderListData                 = omsclient.OrderListData
	OrderListReq                  = omsclient.OrderListReq
	OrderListResp                 = omsclient.OrderListResp
	OrderOperateHistoryAddReq     = omsclient.OrderOperateHistoryAddReq
	OrderOperateHistoryAddResp    = omsclient.OrderOperateHistoryAddResp
	OrderOperateHistoryDeleteReq  = omsclient.OrderOperateHistoryDeleteReq
	OrderOperateHistoryDeleteResp = omsclient.OrderOperateHistoryDeleteResp
	OrderOperateHistoryListData   = omsclient.OrderOperateHistoryListData
	OrderOperateHistoryListReq    = omsclient.OrderOperateHistoryListReq
	OrderOperateHistoryListResp   = omsclient.OrderOperateHistoryListResp
	OrderOperateHistoryUpdateReq  = omsclient.OrderOperateHistoryUpdateReq
	OrderOperateHistoryUpdateResp = omsclient.OrderOperateHistoryUpdateResp
	OrderRefundReq                = omsclient.OrderRefundReq
	OrderRefundResp               = omsclient.OrderRefundResp
	OrderReturnApplyAddReq        = omsclient.OrderReturnApplyAddReq
	OrderReturnApplyAddResp       = omsclient.OrderReturnApplyAddResp
	OrderReturnApplyDeleteReq     = omsclient.OrderReturnApplyDeleteReq
	OrderReturnApplyDeleteResp    = omsclient.OrderReturnApplyDeleteResp
	OrderReturnApplyListData      = omsclient.OrderReturnApplyListData
	OrderReturnApplyListReq       = omsclient.OrderReturnApplyListReq
	OrderReturnApplyListResp      = omsclient.OrderReturnApplyListResp
	OrderReturnApplyUpdateReq     = omsclient.OrderReturnApplyUpdateReq
	OrderReturnApplyUpdateResp    = omsclient.OrderReturnApplyUpdateResp
	OrderReturnReasonAddReq       = omsclient.OrderReturnReasonAddReq
	OrderReturnReasonAddResp      = omsclient.OrderReturnReasonAddResp
	OrderReturnReasonDeleteReq    = omsclient.OrderReturnReasonDeleteReq
	OrderReturnReasonDeleteResp   = omsclient.OrderReturnReasonDeleteResp
	OrderReturnReasonListData     = omsclient.OrderReturnReasonListData
	OrderReturnReasonListReq      = omsclient.OrderReturnReasonListReq
	OrderReturnReasonListResp     = omsclient.OrderReturnReasonListResp
	OrderReturnReasonUpdateReq    = omsclient.OrderReturnReasonUpdateReq
	OrderReturnReasonUpdateResp   = omsclient.OrderReturnReasonUpdateResp
	OrderSettingAddReq            = omsclient.OrderSettingAddReq
	OrderSettingAddResp           = omsclient.OrderSettingAddResp
	OrderSettingDeleteReq         = omsclient.OrderSettingDeleteReq
	OrderSettingDeleteResp        = omsclient.OrderSettingDeleteResp
	OrderSettingListData          = omsclient.OrderSettingListData
	OrderSettingListReq           = omsclient.OrderSettingListReq
	OrderSettingListResp          = omsclient.OrderSettingListResp
	OrderSettingUpdateReq         = omsclient.OrderSettingUpdateReq
	OrderSettingUpdateResp        = omsclient.OrderSettingUpdateResp
	OrderUpdateReq                = omsclient.OrderUpdateReq
	OrderUpdateResp               = omsclient.OrderUpdateResp

	OrderItemService interface {
		OrderItemAdd(ctx context.Context, in *OrderItemAddReq, opts ...grpc.CallOption) (*OrderItemAddResp, error)
		OrderItemList(ctx context.Context, in *OrderItemListReq, opts ...grpc.CallOption) (*OrderItemListResp, error)
		OrderItemUpdate(ctx context.Context, in *OrderItemUpdateReq, opts ...grpc.CallOption) (*OrderItemUpdateResp, error)
		OrderItemDelete(ctx context.Context, in *OrderItemDeleteReq, opts ...grpc.CallOption) (*OrderItemDeleteResp, error)
	}

	defaultOrderItemService struct {
		cli zrpc.Client
	}
)

func NewOrderItemService(cli zrpc.Client) OrderItemService {
	return &defaultOrderItemService{
		cli: cli,
	}
}

func (m *defaultOrderItemService) OrderItemAdd(ctx context.Context, in *OrderItemAddReq, opts ...grpc.CallOption) (*OrderItemAddResp, error) {
	client := omsclient.NewOrderItemServiceClient(m.cli.Conn())
	return client.OrderItemAdd(ctx, in, opts...)
}

func (m *defaultOrderItemService) OrderItemList(ctx context.Context, in *OrderItemListReq, opts ...grpc.CallOption) (*OrderItemListResp, error) {
	client := omsclient.NewOrderItemServiceClient(m.cli.Conn())
	return client.OrderItemList(ctx, in, opts...)
}

func (m *defaultOrderItemService) OrderItemUpdate(ctx context.Context, in *OrderItemUpdateReq, opts ...grpc.CallOption) (*OrderItemUpdateResp, error) {
	client := omsclient.NewOrderItemServiceClient(m.cli.Conn())
	return client.OrderItemUpdate(ctx, in, opts...)
}

func (m *defaultOrderItemService) OrderItemDelete(ctx context.Context, in *OrderItemDeleteReq, opts ...grpc.CallOption) (*OrderItemDeleteResp, error) {
	client := omsclient.NewOrderItemServiceClient(m.cli.Conn())
	return client.OrderItemDelete(ctx, in, opts...)
}
