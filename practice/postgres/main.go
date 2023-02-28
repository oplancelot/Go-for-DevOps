package main

//https://medium.com/geekculture/work-with-go-postgresql-using-pgx-caee4573672#:~:text=pgx%20is%20a%20PostgreSQL%20driver%20and%20toolkit%20written,use%20pgx%20and%20look%20out%20for%20other%20libraries.
import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// create table public.person
// (
//
//	id            integer      not null
//	    constraint person_pk
//	        primary key,
//	first_name    varchar(100) not null,
//	last_name     varchar(100) not null,
//	date_of_birth date         not null
//
// );
type PersonRec struct {
	first_name string
	last_name  string
	id         int
}

func GetUser(ctx context.Context, conn *sql.DB, id int) (PersonRec, error) {
	const query = `SELECT "first_name","last_name" FROM users WHERE "id" = $1`
	u := PersonRec{id: id}
	err := conn.QueryRowContext(ctx, query, id).Scan(&u)
	return u, err
}

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

	_, err = GetUser(context.Background(), conn, 1)

	//使用pgxpool库

	// dbPool, err := pgxpool.Connect(context.Background(), dbURL)

	// if err != nil {
	// 	return
	// 	fmt.Errorf("connect to db error: %s\n", err)
	// }
	// defer dbPool.Close()
	// // execute the select query and get result rows
	// rows, err := dbPool.Query(context.Background(), "select * from public.person")
	// if err != nil {
	// 	log.Fatal("error while executing query")
	// }

	// // iterate through the rows
	// for rows.Next() {
	// 	values, err := rows.Values()
	// 	if err != nil {
	// 		log.Fatal("error while iterating dataset")
	// 	}

	// 	// convert DB types to Go types
	// 	id := values[0].(int32)
	// 	firstName := values[1].(string)
	// 	lastName := values[2].(string)
	// 	dateOfBirth := values[3].(time.Time)

	// 	log.Println("[id:", id, ", first_name:", firstName, ", last_name:", lastName, ", date_of_birth:", dateOfBirth, "]")
	// }

	// 	// id can be taken as a user input.
	// 	// for the time being, let's hard code it
	// 	id := 1

	// 	// execute the query and get result rows
	// 	rows2, err := dbPool.Query(context.Background(), "select * from public.get_person_details($1)", id)
	// 	log.Println("id: ", id)
	// 	if err != nil {
	// 		log.Fatal("error while executing query")
	// 	}

	// 	// iterate through the rows
	// 	for rows2.Next() {
	// 		values, err := rows2.Values()
	// 		if err != nil {
	// 			log.Fatal("error while iterating dataset")
	// 		}

	// 		// convert DB types to Go types
	// 		firstName := values[0].(string)
	// 		lastName := values[1].(string)
	// 		dateOfBirth := values[2].(time.Time)

	// 		log.Println("[first_name:", firstName, ", last_name:", lastName, ", date_of_birth:", dateOfBirth, "]")
	// 	}

}
