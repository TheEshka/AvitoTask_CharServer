package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/TheEshka/AvitoTask_CharServer/model"
	"github.com/gin-gonic/gin"
	"github.com/xeipuuv/gojsonschema"
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
		body, _ := ioutil.ReadAll(c.Request.Body)
		result, err := gojsonschema.Validate(gojsonschema.NewStringLoader(model.AddUserSchema),
			gojsonschema.NewBytesLoader(body))
		if (err != nil) || (!result.Valid()) {
			c.String(http.StatusBadRequest, "")
			return
		}

		var user *model.User
		json.Unmarshal(body, &user)

		id, err := m.CreateUser(user.Username)
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

		c.JSON(http.StatusCreated, gin.H{
			"id": strconv.FormatInt(id, 10),
		})
	}
}

func addChatWithUsers(m *model.Model) gin.HandlerFunc {
	return func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		result, err := gojsonschema.Validate(gojsonschema.NewStringLoader(model.AddChatSchema),
			gojsonschema.NewBytesLoader(body))
		if (err != nil) || (!result.Valid()) {
			c.String(http.StatusBadRequest, "")
			return
		}

		var chatCreate *model.CreateChatRequest
		json.Unmarshal(body, &chatCreate)

		id, err := m.CreateChatWithUsers(chatCreate.Name, chatCreate.Users)
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

		c.JSON(http.StatusCreated, gin.H{
			"id": strconv.FormatInt(id, 10),
		})
	}
}

func addMessage(m *model.Model) gin.HandlerFunc {
	return func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		result, err := gojsonschema.Validate(gojsonschema.NewStringLoader(model.AddMessageSchema),
			gojsonschema.NewBytesLoader(body))
		if (err != nil) || (!result.Valid()) {
			c.String(http.StatusBadRequest, "")
			return
		}

		var message *model.Message
		json.Unmarshal(body, &message)

		id, err := m.CreateMessage(message.ChatID, message.AuthorID, message.Text)
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

		c.JSON(http.StatusCreated, gin.H{
			"id": strconv.FormatInt(id, 10),
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
