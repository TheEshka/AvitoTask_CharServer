package model

type db interface {
	CreateUser(username string) (int64, error)
	CreateChatWithUsers(name string, users []string) (int64, error)
	CreateMessage(chatID, userID, text string) (int64, error)
	UserChats(userID string) ([]*Chat, error)
	ChatMessages(chatID string) ([]*Message, error)
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
