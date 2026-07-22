package paste

import "net/http"

func (ph *PasteHandler) renderEditForm(w http.ResponseWriter, status int, paste Paste, errorMessage string) {
	data := EditPageData{
		Title: "Edit paste · Pastebox",
		Paste: paste,
		Error: errorMessage,
	}

	ph.renderer.Render(w, status, "edit.html", data)
}

func (ph *PasteHandler) renderNewForm(w http.ResponseWriter, status int, paste Paste, errorMessage string) {
	data := NewPageData{
		Title: "New paste · Pastebox",
		Paste: paste,
		Error: errorMessage,
	}

	ph.renderer.Render(w, status, "new.html", data)
}
