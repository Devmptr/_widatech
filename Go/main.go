package main

import (
	"fmt"
	"widatech_interview/golang/core"
	"widatech_interview/golang/database"

	"github.com/alecthomas/kingpin/v2"
)

var (
	cmd = kingpin.Arg("cmd", "Command to Execute.").Required().String()
)

func main() {
	kingpin.Parse()

	app := core.NewApp(".")
	app.Boot()

	switch *cmd {
	case "serve":
		serve(app)
	case "migrate":
		migrate(app)
	default:
		fmt.Println("[ERR] [CMD] Not Found")
	}
}

func serve(app *core.App) {
	app.Serve()
}

func migrate(app *core.App) {
	err := database.Execute(app.Config.DatabaseConfig)

	if err != nil {
		fmt.Printf("[ERR] [DBS] Error Migrate with : %s\n", err.Error())
		return
	}

	fmt.Println("[INF] [DBS] Success Migrate")
}
