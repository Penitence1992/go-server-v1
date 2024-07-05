package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/penitence1992/go-server-v1/pkg/api"
	"github.com/penitence1992/go-server-v1/pkg/capture"
	"github.com/penitence1992/go-server-v1/pkg/errors"
	"github.com/sirupsen/logrus"
)

func NewRecover(catcher *capture.Catcher) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logrus.Errorf("捕获异常, 异常信息: %v", r)
				if cwE, ok := r.(errors.CwError); ok {
					c.JSON(cwE.Code(), &api.CwResponse{
						Code:    cwE.Code(),
						Data:    cwE.Data(),
						BizCode: cwE.BizCode(),
						Msg:     cwE.Error(),
					})
				} else {
					if e, ok := r.(error); ok {
						if o, res := catcher.TryConvert(e); o {
							c.JSON(res.Code, res)
						} else {
							c.JSON(500, unknownErrorResponse())
						}
					} else {
						c.JSON(500, unknownErrorResponse())
					}
				}
			}
		}()
		c.Next()
	}
}

func unknownErrorResponse() *api.CwResponse {
	return &api.CwResponse{
		Code:    500,
		Data:    nil,
		BizCode: errors.ServerError,
		Msg:     "服务器错误",
	}
}
