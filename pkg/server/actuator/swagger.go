package actuator

import (
	"github.com/gin-gonic/gin"
	"os"
	"sigs.k8s.io/yaml"
)

var SwaggerDoc = "{}"

func init() {
	blob, err := os.ReadFile("swagger.yml")
	if err != nil {
		blob, err = os.ReadFile("swagger.yaml")
		if err != nil {
			return
		}
	}
	if data, err := yaml.YAMLToJSON(blob); err == nil {
		SwaggerDoc = string(data)
	}
}

func swagger(context *gin.Context) {
	context.Status(200)
	context.Writer.WriteString(SwaggerDoc)
	context.Header("Content-Type", "application/json")
}
