package web

import "net/http"

type Handler struct {
	templates *TemplateManager
}

func (h *Handler) Dashboard(w http.ResponseWriter, r *http.Request) {
	err := h.templates.Render(w, "dashbord.html", nil)
	if err != nil {
		http.Error(w, "Template don't render", http.StatusNotFound)
		return
	}
}