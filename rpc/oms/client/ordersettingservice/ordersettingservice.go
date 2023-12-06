// Code generated by goctl. DO NOT EDIT.
// Source: oms.proto

package ordersettingservice

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

	OrderSettingService interface {
		OrderSettingAdd(ctx context.Context, in *OrderSettingAddReq, opts ...grpc.CallOption) (*OrderSettingAddResp, error)
		OrderSettingList(ctx context.Context, in *OrderSettingListReq, opts ...grpc.CallOption) (*OrderSettingListResp, error)
		OrderSettingUpdate(ctx context.Context, in *OrderSettingUpdateReq, opts ...grpc.CallOption) (*OrderSettingUpdateResp, error)
		OrderSettingDelete(ctx context.Context, in *OrderSettingDeleteReq, opts ...grpc.CallOption) (*OrderSettingDeleteResp, error)
	}

	defaultOrderSettingService struct {
		cli zrpc.Client
	}
)

func NewOrderSettingService(cli zrpc.Client) OrderSettingService {
	return &defaultOrderSettingService{
		cli: cli,
	}
}

func (m *defaultOrderSettingService) OrderSettingAdd(ctx context.Context, in *OrderSettingAddReq, opts ...grpc.CallOption) (*OrderSettingAddResp, error) {
	client := omsclient.NewOrderSettingServiceClient(m.cli.Conn())
	return client.OrderSettingAdd(ctx, in, opts...)
}

func (m *defaultOrderSettingService) OrderSettingList(ctx context.Context, in *OrderSettingListReq, opts ...grpc.CallOption) (*OrderSettingListResp, error) {
	client := omsclient.NewOrderSettingServiceClient(m.cli.Conn())
	return client.OrderSettingList(ctx, in, opts...)
}

func (m *defaultOrderSettingService) OrderSettingUpdate(ctx context.Context, in *OrderSettingUpdateReq, opts ...grpc.CallOption) (*OrderSettingUpdateResp, error) {
	client := omsclient.NewOrderSettingServiceClient(m.cli.Conn())
	return client.OrderSettingUpdate(ctx, in, opts...)
}

func (m *defaultOrderSettingService) OrderSettingDelete(ctx context.Context, in *OrderSettingDeleteReq, opts ...grpc.CallOption) (*OrderSettingDeleteResp, error) {
	client := omsclient.NewOrderSettingServiceClient(m.cli.Conn())
	return client.OrderSettingDelete(ctx, in, opts...)
}
