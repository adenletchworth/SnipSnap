package db

import (
	"database/sql"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "./snippets.db")

	if err != nil {
		log.Fatal(err)
	}

}
