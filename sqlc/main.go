package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"business/sqlc/model"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.ExpandEnv("${MYSQL_USER}"), os.ExpandEnv("${MYSQL_PASSWORD}"),
		os.ExpandEnv("${DB_HOST}"), os.ExpandEnv("${DB_PORT}"),
		os.ExpandEnv("${MYSQL_DATABASE}"))

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// ここが修正ポイント
	queries := model.New(conn)
	ctx := context.Background()
	rows, err := queries.GetPostWithNestedReplies(ctx)
	if err != nil {
		// エラーハンドリング
	}

	for _, user := range rows {
		fmt.Printf("%v \n", user)
	}

}
