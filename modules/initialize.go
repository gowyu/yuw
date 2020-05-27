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
	env *viper.Viper
	util *Utils
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
	if str == "" {
		panic("--env=? is not configured")
	}

	if ok, _ := init.util.StrContains(str, defaultEnvironment ...); ok == false {
		panic("--env=? must be in dev,stg,prd")
	}

	init.env.SetConfigName(".env." + str)

	if err := init.env.ReadInConfig(); err != nil {
		panic(err.Error())
	}

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







