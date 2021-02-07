package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/socket/v1")
	{
		v1.POST("/group", createGroup)
		v1.GET("/group/:id", fetchGroup)
		v1.POST("/group/:id/client", createClientWithGroupId)
		v1.POST("/group/:id/message", sendGroupMessage)
		v1.POST("/client/:id/message", sendMessageToClient)
		v1.DELETE("/group/:id/client/:id", deleteClient)
		v1.DELETE("/group/:id", deleteGroup)
		v1.POST("/auth/user", createAuthUser)
		v1.GET("/ws", func(c *gin.Context) {
			clientId := c.Param("clientId")
			if clientId == "" {
				clientId = uuid.NewString()
			}
			log.Print("clientId:", clientId)
			wsHandler(c.Writer, c.Request, clientId)
		})
	}

	router.Run(":5050")
}

func createAuthUser(context *gin.Context) {

}

func deleteGroup(context *gin.Context) {

}

func deleteClient(context *gin.Context) {

}

func sendMessageToClient(context *gin.Context) {
	bodyBytes, _ := ioutil.ReadAll(context.Request.Body)
	sendMessage(bodyBytes, context.Param("id"))
}

func sendGroupMessage(context *gin.Context) {

}

func createClientWithGroupId(context *gin.Context) {

}

func fetchGroup(context *gin.Context) {

}

func createGroup(context *gin.Context) {

}
