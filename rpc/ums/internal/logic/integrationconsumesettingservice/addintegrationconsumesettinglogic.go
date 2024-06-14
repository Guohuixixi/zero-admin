package integrationconsumesettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddIntegrationConsumeSettingLogic 添加积分消费设置
/*
Author: LiuFeiHua
Date: 2024/6/11 14:21
*/
type AddIntegrationConsumeSettingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddIntegrationConsumeSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddIntegrationConsumeSettingLogic {
	return &AddIntegrationConsumeSettingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddIntegrationConsumeSetting 添加积分消费设置
func (l *AddIntegrationConsumeSettingLogic) AddIntegrationConsumeSetting(in *umsclient.AddIntegrationConsumeSettingReq) (*umsclient.AddIntegrationConsumeSettingResp, error) {
	err := query.UmsIntegrationConsumeSetting.WithContext(l.ctx).Create(&model.UmsIntegrationConsumeSetting{
		DeductionPerAmount: in.DeductionPerAmount,
		MaxPercentPerOrder: in.MaxPercentPerOrder,
		UseUnit:            in.UseUnit,
		CouponStatus:       in.CouponStatus,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.AddIntegrationConsumeSettingResp{}, nil
}