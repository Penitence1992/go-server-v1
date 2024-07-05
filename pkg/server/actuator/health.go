package actuator

import "github.com/gin-gonic/gin"

type Health struct {
	Status string   `json:"status"`
	Groups []string `json:"groups,omitempty"`
}

func health(context *gin.Context) {
	context.JSON(200, &Health{
		Status: "UP",
		Groups: []string{
			"liveness",
			"readiness",
		},
	})
}

func healthLiveness(context *gin.Context) {
	context.JSON(200, &Health{
		Status: "UP",
	})
}

func healthReadiness(context *gin.Context) {
	context.JSON(200, &Health{
		Status: "UP",
	})
}
