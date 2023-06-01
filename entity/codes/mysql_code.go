package codes

const (
	MysqlConnection = `package database

import (
	"database/sql"
	"fmt"
	"time"

	"{{.PackageName}}/pkg/config"
	_ "github.com/go-sql-driver/mysql"
)
	
func MysqlConnect() (*sql.DB, error) {
	// Load config .env
	config.LoadEnv()
	
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Username, config.Password, config.Db_url, config.Db_port, config.Db_name)
	db, err := sql.Open(config.Driver, dataSource)
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
