package model

type db interface {
	CreateUser(username string) (int64, error)
	CreateChat(name string) (int64, error)
	CreateMessage(chatID int, userID int, text string) (int64, error)
	InsertUserToChat(chatID int, userID int) error
	UserChats(userID int) ([]*Chat, error)
	ChatMessages(chatID int) ([]*Message, error)
}

//Model : wrapper struct
type Model struct {
	db
}

//New : incert object that realize interface db into wrapper struct
func New(db db) *Model {
	return &Model{
		db: db,
	}
}
