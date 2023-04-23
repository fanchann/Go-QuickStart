package layer

import (
	"Go-QuickStart/utils"
	"fmt"
)

type (
	FileSpec struct {
		Location string
		PkgName  string
		ScCode   string
		FileName string
	}
)

var (
	Keys = []string{
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
		"domain":                   fmt.Sprintf("%s/%s", utils.GetTopDirectory(), DomainLayer),
		"modelsInDomain":           fmt.Sprintf("%s/%s/%s", utils.GetTopDirectory(), DomainLayer, ModelsLayer),
		"domainInDomain":           fmt.Sprintf("%s/%s/%s/%s", utils.GetTopDirectory(), DomainLayer, ModelsLayer, DomainLayer),
		"webInDomain":              fmt.Sprintf("%s/%s/%s/%s", utils.GetTopDirectory(), DomainLayer, ModelsLayer, WebLayer),
		"repoInDomain":             fmt.Sprintf("%s/%s/%s", utils.GetTopDirectory(), DomainLayer, RepoLayer),
		"usecaseInDomain":          fmt.Sprintf("%s/%s/%s", utils.GetTopDirectory(), DomainLayer, ServicesLayer),
		"infrastucture":            fmt.Sprintf("%s/%s", utils.GetTopDirectory(), InsfratuctureLayer),
		"utilsInInfrastucture":     fmt.Sprintf("%s/%s/%s", utils.GetTopDirectory(), InsfratuctureLayer, UtilsLayer),
		"databaseInInfrastructure": fmt.Sprintf("%s/%s/%s", utils.GetTopDirectory(), InsfratuctureLayer, DatabaseLayer),
		"envInInfrastructure":      fmt.Sprintf("%s/%s/%s", utils.GetTopDirectory(), InsfratuctureLayer, EnvLayer),
		"interface":                fmt.Sprintf("%s/%s", utils.GetTopDirectory(), InterfaceLayer),
		"controllerInInterface":    fmt.Sprintf("%s/%s/%s", utils.GetTopDirectory(), InterfaceLayer, ControllerLayer),
		"middlewareInInterface":    fmt.Sprintf("%s/%s/%s", utils.GetTopDirectory(), InterfaceLayer, MiddlewareLayer),
		"test":                     fmt.Sprintf("%s/%s", utils.GetTopDirectory(), TestLayer),
	}
)
