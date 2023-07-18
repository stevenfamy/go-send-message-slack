package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/stevenfamy/go-send-message-slack/controllers"
	"github.com/stevenfamy/go-send-message-slack/middleware"
)

func SlackRoutes(api *gin.RouterGroup) {
	const routesGroup string = "/slack"
	api.POST(routesGroup+"/send-message", middleware.AuthorizationMiddleware(), controllers.ReceivedAPI)
}
