package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"service_history/internal/app/api"
	"service_history/pkg/config"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("error loading .env file:", err, "will check `go env`")
	}
}

func main() {
	conf, err := config.NewConfig(os.Getenv("CONFIGURE"))
	if err != nil {
		log.Println(err)
		return
	}

	serv := api.NewApiServer(conf)
	err = serv.StartServ()
	if err != nil {
		log.Println(err)
		return
	}

}
