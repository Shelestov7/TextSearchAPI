package main

import (
	"log"
	"net/http"
	"os"

	"TextSearchAPI/internal/api/routes"
	"TextSearchAPI/internal/api/v0.1"
	"TextSearchAPI/internal/config"
)

func main() {

	// load ENV to configure app
	conf, err := config.LoadENV()
	if err != nil {
		log.Fatalf("Cant load Enviroment variables: %v", err)
	}

	text, err := os.ReadFile(conf.FilePath)
	if err != nil {
		log.Fatalf("Failed to read document with error: %v", err)
	}

	v0_1.InitTextSearchSource(text)

	r := routes.InitRoutes()

	err = http.ListenAndServe(":"+conf.HTTPPort, r)
	if err != nil {
		log.Fatalf("Cant start server on port %v: %v ", conf.HTTPPort, err)
	}

}
