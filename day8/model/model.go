package model

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func init() {
	newDb, err := sql.Open("sqlite3", "./day8/db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	_, err = newDb.Exec(`create table if not exists phones (
    	id integer primary key,
    	name varchar(30),
    	phone varchar(30)
	)`)
	if err != nil {
		log.Fatal(err)
	}
	db = newDb
}

func CloseDB() {
	db.Close()
}

func Create(name, phone string) int {
	res, err := db.Exec(`insert into phones (name, phone) values (?, ?)`, name, phone)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return int(id)
}

func ReadAll() [][]string {
	rows, err := db.Query(`select name, phone from phones;`)
	if err != nil {
		log.Fatal(err)
	}
	res := make([][]string, 0)
	for rows.Next() {
		c := make([]string, 2)
		rows.Scan(&c[0], &c[1])
		res = append(res, c)
	}
	return res
}

func Read(name string) (phone string) {
	rows, err := db.Query(`select phone from phones where name = ?`, name)
	if err != nil {
		log.Fatal(err)
	}
	if rows.Next() {
		rows.Scan(&phone)
	}
	return
}
