package paste

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/lucasmirandalm/pastebox-web/internal/render"
)

type PasteHandler struct {
	renderer *render.Renderer
	service  *PasteService
}

func NewPasteHandler(renderer *render.Renderer, service *PasteService) *PasteHandler {
	return &PasteHandler{
		renderer: renderer,
		service:  service,
	}
}

func (ph *PasteHandler) Home(w http.ResponseWriter, r *http.Request) {
	const userID int64 = 1

	totalPastes, err := ph.service.CountByUserID(r.Context(), userID)
	if err != nil {
		http.Error(w, "failed to count pastes", http.StatusInternalServerError)
		return
	}

	pastes, err := ph.service.ListByUserID(r.Context(), userID)
	if err != nil {
		http.Error(w, "failed to load pastes", http.StatusInternalServerError)
		return
	}

	data := HomePageData{
		Title:       "Pastebox Web",
		TotalPastes: totalPastes,
		Pastes:      pastes,
	}

	ph.renderer.Render(w, http.StatusOK, "home.html", data)
}

func (ph *PasteHandler) Edit(w http.ResponseWriter, r *http.Request) {
	pasteID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid paste id", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Edit paste %d", pasteID)
}
