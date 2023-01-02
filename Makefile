user-proto:
	protoc --proto_path=${PWD} --go_out=${PWD}/user --go-grpc_out=${PWD}/user  --grpc-gateway_out=${PWD}/user  --grpc-gateway_opt allow_delete_body=true --grpc-gateway_opt generate_unbound_methods=true --openapiv2_out=allow_delete_body=true:${PWD}/user -I ${PWD}/user user.proto

schedule-tracking-proto:
	protoc --proto_path=${PWD} --go_out=${PWD}/schedule-tracking --go-grpc_out=${PWD}/schedule-tracking  --grpc-gateway_out=${PWD}/schedule-tracking  --grpc-gateway_opt allow_delete_body=true --grpc-gateway_opt generate_unbound_methods=true --openapiv2_out=allow_delete_body=true:${PWD}/schedule-tracking -I ${PWD}/schedule-tracking schedule_tracking.proto

freight-proto:
	protoc --proto_path=${PWD}/freight   --go_out=${PWD}/freight    --go-grpc_out=${PWD}/freight    -I ${PWD}/freight  freight.proto

tracking-proto:
	protoc --proto_path=${PWD} --go_out=${PWD}/tracking --go-grpc_out=${PWD}/tracking  --grpc-gateway_out=${PWD}/tracking  --grpc-gateway_opt allow_delete_body=true --grpc-gateway_opt generate_unbound_methods=true --openapiv2_out=allow_delete_body=true:${PWD}/tracking -I ${PWD}/tracking tracking.proto


proto: user-proto schedule-tracking-proto tracking-proto