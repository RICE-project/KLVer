package mysql

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func Connect(dbname string, username string, password string, host string = "localhost", port string = "3306") (*DB, error){
        dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
        db, err := sql.Open("mysql", dsn)
        return db, err
}
