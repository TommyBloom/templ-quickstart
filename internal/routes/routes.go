package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"encoding/json"

	"log"

	"text/template"
)

type Page struct {
	Title         string `json:"title"`
	PublishedDate string `json:"publishedDate"`
	Slug          string `json:"slug"`
}

func FetchPages(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://your-payloadcms2-api-endpoint/pages")
	if err != nil {
		log.Printf("Error fetching pages: %v", err)
		http.Error(w, "Failed to fetch pages", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {package routes

		import (
			"net/http"
			"io/ioutil"
			"encoding/json"
			"log"
			"github.com/go-chi/chi/v5"
		)
		
		type Page struct {
			Title         string `json:"title"`
			PublishedDate string `json:"publishedDate"`
			Slug          string `json:"slug"`
		}
		
		func FetchPages(w http.ResponseWriter, r *http.Request) {
			resp, err := http.Get("https://your-payloadcms2-api-endpoint/pages")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer resp.Body.Close()
		
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		
			var pages []Page
			if err := json.Unmarshal(body, &pages); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		
			// Pass pages to template
			renderTemplate(w, "pages", pages)
		}
		
		func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
			// Implement your template rendering logic here
		}
		
		func NewRouter() *chi.Mux {
			r := chi.NewRouter()
			r.Get("/", FetchPages)
			return r
		}
		
		log.Printf("Error reading response body: %v", err)
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	var pages []Page
	if err := json.Unmarshal(body, &pages); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		http.Error(w, "Failed to unmarshal response", http.StatusInternalServerError)
		return
	}

	// Pass data to template or respond with JSON
	renderTemplate(w, "pages", pages)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := fmt.Sprintf("templates/%s.html", tmpl)
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", FetchPages).Methods("GET")
	// Add more routes as needed
	return r
}
