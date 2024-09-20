package handler

import (
	"strconv"

	"github.com/friendsofgo/errors"
)

func parseID(idRaw string) (int64, error) {
	id, err := strconv.ParseInt(idRaw, 10, 64)
	if err != nil {
		return 0, errors.Wrap(err, "parsing id")
	}
	return id, nil
}
