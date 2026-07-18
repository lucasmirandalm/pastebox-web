package paste

import (
	"net/http"

	"github.com/lucasmirandalm/pastebox-web/internal/render"
)

type PasteHandler struct {
	renderer *render.Renderer
}

func NewPasteHandler(renderer *render.Renderer) *PasteHandler {
	return &PasteHandler{
		renderer: renderer,
	}
}

func (ph *PasteHandler) Home(w http.ResponseWriter, r *http.Request) {
	data := HomePageData{
		Title: "Pastebox Web",
	}

	ph.renderer.Render(w, http.StatusOK, "home.html", data)
}
