package repository

import "github.com/zhanglindeng/goweb/model"

type MenuRepository struct{}

func (MenuRepository) Exist(id int) (bool, error) {
	c, err := getMysqlConn()
	if err != nil {
		return false, err
	}

	var count int
	if err := c.Model(&model.Menu{}).Where("id=?", id).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (MenuRepository) Submenus(id int) ([]model.Submenu, error) {
	var submenus []model.Submenu
	c, err := getMysqlConn()
	if err != nil {
		return submenus, err
	}
	if err := c.Order("sort asc").Order("id asc").Where("menu_id=?", id).Find(&submenus).Error; err != nil {
		return submenus, err
	}
	return submenus, nil
}

func (MenuRepository) FindById(id int) (*model.Menu, error) {
	m := &model.Menu{}
	c, err := getMysqlConn()
	if err != nil {
		return m, err
	}
	if err := c.First(m, id).Error; err != nil {
		return m, err
	}
	return m, nil
}

func (MenuRepository) Update(m *model.Menu) error {
	c, err := getMysqlConn()
	if err != nil {
		return err
	}
	return c.Save(m).Error
}

func (MenuRepository) All(menus *[]model.Menu) error {
	c, err := getMysqlConn()
	if err != nil {
		return err
	}
	return c.Preload("Submenus").Find(menus).Error
}

func (MenuRepository) Active(menus *[]model.Menu) error {
	c, err := getMysqlConn()
	if err != nil {
		return err
	}
	return c.Preload("Submenus").Where("status=?", 1).Find(menus).Error
}

func (MenuRepository) Add(m *model.Menu) error {
	c, err := getMysqlConn()
	if err != nil {
		return err
	}
	return c.Create(m).Error
}

func (MenuRepository) Del(id int) error {
	c, err := getMysqlConn()
	if err != nil {
		return err
	}
	return c.Delete(&model.Menu{BaseModel: model.BaseModel{ID: uint(id)}}).Error
}

func (MenuRepository) Disable(id int) error {
	c, err := getMysqlConn()
	if err != nil {
		return err
	}
	return c.Model(&model.Menu{}).Where("id=?", id).Update("status", model.MenuStatusDisabled).Error
}

func (MenuRepository) Enable(id int) error {
	c, err := getMysqlConn()
	if err != nil {
		return err
	}
	return c.Model(&model.Menu{}).Where("id=?", id).Update("status", model.MenuStatusEnable).Error
}
