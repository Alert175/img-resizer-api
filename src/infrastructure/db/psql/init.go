package psql

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"img-resizer-api/src/infrastructure/db/psql/models"
	"img-resizer-api/src/infrastructure/pkg/utils/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbInit() {
	dbUser, dbPassword, dbName, dbPort, dbHost := os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_HOST")
	intPort, err := strconv.Atoi(dbPort)
	if err != nil {
		log.Fatal("db port err: ", err)
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, intPort, dbUser, dbPassword, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("storage err: ", err)
	}

	logger.Success("db connected")
	DB = db
	autoMigrate(DB)
}

func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.ImageLog{},
	)
	if err != nil {
		logger.Error("models is not migrated")
		log.Fatal(err)
	}
	logger.Success("models is migrated")
}
