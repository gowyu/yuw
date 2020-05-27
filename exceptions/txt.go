package exceptions

import (
	"fmt"
	"github.com/spf13/cast"
)

var (
	txt map[string]interface{} = map[string]interface{}{
		"yum^default":	"unknown error, ",

		"yuw^m": "error module, ",
		"yuw^m_b": "error initialize",
		"yuw^m_init_a": "config environment, go run ... --env=dev|stg|prd",
		"yuw^m_init_b": "config environment, --env=dev|stg|prd",
		"yuw^m_init_c": "missing .env.dev.yaml",
		"yuw^m_init_d": "config environment, ",
	}
)

func init() {

}

func TxT(tag string, content ... interface{}) (str string) {
	s, ok := txt[tag]
	if ok {
		str = cast.ToString(s)
	} else {
		str = cast.ToString(txt["yum^default"])
	}

	if len(content) > 0 {
		str = str + ", " + fmt.Sprint(content ...)
	}

	return
}


