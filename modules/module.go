package modules

import (
	c "github.com/gowyu/yuw/configs"
	E "github.com/gowyu/yuw/exceptions"
	"github.com/spf13/cast"
	"github.com/spf13/pflag"
	"time"
)

type (
	YuwInitialize struct {
		Table bool
		Redis bool
		I18nT bool
		Email bool
	}

	module struct {
		util *Utils
	}
)

func init() {
	m := New()
	m.cfg()

	var YuwInitialized *YuwInitialize
	E.ErrPanic(m.util.MapToStruct(
		I.Get("YuwInitialize", []interface{}{}),
		&YuwInitialized,
	))

	if YuwInitialized.Table {
		strTimeLocation := cast.ToString(I.Get("Yuw.TimeLocation", c.LocationAsiaShanghai))
		timeLocation, err := m.util.SetTimeLocation(strTimeLocation)
		E.ErrPanic(err)

		m.db(timeLocation)
	}
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
	E.ErrPanic(init.Env.BindPFlags(pflag.CommandLine))

	I = init.LoadInitializedFromYaml()
	E.ErrArray(&E.ErrType{"yuw^m_a":I == nil})
}

func (module *module) db(timeLocation *time.Location) {
	sysTimeLocation = timeLocation

	var configs *dbConfigs

	cfg := I.Get("DBClusters.Configure", map[string]interface{}{}).(map[string]interface{})
	if len(cfg) == 0 {
		configs = &dbConfigs {
			DriverName: "mysql",
			MaxOpen: 1000,
			MaxIdle: 500,
			ShowedSQL: false,
			CachedSQL: false,
		}
	} else {
		E.ErrPanic(module.util.MapToStruct(cfg, &configs))
	}

	env := I.Get("DBClusters.Databases", map[string]interface{}{}).(map[string]interface{})
	E.ErrArray(&E.ErrType{"yuw^m_c":len(env) == 0})

	masterDB = make(map[string][]*database, 0)
	slaverDB = make(map[string][]*database, 0)

	for table, db := range env {
		for method, databases := range db.(map[string]interface{}) {
			var dbEngines []*database = make([]*database, len(databases.([]interface{})))
			for key, database := range databases.([]interface{}) {
				var cluster *dbCluster

				toMap := module.util.InterfaceToStringInMap(database.(map[interface{}]interface{}))
				toMap["Table"] = table

				E.ErrPanic(module.util.MapToStruct(toMap, &cluster))

				dbEngine := NewDatabase()
				dbEngine.dbCluster = cluster
				dbEngine.dbConfigs = configs

				dbEngines[key] = dbEngine.instance()
			}

			switch method {
			case "master":
				masterDB[table] = dbEngines
			case "slaver":
				slaverDB[table] = dbEngines
			default:
				continue
			}
		}
	}
}
