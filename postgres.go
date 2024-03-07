package postgres

import (
	"fmt"
	"github.com/0xJacky/cosy/settings"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open(dbs *settings.DataBase) gorm.Dialector {
	return postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbs.Host, dbs.User, dbs.Password, dbs.Name, dbs.Port))
}