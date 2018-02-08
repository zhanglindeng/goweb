package repository

import "github.com/zhanglindeng/goweb/model"

type SubmenuRepository struct{}

func (SubmenuRepository) FindById(id int) (*model.Submenu, error) {
	s := &model.Submenu{}
	c, err := getMysqlConn()
	if err != nil {
		return s, err
	}
	if err := c.First(s, id).Error; err != nil {
		return s, err
	}
	return s, nil
}


func (SubmenuRepository) Update(m *model.Submenu) error {
	c, err := getMysqlConn()
	if err != nil {
		return err
	}
	return c.Save(m).Error
}

func (SubmenuRepository) All(submenus []model.Submenu) error {
	c, err := getMysqlConn()
	if err != nil {
		return err
	}
	return c.Find(&submenus).Error
}

func (SubmenuRepository) Add(s *model.Submenu) error {
	c, err := getMysqlConn()
	if err != nil {
		return err
	}
	return c.Create(s).Error
}

func (SubmenuRepository) Del(id int) error {
	c, err := getMysqlConn()
	if err != nil {
		return err
	}
	return c.Delete(&model.Submenu{BaseModel: model.BaseModel{ID: uint(id)}}).Error
}

func (SubmenuRepository) Disable(id int) error {
	c, err := getMysqlConn()
	if err != nil {
		return err
	}
	return c.Model(&model.Submenu{}).Where("id=?", id).Update("status", model.SubmenuStatusDisabled).Error
}

func (SubmenuRepository) Enable(id int) error {
	c, err := getMysqlConn()
	if err != nil {
		return err
	}
	return c.Model(&model.Submenu{}).Where("id=?", id).Update("status", model.SubmenuStatusEnable).Error
}
