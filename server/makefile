## 启动前端服务器
run_server_frontend:
	@go run main.go --port=3250 --type=ServerConnector --frontend=true

## 启动后端服务器之中国象棋服务器
run_server_chinese_chess:
	@go run main.go --type=ServerChineseChess --frontend=false

## 启动后端服务器之http服务器
run_server_http:
	@go run main.go --type=ServerHttp --frontend=false

## 编译proto,需要安装的插件如下
# 插件1:protoc,参考：
# 插件2：protoc-gen-gohttp,参考：https://github.com/nametake/protoc-gen-gohttp
build_proto:
	@echo "step1:先清掉pb文件夹"
	@rm -rf pb
	@mkdir pb
	@echo "step2:再生成pb文件"
	@protoc --go_out=../ -I=../proto  ../proto/common/*.proto
	@protoc --go_out=plugins=grpc:../ --gohttp_out=../ -I=../proto ../proto/http/*.proto

