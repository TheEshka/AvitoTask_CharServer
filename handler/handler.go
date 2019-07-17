package handler

import (
	"log"
	"net/http"

	"github.com/TheEshka/AvitoTask_CharServer/model"
	"github.com/gin-gonic/gin"
)

//Start :
func Start(m *model.Model, listenPort string) {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/users/add", addUser(m))

	// chats := r.Group("/chats")
	// {
	// 	chats.POST("/add", addChat(m))
	// 	//chats.POST("/get")
	// }

	// messages := r.Group("/messages")
	// {
	// 	messages.POST("/add", addMessage(m))
	// 	//messages.POST("/get")
	// }

	// r.POST("/user", createUser(m))
	// r.PATCH("/user", alterUser(m))
	// r.DELETE("/user", deleteUser(m))
	// r.GET("/user", getUser(m))

	// r.POST("/auth", verifyUser(m))

	r.Run(listenPort)
}

func addUser(m *model.Model) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json *model.User
		log.Println("First JsSON", json)

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := m.CreateUser(json.Username)

		switch err {
		case model.ErrOnDatabase:
			c.JSON(http.StatusServiceUnavailable, gin.H{"error code": err.Error()})
			log.Println(err)
			return
		case model.ErrAlreadyExist:
			c.JSON(http.StatusBadRequest, gin.H{"error code": err.Error()})
			log.Println(err)
			return
		case model.ErrPasswordFormat:
			c.JSON(http.StatusBadRequest, gin.H{"error code": err.Error()})
			log.Println(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	}
}

func addChat(m *model.Model) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json *model.Chat

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := m.CreateUser(json.Name)

		switch err {
		case model.ErrOnDatabase:
			c.JSON(http.StatusServiceUnavailable, gin.H{"error code": err.Error()})
			log.Println(err)
			return
		case model.ErrAlreadyExist:
			c.JSON(http.StatusBadRequest, gin.H{"error code": err.Error()})
			log.Println(err)
			return
		case model.ErrPasswordFormat:
			c.JSON(http.StatusBadRequest, gin.H{"error code": err.Error()})
			log.Println(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	}
}

func addMessage(m *model.Model) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json *model.Message

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := m.CreateMessage(json.ChatID, json.AuthorID, json.Text)

		switch err {
		case model.ErrOnDatabase:
			c.JSON(http.StatusServiceUnavailable, gin.H{"error code": err.Error()})
			log.Println(err)
			return
		case model.ErrAlreadyExist:
			c.JSON(http.StatusBadRequest, gin.H{"error code": err.Error()})
			log.Println(err)
			return
		case model.ErrPasswordFormat:
			c.JSON(http.StatusBadRequest, gin.H{"error code": err.Error()})
			log.Println(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	}
}
