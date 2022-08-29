DB_URL=http://127.0.0.1:3000
network:
	docker network create bank-network
postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root1234 -e POSTGRES_DB=bank -d postgres:14-alpine
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


.PHONY: postgres createdb dropdb migrateup migratedown sqlc