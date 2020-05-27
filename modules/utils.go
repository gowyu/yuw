package modules

import (
	"fmt"
	"github.com/gowyu/yuw/exceptions"
	"github.com/spf13/cast"
	"strings"
)

type Utils struct {

}

func NewUtils() *Utils {
	return &Utils{

	}
}

func (util *Utils) Panic(condition bool, content ... interface{}) {
	if condition {
		panic(fmt.Sprint(content ...))
	}
}

func (util *Utils) StrContains(str string, src ...interface{}) (ok bool, err error) {
	if len(src) < 1 {
		err = exceptions.Err("yuw^mod_util_a", exceptions.ErrPosition())
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
