package actuator

import "github.com/gin-gonic/gin"

func RegistryActuatorEndpoint(engine *gin.Engine) {
	engine.GET("swagger", swagger)
	g := engine.Group("actuator")
	g.GET("info", info)
	g.GET("health", health)
	h := g.Group("health")
	h.GET("liveness", healthLiveness)
	h.GET("readiness", healthReadiness)
}
