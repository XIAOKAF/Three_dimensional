package service

import (
	"Three_dimensional/dao"
	"Three_dimensional/model"
)

func TransferMoney(transfer model.Transfer) (error,bool) {
	err, flag := dao.TransferMoney(transfer)
	if err != nil {
		return err,true
	}
	if flag != true {
		return nil,false
	}
	return nil,true
}

func SelectDetails(transfer model.Transfer) (model.Transfer,error) {
	t,err := dao.SelectDetails(transfer)
	if err != nil {
		return t, err
	}
	return t,nil
}