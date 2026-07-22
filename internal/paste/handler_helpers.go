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
