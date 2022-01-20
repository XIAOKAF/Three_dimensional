package dao

import (
	"Three_dimensional/model"
)

func Register(user model.User) (int, error) {
	sqlStr := "insert into user(name,password) values(?,?)"
	_, err := DB.Exec(sqlStr, user.Name, user.Password)
	if err != nil {
		return 0, err
	}
	rows, err := DB.Query("select id from user where name = ?", user.Name)
	if err != nil {
		return 0, err
	}
	for rows.Next() {
		err = rows.Scan(&user.Id)
		if err != nil {
			return user.Id, err
		}
	}
	return user.Id, nil
}

func IsNameRepeated(name string) (bool, error) {
	var n string
	rows, err := DB.Query("select name from user ")
	if err != nil {
		return false, err
	}
	for rows.Next() {
		err := rows.Scan(&n)
		if err != nil {
			return false, err
		}
		if n == name {
			return false, nil
		}
	}
	return true, nil
}

func IsNameExist(name string) (bool, error) {
	var n string
	rows, err := DB.Query("select name from user ")
	if err != nil {
		return false, err
	}
	for rows.Next() {
		err := rows.Scan(&n)
		if err != nil {
			return false, err
		}
		if n == name {
			return true, nil
		}
	}
	return false, nil
}

func SelectPasswordByUserName(user model.User) (string, error) {
	rows, err := DB.Query("select password from user where name = ?", user.Name)
	if err != nil {
		return "", err
	}
	for rows.Next() {
		err := rows.Scan(&user.Password)
		if err != nil {
			return user.Password, err
		}
	}
	return user.Password, nil
}
