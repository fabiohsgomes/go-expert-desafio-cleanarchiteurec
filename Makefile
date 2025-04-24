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
	make go-tidy
docker-up:
	@echo "Starting Docker containers"
	docker-compose up -d
	@echo "Docker containers started."
go-tidy:
	@echo "Go installing dependencies"
	go mod tidy
	@echo "Go dependencies installed."
go-run:
	make go-tidy
	make docker-up
	cd cmd/ordersystem && go run main.go wire_gen.go
grpc-gen:
	protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto
migrate-create:
	migrate create -ext=sql -dir=internal/infra/database/migration -seq init  
migrate-mysql-up:
	make docker-up
	@echo "Executing migrations"
	go install -tags "mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	migrate -path=internal/infra/database/migration -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up	
	@echo "Migrations executed."	
PHONY: go-install docker-up go-tidy go-run grpc-gen migrate-create migrate-mysql-up