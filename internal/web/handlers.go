package web

import (
	"log/slog"
	"net/http"
)

type Handler struct {
	templates *TemplateManager
}

func (h *Handler) Dashboard(w http.ResponseWriter, r *http.Request) {
	err := h.templates.Render(w, "dashboard.html", nil)
	if err != nil {
		slog.Error("Erreur lors du rendu du template", "error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}