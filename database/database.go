package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/TheEshka/AvitoTask_CharServer/model"
	_ "github.com/lib/pq"
)

//Config : ConnectString - Postgres connecting settings
type Config struct {
	ConnectString string
}

//InitDb : creating connecton with database
func InitDb(cfg Config) (*pgDb, error) {
	var dbConn *sql.DB
	var err error
	for i := 0; ; {
		i++
		dbConn, err = sql.Open("postgres", cfg.ConnectString)
		if err != nil {
			log.Println("Database connecting error")
			if i == 20 {
				return nil, err
			}
			time.Sleep(time.Second * 5)
		}
		break
	}
	p := &pgDb{dbConn: dbConn}
	err = p.prepareSQLStatements()
	if err != nil {
		return nil, err
	}
	return p, nil

}

type pgDb struct {
	dbConn *sql.DB

	sqlCreateUser *sql.Stmt
	sqlCreateChat       *sql.Stmt
	sqlInsertUserToChat *sql.Stmt
	sqlCreateMessage    *sql.Stmt
	sqlUserChats        *sql.Stmt
	sqlChatMessages     *sql.Stmt
}

func (p *pgDb) prepareSQLStatements() (err error) {
	request := "INSERT INTO users (username, created_at) VALUES ($1, $2) RETURNING id"
	if p.sqlCreateUser, err = p.dbConn.Prepare(request); err != nil {
		log.Printf("Error preparing sqlCreateUser: %v", err)
		return err
	}

	request = "INSERT INTO chats (name, created_at) VALUES ($1, $2) RETURNING id"
	if p.sqlCreateChat, err = p.dbConn.Prepare(request); err != nil {
		log.Printf("Error preparing sqlCreateChat: %v", err)
		return err
	}

	request = "INSERT INTO chat_users (chat_id, user_id) VALUES ($1, $2)"
	if p.sqlInsertUserToChat, err = p.dbConn.Prepare(request); err != nil {
		log.Printf("Error preparing sqlInsertUserToChat: %v", err)
		return err
	}

	request = "INSERT INTO messages (chat_id, author_id, text, created_at) VALUES ($1, $2, $3, $4) RETURNING id"
	if p.sqlCreateMessage, err = p.dbConn.Prepare(request); err != nil {
		log.Printf("Error preparing sqlCreateMessage: %v", err)
		return err
	}

	request = "SELECT (chats.id), (chats.name),(chats.created_at) FROM chats " +
		"INNER JOIN chat_users ON chats.id = chat_users.chat_id "+
		"INNER JOIN users ON users.id = chat_users.user_id WHERE users.id = $1 ORDER BY chats.created_at;";
	if p.sqlUserChats, err = p.dbConn.Prepare(request); err != nil {
		log.Printf("Error preparing sqlUserChats: %v", err)
		return err
	}

	request = "SELECT (messages.id), (messages.chat_id), (messages.author_id), (messages.text), (messages.created_at) FROM messages " +
		"INNER JOIN chats ON messages.chat_id = chats.id " +
		"WHERE chats.id = $1 ORDER BY messages.created_at;"
	if p.sqlChatMessages, err = p.dbConn.Prepare(request); err != nil {
		log.Printf("Error preparing sqlChatsMessages: %v", err)
		return err
	}

	return nil
}

func (p *pgDb) CreateUser(username string) (int64, error) {
	var lastInsertID int64
	err := p.sqlCreateUser.QueryRow(username, time.Now()).Scan(&lastInsertID)
	if err != nil {
		log.Printf("CreateUser database error / %v", err)
		return 0, err
	}
	return lastInsertID, nil
}

func (p *pgDb) CreateChat(name string) (int64, error) {
	var lastInsertID int64
	err := p.sqlCreateChat.QueryRow(name, time.Now()).Scan(&lastInsertID)
	if err != nil {
		log.Printf("CreateChat database error / %v", err)
		return 0, err
	}
	return lastInsertID, nil
}

func (p *pgDb) CreateMessage(chatID int, userID int, text string) (int64, error) {
	var lastInsertID int64
	err := p.sqlCreateMessage.QueryRow(chatID, userID, text, time.Now()).Scan(&lastInsertID)
	if err != nil {
		log.Printf("CreateMessage database error / %v", err)
		return 0, err
	}
	return lastInsertID, nil
}

func (p *pgDb) InsertUserToChat(chatID int, userID int) error {
	_, err := p.sqlInsertUserToChat.Exec(chatID, userID, time.Now())
	if err != nil {
		log.Printf("CreateMessage database error / %v", err)
		return err
	}
	return nil
}

func (p *pgDb) UserChats(userID int) ([]*model.Chat, error) {
	rows, err := p.sqlUserChats.Query(userID)
	if err != nil {
		log.Printf("UserChats database error / %v", err)
		return nil, err
	}
	defer rows.Close()

	chats := make([]*model.Chat, 0)
	for rows.Next() {
		chat := &model.Chat{}
		rows.Scan(&chat.ID, &chat.Name, &chat.CreatedAt)
		chats = append(chats, chat)
	}

	return chats, nil
}

func (p *pgDb) ChatMessages(chatID int) ([]*model.Message, error) {
	rows, err := p.sqlChatMessages.Query(chatID)
	if err != nil {
		log.Printf("ChatMessages database error / %v", err)
		return nil, err
	}
	defer rows.Close()

	messages := make([]*model.Message, 0)
	for rows.Next() {
		message := &model.Message{}
		rows.Scan(&message.ID, &message.ChatID, &message.AuthorID, &message.Text, &message.CreatedAt)
		messages = append(messages, message)
	}

	return messages, nil
}
