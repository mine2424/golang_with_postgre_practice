package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	// postgres接続情報
	conn = "host=postgres port=5432 user=postgres password=secret dbname=simplebank sslmode=disable"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func slowQuery() error {
	_, err := db.Exec("SELECT pg_sleep(5)")
	return err
}

func slowHandler(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	err := slowQuery()
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return
	}
	fmt.Fprintln(w, "OK")
	fmt.Printf("slowHandler took: %v\n", time.Since(start))
}

func main() {
	var err error

	// SQLに接続
	db, err = sql.Open("postgres", conn)
	logFatal(err)

	// 応答確認
	err = db.Ping()
	logFatal(err)

	// serverの基本設定
	srv := http.Server{
		Addr:         ":8080",
		WriteTimeout: 2 * time.Second,
		Handler:      http.HandlerFunc(slowHandler),
	}

	// serverスタート
	log.Println("Start Http Server...")
	log.Fatal(srv.ListenAndServe())
}
