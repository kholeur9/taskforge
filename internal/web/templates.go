package web

import (
	"fmt"
	"html/template"
	"io"
	//"net/http"
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

func (tm *TemplateManager) Render(w io.Writer, name string, data any) error {
	t, ok := tm.templates[name]
	if !ok {
		return fmt.Errorf("template inconnu %q", name)
	}
	return t.ExecuteTemplate(w, "base.html", data)
}
