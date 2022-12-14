package config

import (
	"fmt"
	"log"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/joho/godotenv"

	"img-resizer-api/src/infrastructure/adapter/http"
	"img-resizer-api/src/infrastructure/db/psql"
)

func InitConfig() error {
	envInit()
	psql.DbInit()
	vips.Startup(nil)
	defer http.InitHttpAdapter()
	return nil
}

func envInit() {
	if err := godotenv.Load("./config/.env"); err != nil {
		log.Fatal("Error loading .env file. Error - ", err)
	} else {
		fmt.Println("env loaded")
	}
}
