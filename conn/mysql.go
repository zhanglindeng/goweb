package conn

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/zhanglindeng/goweb/config"
)

var mysqlConn *gorm.DB

func GetMysqlConn() (*gorm.DB, error) {

	if mysqlConn != nil {
		return mysqlConn, nil
	}

	conn, err := gorm.Open(config.DbDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&collation=%s&parseTime=true&loc=PRC",
		config.DbUsername, config.DbPassword, config.DbHost, config.DbPort, config.DbDatabase, config.DbCharset,
		config.DbCollation))

	if err != nil {
		return mysqlConn, err
	}

	// test conn
	if err := conn.DB().Ping(); err != nil {
		return mysqlConn, err
	}

	mysqlConn = conn
	mysqlConn.LogMode(true)

	return mysqlConn, nil
}
