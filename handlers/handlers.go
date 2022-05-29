package handlers

import (
	"github.com/koopa0/Celeritas"
	"github.com/koopa0/myapp/data"
	"net/http"
)

type Handlers struct {
	App    *Celeritas.Celeritas
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
