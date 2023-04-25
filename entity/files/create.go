package files

import (
	"flag"
	"os"

	"github.com/mitchellh/colorstring"

	"Go-QuickStart/entity/codes"
	"Go-QuickStart/entity/layer"
	"Go-QuickStart/utils"
	"Go-QuickStart/utils/helpers"
)

func GenerateLayer() {
	packageName := flag.String("pkg", "", "package name is required")
	flag.Parse()

	if *packageName == "" {
		colorstring.Println("[red][-] [white]please fill in the package name, example : `-pkg=<awesome project>`")
		os.Exit(1)
	}

	files := []layer.FileSpec{
		{Location: layer.Domain["databaseInPackage"] + "/", PackageName: *packageName, ScCode: codes.MysqlConnection, FileName: "mysql.go"},
		{Location: layer.Domain["utilsInPackage"] + "/", PackageName: *packageName, ScCode: codes.UtilsError, FileName: "error.go"},
		{Location: helpers.GetWorkingDirectory(), PackageName: *packageName, GoVersion: helpers.GetGoVersion(), ScCode: codes.GoMod, FileName: "go.mod"},
		{Location: helpers.GetWorkingDirectory(), PackageName: *packageName, ScCode: codes.GoSum, FileName: "go.sum"},
		{Location: layer.Domain["envInPackage"] + "/", PackageName: *packageName, ScCode: codes.EnvVariable, FileName: "database.go"},
		{Location: helpers.GetWorkingDirectory(), PackageName: "", ScCode: codes.Readme, FileName: "README.md"},
		{Location: helpers.GetWorkingDirectory(), PackageName: *packageName, ScCode: codes.MainCode, FileName: "main.go"},
	}

	newWrite := utils.NewWriteToFiles()
	defer newWrite.CreateFile(files)

	for _, layered := range layer.DomainKeys {
		newWrite.CreateFolderProject(layer.Domain[layered], layer.Perm)
	}
}
