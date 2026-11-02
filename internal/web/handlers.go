package web

import (
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
	data := DashboardData{
		EventsReceived: 10,
		TasksExecuted: 7,
		TasksFailed: 3,
	}
	err := h.templates.Render(w, "dashboard", data)
	if err != nil {
		slog.Error("Erreur lors du rendu du template dashboard", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Events(w http.ResponseWriter, r *http.Request) {
	err := h.templates.Render(w, "events", nil)
	if err != nil {
		slog.Error("Erreur lors du rendu du template événements", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}