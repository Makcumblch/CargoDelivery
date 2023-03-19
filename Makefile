postgres-run:
	docker run --name=cargo-delivery-db -e POSTGRES_PASSWORD=$(password) -p 5432:5432 -d --rm postgres

migrate-create:
	migrate create -ext sql -dir ./schema -seq init

migrate-up:
	migrate -path ./schema -database 'postgres://postgres:admin@localhost:5432/postgres?sslmode=disable' up

migrate-down:
	migrate -path ./schema -database 'postgres://postgres:admin@localhost:5432/postgres?sslmode=disable' down

run:
	go run cmd/main.go