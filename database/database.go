package database

import (
	"database/sql"
	"fmt"
)
var db *sql.DB
var Queries map[string]*sql.Stmt

func Connect (pg_user, pg_pass, pg_base string ) error {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", pg_user, pg_pass, pg_base))
	if err != nil {
		return err
	}
	return err
}

func Insert_football(val float32) error  {
	_, err := db.Exec(`DELETE FROM football`)
	if err != nil {
		return err
	}
	_, error := db.Exec(`INSERT INTO football (value) VALUES($1)`,val)
	if error != nil {
		return error
	}
	return nil
}

func Insert_baseball(val float32) error  {
	_, err := db.Exec(`DELETE FROM baseball`)
	if err != nil {
		return err
	}
	_, error := db.Exec(`INSERT INTO baseball (value) VALUES($1)`,val)
	if error != nil {
		return error
	}
	return nil
}

func Insert_soccer(val float32) error  {
	_, err := db.Exec(`DELETE FROM soccer`)
	if err != nil {
		return err
	}
	_, error := db.Exec(`INSERT INTO soccer (value) VALUES($1)`,val)
	if error != nil {
		return error
	}
	return nil
}

func Select_sport(sport string) float32  {
	Queries = make(map[string]*sql.Stmt)
	Queries["football"], _ = db.Prepare(`SELECT "value" FROM "football"`)
	Queries["baseball"], _ = db.Prepare(`SELECT "value" FROM "baseball"`)
	Queries["soccer"], _ = db.Prepare(`SELECT "value" FROM "soccer"`)
	stmt, _ := Queries[sport]
	val := stmt.QueryRow()
	var m float32
	val.Scan(&m)
	return m
}



