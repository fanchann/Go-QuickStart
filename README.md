## Go QuickStart v2.0.3
Accelerate the development of your REST API!\
This cutting-edge tool empowers you to swiftly generate a blazing-fast REST API without the hassle of manual database configuration.\
Kickstart with `Go-QuickStart -pkg=awesomeProject` and behold the extraordinary results unfold before your eyes.
## Installation
`go install github.com/fanchann/Go-QuickStart`

## Configuration
Set your environment variable
#### Open .env to configuration
```
DB_DRIVER=mysql
DB_AUTH_USERNAME=root
DB_AUTH_PASSWORD=root
DB_NAME=local_production
DB_URL=127.0.0.1
DB_PORT=3306
```
## Project Structure
```
.
├── .env
├── go.mod
├── go.sum
├── interface
│   ├── controller
│   └── middleware
├── internal
│   ├── models
│   │   ├── domain
│   │   └── web
│   ├── repositories
│   └── services
├── main.go
├── pkg
│   ├── config
│   │   └── config.go
│   ├── database
│   │   └── mysql.go
│   └── utils
│       └── error.go
├── README.md
└── tests


```

# Explanation for each layer in the Go-QuickStart structure:

#### Interface
Contains the sub-layers "controller" and "middleware".

*   Controller\
Handles user's HTTP requests and connects them to the application logic.
*   Middleware\
Adds additional features like authentication, authorization, logging, etc.

#### Internal
Contains the sub-layers "models", "repositories", and "services".

*   Models\
Provides definitions for the data structures used in the application.
*   Repositories\
Interacts with the data storage, such as a database.
*   Services\
Provides the business logic of the application.

#### Pkg 
Contains utilities, configuration files, and connections to external resources.
*   Config\
Application configuration files.
*   Database\
Implementation of the database connection.
*   Utils\
Common utilities used in the application.

#### Tests
Contains unit tests or automated tests to ensure the functionality of the application

## Thanks !
