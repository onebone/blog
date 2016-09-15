package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

var provider *SQLiteProvider

type Provider interface {
	init()
	Register(string, string, string) bool
	SignIn(string, string) bool
	SignOut(string) bool
}

type SQLiteProvider struct {
	db *sql.DB
}

func (s *SQLiteProvider) init() {
	var err error
	s.db, err = sql.Open("sqlite3", "blog.db")
	checkErr(err)

	stmt, err := s.db.Prepare(`CREATE TABLE IF NOT EXISTS accounts (
		username TEXT,
		password TEXT,
		email TEXT,
		registered INTEGER,
		lastLogin INTEGER,
		alerts TEXT
		)`)
	checkErr(err)

	_, err = stmt.Exec()
	checkErr(err)
}

func (s *SQLiteProvider) Register(username, password, email string) bool {
	stmt, err := s.db.Prepare("INSERT INTO accounts (username, password, email, registered, lastLogin, alerts) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return false
	}

	_, err = stmt.Exec(username, password, email, time.Now().Unix(), time.Now().Unix(), "{}")
	if err != nil {
		return false
	}
	return true
}

func (s *SQLiteProvider) SignIn(username, password string) bool {
	return false
}

func (s *SQLiteProvider) SignOut(username string) bool {
	return false
}
