package main

import (
	"Go-BoilerPlateAPI/entity"
	"Go-BoilerPlateAPI/utils"
	"flag"
	"os"

	"github.com/mitchellh/colorstring"
)

func main() {
	packageName := flag.String("pkg", "", "need name!")
	flag.Parse()

	files := []entity.FileSpec{
		{Location: entity.Domain["databaseInInfrastructure"].(string) + "/", PkgName: *packageName, ScCode: entity.MysqlConnect, FileName: "mysql.go"},
		{Location: entity.Domain["envInInfrastructure"].(string) + "/", PkgName: *packageName, ScCode: entity.Env, FileName: ".env"},
		{Location: entity.Domain["utilsInInfrastucture"].(string) + "/", PkgName: *packageName, ScCode: entity.CacthError, FileName: "error.go"},
		{Location: entity.GetTopDirectory(), PkgName: *packageName, ScCode: entity.GoMod, FileName: "go.mod"},
		{Location: entity.GetTopDirectory(), PkgName: *packageName, ScCode: entity.GoSum, FileName: "go.sum"},
		{Location: entity.GetTopDirectory(), PkgName: "", ScCode: entity.README, FileName: "README.md"},
		{Location: entity.GetTopDirectory(), PkgName: *packageName, ScCode: entity.MainGo, FileName: "main.go"},
	}

	if *packageName == "" {
		colorstring.Println("[red][-] [white]please fill in the package name, example : `-pkg=<awesome project>`")
		os.Exit(1)
	}

	for _, layer := range entity.Keys {

		utils.CreateFolderProject(entity.Domain[layer].(string), entity.Perm)

	}

	utils.CreateFiles(files)

	colorstring.Printf(`[red]
	╋╋╋╋╋╋╋┏━━━┓╋╋╋╋╋╋┏┓╋┏━━━┓┏┓╋╋╋╋╋┏┓
	╋╋╋╋╋╋╋┃┏━┓┃╋╋╋╋╋╋┃┃╋┃┏━┓┣┛┗┓╋╋╋┏┛┗┓
	┏━━┳━━┓┃┃╋┃┣┓┏┳┳━━┫┃┏┫┗━━╋┓┏╋━━┳┻┓┏┛
	┃┏┓┃┏┓┃┃┃╋┃┃┃┃┣┫┏━┫┗┛┻━━┓┃┃┃┃┏┓┃┏┫┃
	┃┗┛┃┗┛┃┃┗━┛┃┗┛┃┃┗━┫┏┓┫┗━┛┃┃┗┫┏┓┃┃┃┗┓
	┗━┓┣━━┛┗━━┓┣━━┻┻━━┻┛┗┻━━━┛┗━┻┛┗┻┛┗━┛
	┏━┛┃╋╋╋╋╋╋┗┛
	┗━━┛ [cyan]v.1.0
[white]INFO : [green]%s `, "success generate template!")
}
