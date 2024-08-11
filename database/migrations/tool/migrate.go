package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
	"log"
	"os"
	_ "pince/database"
)

func main() {
	args := os.Args
	command := args[2]
	migrationsDirRelativePath := args[1]
	MigrateDatabase(command, migrationsDirRelativePath+os.Getenv("SERVICE_NAME"))
}

func MigrateDatabase(command string, migrationsDirRelativePath string) {
	dataSourceString := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USERNAME") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable"
	db, err := goose.OpenDBWithDriver(os.Getenv("DB_TYPE"), dataSourceString)
	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}

	if err := goose.RunContext(context.Background(), command, db, migrationsDirRelativePath); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
}
