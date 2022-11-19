package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/artmadar/golang-monolith-web-app/pkg/config"
	"github.com/artmadar/golang-monolith-web-app/pkg/handlers"
	"github.com/artmadar/golang-monolith-web-app/pkg/render"
)

const appPort = 8080

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.GetTemplateCache()

	if err != nil {
		log.Fatal("Can't create template cache")
	}

	app.TempateCache = tc
	app.UserCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", appPort),
		Handler: routes(&app),
	}
	forever := make(chan any)
	go srv.ListenAndServe()
	log.Printf("Golang-monolith-web-app started at port %d\n", appPort)
	<-forever
}
