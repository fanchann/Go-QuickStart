package codes

const (
	ConfigGo = `package config

import (
	"os"

	"{{.PackageName}}/pkg/utils"
	"github.com/joho/godotenv"
)
	
var (
	Driver   string
	Username string
	Password string
	Db_name  string
	Db_url   string
	Db_port  string
)
	
func LoadEnv() {
	errEnv := godotenv.Load(".env")
	utils.LogErrorWithPanic(errEnv)
	
	Driver = os.Getenv("DB_DRIVER")
	Username = os.Getenv("DB_AUTH_USERNAME")
	Password = os.Getenv("DB_AUTH_PASSWORD")
	Db_name = os.Getenv("DB_NAME")
	Db_url = os.Getenv("DB_URL")
	Db_port = os.Getenv("DB_PORT")
}`

	EnvConfiguration = `DB_DRIVER = mysql
DB_AUTH_USERNAME =
DB_AUTH_PASSWORD =
DB_NAME =
DB_URL =
DB_PORT =`
)
