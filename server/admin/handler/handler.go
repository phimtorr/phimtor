package handler

import "net/http"

type Handler struct {
	repo Repository
}

func New(repo Repository) *Handler {
	if repo == nil {
		panic("nil repository")

	}
	return &Handler{
		repo: repo,
	}
}

func redirect(w http.ResponseWriter, r *http.Request, url string) {
	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("HX-Redirect", url)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Redirect(w, r, url, http.StatusSeeOther)
	}
}
