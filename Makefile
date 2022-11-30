user-proto:
	protoc --proto_path=${PWD}/protobuf/user  --go_out=${PWD}/protobuf/user   --go-grpc_out=${PWD}/protobuf/user    -I ${PWD}/protobuf/user user.proto

schedule-tracking-proto:
	protoc --proto_path=${PWD}/protobuf/schedule-tracking   --go_out=${PWD}/protobuf/schedule-tracking    --go-grpc_out=${PWD}/protobuf/schedule-tracking    -I ${PWD}/protobuf/schedule-tracking  schedule_tracking.proto

freight-proto:
	protoc --proto_path=${PWD}/protobuf/freight   --go_out=${PWD}/protobuf/freight    --go-grpc_out=${PWD}/protobuf/freight    -I ${PWD}/protobuf/freight  freight.proto

tracking-proto:
	PROTOC_GEN_TS_PATH="node_modules/.bin/protoc-gen-ts" \
    protoc --proto_path=${PWD}/tracking  --go_out=${PWD}/tracking --go-grpc_out=${PWD}/tracking --js_out="import_style=commonjs,binary:${PWD}/tracking/" --ts_out="${PWD}/tracking" -I ${PWD}/tracking tracking.proto \
    STRING1='import * as grpc from "grpc";' STRING2='import * as grpc from "@grpc/grpc-js";'  awk 'NR==7 { sub(${STRING1}, ${STRING2}) }'

