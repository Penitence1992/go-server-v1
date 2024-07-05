package capture

import (
	"github.com/penitence1992/go-server-v1/pkg/api"
	"github.com/penitence1992/go-server-v1/pkg/errors"
	"reflect"
	"testing"
)

func TestRegisterErrorCapture(t *testing.T) {
	e := errors.NewResourceNotFoundError("资源不存在")
	c := NewCatcher()
	c.RegisterErrorCapture(reflect.TypeOf(e), func(err error) *api.CwResponse {
		return &api.CwResponse{
			Code:    404,
			Msg:     "资源不存在",
			BizCode: errors.RequestError,
		}
	})
	b, res := c.TryConvert(e)
	if !b {
		t.Fatalf("转换不应该失败")
	}
	if res.Code != e.Code() &&
		res.Msg != e.Error() &&
		res.BizCode != e.BizCode() {
		t.Fatalf("")
	}
}
