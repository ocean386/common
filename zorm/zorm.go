package zorm

import (
	"fmt"
	"github.com/ocean386/common/sharding"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"gorm.io/plugin/prometheus"
)

func NewClient(config *DBConfig) *gorm.DB {

	dbMain := getConn(config.MainSource.SchemaConfig, config.Mode)
	for _, shading := range config.Sharding {
		var args []interface{}
		for _, table := range shading.Tables {
			args = append(args, table)
		}
		middleware := sharding.Register(sharding.Config{
			ShardingKey:    shading.ShardingKey,
			NumberOfShards: uint(shading.NumberOfShards),
			Name:           shading.Name,
		}, args...)
		err := dbMain.Use(middleware)
		if err != nil {
			panic(err)
		}
	}
	var dbMainRes []gorm.Dialector
	for _, replica := range config.MainSource.Replicas {
		dbMainRes = append(dbMainRes, getConn(replica, config.Mode).Dialector)
	}

	mainSource := dbresolver.Register(
		dbresolver.Config{
			Replicas: dbMainRes,
			Policy:   dbresolver.RandomPolicy{},
		},
	)
	for _, source := range config.OtherSource {

		var res []gorm.Dialector
		for _, replica := range source.Replicas {
			db := getConn(replica, config.Mode)
			res = append(res, db.Dialector)
		}
		var s []gorm.Dialector
		db := getConn(source.SchemaConfig, config.Mode)
		s = append(s, db.Dialector)
		var args []interface{}
		for _, table := range source.Tables {
			args = append(args, table)
		}
		mainSource.Register(
			dbresolver.Config{
				Sources:  s,
				Replicas: res,
				Policy:   dbresolver.RandomPolicy{},
			}, args...)
	}

	err := dbMain.Use(mainSource)

	if err != nil {
		panic(err)
	}

	err = dbMain.Use(prometheus.New(prometheus.Config{
		DBName:          config.Prometheus.Name,                    // 使用 `DBName` 作为指标 label
		RefreshInterval: uint32(config.Prometheus.RefreshInterval), // 指标刷新频率(默认15秒)
		//	PushAddr:        "prometheus pusher address",           //  `PushAddr`推送指标
		StartServer:    true,                                     // 启用一个 http 服务来暴露指标
		HTTPServerPort: uint32(config.Prometheus.HTTPServerPort), // HTTP服务监听端口
		MetricsCollector: []prometheus.MetricsCollector{
			&prometheus.MySQL{
				VariableNames: []string{"Threads_running"},
			},
		}, // 用户自定义指标
	}))

	if err != nil {
		panic(err)
	}

	return dbMain
}

func getConn(config SchemaConfig, mode string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host,
		config.Port, config.DBName,
	)
	Db, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{Logger: NewGormLogger(mode)})
	if err != nil {
		panic(err)
	}
	db, err := Db.DB()
	db.SetMaxIdleConns(config.MaxIdelConn)
	db.SetMaxOpenConns(config.MaxOpenConn)
	if err != nil {
		panic(err)
	}
	return Db
}
