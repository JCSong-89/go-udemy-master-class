DB_URL=postgresql://postgres:root1234@127.0.0.1:5433/bank?sslmode=disable
network:
	docker network create bank-network
postgres:
	docker run -d -p 5432:5432 --name pgsql  -it -v  pgdate:/var/lib/postgresql/data -e POSTGRES_PASSWORD=1234 postgres
createdb:
	docker exec -it postgres createdb --username=root --owner=root bank
dropdb:
	docker exec -it postgres dropdb --bank
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up
migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test