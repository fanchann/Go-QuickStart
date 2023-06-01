package layer

import (
	"fmt"

	"github.com/Go-QuickStart/utils/helpers"
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
		"serviceInDomain",
		"package",
		"utilsInPackage",
		"databaseInPackage",
		"configInPackage",
		"interface",
		"controllerInInterface",
		"middlewareInInterface",
		"tests",
	}

	Perm = 0755

	InternalLayer = "internal"
	ModelsLayer   = "models"
	DomainLayer   = "domain"
	WebLayer      = "web"
	RepoLayer     = "repositories"
	ServicesLayer = "services"

	PackageLayer  = "pkg"
	UtilsLayer    = "utils"
	DatabaseLayer = "database"
	ConfigLayer   = "config"

	InterfaceLayer  = "interface"
	ControllerLayer = "controller"
	MiddlewareLayer = "middleware"

	TestLayer = "tests"

	Domain = map[string]string{
		"domain":                fmt.Sprintf("%s/%s", helpers.GetWorkingDirectory(), InternalLayer),
		"modelsInDomain":        fmt.Sprintf("%s/%s/%s", helpers.GetWorkingDirectory(), InternalLayer, ModelsLayer),
		"domainInDomain":        fmt.Sprintf("%s/%s/%s/%s", helpers.GetWorkingDirectory(), InternalLayer, ModelsLayer, DomainLayer),
		"webInDomain":           fmt.Sprintf("%s/%s/%s/%s", helpers.GetWorkingDirectory(), InternalLayer, ModelsLayer, WebLayer),
		"repoInDomain":          fmt.Sprintf("%s/%s/%s", helpers.GetWorkingDirectory(), InternalLayer, RepoLayer),
		"serviceInDomain":       fmt.Sprintf("%s/%s/%s", helpers.GetWorkingDirectory(), InternalLayer, ServicesLayer),
		"package":               fmt.Sprintf("%s/%s", helpers.GetWorkingDirectory(), PackageLayer),
		"utilsInPackage":        fmt.Sprintf("%s/%s/%s", helpers.GetWorkingDirectory(), PackageLayer, UtilsLayer),
		"databaseInPackage":     fmt.Sprintf("%s/%s/%s", helpers.GetWorkingDirectory(), PackageLayer, DatabaseLayer),
		"configInPackage":       fmt.Sprintf("%s/%s/%s", helpers.GetWorkingDirectory(), PackageLayer, ConfigLayer),
		"interface":             fmt.Sprintf("%s/%s", helpers.GetWorkingDirectory(), InterfaceLayer),
		"controllerInInterface": fmt.Sprintf("%s/%s/%s", helpers.GetWorkingDirectory(), InterfaceLayer, ControllerLayer),
		"middlewareInInterface": fmt.Sprintf("%s/%s/%s", helpers.GetWorkingDirectory(), InterfaceLayer, MiddlewareLayer),
		"tests":                 fmt.Sprintf("%s/%s", helpers.GetWorkingDirectory(), TestLayer),
	}
)
