// Package rpc provides ...
package rpc

import (
	"context"

	cmd_demo "github.com/deepzz0/appdemo/api/cmd-demo"
)

// UserSrv user server
type UserSrv struct {
	cmd_demo.UnimplementedUserServer
}

// UserInfo implements user server
func (s *UserSrv) UserInfo(ctx context.Context, in *cmd_demo.UserInfoReq) (*cmd_demo.UserInfoResp, error) {

	return nil, nil
}
