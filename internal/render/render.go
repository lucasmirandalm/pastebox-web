package render

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type Renderer struct {
	templates map[string]*template.Template
}

func New(templatesDir string) (*Renderer, error) {
	pages, err := filepath.Glob(filepath.Join(templatesDir, "pages", "*.html"))
	if err != nil {
		return nil, err
	}

	templates := make(map[string]*template.Template)

	for _, page := range pages {
		name := filepath.Base(page)

		files := []string{
			filepath.Join(templatesDir, "layouts", "base.html"),
			page,
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		templates[name] = tmpl
	}

	return &Renderer{
		templates: templates,
	}, nil
}

func (r *Renderer) Render(w http.ResponseWriter, status int, page string, data any) {
	tmpl, ok := r.templates[page]
	if !ok {
		http.Error(w, fmt.Sprintf("template %s not found", page), http.StatusInternalServerError)
		return
	}

	var buffer bytes.Buffer

	err := tmpl.ExecuteTemplate(&buffer, "base", data)
	if err != nil {
		http.Error(w, "failed to render template", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)

	_, _ = buffer.WriteTo(w)
}
