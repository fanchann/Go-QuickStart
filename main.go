package main

import (
	"github.com/mitchellh/colorstring"

	"github.com/Go-QuickStart/entity/files"
)

func main() {
	files.GenerateLayer()
	colorstring.Printf(`[red]
	╋╋╋╋╋╋╋┏━━━┓╋╋╋╋╋╋┏┓╋┏━━━┓┏┓╋╋╋╋╋┏┓
	╋╋╋╋╋╋╋┃┏━┓┃╋╋╋╋╋╋┃┃╋┃┏━┓┣┛┗┓╋╋╋┏┛┗┓
	┏━━┳━━┓┃┃╋┃┣┓┏┳┳━━┫┃┏┫┗━━╋┓┏╋━━┳┻┓┏┛
	┃┏┓┃┏┓┃┃┃╋┃┃┃┃┣┫┏━┫┗┛┻━━┓┃┃┃┃┏┓┃┏┫┃
	┃┗┛┃┗┛┃┃┗━┛┃┗┛┃┃┗━┫┏┓┫┗━┛┃┃┗┫┏┓┃┃┃┗┓
	┗━┓┣━━┛┗━━┓┣━━┻┻━━┻┛┗┻━━━┛┗━┻┛┗┻┛┗━┛
	┏━┛┃╋╋╋╋╋╋┗┛
	┗━━┛ [cyan]v.2.0
[white]INFO : [green]%s `, "success generate template!")
}
