package middleware

import (
	"github.com/koopa0/Celeritas"
	"github.com/koopa0/myapp/data"
)

type Middleware struct {
	App    *Celeritas.Celeritas
	Models data.Models
}
