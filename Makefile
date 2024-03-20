run:
	go run cmd/main.go

proto-gen:
	./scripts/gen-proto.sh

migrate-up:
	migrate -path migrations -database "postgres://postgres:1234@localhost:5432/userdb?sslmode=disable" -verbose up

migrate-down:
	migrate -path migrations -database "postgres://postgres:1234@localhost:5432/userdb?sslmode=disable" -verbose down

migrate_file:
	migrate create -ext sql -dir migrations/ -seq users

migrate-dirty:
	migrate -path ./migrations/ -database "postgresql://postgres:1234@localhost:5432/userdb?sslmode=disable" force 1