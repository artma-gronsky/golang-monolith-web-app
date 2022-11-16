package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/artmadar/golang-monolith-web-app/pkg/config"
	"github.com/artmadar/golang-monolith-web-app/pkg/handlers"
	"github.com/artmadar/golang-monolith-web-app/pkg/render"
)

const appPort = 8080

func main() {
	var app config.AppConfig
	tc, err := render.GetTemplateCache()

	if err != nil {
		log.Fatal("Can't create template cache")
	}

	app.TempateCache = tc
	app.UserCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	forever := make(chan any)

	go http.ListenAndServe(fmt.Sprintf(":%d", appPort), nil)
	log.Printf("Golang-monolith-web-app started at port %d\n", appPort)

	<-forever
}
