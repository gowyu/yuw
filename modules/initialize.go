package modules

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const defaultType = "yaml"

var (
	I *initialize = nil
	defaultEnvironment []interface{} = []interface{}{"dev", "stg", "prd"}
)

type initialize struct {
	fs *file
	env *viper.Viper
	util *utils
}

func NewInitialize() *initialize {
	return &initialize {
		env: viper.New(),
		util: NewUtils(),
	}
}

func (init *initialize) LoadInitializedFromYaml() *initialize {
	init.env.AddConfigPath(".")
	init.env.SetConfigType(defaultType)

	str := init.env.GetString("env")
	init.util.Panic(str == "", []interface{}{"--env=? is not configured"})

	var ok bool
	ok, _ = init.util.StrContains(str, defaultEnvironment ...)
	init.util.Panic(ok == false, []interface{}{"--env=? must be in dev,stg,prd"})

	ok = init.fs.IsExist(".env." + str + ".yaml")
	init.util.Panic(ok, []interface{}{".env is not exist"})

	init.env.SetConfigName(".env." + str)

	err := init.env.ReadInConfig()
	init.util.Panic(err != nil, []interface{}{})

	init.env.WatchConfig()
	init.env.OnConfigChange(func (e fsnotify.Event){})

	return init
}

func (init *initialize) Get(key string, val interface{}) (res interface{}) {
	res = val
	if init.env.IsSet(key) {
		res = init.env.Get(key)
	}

	return
}







