.PHONY: all admin-svc client-svc api-gw auth-svc proto run stop debug-api
proto:
	protoc  --proto_path=admin-svc/pb/ --go_out=./admin-svc --go-grpc_out=./admin-svc/admin_service.proto
	protoc  --proto_path=client-svc/pb/ --go_out=./client-svc --go-grpc_out=./client-svc/client_service.proto
	protoc  --proto_path=auth-svc/pb/ --go_out=./auth-svc --go-grpc_out=./auth-svc/auth_service.proto

auth-svc:
	 go build -o ./auth-svc/auth2-svc ./auth-svc/main.go

admin-svc:
	 go build -o ./admin-svc/admin-svc ./admin-svc/main.go

client-svc:
	 go build -o ./client-svc/client-svc ./client-svc/main.go

api-gw:
	 go build -o ./api-gw/api-shop ./api-gw/main.go

all: admin-svc client-svc api-gw auth-svc

run:
	./admin-svc/admin-svc &
	./client-svc/client-svc &
	./auth-svc/auth2-svc &
	./api-gw/api-shop &

debug-api:
	 dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient ./api-gw/main.go --check-go-version=false

stop:
	killall -q -w -s SIGTERM api-shop client-svc admin-svc auth2-svc

rerun: stop all run