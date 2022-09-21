package main

import (
	"database/sql"
	"github.com/JCSong-89/go-udemy-master-class/api"
	db "github.com/JCSong-89/go-udemy-master-class/db/sqlc"
	"github.com/JCSong-89/go-udemy-master-class/utill"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := utill.LoadConfig("../.")

	// DB 커넥션
	conn, err := sql.Open(config.DBdDiver, config.DBDNS)
	if err != nil {
		log.Fatal(err)
	}

	//서버 인스턴스 생성
	store := db.NewStore(conn)
	server := api.NewServer(store)

	//서버 시작
	err = server.Start(config.ServerHost)
	if err != nil {
		log.Fatal(err)
	}
}
