package codes

const (
	MainCode = `package main

import (
	"fmt"

	"{{.PackageName}}/pkg/database"
	"{{.PackageName}}/pkg/utils"
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
