package storage

import (
	"errors"
	log "github.com/sirupsen/logrus"
)

type PluginType string

var PLUGIN_EXISTS = errors.New("插件已存在")

const (
	POSTGRES PluginType = "postgres"
	SQLITE   PluginType = "sqlite"
)

var pluginRegistry = make(map[PluginType]ClientCreateFunc)

func Create(setting DatabaseSetting) (JDBC, error) {
	name := setting.DbType
	if f, err := GetStoragePlugins(name); err != nil && f != nil {
		return nil, errors.New("数据库存储插件:" + string(name) + "不存在")
	} else {
		log.Infof("创建:%s数据库, 地址: %s", name, setting.Host)
		return f(setting), nil
	}
}

func CreateWithPool(setting DatabaseSetting, pool Pool) (JDBC, error) {
	jdbc, err := Create(setting)
	if err != nil {
		return jdbc, err
	}
	sqlDb, err := jdbc.MustGetDB().DB()
	if err != nil {
		return jdbc, err
	}
	if pool.MaxIdle != 0 {
		sqlDb.SetMaxIdleConns(pool.MaxIdle)
	}
	if pool.MaxOpen != 0 {
		sqlDb.SetMaxOpenConns(pool.MaxOpen)
	}
	if pool.MaxLifetime != 0 {
		sqlDb.SetConnMaxLifetime(pool.MaxLifetime)
	}
	if pool.MaxIdleTime != 0 {
		sqlDb.SetConnMaxIdleTime(pool.MaxIdleTime)
	}
	return jdbc, err
}

func GetStoragePlugins(name PluginType) (ClientCreateFunc, error) {
	p := pluginRegistry[name]
	if p == nil {
		return nil, errors.New("数据库存储插件:" + string(name) + "不存在")
	}
	return p, nil
}

func Register(name PluginType, f ClientCreateFunc) error {
	if pluginRegistry[name] == nil {
		pluginRegistry[name] = f
		return nil
	} else {
		return PLUGIN_EXISTS
	}
}
