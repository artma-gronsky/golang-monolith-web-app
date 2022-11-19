package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds application configuration
type AppConfig struct {
	UserCache    bool
	TempateCache map[string]*template.Template
	InfoLog      *log.Logger
	InProduction bool
	Session      *scs.SessionManager
}
