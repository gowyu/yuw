package modules

import (
	"github.com/gowyu/yuw/exceptions"
	"github.com/spf13/pflag"
)

func init() {
	module := New()
	module.cfg()
}

type YuwInitialize struct {
	db bool
	Redis bool
	I18nT bool
	Email bool
}

type module struct {
	util *Utils
}

func New() *module {
	return &module {
		util: NewUtils(),
	}
}

func (module *module) cfg() {
	if I != nil {
		return
	}

	pflag.String("env", "", "environment configure")
	pflag.Parse()

	init := NewInitialize()
	err := init.env.BindPFlags(pflag.CommandLine)
	module.util.Panic(
		err != nil,
		exceptions.TxT("yum^m"), err.Error(),
		exceptions.ErrPosition(),
	)

	I = init.LoadInitializedFromYaml()
	module.util.Panic(
		I == nil,
		exceptions.TxT("yum^m_a"),
		exceptions.ErrPosition(),
	)
}
