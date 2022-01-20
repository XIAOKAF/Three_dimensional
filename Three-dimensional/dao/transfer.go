package dao

import (
	"Three_dimensional/model"
	"fmt"
)

func TransferMoney(transfer model.Transfer) (error, bool) {
	var income, expend, sum1, sum2 float32

	rows, err := DB.Query("select expend,sum from account where accountName = ?", transfer.Payer)
	if err != nil {
		fmt.Println(err)
		return err, true
	}
	for rows.Next() {
		err = rows.Scan(&expend, &sum1)
		if err != nil {
			return err, true
		}
		//余额不足
		if sum1 < transfer.Amount {
			return nil, false
		}
	}

	rows, err = DB.Query("select income, sum from account where accountName = ?", transfer.Payee)
	if err != nil {
		return err, true
	}
	for rows.Next() {
		err = rows.Scan(&income, &sum2)
		if err != nil {
			fmt.Println(err)
			return err, true
		}
	}

	sqlStr1 := "update account set expend = ?, sum = ? where accountName = ?" //更新转账人的信息
	sqlStr2 := "update account set income = ?, sum = ? where accountName = ?" //更新被转账人的信息
	_, err = DB.Exec(sqlStr1, expend+transfer.Amount, sum1-transfer.Amount, transfer.Payer)
	if err != nil {
		return err, true
	}
	_, err = DB.Exec(sqlStr2, income+transfer.Amount, sum2+transfer.Amount, transfer.Payee)
	if err != nil {
		return err, true
	}

	sqlStr3 := "insert into details(payer, payee, amount, postscript, time)values(?,?,?,?,?)"//记录转账信息
	_,err = DB.Exec(sqlStr3,transfer.Payer,transfer.Payee,transfer.Amount,transfer.Postscript,transfer.Time)
	if err != nil {
		return err,true
	}
	return nil, true
}

func SelectDetails(transfer model.Transfer) (model.Transfer,error) {
	rows, err := DB.Query("select * from details where postscript like '%"+transfer.Postscript+"%'")
	if err!= nil{
		return model.Transfer{},err
	}
	for rows.Next() {
		err = rows.Scan(&transfer.PaymentId,&transfer.Payer,&transfer.Payee,&transfer.Amount,&transfer.Postscript,&transfer.Time)
		if err != nil {
			return model.Transfer{},err
		}
	}
	return transfer,nil
}
