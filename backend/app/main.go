package main

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"backend/server"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	utils.LoadEnv(".env")

	DB_NAME := os.Getenv("DB_NAME")
	MIGRATION_DIR := os.Getenv("MIGRATION_DIR")

	models.OrmInstance.InitDB(DB_NAME, MIGRATION_DIR)
	models.MigrateModels(MIGRATION_DIR, *models.OrmInstance)

	s := server.NewRouter()

	PORT := os.Getenv("PORT")

	s.StartServer(PORT)
	fmt.Println("listening.....")
}
