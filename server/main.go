package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/acceptor"
	"github.com/topfreegames/pitaya/cluster"
	"github.com/topfreegames/pitaya/component"
	"github.com/topfreegames/pitaya/constants"
	"github.com/topfreegames/pitaya/modules"
	"github.com/topfreegames/pitaya/serialize/protobuf"
	"server/module_db"
	"server/server_http"
)

//
//func configureLobby() {
//	componentLobby := server_lobby.NewComponentLobby()
//	pitaya.Register(componentLobby, component.WithName("ComponentLobby"))
//	pitaya.RegisterRemote(componentLobby, component.WithName("ComponentLobby"))
//}
//
//func configureBjl() {
//	componentLobby := server_bjl.NewComponentBjl()
//	pitaya.Register(componentLobby, component.WithName("ComponentBjl"))
//	pitaya.RegisterRemote(componentLobby, component.WithName("ComponentBjl"))
//}
//
func configureHttpSever() {
	componentHttp := server_http.NewComponentHttp()
	pitaya.Register(componentHttp, component.WithName("ComponentHttp"))
	pitaya.RegisterRemote(componentHttp, component.WithName("ComponentHttp"))
}

//配置前端服务器
func configureFrontend(port int) {
	tcp := acceptor.NewTCPAcceptor(fmt.Sprintf(":%d", port))
	pitaya.AddAcceptor(tcp)
}

func main() {
	frontedPort := flag.Int("port", 3250, "the port to listen")
	serverType := flag.String("type", "connector", "the server type")
	isFrontend := flag.Bool("frontend", true, "if server is frontend")
	flag.Parse()
	defer pitaya.Shutdown()
	pitaya.SetSerializer(protobuf.NewSerializer())
	switch *serverType {
	case "ServerConnector": //前端服务器
		configureFrontend(*frontedPort)
		break
	case "ServerHttp": //http服务器
		configureHttpSever()
		break
	case "ServerChineseChess": //中国象棋服务器

		break

	default:
		fmt.Printf("error serverType = %s\n", *serverType)
		return
	}
	configs := setConfig()

	meta := map[string]string{
		constants.GRPCHostKey: "127.0.0.1", //正式项目中，需要设置grpc的host
		constants.GRPCPortKey: "3434",      //正式项目中，需要设置grpc的port
	}

	pitaya.Configure(*isFrontend, *serverType, pitaya.Cluster, meta, configs)

	registerModule()

	//使用nats的rpc-server
	gs, err := cluster.NewNatsRPCServer(pitaya.GetConfig(), pitaya.GetServer(), pitaya.GetMetricsReporters(), pitaya.GetDieChan())
	if err != nil {
		panic(err)
	}
	pitaya.SetRPCServer(gs)

	//使用nats的rpc-client
	gc, err := cluster.NewNatsRPCClient(
		pitaya.GetConfig(),
		pitaya.GetServer(),
		pitaya.GetMetricsReporters(),
		pitaya.GetDieChan(),
	)
	if err != nil {
		panic(err)
	}

	pitaya.SetRPCClient(gc)
	pitaya.Start()
}

//使用viper这个库，配置该游戏的一些配置项
func setConfig() *viper.Viper {
	//具体默认配置见：https://pitaya.readthedocs.io/en/latest/configuration.html
	configs := viper.New()

	//1.redis各项配置
	//1.1 redis的host
	configs.Set("pitaya.modules.redis.default.client.host", "127.0.0.1")
	//1.2 redis的port
	configs.Set("pitaya.modules.redis.default.client.port", 6379)
	//1.3 redis的哪个数据库，这里使用默认0的
	configs.Set("pitaya.modules.redis.default.client.db", 0)
	//1.4 redis的重试次数
	configs.Set("pitaya.modules.redis.default.client.retry", 5)
	//1.5 redis的size,用于初始化redis
	configs.Set("pitaya.modules.redis.default.client.size", 500)
	//1.6
	configs.Set("pitaya.modules.redis.default.client.idle", 10)

	//2.grpc的配置
	//2.1.grpc的端口（实现上在默认的设置中，端口即是3434，见上面的配置地址）
	configs.Set("pitaya.cluster.rpc.server.grpc.port", 3434)
	return configs
}

func registerModule() {
	//绑定etcd的模块
	bs := modules.NewETCDBindingStorage(pitaya.GetServer(), pitaya.GetConfig())
	pitaya.RegisterModule(bs, "bindingsStorage")

	//redis模块
	redisModule := module_db.NewRedisModule()
	pitaya.RegisterModule(redisModule, "redisModule")
}
