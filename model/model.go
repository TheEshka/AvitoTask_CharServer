package model

import (
	"errors"
)

//User : general information about chat
type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
}

//Chat : general information about chat
type Chat struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

//Message : general information about message
type Message struct {
	ID        string `json:"id"`
	ChatID    string `json:"chat"`
	AuthorID  string `json:"author"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

//CreateChatRequest : request body for binding
type CreateChatRequest struct {
	Name  string   `json:"name"`
	Users []string `json:"users"`
}

//GetUserChatsRequest : request body for binding
type GetUserChatsRequest struct {
	UserID string `json:"user"`
}

//GetChatsMessagesRequest : request body for binding
type GetChatsMessagesRequest struct {
	ChatID string `json:"chat"`
}

//ErrOnDatabase : error after sending request  to database
var ErrOnDatabase = errors.New("Database connection error")

//ErrRequest : error after sending request  to database
var ErrRequest = errors.New("Not rught request")

//AddUserSchema : JSON Schema for add user request
var AddUserSchema = `
{
	"title": "Add User Schema",
	"type": "object",
	"properties": {
		"username": {
			"type": "string"
		}
	},
	"required": ["username"]
}
`

//AddChatSchema : JSON Schema for add user request
var AddChatSchema = `
{
	"title": "Add Chat Schema",
	"type": "object",
	"properties": {
		"name": {
			"type": "string"
		},
		"users": {
			"type": "array",
			"items":{
				"type": "string"
			}
		}
	},
	"required": ["name", "users"]
}
`

//AddMessageSchema : JSON Schema for add user request
var AddMessageSchema = `
{
	"title": "Add Chat Schema",
	"type": "object",
	"properties": {
		"chat": {
			"type": "string"
		},
		"author": {
			"type": "string"
		},
		"text": {
			"type": "string"
		}
	},
	"required": ["chat", "author", "text"]
}
`
