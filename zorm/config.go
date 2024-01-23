package zorm

type SchemaConfig struct {
	Host         string `json:",env=DATABASE_HOST"`
	Port         int    `json:",env=DATABASE_PORT"`
	Username     string `json:",default=root,env=DATABASE_USERNAME"`
	Password     string `json:",optional,env=DATABASE_PASSWORD"`
	DBName       string `json:",env=DATABASE_DBNAME"`
	SSLMode      string `json:",optional,default=disable,env=DATABASE_SSL_MODE"`
	Type         string `json:",default=mysql,options=[mysql,postgres,sqlite3],env=DATABASE_TYPE"`
	MaxOpenConn  int    `json:",optional,default=250,env=DATABASE_MAX_OPEN_CONN"`
	MaxIdelConn  int    `json:",optional,default=250,env=DATABASE_MAX_IDEL_CONN"`
	CacheTime    int    `json:",optional,default=10,env=DATABASE_CACHE_TIME"`
	DBPath       string `json:",optional,env=DATABASE_DBPATH"`
	MysqlConfig  string `json:",optional,env=DATABASE_MYSQL_CONFIG"`
	PGConfig     string `json:",optional,env=DATABASE_PG_CONFIG"`
	SqliteConfig string `json:",optional,env=DATABASE_SQLITE_CONFIG"`
}

type DBConfig struct {
	MainSource  Source
	OtherSource []Source     `json:",optional"`
	Prometheus  DBPrometheus `json:",optional"`
	Sharding    []DBSharding `json:",optional"`
	Mode        string       `json:",optional"`
	DisableLog  bool         `json:",optional"`
}

type Source struct {
	SchemaConfig
	Replicas []SchemaConfig `json:",optional"`
	Tables   []string       `json:",optional"`
}

type DBPrometheus struct {
	Name            string `json:",default=zorm"`
	RefreshInterval int    `json:",default=5"`
	HTTPServerPort  int    `json:",default=8084"`
}

type DBSharding struct {
	Name           string
	ShardingKey    string
	NumberOfShards int
	Tables         []string
}
