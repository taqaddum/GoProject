package component

import (
	"GoProject/main/config"
	_ "github.com/lib/pq"
	"log"
	"sync"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var oncePostgres sync.Once
var engine *xorm.Engine

func NewPostgres() *xorm.Engine {
	oncePostgres.Do(
		func() {
			var err error
			engine, err = xorm.NewEngine(config.GetDBMS(), config.GetDSN())
			if err != nil || engine == nil {
				log.Fatalf("数据库引擎初始化错误 %s", err.Error())
			}

			//设置数据库连接参数
			engine.SetMapper(names.GonicMapper{})
			engine.ShowSQL(false)
			engine.SetMaxOpenConns(100)
			engine.SetConnMaxIdleTime(5 * time.Minute)
			engine.SetMaxIdleConns(10)
		})

	if err := engine.Ping(); err != nil {
		log.Fatalf("数据库连接失败 %s", err.Error())
	}

	return engine
}
