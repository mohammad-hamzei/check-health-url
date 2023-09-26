package rest

import (
	"fmt"

	"git.tashilcar.com/core/tashilcar/presentation/rest/routes"
	"github.com/gin-gonic/gin"
)

func Start() error {
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run(":8080")
	return nil
}

func Stop() {
	fmt.Println("* STOPPING GRPC SERVER *")
}
