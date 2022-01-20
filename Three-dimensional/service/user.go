package service

import (
	"Three_dimensional/dao"
	"Three_dimensional/model"
	"database/sql"
)

func Register(user model.User) (int, error) {
	userId, err := dao.Register(user)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func IsNameRepeated(name string) (bool,error) {
	flag, err := dao.IsNameRepeated(name)
	if err != nil {
		return false, err
	}
	if flag != true {
		return false, nil
	}
	return true, nil
}

func IsNameExist(name string) (bool, error) {
	flag, err := dao.IsNameExist(name)
	if err != nil {
		if err == sql.ErrNoRows{
			return false,nil
		}
		return false, err
	}
	if flag != true {
		return false, nil
	}
	return true, nil
}

func SelectPasswordByUserName(user model.User) (string, error) {
	pwd, err := dao.SelectPasswordByUserName(user)
	if err != nil {
		return pwd, err
	}
	return pwd, nil
}
