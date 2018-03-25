package repository

import (
	"github.com/zhanglindeng/goweb/cache"
	"github.com/zhanglindeng/goweb/model"
)

type UserRepository struct{}

func (ur UserRepository) All(users *[]model.User) error {
	c, err := getMysqlConn()
	if err != nil {
		return err
	}
	return c.Select([]string{"id,name,email,created_at,updated_at"}).Find(users).Error
}

func (ur UserRepository) Create(u *model.User) (error) {
	c, err := getMysqlConn()
	if err != nil {
		return err
	}
	return c.Create(u).Error
}

func (ur UserRepository) FindById(id uint) (u model.User, err error) {

	c, err := getMysqlConn()
	if err != nil {
		return u, err
	}

	if err = c.First(&u, id).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (ur UserRepository) FindByEmail(email string) (u model.User, err error) {

	if uid, err := cache.GetUserEmail2Id(email); err == nil {
		return ur.FindById(uid)
	}

	c, err := getMysqlConn()
	if err != nil {
		return u, err
	}

	if err = c.Where("email=?", email).First(&u).Error; err != nil {
		return u, err
	}

	cache.SetUserEmail2Id(u.Email, u.ID)

	return u, nil
}
