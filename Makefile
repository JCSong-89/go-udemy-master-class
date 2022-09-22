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
migratelastes:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1
migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
migraterevert:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run cmd/server/main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/JCSong-89/go-udemy-master-class/db/sqlc Store
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server migratelastes migraterevert mock