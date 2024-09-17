package http2

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/phimtorr/phimtor/server/auth"
)

func (s Server) GetVideo(w http.ResponseWriter, r *http.Request, id int64) {
	user := auth.UserFromCtx(r.Context())
	video, err := s.repo.GetVideo(r.Context(), user, id)
	if err != nil {
		respondError(w, r, err)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"video": video,
	})

}
