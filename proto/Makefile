# yarn 安装的protoc-gen-ts_proto文件路径
PROTOC_GEN_TS_PATH="/Users/even/node_modules/ts-proto/protoc-gen-ts_proto"

# 编译后的文件存放路径
TS_OUT_DIR="../client/assets/script/pb"

# 清理生成ts的输出文件目录
clear_ts:
	rm -rf ${TS_OUT_DIR}/*

# 生成ts的protobuf文件
gen_ts:
	protoc \
	--plugin=${PROTOC_GEN_TS_PATH} \
	--ts_proto_out=${TS_OUT_DIR} \
	./*.proto

# 生成所有的protobuf文件
gen:
	make clear_ts 
	make gen_ts	
	