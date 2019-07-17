package handler

import (
	"fmt"
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

	chats := r.Group("/chats")
	{
		chats.POST("/add", addChatWithUsers(m))
		chats.POST("/get", getUserChats(m))
	}

	messages := r.Group("/messages")
	{
		messages.POST("/add", addMessage(m))
		messages.POST("/get", getChatMessages(m))
	}

	r.Run(listenPort)
}

func addUser(m *model.Model) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json *model.User

		if err := c.ShouldBindJSON(&json); err != nil {
			c.String(http.StatusBadRequest, "")
			return
		}

		id, err := m.CreateUser(json.Username)

		switch err {
		case model.ErrOnDatabase:
			c.String(http.StatusServiceUnavailable, "")
			log.Println(err)
			return
		case model.ErrRequest:
			c.String(http.StatusBadRequest, "")
			log.Println(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	}
}

func addChatWithUsers(m *model.Model) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json *model.CreateChatRequest

		if err := c.ShouldBindJSON(&json); err != nil {
			c.String(http.StatusBadRequest, "")
			return
		}
		fmt.Println(json)

		id, err := m.CreateChatWithUsers(json.Name, json.Users)

		switch err {
		case model.ErrOnDatabase:
			c.String(http.StatusServiceUnavailable, "")
			log.Println(err)
			return
		case model.ErrRequest:
			c.String(http.StatusBadRequest, "")
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
			c.String(http.StatusBadRequest, "")
			return
		}

		id, err := m.CreateMessage(json.ChatID, json.AuthorID, json.Text)

		switch err {
		case model.ErrOnDatabase:
			c.String(http.StatusServiceUnavailable, "")
			log.Println(err)
			return
		case model.ErrRequest:
			c.String(http.StatusBadRequest, "")
			log.Println(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	}
}

func getUserChats(m *model.Model) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json *model.GetUserChatsRequest

		if err := c.ShouldBindJSON(&json); err != nil {
			c.String(http.StatusBadRequest, "")
			return
		}

		chats, err := m.UserChats(json.UserID)

		switch err {
		case model.ErrOnDatabase:
			c.String(http.StatusServiceUnavailable, "")
			log.Println(err)
			return
		case model.ErrRequest:
			c.String(http.StatusBadRequest, "")
			log.Println(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"chats": chats,
		})
	}
}

func getChatMessages(m *model.Model) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json *model.GetChatsMessagesRequest

		if err := c.ShouldBindJSON(&json); err != nil {
			c.String(http.StatusBadRequest, "")
			return
		}

		messages, err := m.ChatMessages(json.ChatID)

		switch err {
		case model.ErrOnDatabase:
			c.String(http.StatusServiceUnavailable, "")
			log.Println(err)
			return
		case model.ErrRequest:
			c.String(http.StatusBadRequest, "")
			log.Println(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"messages": messages,
		})
	}
}
