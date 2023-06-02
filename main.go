package main

import (
	"github.com/mitchellh/colorstring"

	"github.com/fanchann/Go-QuickStart/entity/files"
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
	┗━━┛ [cyan]v.2.0.3
[white]INFO : [green]%s `, "success generate template!")
}
