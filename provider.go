package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Provider interface {
	init()
	Register(string, string) bool
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

	s.db.Prepare(`CREATE TABLE IF NOT EXISTS accounts (
		username TEXT,
		password TEXT,
		alerts TEXT,
		)`)
}

func (s *SQLiteProvider) Register(username, password string) bool {
	stmt, err := s.db.Prepare("INSERT INTO accounts (username, password, alerts) VALUES (?, ?, ?)")
	if err != nil {
		return false
	}

	_, err = stmt.Exec(username, password, "")
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
