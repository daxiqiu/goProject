postgres:
	docker run --name postgres  -e POSTGRES_USER=root  -e POSTGRES_PASSWORD=daxi1982 -p 5432:5432 -d postgres:alpine

createdb: 
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:daxi1982@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:daxi1982@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:daxi1982@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:daxi1982@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb  -destination db/mock/store.go goProject/db/sqlc Store

.PHONY:	postgres	createdb	dropdb migrateup migratedown sqlc test server mock migrateup1 migratedown1