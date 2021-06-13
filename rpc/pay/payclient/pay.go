// Code generated by goctl. DO NOT EDIT!
// Source: pay.proto

//go:generate mockgen -destination ./pay_mock.go -package payclient -source $GOFILE

package payclient

import (
	"context"

	"go-zero-admin/rpc/pay/pay"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	UnifiedOrderReq    = pay.UnifiedOrderReq
	UnifiedOrderResp   = pay.UnifiedOrderResp
	H5UnifiedOrderResp = pay.H5UnifiedOrderResp
	OrderQueryReq      = pay.OrderQueryReq
	OrderQueryResp     = pay.OrderQueryResp

	Pay interface {
		AppUnifiedOrder(ctx context.Context, in *UnifiedOrderReq) (*UnifiedOrderResp, error)
		H5UnifiedOrder(ctx context.Context, in *UnifiedOrderReq) (*H5UnifiedOrderResp, error)
		JsUnifiedOrder(ctx context.Context, in *UnifiedOrderReq) (*UnifiedOrderResp, error)
		OrderQuery(ctx context.Context, in *OrderQueryReq) (*OrderQueryResp, error)
	}

	defaultPay struct {
		cli zrpc.Client
	}
)

func NewPay(cli zrpc.Client) Pay {
	return &defaultPay{
		cli: cli,
	}
}

func (m *defaultPay) AppUnifiedOrder(ctx context.Context, in *UnifiedOrderReq) (*UnifiedOrderResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.AppUnifiedOrder(ctx, in)
}

func (m *defaultPay) H5UnifiedOrder(ctx context.Context, in *UnifiedOrderReq) (*H5UnifiedOrderResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.H5UnifiedOrder(ctx, in)
}

func (m *defaultPay) JsUnifiedOrder(ctx context.Context, in *UnifiedOrderReq) (*UnifiedOrderResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.JsUnifiedOrder(ctx, in)
}

func (m *defaultPay) OrderQuery(ctx context.Context, in *OrderQueryReq) (*OrderQueryResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.OrderQuery(ctx, in)
}