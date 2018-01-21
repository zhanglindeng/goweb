package model

import (
	"github.com/jinzhu/gorm"
	"github.com/zhanglindeng/goweb/config"
	conn2 "github.com/zhanglindeng/goweb/conn"
)

func Migrate() error {

	conn, err := conn2.GetMysqlConn()
	if err != nil {
		return err
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.DbPrefix + defaultTableName
	}

	return conn.Set("gorm:table_options", "ENGINE="+config.DbEngine).AutoMigrate(
		&User{},
		&AccessLog{},
	).Error
}
