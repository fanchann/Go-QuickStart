package codes

const (
	MainCode = `package main

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
