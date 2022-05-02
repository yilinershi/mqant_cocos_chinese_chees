package server_http

import (
	"context"
	"fmt"
	"github.com/topfreegames/pitaya/logger"
	"server/pb/http/pb_login"
)

type LoginServer struct {
}

func (l *LoginServer) LoginHandle(ctx context.Context, req *pb_login.LoginReq) (*pb_login.LoginResp, error) {
	logger.Log.Infof("req=%s\n", req.IMEI)
	return &pb_login.LoginResp{
		Name: fmt.Sprintf("Hello, %s", req.IMEI),
	}, nil
}
