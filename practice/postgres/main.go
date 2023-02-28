package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {

	dbURL := "postgres://postgres:postgres@localhost:5432/exampledb"
	/*dbURL might look like:"postgres://username:password@localhost:5432/database_name"*/
	conn, err := sql.Open("pgx", dbURL)
	if err != nil {
		return
		fmt.Errorf("connect to db error: %s\n", err)
	} else {

		fmt.Printf("connect to db \n")

	}

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	if err := conn.PingContext(ctx); err != nil {
		return
	} else {
		fmt.Println("pong \n")
	}
	cancel()

}
