package web

import (
	"net/http"
	"text/template"
)

type TemplateManager struct {
	templates *template.Template
}

func NewTemplateManger() (*TemplateManager, error) {
	tmpl, err := template.ParseFiles("templates/dashboard.html")
	if err != nil {
		return nil, err
	}
	return &TemplateManager{
		templates: tmpl,
	}, nil
}

func (tm *TemplateManager) Render(w http.ResponseWriter, name string, data any) error {
	return tm.templates.ExecuteTemplate(w, name, data)
}