grpc-gen:
	protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto
migrate-create:
	migrate create -ext=sql -dir=internal/infra/database/migration -seq init  
migrate-mysql-up:
	go install -tags "mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	migrate -path=internal/infra/database/migration -database "mysql://root:root@tcp(mysql:3306)/orders" -verbose up
PHONY: migrate-create migrate-mysql-up grpc-gen