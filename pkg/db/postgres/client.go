package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func NewClient(host, port, user, password, dbname string) (*sql.DB, error) {
	psqlConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// open database
	db, err := sql.Open("postgres", psqlConnectionString)
	if err != nil {
		return nil, err
	}
	//defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	//CheckError(err)
	return db, nil
}
