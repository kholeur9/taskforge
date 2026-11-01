package web

import (
	"log/slog"
	"net/http"
)

type Handler struct {
	templates *TemplateManager
}

func NewHandler(templates *TemplateManager) *Handler{
	return &Handler{
		templates: templates,
	}
}

func (h *Handler) Dashboard(w http.ResponseWriter, r *http.Request) {
	err := h.templates.Render(w, "base.html", nil)
	if err != nil {
		slog.Error("Erreur lors du rendu du template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}