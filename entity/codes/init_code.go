package codes

import "fmt"

const (
	GoMod = `module {{.PackageName}}

go {{.GoVersion}}

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
)

var Readme = fmt.Sprintf("# Read This\n- add package with\\\n```go mod download```\n# Configuration\n#### Open .env to configuration\n```\nDB_DRIVER=mysql\nDB_AUTH_USERNAME=<your username here>\nDB_AUTH_PASSWORD=<your password here>\nDB_NAME=<database target>\nDB_URL=<database url>\nDB_PORT=<database port>\n```\n\n#### Example\n```\nDB_DRIVER=mysql\nDB_AUTH_USERNAME=root\nDB_AUTH_PASSWORD=root\nDB_NAME=local_production\nDB_URL=127.0.0.1\nDB_PORT=3306\n```\nhappy coding with MySQL and Go!\n")
