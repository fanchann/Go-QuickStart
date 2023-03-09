package entity

const (
	Env = `#HERE IS CONFIGURATION TO CONNECT DATABASE
DB_DRIVER = 
DB_AUTH_USERNAME = 
DB_AUTH_PASSWORD = 
DB_NAME = 
DB_URL = 
DB_PORT =`

	MysqlConnect = `package database

import (
	helper "{{.PackageName}}/insfrastructure/utils"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)
	
func MysqlConnect() (*sql.DB,error) {
	curentDir, err := os.Getwd()
	helper.LogErrorWithPanic(err)

	envDir := filepath.Join(curentDir, "insfrastructure", "environments")
	err = godotenv.Load(filepath.Join(envDir, ".env"))
	helper.LogErrorWithPanic(err)
	
	driver := os.Getenv("DB_DRIVER")
	username := os.Getenv("DB_AUTH_USERNAME")
	password := os.Getenv("DB_AUTH_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_url := os.Getenv("DB_URL")
	db_port := os.Getenv("DB_PORT")
	
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, db_url, db_port, db_name)
	
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(30 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(50)
	
	return db,nil
}
`

	CacthError = `package utils

import "log"
	
func LogErrorWithPanic(err error) {
	if err != nil {
		log.Panicf("error : %s", err)
	}
}
`

	GoMod = `module {{.PackageName}}

go 1.20

require (
	github.com/go-sql-driver/mysql v1.7.0
	github.com/joho/godotenv v1.5.1
)
`
	GoSum = `github.com/go-sql-driver/mysql v1.7.0 h1:ueSltNNllEqE3qcWBTD0iQd3IpL/6U+mJxLkazJ7YPc=
github.com/go-sql-driver/mysql v1.7.0/go.mod h1:OXbVy3sEdcQ2Doequ6Z5BW6fXNQTmx+9S1MCJN5yJMI=
github.com/joho/godotenv v1.5.1 h1:7eLL/+HRGLY0ldzfGMeQkb7vMd0as4CfYvUVzLqw0N0=
github.com/joho/godotenv v1.5.1/go.mod h1:f4LDr5Voq0i2e/R5DDNOoa2zzDfwtkZa6DnEwAbqwq4=
`

	README = `# Read This
if you have not installed _mysql_ driver use:
- _go get -u github.com/go-sql-driver/mysql_
install additional packages to load _.env_
- _go get github.com/joho/godotenv_
	
~ enjoy
`

	MainGo = `package main

import (
	"{{.PackageName}}/insfrastructure/database"
	"{{.PackageName}}/insfrastructure/utils"
	"fmt"
)
	
func main() {
	db, err := database.MysqlConnect()
	utils.LogErrorWithPanic(err)
	err = db.Ping()
	utils.LogErrorWithPanic(err)
	
	//success
	fmt.Println("Success connected to database")
	}
`
)
