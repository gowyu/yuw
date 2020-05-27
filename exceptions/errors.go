package exceptions

import (
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"runtime"
)

var (
	msg map[string]interface{} = map[string]interface{}{
		"yuw^mod_util_a": "utils contains, source is nil",
	}
)

func Err(tag string, content ...interface{}) (err error) {
	var str string = "Unknown Error!"

	s, ok := msg[tag]
	if ok {
		str = cast.ToString(s)
	}

	if len(content) > 0 {
		str = str + ", " + fmt.Sprint(content ...)
	}

	err = errors.New(str)
	return
}

func ErrPosition() interface{} {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%v:%v", file, line)
}
