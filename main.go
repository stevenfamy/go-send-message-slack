package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stevenfamy/go-send-message-slack/routes"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")
	{
		routes.SlackRoutes(api)
	}
	r.Run(":8080")

}

//hehe
