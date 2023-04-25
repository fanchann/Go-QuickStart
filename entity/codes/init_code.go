package codes

import "fmt"

const (
	GoMod = `module {{.PackageName}}

go {{.GoVersion}}

require (
	github.com/go-sql-driver/mysql v1.7.0
)
`
	GoSum = `github.com/go-sql-driver/mysql v1.7.0 h1:ueSltNNllEqE3qcWBTD0iQd3IpL/6U+mJxLkazJ7YPc=
github.com/go-sql-driver/mysql v1.7.0/go.mod h1:OXbVy3sEdcQ2Doequ6Z5BW6fXNQTmx+9S1MCJN5yJMI=
`
)

var Readme = fmt.Sprintf("# Read This\n- if you have not installed _mysql_ driver\\\n```go get -u github.com/go-sql-driver/mysql ```\n# Configuration\n#### Open terminal and set your environment variabel\n```\nexport DB_DRIVER=mysql\nexport DB_AUTH_USERNAME=<your username here>\nexport DB_AUTH_PASSWORD=<your password here>\nexport DB_NAME=<database target>\nexport DB_URL=<database url>\nexport DB_PORT=<database port>\n```\n\n#### Example\n```\nexport DB_DRIVER=mysql\nexport DB_AUTH_USERNAME=root\nexport DB_AUTH_PASSWORD=root\nexport DB_NAME=local_production\nexport DB_URL=127.0.0.1\nexport DB_PORT=3306\n```\nhappy coding with MySQL and Go!\n")
