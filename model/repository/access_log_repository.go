package repository

import (
	"github.com/zhanglindeng/goweb/model"
	"github.com/zhanglindeng/goweb/util"
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

func (ar AccessLogRepository) Find(offset, limit int) ([]model.AccessLog, int, error) {
	var logs []model.AccessLog
	var count int



	c, err := getMysqlConn()
	if err != nil {
		return logs, count, err
	}

	if err := c.Model(&model.AccessLog{}).Count(&count).Error; err != nil {
		return logs, count, err
	}

	if err := c.Order("request_time_float desc").Offset(offset).Limit(limit).Find(&logs).Error; err != nil {
		return logs, count, err
	}

	return logs, count, nil
}
