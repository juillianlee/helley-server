package main

import (
	"app-helley/src/infrastructure/config"
	"app-helley/src/infrastructure/http/router"
	"fmt"
)

func main() {
	config.LoadEnviorments()
	db := config.ConnectionDatabase()
	r := router.NewRouter(db)
	fmt.Printf("App running on http://localhost:%d", config.PORT)
	r.Logger.Fatal(r.Start(fmt.Sprintf(":%d", config.PORT)))
}
