package view

import (
	"net/http"
	"path/filepath"
	"snake-scape/internal/middleware"
	"snake-scape/internal/payload"
	"snake-scape/internal/template"
)

func ServeFavicon(w http.ResponseWriter, r *http.Request) {
	filePath := "favicon.ico"
	fullPath := filepath.Join(".", "static", filePath)
	http.ServeFile(w, r, fullPath)
}

func ServeStaticFiles(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[len("/static/"):]
	fullPath := filepath.Join(".", "static", filePath)
	http.ServeFile(w, r, fullPath)
}

func Home(ctx *middleware.CustomContext, w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return

	}
	template.Home("Templ Quickstart").Render(ctx, w)
}

func PageHandler(ctx *middleware.CustomContext, w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Query().Get("slug")
	page, err := payload.FetchPage(slug)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	template.RenderPage(page).Render(ctx, w)
}
