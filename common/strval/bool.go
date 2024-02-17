package strval

import (
	"strconv"
)

func MustBool(str string) bool {
	r, err := strconv.ParseBool(str)
	if err != nil {
		panic(err)
	}
	return r
}
