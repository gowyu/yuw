package modules

import (
	"encoding/json"
	c "github.com/gowyu/yuw/configs"
	E "github.com/gowyu/yuw/exceptions"
	"github.com/spf13/cast"
	"math/rand"
	"strings"
	"time"
)

type Utils struct {

}

func NewUtils() *Utils {
	return &Utils {

	}
}

func (util *Utils) SetTimeLocation(toTime string) (*time.Location, error) {
	if toTime == "" {
		toTime = c.LocationAsiaShanghai
	}

	return time.LoadLocation(toTime)
}

func (util *Utils) InterfaceToStringInMap(data map[interface{}]interface{}) (toMap map[string]interface{}) {
	if len(data) < 1 {
		return
	}

	toMap = make(map[string]interface{}, len(data))
	for key, val := range data {
		toMap[cast.ToString(key)] = val
	}

	return
}

func (util *Utils) MapToStruct(src interface{}, data interface{}) (err error) {
	strJson, err := json.Marshal(src)
	if err != nil {
		return
	}

	return json.Unmarshal(strJson, &data)
}

func (util *Utils) StrContains(str string, src ...interface{}) (ok bool, err error) {
	if len(src) < 1 {
		err = E.Err("yuw^mod_util_a", E.ErrPosition())
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

func (util *Utils) IntRand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
