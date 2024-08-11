package database

import (
	"database/sql"
	"gorm.io/gorm"
	"log"
	"pince/common"
	"pince/database/drivers"
)

type Connection struct {
	Driver drivers.IDriver
	Db     *sql.DB
	GormDb *gorm.DB
}

func (connection *Connection) Connect() {
	var err error
	connection.GormDb, err = connection.Driver.InitConnection()
	if err != nil {
		log.Println(common.Red + err.Error() + common.Reset)
	}
}
