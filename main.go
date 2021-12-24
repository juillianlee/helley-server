package main

import (
	"app-helley/src/config"
	"app-helley/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnviorments()
	db := config.ConnectionDatabase()
	r := router.NewRouter(db)

	log.Printf("listen on http://localhost:%d", config.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), r))
}
