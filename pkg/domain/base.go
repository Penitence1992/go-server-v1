package domain

import "time"

type BaseColumn struct {
	Created time.Time `gorm:"autoCreateTime" json:"created"`
	Updated time.Time `gorm:"autoUpdateTime" json:"updated"`
	Del     bool
}
