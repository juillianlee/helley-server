package main

import (
	"app-helley/src/config"
	"app-helley/src/router"
	"fmt"
)

func main() {
	config.LoadEnviorments()
	db := config.ConnectionDatabase()
	r := router.NewRouter(db)
	r.Logger.Fatal(r.Start(fmt.Sprintf(":%d", config.PORT)))
}
