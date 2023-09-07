package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stevenfamy/go-send-message-slack/models"
)

type UpdateBuild struct {
	FeBranch string `json:"feBranch"`
	BeBranch string `json:"beBranch"`
}

func DoUpdateServerBuild(c *gin.Context) {
	var updateBuild UpdateBuild

	projectId := c.Param("projectId")
	serverId := c.Param("serverId")
	name := c.Param("name")

	if err := c.ShouldBindJSON(&updateBuild); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.UpdateServerBuild(projectId, serverId, name, updateBuild.FeBranch, updateBuild.BeBranch)

	c.JSON(http.StatusOK, gin.H{"message": "success"})

}
