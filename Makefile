go-install:
	@echo "Installing Go tools"
	go install github.com/google/wire/cmd/wire@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/ktr0731/evans@latest
	go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/air-verse/air@latest
	@echo "Go tools installed."
grpc-gen:
	protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto
migrate-create:
	migrate create -ext=sql -dir=internal/infra/database/migration -seq init  
migrate-mysql-up:
	go install -tags "mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	migrate -path=internal/infra/database/migration -database "mysql://root:root@tcp(mysql:3306)/orders" -verbose up	
PHONY: go-install grpc-gen migrate-create migrate-mysql-up