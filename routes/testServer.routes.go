package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/stevenfamy/go-send-message-slack/controllers"
	"github.com/stevenfamy/go-send-message-slack/middleware"
)

func TestServerRoutes(api *gin.RouterGroup) {
	const routesGroup string = "/test-server"
	api.POST(routesGroup+"/update-build/:projectId/:serverId/:name", middleware.AuthorizationMiddleware(), controllers.DoUpdateServerBuild)
}
