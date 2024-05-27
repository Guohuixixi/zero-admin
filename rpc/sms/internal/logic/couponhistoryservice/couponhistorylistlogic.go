package couponhistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// CouponHistoryListLogic
/*
Author: LiuFeiHua
Date: 2023/11/29 11:33
*/
type CouponHistoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponHistoryListLogic {
	return &CouponHistoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CouponHistoryList 根据优惠券id，使用状态，订单编号分页获取领取记录
func (l *CouponHistoryListLogic) CouponHistoryList(in *smsclient.CouponHistoryListReq) (*smsclient.CouponHistoryListResp, error) {
	q := query.SmsCouponHistory.WithContext(l.ctx)

	if in.CouponId != 0 {
		q = q.Where(query.SmsCouponHistory.CouponID.Eq(in.CouponId))
	}
	if in.MemberId != 0 {
		q = q.Where(query.SmsCouponHistory.MemberID.Eq(in.MemberId))
	}
	if len(in.OrderSn) > 0 {
		q = q.Where(query.SmsCouponHistory.OrderSn.Eq(in.OrderSn))
	}
	if in.UseStatus != 3 {
		q = q.Where(query.SmsCouponHistory.UseStatus.Eq(in.UseStatus))
	}

	result, count, err := q.FindByPage(int((in.Current-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "根据优惠券id，使用状态，订单编号分页获取领取记录失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*smsclient.CouponHistoryListData
	for _, item := range result {

		list = append(list, &smsclient.CouponHistoryListData{
			Id:             item.ID,
			CouponId:       item.CouponID,
			MemberId:       item.MemberID,
			CouponCode:     item.CouponCode,
			MemberNickname: item.MemberNickname,
			GetType:        item.GetType,
			CreateTime:     item.CreateTime.Format("2006-01-02 15:04:05"),
			UseStatus:      item.UseStatus,
			UseTime:        item.UseTime.Format("2006-01-02 15:04:05"),
			OrderId:        item.OrderID,
			OrderSn:        item.OrderSn,
		})
	}

	logc.Infof(l.ctx, "根据优惠券id，使用状态，订单编号分页获取领取记录,参数：%+v,响应：%+v", in, list)
	return &smsclient.CouponHistoryListResp{
		Total: count,
		List:  list,
	}, nil
}
