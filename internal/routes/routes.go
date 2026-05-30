package routes

import "net/http"

func AddRoutes(
	mux *http.ServeMux,
) {

	mux.Handle("/", handlers.HomeHandler)

	mux.Handle("/about", handlers.AboutHandler)

	mux.Handle("/projects", handlers.ProjectsHandler)

	mux.Handle("/project", handlers.ProjectHandler)

	mux.Handle("/blog", handlers.BlogHandler)

	mux.Handle("/blog/{id}", handlers.PostHandler)

	mux.Handle("/contact", handlers.ContactHandler)
}
