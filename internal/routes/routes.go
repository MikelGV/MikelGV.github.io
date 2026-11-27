package routes

import (
	"net/http"

	"github.com/MikelGV/Personal-Website/internal/config"
	"github.com/MikelGV/Personal-Website/internal/handlers"
)

func AddRoutes(
	mux *http.ServeMux,
	cfg *config.Config,
) {

	mux.Handle("/", handlers.HomeHandler())

	mux.Handle("/about", handlers.AboutHandler())

	mux.Handle("/projects", handlers.ProjectsHandler())

	mux.Handle("GET /projects/preview/{id}", handlers.ProjectPreviewHandler())

	mux.Handle("/project", handlers.ProjectHandler())

	mux.Handle("/blog", handlers.BlogHandler())

	mux.Handle("GET /blog/more/{page}", handlers.BlogMoreHandler())

	mux.Handle("/blog/{id}", handlers.PostHandler())

	mux.Handle("/contact", handlers.ContactHandler())

	mux.Handle("/cv", handlers.ResumeHandler())

	mux.Handle("/resume", http.RedirectHandler("/cv", http.StatusMovedPermanently))
}
