package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stevenfamy/go-send-message-slack/models"
	"github.com/stevenfamy/go-send-message-slack/routes"
)

func main() {
	models.ConnectDatabase()

	r := gin.Default()
	api := r.Group("/api")
	{
		routes.SlackRoutes(api)
		routes.TestServerRoutes(api)
	}
	r.Run(":8080")

}

//hehe
