package data

import (
	_ "embed"
	"encoding/json"
)

//go:embed projects.json
var projectsJSON []byte

type BlogRef struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type Project struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Field       string    `json:"field"`
	Blogs       []BlogRef `json:"blogs_about_it"`
	Image       string    `json:"image"`
	GitHub      string    `json:"github"`
	Year        string    `json:"year"`
	Techs       []string  `json:"techs"`
}

func LoadProjects() ([]Project, error) {
	var projects []Project
	if err := json.Unmarshal(projectsJSON, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

func FindProject(projects []Project, id string) *Project {
	for i := range projects {
		if projects[i].ID == id {
			return &projects[i]
		}
	}
	return nil
}
