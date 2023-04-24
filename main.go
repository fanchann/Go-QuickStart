package main

import (
	"Go-QuickStart/entity/files"

	"github.com/mitchellh/colorstring"
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
