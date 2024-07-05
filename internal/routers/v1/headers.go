package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/penitence1992/go-gin-server/pkg/utils"
	"strings"
)

func RegistryHeadersRoute(g *gin.Engine) error {
	group := g.Group("v1").Group("headers")
	group.Any("", utils.Wrap(printHeader))
	return nil
}

func printHeader(c *gin.Context) interface{} {
	h := make(map[string]string)
	for k, vs := range c.Request.Header {
		h[k] = strings.Join(vs, ",")
	}
	return h
}
