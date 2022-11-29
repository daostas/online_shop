.PHONY: all prod-svc user-svc setting-svc api-gw proto run stop debug-api
proto:
	protoc  --proto_path=prod-svc/pb/ --go_out=./prod-svc --go-grpc_out=./prod-svc/prod_service.proto
	protoc  --proto_path=user-svc/pb/ --go_out=./user-svc --go-grpc_out=./user-svc/user_service.proto
	protoc  --proto_path=setting-svc/pb/ --go_out=./setting-svc --go-grpc_out=./setting-svc/setting_service.proto

prod-svc:
	 go build -o ./prod-svc/prod-svc ./prod-svc/main.go

user-svc:
	 go build -o ./user-svc/user-svc ./user-svc/main.go

setting-svc:
	 go build -o ./setting-svc/setting-svc ./setting-svc/main.go


api-gw:
	 go build -o ./api-gw/api-shop ./api-gw/main.go

all: prod-svc user-svc setting-svc api-gw

run:
	./prod-svc/prod-svc &
	./user-svc/user-svc &
	./setting-svc/setting-svc &
	./api-gw/api-shop &

debug-api:
	 dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient ./api-gw/main.go --check-go-version=false

stop:
	killall -q -w -s SIGTERM api-shop user-svc prod-svc setting-svc