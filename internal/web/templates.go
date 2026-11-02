package web

import (
	"html/template"
	"net/http"
)

type TemplateManager struct {
	templates map[string]*template.Template
}

func NewTemplateManager() (*TemplateManager, error) {
	tmplBase, err := template.ParseFiles("templates/layouts/base.html", "templates/partials/navbar.html")
	if err != nil {
		return nil, err
	}
	tmplCloneDash, err := tmplBase.Clone()
	if err != nil {
		return nil, err
	}
	tmplDash, err := tmplCloneDash.ParseFiles("templates/pages/dashboard.html")
	if err != nil {
		return nil, err
	}
	tmplCloneEvents, err := tmplBase.Clone()
	if err != nil {
		return nil, err
	}
	tmplEvent, err := tmplCloneEvents.ParseFiles("templates/pages/events.html")
	if err != nil {
		return nil, err
	}
	return &TemplateManager{
		templates: map[string]*template.Template{
			"dashboard": tmplDash,
			"events": tmplEvent,
		},
	}, nil
}

func (tm *TemplateManager) Render(w http.ResponseWriter, name string, data any) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return tm.templates[name].ExecuteTemplate(w, "base.html", data)
}
