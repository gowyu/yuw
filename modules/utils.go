package modules

import (
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"strings"
)

type utils struct {

}

func NewUtils() *utils {
	return &utils{

	}
}

func (util *utils) Panic(condition bool, content ... interface{}) {
	if condition {
		panic(fmt.Sprint(content ...))
	}
}

func (util *utils) StrContains(str string, src ...interface{}) (ok bool, err error) {
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
