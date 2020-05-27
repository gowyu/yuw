package modules

import (
	"github.com/spf13/pflag"
)

func init() {
	module := new(module)
	module.cfg()
}

type YuwInitialize struct {
	db bool
	Redis bool
	I18nT bool
	Email bool
}

type module struct {

}

func (module *module) cfg() {
	if I != nil {
		return
	}

	pflag.String("env", "", "environment configure")
	pflag.Parse()

	init := NewInitialize()
	err := init.env.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err.Error())
	}

	I = init.LoadInitializedFromYaml()
	if I == nil {
		panic("error initialize")
	}
}
