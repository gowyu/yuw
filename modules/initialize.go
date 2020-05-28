package modules

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gowyu/yuw/exceptions"
	"github.com/spf13/viper"
)

const defaultType = "yaml"

var (
	I *initialize = nil
	defaultEnvironment []interface{} = []interface{}{"dev", "stg", "prd"}
)

type initialize struct {
	fs *File
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
	init.util.Panic(
		str == "",
		exceptions.TxT("yuw^m_init_a"), exceptions.ErrPosition(),
	)

	var ok bool
	ok, _ = init.util.StrContains(str, defaultEnvironment ...)
	init.util.Panic(
		ok == false,
		exceptions.TxT("yuw^m_init_b"), exceptions.ErrPosition(),
	)

	ok = init.fs.IsExist(".env." + str + ".yaml")
	init.util.Panic(
		ok == false,
		exceptions.TxT("yuw^m_init_c"), exceptions.ErrPosition(),
	)

	init.env.SetConfigName(".env." + str)

	err := init.env.ReadInConfig()
	init.util.Panic(
		err != nil,
		exceptions.TxT("yuw^m_init_d"), exceptions.ErrPosition(),
	)

	init.env.WatchConfig()
	init.env.OnConfigChange(func (e fsnotify.Event){

	})

	return init
}

func (init *initialize) Get(key string, val interface{}) (res interface{}) {
	res = val
	if init.env.IsSet(key) {
		res = init.env.Get(key)
	}

	return
}







