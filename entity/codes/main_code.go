package codes

const (
	MainCode = `package main

import (
<<<<<<< HEAD
	"{{.PackageName}}/infrastructure/database"
	"{{.PackageName}}/infrastructure/utils"
=======
	"{{.PackageName}}/pkg/database"
	"{{.PackageName}}/pkg/utils"
>>>>>>> edit/utils
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
