package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/penitence1992/go-server-v1/pkg/utils"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type GinRegisterFunc = func(engine *gin.Engine)

type Server struct {
	Port        uint16
	BindAddress string
	httpServer  *http.Server
	route       *gin.Engine
	sslMode     bool
}

func CreateServer(port uint16, bindAddress string, tlsConfig *tls.Config) *Server {
	address := fmt.Sprintf("%s:%d", bindAddress, port)
	httpServerConfig := &Server{
		Port:        port,
		BindAddress: bindAddress,
		httpServer: &http.Server{
			Addr:    address,
			Handler: nil,
		},
		route: gin.Default(),
	}
	if tlsConfig != nil {
		httpServerConfig.httpServer.TLSConfig = tlsConfig
		httpServerConfig.sslMode = true
	}
	return httpServerConfig
}

func (s Server) RegisterRoute(f GinRegisterFunc) {
	f(s.route)
}

func (s Server) StartListen() {
	done := utils.CreateInterruptChan()
	go func() {
		sg := <-done
		log.Infof("accept signal %s", sg.String())
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := s.httpServer.Shutdown(ctx); err != nil {
			log.Error("Shutdown server:", err)
		} else {
			log.Info("Shutdown server")
		}
	}()

	log.Infof("Start http server with %s:%v ...\n", s.BindAddress, s.Port)
	s.httpServer.Handler = s.route
	var err error
	if s.sslMode {
		err = s.httpServer.ListenAndServeTLS("", "")
	} else {
		err = s.httpServer.ListenAndServe()
	}
	if err != nil {
		if err == http.ErrServerClosed {
			log.Warn("Server closed under request")
		} else {
			log.Info("Server closed unexpected")
		}
	}
}
