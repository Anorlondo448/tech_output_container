package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Member is test table
type Member struct {
	ID      int
	Name    string
	Point   int
	Company string
}

func getMember() {
	mysqlEndpoint := os.Getenv("MYSQL_ENDPOINT")

	db, err := sql.Open("mysql", mysqlEndpoint)
	log.Println("Connected to mysql.")

	//接続でエラーが発生した場合の処理
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//データベースへクエリを送信。引っ張ってきたデータがrowsに入る。
	rows, err := db.Query("SELECT * FROM member")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	//レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
	for rows.Next() {
		var member Member
		err := rows.Scan(&member.ID, &member.Name, &member.Point, &member.Company)

		if err != nil {
			panic(err.Error())
		}
		fmt.Println(member.ID, member.Name, member.Point, member.Company)
	}
}

func topHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "app-top")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "app-pong")
}

func main() {
	http.HandleFunc("/", topHandler)
	http.HandleFunc("/ping", pingHandler)
	http.ListenAndServe(":8081", nil)
}
