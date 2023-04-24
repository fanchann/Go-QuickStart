package codes

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
	Readme = `# Read This
if you have not installed _mysql_ driver use:
- _go get -u github.com/go-sql-driver/mysql_
install additional packages to load _.env_
- _go get github.com/joho/godotenv_
	
~ enjoy
`
)
