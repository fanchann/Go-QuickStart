## Go QuickStart v 2.0
Speed up to create your REST API!\
This tool is to help you create REST API faster, you don't need to  manually configure the connection to the database anymore.\
start with `go-start -pkg=awesomeProject` and see a miracle come out

## Installation
for linux:\
`git clone https://github.com/fanchann/Go-QuickStart && cd Go-QuickStart && go build -o bin/go-start main.go && sudo mv bin/go-start /bin`

## Configuration
Set your environment variable
### Linux
`this tool is still the same as the previous version, still using mysql, maybe in xxxxx version will adopt some database`
```
export DB_DRIVER=mysql
export DB_AUTH_USERNAME=<your username here>
export DB_AUTH_PASSWORD=<your password here>
export DB_NAME=<database target>
export DB_URL=<database url>
export DB_PORT=<database port>
```
### Example

```
export DB_DRIVER=mysql
export DB_AUTH_USERNAME=root
export DB_AUTH_PASSWORD=root
export DB_NAME=local_production
export DB_URL=127.0.0.1
export DB_PORT=3306
```
## Structure folder
```
.
|-- README.md
|-- go.mod
|-- go.sum
|-- interface
|   |-- controller
|   `-- middleware
|-- internal
|   |-- models
|   |   |-- domain
|   |   `-- web
|   |-- repositories
|   `-- services
|-- main.go
|-- pkg
|   |-- database
|   |   `-- mysql.go
|   |-- environments
|   |   `-- database.go
|   `-- utils
|       `-- error.go
`-- tests

```


## Thanks !
