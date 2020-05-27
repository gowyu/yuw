package modules

import (
	"errors"
	"strings"
	"github.com/spf13/cast"
)

type Utils struct {

}

func NewUtils() *Utils {
	return &Utils{

	}
}

func (util *Utils) StrContains(str string, src ...interface{}) (ok bool, err error) {
	if len(src) < 1 {
		err = errors.New("source is nil")
		return
	}

	for _, val := range src {
		if strings.Contains(str, cast.ToString(val)) {
			ok = true
			return
		}
	}

	return
}
