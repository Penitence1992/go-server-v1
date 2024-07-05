package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/penitence1992/go-server-v1/internal/model"
	"github.com/penitence1992/go-server-v1/pkg/api"
	cerrors "github.com/penitence1992/go-server-v1/pkg/errors"
	"github.com/penitence1992/go-server-v1/pkg/storage"
	"github.com/penitence1992/go-server-v1/pkg/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const Group = "v1"

type Hello struct {
	client storage.JDBC
	db     *gorm.DB
}

func RegistryTestTableRoute(g *gin.Engine, client storage.JDBC) error {
	if client == nil {
		return errors.New("client不能为空")
	}
	h := &Hello{
		client: client,
		db:     client.MustGetDB(),
	}
	group := g.Group(Group).Group("test")
	group.GET("/:id", utils.Wrap(h.getById))
	group.GET("", utils.Wrap(h.getPage))
	group.POST("", utils.Wrap(h.create))
	return nil
}

// getById 通过id获取数据
// swagger:route GET /:id test getTestById
// 通过id来获取数据
// Consumes:
// - application/json
// Produces:
// - application/json
// Schemes: http
// Parameters:
//   - name: id
//     in: path
//     description: id
//     type: string
//     required: true
//
// Responses:
//
//	200:
//	  message: ok
//	  schema: $/definitions/testTable
func (h *Hello) getById(c *gin.Context) interface{} {
	id := c.Param("id")
	if id == "" {
		panic(cerrors.ErrResourceNotFound)
	}
	logrus.Infof("获取%s的数据", id)
	db := h.db
	var u model.TestTable
	err := db.Where("id = ?", id).First(&u).Error
	if err == gorm.ErrRecordNotFound {
		panic(cerrors.ErrResourceNotFound)
	}
	utils.PanicIfNotNil(err)
	return u
}

func (h *Hello) getPage(c *gin.Context) interface{} {
	db := h.db
	var u []model.TestTable
	db.Limit(10).Find(&u)
	return api.NewPage(api.Pageable{
		Page: 0, Size: 10,
	}, 100, len(u), u)
}

func (h *Hello) create(c *gin.Context) interface{} {
	var u model.TestTable
	utils.PanicIfNotNil(c.Bind(&u))
	db := h.db
	u.ID = uuid.NewString()
	db.Save(&u)
	return u.ID
}
