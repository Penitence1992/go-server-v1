package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/penitence1992/go-gin-server/pkg/api"
)

func PanicIfNotNil(e error) {
	if e != nil {
		panic(e)
	}
}

func Wrap(f func(*gin.Context) interface{}) gin.HandlerFunc {
	return func(c2 *gin.Context) {
		data := f(c2)
		c2.JSON(200, api.Ok(data))
	}
}
