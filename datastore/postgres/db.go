package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// db details
const (
	postgres_host     = "dpg-d81urjmk1jcs73dt5700-a.singapore-postgres.render.com"
	postgres_port     = 5432
	postgres_user     = "postgres_admin"
	postgres_password = "aOONHfHIvqZew3QyPYIcKS5xAroqPfas"
	postgres_dbname   = "my_db_h7vk"
)

// create pointer variable Db which points to sql driver
var Db *sql.DB

// init() is always called before main() by Go compiler
func init() {
	// creating a database connection string
	db_info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		postgres_host, postgres_port, postgres_user, postgres_password, postgres_dbname)
	// fmt.Println(db_info)
	var err error
	// establish a connection to postgresql server using the driver
	Db, err = sql.Open("postgres", db_info)
	// handle error
	if err != nil {
		panic(err)
	} else {
		log.Println("Database successfully connected")
	}
}
