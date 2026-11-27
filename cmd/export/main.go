package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"

	"github.com/MikelGV/Personal-Website/internal/config"
	"github.com/MikelGV/Personal-Website/internal/data"
	"github.com/MikelGV/Personal-Website/internal/handlers"
	"github.com/MikelGV/Personal-Website/internal/server"
	"github.com/MikelGV/Personal-Website/views"

	"github.com/a-h/templ"
)

type page struct {
	url  string
	file string
}

func main() {
	out := flag.String("out", "dist", "output directory")
	flag.Parse()

	if err := run(*out); err != nil {
		log.Fatalf("export failed: %v", err)
	}
}

func run(out string) error {
	if out == "" || out == "/" {
		return fmt.Errorf("invalid output directory %q", out)
	}

	if err := os.RemoveAll(out); err != nil {
		return fmt.Errorf("remove %s: %w", out, err)
	}
	if err := os.MkdirAll(out, 0o755); err != nil {
		return fmt.Errorf("create %s: %w", out, err)
	}

	srv := server.NewServer(&config.Config{})

	blogs, err := data.PublishedBlogs()
	if err != nil {
		return fmt.Errorf("load blogs: %w", err)
	}

	projects, err := data.LoadProjects()
	if err != nil {
		return fmt.Errorf("load projects: %w", err)
	}

	pages := []page{
		{url: "/", file: "index.html"},
		{url: "/about", file: "about.html"},
		{url: "/projects", file: "projects.html"},
		{url: "/blog", file: "blog.html"},
		{url: "/contact", file: "contact.html"},
		{url: "/cv", file: "cv.html"},
	}

	for _, b := range blogs {
		pages = append(pages, page{
			url:  "/blog/" + b.ID,
			file: "blog/" + b.ID + ".html",
		})
	}

	for _, p := range projects {
		pages = append(pages, page{
			url:  "/projects/preview/" + p.ID,
			file: "projects/preview/" + p.ID + ".html",
		})
	}

	totalBlogPages := (len(blogs) + handlers.BlogPageSize - 1) / handlers.BlogPageSize
	for n := 2; n <= totalBlogPages; n++ {
		pages = append(pages, page{
			url:  fmt.Sprintf("/blog/more/%d", n),
			file: fmt.Sprintf("blog/more/%d.html", n),
		})
	}

	for _, p := range pages {
		body, err := renderPage(srv, p.url)
		if err != nil {
			return err
		}
		if err := writeFile(out, p.file, body); err != nil {
			return err
		}
		fmt.Printf("wrote %s\n", p.file)
	}

	if err := writeFile(out, "resume.html", redirectPage("/cv")); err != nil {
		return err
	}
	fmt.Println("wrote resume.html")

	notFound, err := renderComponent(views.NotFound())
	if err != nil {
		return fmt.Errorf("render 404 page: %w", err)
	}
	if err := writeFile(out, "404.html", notFound); err != nil {
		return err
	}
	fmt.Println("wrote 404.html")

	if err := os.CopyFS(filepath.Join(out, "static"), os.DirFS("static")); err != nil {
		return fmt.Errorf("copy static assets: %w", err)
	}
	fmt.Println("copied static/")

	return nil
}

func renderPage(srv http.Handler, url string) ([]byte, error) {
	req := httptest.NewRequest(http.MethodGet, url, nil)
	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		return nil, fmt.Errorf("GET %s: status %d: %s", url, rec.Code, strings.TrimSpace(rec.Body.String()))
	}

	return rec.Body.Bytes(), nil
}

func renderComponent(component templ.Component) ([]byte, error) {
	var buf bytes.Buffer
	if err := component.Render(context.Background(), &buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func writeFile(root, name string, body []byte) error {
	path := filepath.Join(root, filepath.FromSlash(name))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create directory for %s: %w", name, err)
	}
	if err := os.WriteFile(path, body, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", name, err)
	}
	return nil
}

func redirectPage(target string) []byte {
	return []byte(fmt.Sprintf(`<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta http-equiv="refresh" content="0; url=%s">
<link rel="canonical" href="%s">
<title>Redirecting…</title>
</head>
<body>
<p>Redirecting to <a href="%s">%s</a>…</p>
<script>location.replace(%q)</script>
</body>
</html>
`, target, target, target, target, target))
}
