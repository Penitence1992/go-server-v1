package actuator

import (
	"github.com/gin-gonic/gin"
	"os"
)

type Info struct {
	Git GitInfo `json:"git"`
}

type GitInfo struct {
	Branch string    `json:"branch"`
	Commit GitCommit `json:"commit"`
}

type GitCommit struct {
	Id   string `json:"id"`
	Time string `json:"time"`
}

var i = NewInfo()

func NewInfo() Info {
	id := os.Getenv("COMMIT_ID")
	buildTime := os.Getenv("BUILD_TIME")
	return Info{
		Git: GitInfo{
			Branch: "HEAD",
			Commit: GitCommit{
				Id:   id,
				Time: buildTime,
			},
		},
	}
}

func info(context *gin.Context) {
	context.JSON(200, i)
}
