package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/TheEshka/AvitoTask_CharServer/daemon"
	//_ "github.com/lib/pq"
)

func processFlags() *daemon.Config {
	cfg := &daemon.Config{}

	flag.StringVar(&cfg.ListenSpec, "port", ":8080", "HTTP listen port")
	flag.StringVar(&cfg.Db.ConnectString, "db", "user=postgres password=mysecretpassword dbname=chat_db sslmode=disable",
		"DB connecting parameters. Minimal: user and dbname. For more information see lib/pq")

	flag.Parse()
	return cfg
}

func prepareLogFile() {
	f, err := os.OpenFile("testlogfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error opening file for logs: %v", err)
	}
	log.SetFlags(log.LstdFlags)
	//defer f.Close()
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
}

func main() {
	cfg := processFlags()

	prepareLogFile()

	if err := daemon.Start(cfg); err != nil {
		fmt.Printf("Error in main(): %v", err)
	}

	/*var dbconf = database.Config{ConnectString: "user=postgres password=mysecretpassword dbname=chat_db sslmode=disable"}
	//var qwe = &Config{ListenSpec: "8080", Db: dbconf}
	fmt.Println(time.Now().Format(time.RFC3339))
	db, err := database.InitDb(dbconf)
	fmt.Printf("Connection: %s\n", err)
	res, err := db.ChatMessages(2)
	fmt.Println(err)
	fmt.Println(res[0])*/
}
