package config

import "html/template"

// AppConfig holds application configuration
type AppConfig struct {
	UserCache    bool
	TempateCache map[string]*template.Template
}
