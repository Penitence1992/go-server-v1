package storage

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type ClientCreateFunc func(setting DatabaseSetting) JDBC

type JDBC interface {
	GetDB() (*gorm.DB, error)
	MustGetDB() *gorm.DB
	GetMigrateUrl(migrationsTableName string) string
}

type DatabaseSetting struct {
	DbType     PluginType
	Host       string
	Username   string
	Password   string
	Database   string
	Port       int
	ExtSetting map[string]string
}

func NewDatabaseSetting(host, username, password, database string, port int) DatabaseSetting {
	return DatabaseSetting{
		Host:     host,
		Username: username,
		Password: password,
		Database: database,
		Port:     port,
	}
}

func (d *DatabaseSetting) Validate() error {
	if d.DbType == "" {
		return errors.New("dbT不能为空")
	}
	if d.Host == "" {
		return errors.New("host不能为空")
	}
	if d.Username == "" {
		return errors.New("username不能为空")
	}
	if d.Password == "" {
		return errors.New("password不能为空")
	}
	if d.Database == "" {
		return errors.New("database不能为空")
	}
	return nil
}

type Pool struct {
	MaxIdle     int
	MaxOpen     int
	MaxLifetime time.Duration
	MaxIdleTime time.Duration
}

// NewDefaultPool 创建默认连接池配置
func NewDefaultPool() *Pool {
	return &Pool{
		MaxIdle:     4,
		MaxOpen:     16,
		MaxLifetime: 30 * time.Minute,
	}
}
