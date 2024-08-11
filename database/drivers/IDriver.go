package drivers

import "gorm.io/gorm"

type IDriver interface {
	InitConnection() (*gorm.DB, error)
}
