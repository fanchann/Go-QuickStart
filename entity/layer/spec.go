package layer

import (
	"Go-QuickStart/utils/helpers"
	"fmt"
)

type (
	FileSpec struct {
		Location    string
		PackageName string
		GoVersion   string
		ScCode      string
		FileName    string
	}
)

var (
	DomainKeys = []string{
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
		"tests",
	}
	Perm = 0755

	ModelsLayer   = "models"
	DomainLayer   = "domain"
	WebLayer      = "web"
	RepoLayer     = "repositories"
	ServicesLayer = "services"

	InsfratuctureLayer = "infrastructure"
	UtilsLayer         = "utils"
	DatabaseLayer      = "database"
	EnvLayer           = "environments"

	InterfaceLayer  = "interface"
	ControllerLayer = "controller"
	MiddlewareLayer = "middleware"

	TestLayer = "tests"

	Domain = map[string]interface{}{
		"domain":                   fmt.Sprintf("%s/%s", helpers.GetWorkingDirectory(), DomainLayer),
		"modelsInDomain":           fmt.Sprintf("%s/%s/%s", helpers.GetWorkingDirectory(), DomainLayer, ModelsLayer),
		"domainInDomain":           fmt.Sprintf("%s/%s/%s/%s", helpers.GetWorkingDirectory(), DomainLayer, ModelsLayer, DomainLayer),
		"webInDomain":              fmt.Sprintf("%s/%s/%s/%s", helpers.GetWorkingDirectory(), DomainLayer, ModelsLayer, WebLayer),
		"repoInDomain":             fmt.Sprintf("%s/%s/%s", helpers.GetWorkingDirectory(), DomainLayer, RepoLayer),
		"usecaseInDomain":          fmt.Sprintf("%s/%s/%s", helpers.GetWorkingDirectory(), DomainLayer, ServicesLayer),
		"infrastucture":            fmt.Sprintf("%s/%s", helpers.GetWorkingDirectory(), InsfratuctureLayer),
		"utilsInInfrastucture":     fmt.Sprintf("%s/%s/%s", helpers.GetWorkingDirectory(), InsfratuctureLayer, UtilsLayer),
		"databaseInInfrastructure": fmt.Sprintf("%s/%s/%s", helpers.GetWorkingDirectory(), InsfratuctureLayer, DatabaseLayer),
		"envInInfrastructure":      fmt.Sprintf("%s/%s/%s", helpers.GetWorkingDirectory(), InsfratuctureLayer, EnvLayer),
		"interface":                fmt.Sprintf("%s/%s", helpers.GetWorkingDirectory(), InterfaceLayer),
		"controllerInInterface":    fmt.Sprintf("%s/%s/%s", helpers.GetWorkingDirectory(), InterfaceLayer, ControllerLayer),
		"middlewareInInterface":    fmt.Sprintf("%s/%s/%s", helpers.GetWorkingDirectory(), InterfaceLayer, MiddlewareLayer),
		"tests":                    fmt.Sprintf("%s/%s", helpers.GetWorkingDirectory(), TestLayer),
	}
)
