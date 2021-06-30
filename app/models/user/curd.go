package user

import (
	"github.com/sirupsen/logrus"
	"goto/database/mysql"
)

func (user *User) First(sort string) (err error) {
	switch sort {
	case "asc":
		if err = mysql.DB.First(user).Error; err != nil {
			logrus.Error(err)
			return err
		}
	case "desc":
		if err = mysql.DB.Last(user).Error; err != nil {
			logrus.Error(err)
			return err
		}
	}
	return nil
}

func (user *User) FindOne(id int) (err error) {
	if err = mysql.DB.First(user, id).Error; err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func Get(users *[]User, page int, limit int) (err error) {
	if err = mysql.DB.Offset((page - 1) * limit).Limit(limit).Find(users).Error; err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func GetAll(users *[]User) (err error) {
	if err = mysql.DB.Find(users).Error; err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (user *User) Create() (err error) {
	if err = mysql.DB.Create(user).Error; err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (user *User) Update() (err error) {
	if err = mysql.DB.Save(user).Error; err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (user *User) Delete() (err error) {
	if err = mysql.DB.Delete(user).Error; err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func DeleteById(id int) (err error) {
	if err = mysql.DB.Delete(User{}, id).Error; err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func DeleteByIds(users []User, ids []int) (err error) {
	if err = mysql.DB.Delete(users, ids).Error; err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
