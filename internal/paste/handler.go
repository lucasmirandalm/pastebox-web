package paste

import (
	"errors"
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
	const userID int64 = 1

	pasteID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid paste id", http.StatusBadRequest)
		return
	}

	paste, err := ph.service.FindByID(r.Context(), userID, pasteID)
	if err != nil {
		if errors.Is(err, ErrPasteNotFound) {
			http.Error(w, "paste not found", http.StatusNotFound)
			return
		}

		http.Error(w, "failed to load paste", http.StatusInternalServerError)
		return
	}

	ph.renderer.Render(w, http.StatusOK, "edit.html", paste)
}

func (ph *PasteHandler) Update(w http.ResponseWriter, r *http.Request) {
	const userID int64 = 1

	pasteID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid paste id", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	isFavorite := r.FormValue("is_favorite") == "on"

	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form data", http.StatusBadRequest)
		return
	}

	_, err = ph.service.Update(r.Context(), userID, pasteID, title, content, isFavorite)
	if err != nil {
		if errors.Is(err, ErrPasteNotFound) {
			http.Error(w, "paste not found", http.StatusNotFound)
			return
		}

		if errors.Is(err, ErrPasteTitleRequired) {
			http.Error(w, "title is required", http.StatusBadRequest)
			return
		}

		if errors.Is(err, ErrPasteContentRequired) {
			http.Error(w, "content is required", http.StatusBadRequest)
			return
		}

		http.Error(w, "failed to update paste", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/pastes/"+strconv.FormatInt(pasteID, 10)+"/edit", http.StatusSeeOther)
}
