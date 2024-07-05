package starter

import (
	"crypto/tls"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	_ "github.com/penitence1992/go-gin-server/pkg/app/config/logger"
	"github.com/penitence1992/go-gin-server/pkg/capture"
	"github.com/penitence1992/go-gin-server/pkg/server"
	"github.com/penitence1992/go-gin-server/pkg/server/actuator"
	"github.com/penitence1992/go-gin-server/pkg/server/middleware"
	log "github.com/sirupsen/logrus"
	"reflect"
)

type ServerStarter struct {
	Port                     uint16
	Bind                     string
	TlsConfig                *tls.Config
	f                        server.GinRegisterFunc
	UsePprof                 bool
	catcher                  *capture.Catcher
	registryDefaultEndpoint  bool
	commonMiddlewareRegister bool
}

func NewServerStart(port uint16, bind string, f server.GinRegisterFunc) *ServerStarter {
	return &ServerStarter{
		Port:                     port,
		Bind:                     bind,
		f:                        f,
		catcher:                  newCatcher(),
		commonMiddlewareRegister: true,
	}
}

func (s *ServerStarter) SetTlsConfig(tls *tls.Config) *ServerStarter {
	s.TlsConfig = tls
	return s
}

func (s *ServerStarter) PprofEnabled(enabled bool) *ServerStarter {
	s.UsePprof = enabled
	return s
}

func (s *ServerStarter) SetCommandMiddlewareEnabled(enabled bool) *ServerStarter {
	s.commonMiddlewareRegister = enabled
	return s
}

func (s *ServerStarter) IsRegistryActuatorEndpoint(open bool) *ServerStarter {
	s.registryDefaultEndpoint = open
	return s
}

// AddErrorCapture 添加异常捕捉函数
func (s *ServerStarter) AddErrorCapture(i error, f capture.ErrorCapture) {
	s.catcher.RegisterErrorCapture(reflect.TypeOf(i), f)
}

// Start 启动http server
func (s *ServerStarter) Start() {
	httpServer := server.CreateServer(s.Port, s.Bind, s.TlsConfig)
	httpServer.RegisterRoute(func(engine *gin.Engine) {
		if s.commonMiddlewareRegister {
			engine.Use(middleware.NewRecover(s.catcher))
			engine.Use(middleware.Cors())
		}
		s.f(engine)
		if s.UsePprof {
			log.Infof("开启Pprof内存分析")
			pprof.Register(engine)
		}
		if s.registryDefaultEndpoint {
			actuator.RegistryActuatorEndpoint(engine)
		}
	})
	httpServer.StartListen()
}

func newCatcher() *capture.Catcher {
	c := capture.NewCatcher()
	// c.RegisterErrorCapture() 添加默认捕捉
	return c
}

// StartDaemon 启动http server 服务
// Deprecated : 请使用 NewServerStart 来创建Starter对象, 然后配置后使用 Start 方法来启动
func StartDaemon(port uint16, bind string, tlsConfig *tls.Config, f server.GinRegisterFunc) {
	httpServer := server.CreateServer(port, bind, tlsConfig)
	httpServer.RegisterRoute(f)
	httpServer.StartListen()
}
