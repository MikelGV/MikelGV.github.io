package data

import (
	_ "embed"
	"encoding/json"
)

//go:embed blogs.json
var blogsJSON []byte

type Blog struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Date        string `json:"date"`
	Intro       string `json:"intro"`
	Lead        string `json:"lead"`
	Body        string `json:"body"`
	Outro       string `json:"outro"`
	Image1      string `json:"image1"`
	Image2      string `json:"image2"`
	Image3      string `json:"image3"`
	Draft       bool   `json:"draft"`
}

func LoadBlogs() ([]Blog, error) {
	var blogs []Blog
	if err := json.Unmarshal(blogsJSON, &blogs); err != nil {
		return nil, err
	}
	return blogs, nil
}

func FindBlog(blogs []Blog, id string) *Blog {
	for i := range blogs {
		if blogs[i].ID == id {
			return &blogs[i]
		}
	}
	return nil
}

// PublishedBlogs returns only non-draft entries for public rendering.
// Drafts stay in blogs.json as formatting examples and never reach a view.
func PublishedBlogs() ([]Blog, error) {
	blogs, err := LoadBlogs()
	if err != nil {
		return nil, err
	}
	published := blogs[:0:0]
	for _, b := range blogs {
		if !b.Draft {
			published = append(published, b)
		}
	}
	return published, nil
}
