package model

import (
	"errors"
)

//User : general information about chat
type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
}

//Chat : general information about chat
type Chat struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

//Message : general information about message
type Message struct {
	ID        int    `json:"id"`
	ChatID    int    `json:"chat_id"`
	AuthorID  int    `json:"author_id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

//ErrOnDatabase : error after sending request  to database
var ErrOnDatabase = errors.New("Database connection error")

//ErrIncorrectInput :  request doesn't change anything in database since incorrect verification
var ErrIncorrectInput = errors.New("Incorrect username/email or password")

//ErrAleadyExist : request doesn't change anything in database since there are already exist the ecludind record
var ErrAlreadyExist = errors.New("User with this username/email already exist")

//ErrAleadyExist : request doesn't change anything in database since there are already exist the ecludind record
var ErrPasswordFormat = errors.New("Password must be in base64 format")
