package repository

import (
	"github.com/zhanglindeng/goweb/model"
	"github.com/zhanglindeng/util"
)

type AccessLogRepository struct{}

func (AccessLogRepository) Create(al *model.AccessLog) error {
	c, err := getMysqlConn()
	if err != nil {
		return err
	}
	// TODO access log ip location
	al.IPLocation = ""
	al.Group = ""
	al.ClientOs = ""
	al.ClientBrowser = ""
	al.ContentLengthFormat = util.SizeFormat(al.ContentLength)
	return c.Create(al).Error
}

func (AccessLogRepository) Find(offset, limit int) ([]model.AccessLog, error) {
	var logs []model.AccessLog

	c, err := getMysqlConn()
	if err != nil {
		return logs, err
	}

	if err := c.Order("request_time desc").Offset(offset).Limit(limit).Find(&logs).Error; err != nil {
		return logs, err
	}

	return logs, nil
}
