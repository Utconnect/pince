package drivers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type PostgresDriver struct {
}

func (driver *PostgresDriver) InitConnection() (*gorm.DB, error) {
	var dsn string
	if os.Getenv("IS_DOCKERIZE") == "" {
		dsn = "host=localhost user=postgres password=password dbname=TSS port=5432 sslmode=disable"
	} else {
		dsn = "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USERNAME") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable"
	}

	return gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
}
