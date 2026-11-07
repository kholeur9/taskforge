package web

import (
	"bytes"
	"log/slog"
	"net/http"
)

type Handler struct {
	templates *TemplateManager
}

func NewHandler(templates *TemplateManager) *Handler {
	return &Handler{
		templates: templates,
	}
}

func (h *Handler) Dashboard(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	if err := h.templates.Render(&buf, "dashboard", nil); err != nil {
		slog.Error("Erreur lors du rendu du template dashboard", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
}

func (h *Handler) Events(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	if err := h.templates.Render(&buf, "events", nil); err != nil {
		slog.Error("Erreur lors du rendu du template événements", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
}