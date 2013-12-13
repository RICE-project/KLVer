package db

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func Connect(dsn string) (*sql.DB, error){
        db, err := sql.Open("mysql", dsn)
        return db, err
}
