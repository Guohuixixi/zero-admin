package memberreportservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMemberReportStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberReportStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberReportStatusLogic {
	return &UpdateMemberReportStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户举报表状态
func (l *UpdateMemberReportStatusLogic) UpdateMemberReportStatus(in *cmsclient.UpdateMemberReportStatusReq) (*cmsclient.UpdateMemberReportStatusResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.UpdateMemberReportStatusResp{}, nil
}
