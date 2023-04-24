package layer

import (
	"Go-QuickStart/utils"
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
		"domain":                   utils.CreateLayer(utils.GetWorkingDirectory(), DomainLayer),
		"modelsInDomain":           utils.CreateLayer(utils.GetWorkingDirectory(), DomainLayer, ModelsLayer),
		"domainInDomain":           utils.CreateLayer(utils.GetWorkingDirectory(), DomainLayer, ModelsLayer, DomainLayer),
		"webInDomain":              utils.CreateLayer(utils.GetWorkingDirectory(), DomainLayer, ModelsLayer, WebLayer),
		"repoInDomain":             utils.CreateLayer(utils.GetWorkingDirectory(), DomainLayer, RepoLayer),
		"usecaseInDomain":          utils.CreateLayer(utils.GetWorkingDirectory(), DomainLayer, ServicesLayer),
		"infrastucture":            utils.CreateLayer(utils.GetWorkingDirectory(), InsfratuctureLayer),
		"utilsInInfrastucture":     utils.CreateLayer(utils.GetWorkingDirectory(), InsfratuctureLayer, UtilsLayer),
		"databaseInInfrastructure": utils.CreateLayer(utils.GetWorkingDirectory(), InsfratuctureLayer, DatabaseLayer),
		"envInInfrastructure":      utils.CreateLayer(utils.GetWorkingDirectory(), InsfratuctureLayer, EnvLayer),
		"interface":                utils.CreateLayer(utils.GetWorkingDirectory(), InterfaceLayer),
		"controllerInInterface":    utils.CreateLayer(utils.GetWorkingDirectory(), InterfaceLayer, ControllerLayer),
		"middlewareInInterface":    utils.CreateLayer(utils.GetWorkingDirectory(), InterfaceLayer, MiddlewareLayer),
		"test":                     utils.CreateLayer(utils.GetWorkingDirectory(), TestLayer),
	}
)
