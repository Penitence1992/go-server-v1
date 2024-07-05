package model

import (
	"github.com/penitence1992/go-server-v1/pkg/domain"
)

// 测试表数据
// swagger:model testTable
type TestTable struct {
	domain.BaseColumn

	// id
	// required: true
	ID string `json:"id"`
	// name
	// required: true
	Name string `json:"name"`
}

func (t TestTable) TableName() string {
	return "test_table"
}
