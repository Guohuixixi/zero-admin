// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package syslogservice

import (
	"context"

	"zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ConfigAddReq           = sysclient.ConfigAddReq
	ConfigAddResp          = sysclient.ConfigAddResp
	ConfigDeleteReq        = sysclient.ConfigDeleteReq
	ConfigDeleteResp       = sysclient.ConfigDeleteResp
	ConfigListData         = sysclient.ConfigListData
	ConfigListReq          = sysclient.ConfigListReq
	ConfigListResp         = sysclient.ConfigListResp
	ConfigUpdateReq        = sysclient.ConfigUpdateReq
	ConfigUpdateResp       = sysclient.ConfigUpdateResp
	DeptAddReq             = sysclient.DeptAddReq
	DeptAddResp            = sysclient.DeptAddResp
	DeptDeleteReq          = sysclient.DeptDeleteReq
	DeptDeleteResp         = sysclient.DeptDeleteResp
	DeptListData           = sysclient.DeptListData
	DeptListReq            = sysclient.DeptListReq
	DeptListResp           = sysclient.DeptListResp
	DeptUpdateReq          = sysclient.DeptUpdateReq
	DeptUpdateResp         = sysclient.DeptUpdateResp
	DictAddReq             = sysclient.DictAddReq
	DictAddResp            = sysclient.DictAddResp
	DictDeleteReq          = sysclient.DictDeleteReq
	DictDeleteResp         = sysclient.DictDeleteResp
	DictListData           = sysclient.DictListData
	DictListReq            = sysclient.DictListReq
	DictListResp           = sysclient.DictListResp
	DictUpdateReq          = sysclient.DictUpdateReq
	DictUpdateResp         = sysclient.DictUpdateResp
	InfoReq                = sysclient.InfoReq
	InfoResp               = sysclient.InfoResp
	JobAddReq              = sysclient.JobAddReq
	JobAddResp             = sysclient.JobAddResp
	JobDeleteReq           = sysclient.JobDeleteReq
	JobDeleteResp          = sysclient.JobDeleteResp
	JobListData            = sysclient.JobListData
	JobListReq             = sysclient.JobListReq
	JobListResp            = sysclient.JobListResp
	JobUpdateReq           = sysclient.JobUpdateReq
	JobUpdateResp          = sysclient.JobUpdateResp
	LoginLogAddReq         = sysclient.LoginLogAddReq
	LoginLogAddResp        = sysclient.LoginLogAddResp
	LoginLogDeleteReq      = sysclient.LoginLogDeleteReq
	LoginLogDeleteResp     = sysclient.LoginLogDeleteResp
	LoginLogListData       = sysclient.LoginLogListData
	LoginLogListReq        = sysclient.LoginLogListReq
	LoginLogListResp       = sysclient.LoginLogListResp
	LoginReq               = sysclient.LoginReq
	LoginResp              = sysclient.LoginResp
	MenuAddReq             = sysclient.MenuAddReq
	MenuAddResp            = sysclient.MenuAddResp
	MenuDeleteReq          = sysclient.MenuDeleteReq
	MenuDeleteResp         = sysclient.MenuDeleteResp
	MenuListData           = sysclient.MenuListData
	MenuListReq            = sysclient.MenuListReq
	MenuListResp           = sysclient.MenuListResp
	MenuListTree           = sysclient.MenuListTree
	MenuUpdateReq          = sysclient.MenuUpdateReq
	MenuUpdateResp         = sysclient.MenuUpdateResp
	QueryMenuByRoleIdReq   = sysclient.QueryMenuByRoleIdReq
	QueryMenuByRoleIdResp  = sysclient.QueryMenuByRoleIdResp
	ReSetPasswordReq       = sysclient.ReSetPasswordReq
	ReSetPasswordResp      = sysclient.ReSetPasswordResp
	RoleAddReq             = sysclient.RoleAddReq
	RoleAddResp            = sysclient.RoleAddResp
	RoleDeleteReq          = sysclient.RoleDeleteReq
	RoleDeleteResp         = sysclient.RoleDeleteResp
	RoleListData           = sysclient.RoleListData
	RoleListReq            = sysclient.RoleListReq
	RoleListResp           = sysclient.RoleListResp
	RoleUpdateReq          = sysclient.RoleUpdateReq
	RoleUpdateResp         = sysclient.RoleUpdateResp
	StatisticsLoginLogReq  = sysclient.StatisticsLoginLogReq
	StatisticsLoginLogResp = sysclient.StatisticsLoginLogResp
	SysLogAddReq           = sysclient.SysLogAddReq
	SysLogAddResp          = sysclient.SysLogAddResp
	SysLogDeleteReq        = sysclient.SysLogDeleteReq
	SysLogDeleteResp       = sysclient.SysLogDeleteResp
	SysLogListData         = sysclient.SysLogListData
	SysLogListReq          = sysclient.SysLogListReq
	SysLogListResp         = sysclient.SysLogListResp
	UpdateConfigRoleReq    = sysclient.UpdateConfigRoleReq
	UpdateConfigRoleResp   = sysclient.UpdateConfigRoleResp
	UpdateMenuRoleReq      = sysclient.UpdateMenuRoleReq
	UpdateMenuRoleResp     = sysclient.UpdateMenuRoleResp
	UserAddReq             = sysclient.UserAddReq
	UserAddResp            = sysclient.UserAddResp
	UserDeleteReq          = sysclient.UserDeleteReq
	UserDeleteResp         = sysclient.UserDeleteResp
	UserListData           = sysclient.UserListData
	UserListReq            = sysclient.UserListReq
	UserListResp           = sysclient.UserListResp
	UserStatusReq          = sysclient.UserStatusReq
	UserStatusResp         = sysclient.UserStatusResp
	UserUpdateReq          = sysclient.UserUpdateReq
	UserUpdateResp         = sysclient.UserUpdateResp

	SysLogService interface {
		SysLogAdd(ctx context.Context, in *SysLogAddReq, opts ...grpc.CallOption) (*SysLogAddResp, error)
		SysLogList(ctx context.Context, in *SysLogListReq, opts ...grpc.CallOption) (*SysLogListResp, error)
		SysLogDelete(ctx context.Context, in *SysLogDeleteReq, opts ...grpc.CallOption) (*SysLogDeleteResp, error)
	}

	defaultSysLogService struct {
		cli zrpc.Client
	}
)

func NewSysLogService(cli zrpc.Client) SysLogService {
	return &defaultSysLogService{
		cli: cli,
	}
}

func (m *defaultSysLogService) SysLogAdd(ctx context.Context, in *SysLogAddReq, opts ...grpc.CallOption) (*SysLogAddResp, error) {
	client := sysclient.NewSysLogServiceClient(m.cli.Conn())
	return client.SysLogAdd(ctx, in, opts...)
}

func (m *defaultSysLogService) SysLogList(ctx context.Context, in *SysLogListReq, opts ...grpc.CallOption) (*SysLogListResp, error) {
	client := sysclient.NewSysLogServiceClient(m.cli.Conn())
	return client.SysLogList(ctx, in, opts...)
}

func (m *defaultSysLogService) SysLogDelete(ctx context.Context, in *SysLogDeleteReq, opts ...grpc.CallOption) (*SysLogDeleteResp, error) {
	client := sysclient.NewSysLogServiceClient(m.cli.Conn())
	return client.SysLogDelete(ctx, in, opts...)
}
