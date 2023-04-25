package codes

const (
	MysqlConnection = `package database

import (
	"database/sql"
<<<<<<< HEAD
	"{{.PackageName}}/infrastructure/environments"
=======
	"{{.PackageName}}/pkg/environments"
>>>>>>> edit/utils
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlConnect() (*sql.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", environments.Username, environments.Password, environments.Db_url, environments.Db_port, environments.Db_name)
	db, err := sql.Open(environments.Driver, dataSource)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(30 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(50)

	return db, nil
}
`
)
