user-proto:
	protoc --proto_path=${PWD}/user  --go_out=${PWD}/user   --go-grpc_out=${PWD}/user    -I ${PWD}/user user.proto

schedule-tracking-proto:
	protoc --proto_path=${PWD}/schedule-tracking   --go_out=${PWD}/schedule-tracking    --go-grpc_out=${PWD}/schedule-tracking    -I ${PWD}/schedule-tracking  schedule_tracking.proto

freight-proto:
	protoc --proto_path=${PWD}/freight   --go_out=${PWD}/freight    --go-grpc_out=${PWD}/freight    -I ${PWD}/freight  freight.proto

tracking-proto:
	PROTOC_GEN_TS_PATH="node_modules/.bin/protoc-gen-ts" \
    protoc --proto_path=${PWD}/tracking  --go_out=${PWD}/tracking --go-grpc_out=${PWD}/tracking --js_out="import_style=commonjs,binary:${PWD}/tracking/" --ts_out="${PWD}/tracking" -I ${PWD}/tracking tracking.proto \

proto: user-proto schedule-tracking-proto tracking-proto