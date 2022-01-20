package service

import (
	"Three_dimensional/dao"
	"Three_dimensional/model"
)

func CreateAccount(accountName string) error {
	err := dao.CreateAccount(accountName)
	if err != nil {
		return err
	}
	return nil
}

func Recharge(account model.Account) (error, float32) {
	err, total := dao.Recharge(account)
	if err != nil {
		return err, 0
	}
	return nil, total
}
