package handlers

import (
	"net/http"

	"github.com/MikelGV/Personal-Website/internal/components/ui"
	"github.com/MikelGV/Personal-Website/internal/data"
	"github.com/MikelGV/Personal-Website/views"
)

func HomeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		projects, err := data.LoadProjects()
		if err != nil {
			http.Error(w, "Failed to load projects", http.StatusInternalServerError)
			return
		}
		views.Home(projects).Render(r.Context(), w)
	})
}

func ContactHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		views.Contact().Render(r.Context(), w)
	})
}

func AboutHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		views.AboutMe().Render(r.Context(), w)
	})
}

func BlogHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		views.Blog().Render(r.Context(), w)
	})
}

func PostHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}

func ProjectHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}

func ResumeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		views.Resume().Render(r.Context(), w)
	})
}

func ProjectsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		projects, err := data.LoadProjects()
		if err != nil {
			http.Error(w, "Failed to load projects", http.StatusInternalServerError)
			return
		}
		views.Projects(projects).Render(r.Context(), w)
	})
}

func ProjectPreviewHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		projects, err := data.LoadProjects()
		if err != nil {
			http.Error(w, "Failed to load projects", http.StatusInternalServerError)
			return
		}
		p := data.FindProject(projects, r.PathValue("id"))
		if p == nil {
			http.Error(w, "Project not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		ui.ProjectPreview(*p).Render(r.Context(), w)
	})
}
