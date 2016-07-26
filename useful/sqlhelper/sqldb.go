package main

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
)

const (
	Conn string = "sa:abcd-1234@tcp(localhost:3306)/test?charset=utf8"
)
