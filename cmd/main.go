package main

import (
	"database/sql"
	"github.com/JCSong-89/go-udemy-master-class/api"
	db "github.com/JCSong-89/go-udemy-master-class/db/sqlc"
	"log"
)

const (
	dbDriver      = "postgres"
	dbDSN         = "postgres://postgres:root1234@localhost:5433/bank?sslmode=disable"
	serverAddress = ":3030"
)

func main() {
	// DB 커넥션
	conn, err := sql.Open(dbDriver, dbDSN)
	if err != nil {
		log.Fatal(err)
	}

	//서버 인스턴스 생성
	store := db.NewStore(conn)
	server := api.NewServer(store)

	//서버 시작
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal(err)
	}
}
