package cache

import (
	"github.com/gin-contrib/cache/persistence"
	"github.com/zhanglindeng/util"
)

var userCachePrefix = "user_"

func SetUserEmail2Id(email string, uid uint) error {
	c, err := getCacheConn()
	if err != nil {
		return err
	}

	return c.Set(keyUserEmail2Id(email), uid, persistence.FOREVER)
}

func GetUserEmail2Id(email string) (uint, error) {
	c, err := getCacheConn()
	if err != nil {
		return 0, err
	}
	var uid uint
	if err := c.Get(keyUserEmail2Id(email), &uid); err != nil {
		return 0, err
	}

	return uid, nil
}

func keyUserEmail2Id(email string) string {
	return addKeyPrefix(userCachePrefix + util.Md5("keyUserEmail2Id"+email))
}
