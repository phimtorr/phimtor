package http2

import (
	"net/http"
)

const (
	unsupportedMsg = "Hiện tại không còn hỗ trợ phiên bản này. Mời bạn vào trang chủ phimtor.net để tải bản mới nhất."
)

func (s Server) Unsupported(w http.ResponseWriter, r *http.Request) {
	handleError(w, r, unsupportedMsg, "unsupported", nil, http.StatusNotImplemented)
}
