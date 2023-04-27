package mongox

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"golang_log/library/api-demo/core/logx"
)

type Config struct {
	Uri             string `mapstructure:"uri"`                // 连接地址
	Database        string `mapstructure:"database"`           // 指定数据库
	MaxPoolSize     uint64 `mapstructure:"max_pool_size"`      // 最大连接的数量
	MinPoolSize     uint64 `mapstructure:"min_pool_size"`      // 最小连接的数量
	MaxConnIdleTime uint   `mapstructure:"max_conn_idle_time"` // 最大连接的空闲时间(设置了连接可复用的最大时间)(单位秒)
}

// Mongox 客户端
type Mongox struct {
	client          *mongo.Client
	defaultDatabase *mongo.Database // 默认指定的数据库
	databaseMap     *sync.Map       // 选择其他指定的数据库
}

// Ping 判断服务是否可用
func (m *Mongox) Ping() error {
	return m.client.Ping(context.Background(), readpref.Primary())
}

// Client get mongo client
func (m *Mongox) Client() *mongo.Client {
	return m.client
}

// Database 指定数据库
func (m *Mongox) Database(name string, opts ...*options.DatabaseOptions) *mongo.Database {
	loadDatabase, ok := m.databaseMap.Load(name)
	if ok {
		return loadDatabase.(*mongo.Database)
	}
	database := m.client.Database(name)
	m.databaseMap.Store(name, database)
	return database
}

// Collection 选择集合(表)
func (m *Mongox) Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection {
	return m.defaultDatabase.Collection(name, opts...)
}

var db *Mongox

// DB get mongox
func DB() *Mongox {
	return db
}

// Ping 判断服务是否可用
func Ping() error {
	return db.Ping()
}

// Database 指定数据库
func Database(name string, opts ...*options.DatabaseOptions) *mongo.Database {
	return db.Database(name, opts...)
}

// Collection 选择集合(表)
func Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection {
	return db.Collection(name, opts...)
}

// Connect 连接数据库
func Connect(c *Config) (*Mongox, error) {
	opts := options.Client()
	opts.ApplyURI(c.Uri)

	// 进行配置
	if c.MaxPoolSize > 0 {
		opts.SetMaxPoolSize(c.MaxPoolSize)
	}
	if c.MinPoolSize > 0 {
		opts.SetMinPoolSize(c.MinPoolSize)
	}
	if c.MaxConnIdleTime > 0 {
		opts.SetMaxConnIdleTime(time.Duration(c.MaxConnIdleTime) * time.Second)
	}

	// 链接 mongo 服务
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	// 判断服务是否可用
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &Mongox{
		client:          client,
		defaultDatabase: client.Database(c.Database),
		databaseMap:     &sync.Map{},
	}, nil
}

// Init 初始化数据库
func Init(c *Config) {
	logx.Debugf("mongo Config=%v", c)
	logx.Debug("mongo 开始初始化...")

	var err error
	db, err = Connect(c)
	if err != nil {
		panic(err)
	}

	logx.Debug("mongo 初始化完毕...")
}
