DB_URL = postgresql://postgres:secret@localhost:5433/runsystem_be_test?sslmode=disable

postgres:
	docker run --name postgres -e POSTGRES_PASSWORD=secret -p 5433:5432 -d postgres

createdb:
	docker exec -it postgres createdb --username=postgres runsystem_be_test

dropdb:
	docker exec -it postgres dropdb runsystem_be_test

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc
