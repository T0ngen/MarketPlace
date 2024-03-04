postgres:
	docker run --name postgresmarkteplace -p 5435:5432  -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -e POSTGRES_DB=marketplace -d postgres

createdb:
	docker exec -it postgresmarkteplace createdb --username=root --owner=root market

dropdb:
	docker exec -it postgresmarkteplace dropdb market

migrateup:
	migrate -path database/migrations -database "postgresql://root:1234@localhost:5435/market?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migrations -database "postgresql://root:1234@localhost:5435/market?sslmode=disable" -verbose down

.PHONY: test
test:
	go test ./database/sqlc

