package storage

import (
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/penitence1992/go-server-v1/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	url2 "net/url"
	"sync"
)

func init() {
	e := Register(POSTGRES, NewPgClient)
	utils.PanicIfNotNil(e)
}

type PgClient struct {
	setting DatabaseSetting
	db      *gorm.DB
	lock    *sync.Mutex
}

func NewPgClient(setting DatabaseSetting) JDBC {
	return &PgClient{
		setting: setting,
		lock:    &sync.Mutex{},
	}
}

func (p PgClient) GetDB() (*gorm.DB, error) {
	if p.db != nil {
		return p.db, nil
	}
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.db != nil {
		return p.db, nil
	}

	db, err := gorm.Open(postgres.Open(p.createDsn()))
	if err != nil {
		return nil, err
	}
	// setting
	p.db = db
	return p.db, nil
}

// MustGetDB 必须获取到db, 否则执行 panic操作
func (p PgClient) MustGetDB() *gorm.DB {
	db, err := p.GetDB()
	if err != nil {
		panic(err)
	}
	return db
}

func (p PgClient) GetMigrateUrl(migrationsTableName string) string {
	return p.createMigrateUrl(migrationsTableName)
}

func (p *PgClient) createDsn() string {
	s := p.setting
	port := s.Port
	if port == 0 {
		port = 5432
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d TimeZone=Asia/Shanghai",
		s.Host, s.Username, s.Password, s.Database, port)
	if s.ExtSetting != nil {
		for k, v := range s.ExtSetting {
			dsn += " " + k + "=" + v
		}
	}
	return dsn
}

func (p *PgClient) createMigrateUrl(migrationsTableName string) string {
	s := p.setting
	port := s.Port
	if port == 0 {
		port = 5432
	}
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?%s=%s",
		s.Username, url2.QueryEscape(s.Password), s.Host, port, s.Database, MigrationsQueryName, migrationsTableName,
	)

	if s.ExtSetting != nil {
		for k, v := range s.ExtSetting {
			url += "&" + k + "=" + v
		}
	}
	return url
}
