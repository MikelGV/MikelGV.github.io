package handlers

import (
	"net/http"
	"strconv"

	"github.com/MikelGV/Personal-Website/internal/components/ui"
	"github.com/MikelGV/Personal-Website/internal/data"
	"github.com/MikelGV/Personal-Website/views"
)

const BlogPageSize = 8

func HomeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			views.NotFound().Render(r.Context(), w)
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
		blogs, err := data.PublishedBlogs()
		if err != nil {
			http.Error(w, "Failed to load blogs", http.StatusInternalServerError)
			return
		}
		end := min(len(blogs), BlogPageSize)
		views.Blog(blogs[:end], len(blogs) > end, 2).Render(r.Context(), w)
	})
}

func BlogMoreHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		blogs, err := data.PublishedBlogs()
		if err != nil {
			http.Error(w, "Failed to load blogs", http.StatusInternalServerError)
			return
		}
		page := 2
		if p, err := strconv.Atoi(r.PathValue("page")); err == nil && p >= 2 {
			page = p
		}
		start := (page - 1) * BlogPageSize
		if start >= len(blogs) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			ui.BlogMoreChunk(nil, false, page+1).Render(r.Context(), w)
			return
		}
		end := min(start+BlogPageSize, len(blogs))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		ui.BlogMoreChunk(blogs[start:end], len(blogs) > end, page+1).Render(r.Context(), w)
	})
}

func PostHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		blogs, err := data.PublishedBlogs()
		if err != nil {
			http.Error(w, "Failed to load blogs", http.StatusInternalServerError)
			return
		}
		b := data.FindBlog(blogs, r.PathValue("id"))
		if b == nil {
			http.Error(w, "Blog post not found", http.StatusNotFound)
			return
		}
		views.Post(*b).Render(r.Context(), w)
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
