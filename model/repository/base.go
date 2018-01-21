package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/zhanglindeng/goweb/conn"
)

func getMysqlConn() (*gorm.DB, error) {
	return conn.GetMysqlConn()
}
