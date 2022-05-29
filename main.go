package main

import (
	"github.com/koopa0/Celeritas"
	"github.com/koopa0/myapp/data"
	"github.com/koopa0/myapp/handlers"
	"github.com/koopa0/myapp/middleware"
)

type application struct {
	App        *Celeritas.Celeritas
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
