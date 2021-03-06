.PHONY: migrate_localhost_up
migrate_localhost_up:
	migrate -database postgres://postgres:postgres@localhost:5432/vk_mini_app?sslmode=disable -path ./internal/db/migrations up

.PHONY: migrate_localhost_down
migrate_localhost_down:
	migrate -database postgres://postgres:postgres@localhost:5432/vk_mini_app?sslmode=disable -path ./internal/db/migrations down

.PHONY: build_localhost
build_localhost:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=mod -o vk_mini_app main.go

.PHONY: run_localhost
run_localhost:
	go run main.go