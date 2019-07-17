package daemon

import (
	"log"

	"github.com/TheEshka/AvitoTask_CharServer/database"
	"github.com/TheEshka/AvitoTask_CharServer/model"
	"github.com/TheEshka/AvitoTask_CharServer/handler"
)

//Config : application configs
type Config struct {
	ListenSpec string

	Db database.Config
}

// Start :
func Start(cfg *Config) error {
	log.Printf("Chat server on port %s started\n", cfg.ListenSpec)

	dab, err := database.InitDb(cfg.Db)
	if err != nil {
		log.Fatal("Fatal error with conecting or preparing databse")
		return err
	}
	log.Println("Database connected successful")

	m := model.New(dab)

	handler.Start(m, cfg.ListenSpec)

	return nil

}
