package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
	"github.com/stevenfamy/go-send-message-slack/config"
)

type SendMessage struct {
	ChannelID string `json:"channelID"`
	Title     string `json:"title"`
	Message   string `json:"message"`
	Types     string `json:"types"`
}

func doSendMessage(client *slack.Client, channelID string, title string, message string, types string) {

	var color string = "#7393B3"
	if types == "success" {
		color = "#36A64F"
	} else if types == "warning" {
		color = "#E9D502"
	} else if types == "error" {
		color = "#FF0000"
	}

	attachment := slack.Attachment{
		Pretext: title,
		Text:    message,
		// Color Styles the Text, making it possible to have like Warnings etc.
		Color: color,
		// Fields are Optional extra data!
		// Fields: []slack.AttachmentField{
		// 	{
		// 		Title: "Date",
		// 		Value: time.Now().Format("2017-09-07 17:06:06"),
		// 	},
		// },
	}

	_, timestamp, err := client.PostMessage(
		channelID,
		// uncomment the item below to add a extra Header to the message, try it out :)
		//slack.MsgOptionText("New message from bot", false),
		slack.MsgOptionAttachments(attachment),
	)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Message sent at %s", timestamp)

}

func connectSlack() *slack.Client {
	// Load config
	token := config.GetConfig("SLACK_AUTH_TOKEN")
	appToken := config.GetConfig("SLACK_APP_TOKEN")

	// Create a new client to slack by giving token
	// Set debug to true while developing
	// Also add a ApplicationToken option to the client
	client := slack.New(token, slack.OptionDebug(false), slack.OptionAppLevelToken(appToken))

	return client
}

func ReceivedAPI(c *gin.Context) {
	var sendMessage SendMessage

	if err := c.ShouldBindJSON(&sendMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := connectSlack()

	doSendMessage(client, sendMessage.ChannelID, sendMessage.Title, sendMessage.Message, sendMessage.Types)

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
