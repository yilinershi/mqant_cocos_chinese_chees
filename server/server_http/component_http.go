package server_http

import (
	"bytes"
	"context"
	"fmt"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/component"
	"github.com/topfreegames/pitaya/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"server/module_db"
	"server/pb/http/pb_login"
	"strings"
)

//持有redis的module
var redisModule *module_db.RedisModule

//ComponentHttp 专门处理http的相关服务
type ComponentHttp struct {
	component.Base
}

func NewComponentHttp() *ComponentHttp {
	return &ComponentHttp{}
}

func (this *ComponentHttp) Init() {
	this.initRedis()
	this.initLoginServer()
}

func (this *ComponentHttp) initRedis() {
	if module, err := pitaya.GetModule("redisModule"); err == nil {
		redisModule = module.(*module_db.RedisModule)
	}
}

//初始化登录服务
func (this *ComponentHttp) initLoginServer() {
	srv := &LoginServer{}
	conv := pb_login.NewLoginHTTPConverter(srv)
	http.Handle("/login", conv.LoginHandle(logCallback))
	http.Handle(restPath(conv.LoginHandleWithName(logCallback)))
	go func() {
		logger.Log.Fatal(http.ListenAndServe(":8080", nil))
	}()
}

func logCallback(ctx context.Context, w http.ResponseWriter, r *http.Request, req, resp proto.Message, err error) {
	logger.Log.Infof("INFO: call %s: req: {%v}, resp: {%s}", r.RequestURI, req, resp)
	// YOU MUST HANDLE ERROR
	if err != nil {
		logger.Log.Infof("ERROR: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		p := status.New(codes.Unknown, err.Error()).Proto()
		switch r.Header.Get("Content-Type") {
		case "application/protobuf", "application/x-protobuf":
			buf, err := proto.Marshal(p)
			if err != nil {
				return
			}
			if _, err := io.Copy(w, bytes.NewBuffer(buf)); err != nil {
				return
			}
		case "application/json":
			buf, err := protojson.Marshal(p)
			if err != nil {
				return
			}
			if _, err := io.Copy(w, bytes.NewBuffer(buf)); err != nil {
				return
			}
		default:
		}
	}
}

func restPath(service, method string, hf http.HandlerFunc) (string, http.HandlerFunc) {
	return fmt.Sprintf("/%s/%s", strings.ToLower(service), strings.ToLower(method)), hf
}
