// Package user provides ...
package user

import (
	"context"

	cmd_demo "github.com/deepzz0/appdemo/pkg/proto/cmd-demo"
)

type UserSrv struct {
	cmd_demo.UnimplementedUserServer
}

func (s *UserSrv) UserInfo(ctx context.Context, in *cmd_demo.UserInfoReq) (*cmd_demo.UserInfoResp, error) {

	return nil, nil
}
