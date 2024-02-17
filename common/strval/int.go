package strval

import (
	"strconv"
)

func MustInt(str string) int {
	r, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return r
}
