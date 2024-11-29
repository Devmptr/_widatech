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

	switch *cmd {
	case "serve":
		serve()
	case "migrate":
		migrate()
	default:
		fmt.Println("[ERR] [CMD] Not Found")
	}
}

func serve() {
	app := core.NewApp(".")

	app.Boot()
}

func migrate() {
	err := database.Execute()

	if err != nil {
		fmt.Printf("[ERR] [DBS] Error Migrate with : %s\n", err.Error())
		return
	}

	fmt.Println("[INF] [DBS] Success Migrate")
}
