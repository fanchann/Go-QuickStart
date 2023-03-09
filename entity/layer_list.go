package entity

import (
	"fmt"
	"os"
)

type FileSpec struct {
	Location string
	PkgName  string
	ScCode   string
	FileName string
}

func GetTopDirectory() string {
	TopDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return TopDir
}

const (
	Perm = 0755

	ModelsLayer  = "models"
	DomainLayer  = "domain"
	WebLayer     = "web"
	RepoLayer    = "repository"
	UsecaseLayer = "usecase"

	InsfratuctureLayer = "insfrastructure"
	UtilsLayer         = "utils"
	DatabaseLayer      = "database"
	EnvLayer           = "environments"

	InterfaceLayer  = "interface"
	ControllerLayer = "controller"
	MiddlewareLayer = "middleware"

	TestLayer = "test"
)

var Domain = map[string]interface{}{
	"domain":                   fmt.Sprintf("%s/%s", GetTopDirectory(), DomainLayer),
	"modelsInDomain":           fmt.Sprintf("%s/%s/%s", GetTopDirectory(), DomainLayer, ModelsLayer),
	"domainInDomain":           fmt.Sprintf("%s/%s/%s/%s", GetTopDirectory(), DomainLayer, ModelsLayer, DomainLayer),
	"webInDomain":              fmt.Sprintf("%s/%s/%s/%s", GetTopDirectory(), DomainLayer, ModelsLayer, WebLayer),
	"repoInDomain":             fmt.Sprintf("%s/%s/%s", GetTopDirectory(), DomainLayer, RepoLayer),
	"usecaseInDomain":          fmt.Sprintf("%s/%s/%s", GetTopDirectory(), DomainLayer, UsecaseLayer),
	"infrastucture":            fmt.Sprintf("%s/%s", GetTopDirectory(), InsfratuctureLayer),
	"utilsInInfrastucture":     fmt.Sprintf("%s/%s/%s", GetTopDirectory(), InsfratuctureLayer, UtilsLayer),
	"databaseInInfrastructure": fmt.Sprintf("%s/%s/%s", GetTopDirectory(), InsfratuctureLayer, DatabaseLayer),
	"envInInfrastructure":      fmt.Sprintf("%s/%s/%s", GetTopDirectory(), InsfratuctureLayer, EnvLayer),
	"interface":                fmt.Sprintf("%s/%s", GetTopDirectory(), InterfaceLayer),
	"controllerInInterface":    fmt.Sprintf("%s/%s/%s", GetTopDirectory(), InterfaceLayer, ControllerLayer),
	"middlewareInInterface":    fmt.Sprintf("%s/%s/%s", GetTopDirectory(), InterfaceLayer, MiddlewareLayer),
	"test":                     fmt.Sprintf("%s/%s", GetTopDirectory(), TestLayer),
}

// Slice untuk menentukan urutan key-nya
var Keys = []string{
	"domain",
	"modelsInDomain",
	"domainInDomain",
	"webInDomain",
	"repoInDomain",
	"usecaseInDomain",
	"infrastucture",
	"utilsInInfrastucture",
	"databaseInInfrastructure",
	"envInInfrastructure",
	"interface",
	"controllerInInterface",
	"middlewareInInterface",
	"test",
}
