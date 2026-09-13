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

	mux.Handle("/blog/{id}", handlers.PostHandler())

	mux.Handle("/contact", handlers.ContactHandler())

	mux.Handle("/resume", handlers.ResumeHandler())
}
