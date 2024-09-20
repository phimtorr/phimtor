package handler

import (
	"fmt"
	"net/http"
	"strconv"

	commonErrors "github.com/phimtorr/phimtor/common/errors"
)

const (
	pageSize = 10
)

func parseID(idRaw string) (int64, error) {
	id, err := strconv.ParseInt(idRaw, 10, 64)
	if err != nil {
		return 0, commonErrors.NewIncorrectInputError("invalid-id",
			fmt.Sprintf("invalid id=%s, err: %v", idRaw, err))
	}
	return id, nil
}

func redirect(w http.ResponseWriter, r *http.Request, url string) {
	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("HX-Redirect", url)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Redirect(w, r, url, http.StatusSeeOther)
	}
}
