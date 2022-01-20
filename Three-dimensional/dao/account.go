package dao

import (
	"Three_dimensional/model"
)

func CreateAccount(accountName string) error {
	sqlStr := "insert into account(accountName,income,expend,sum) values(?,?,?,?)"
	_, err := DB.Exec(sqlStr, accountName,0,0,0)
	if err != nil {
		return err
	}
	return nil
}

func Recharge(account model.Account) (error, float32) {
	var income,sum float32
	rows, err := DB.Query("select income,sum from account where accountName = ? ", account.AccountName)
	if err != nil {
		return err, 0
	}
	for rows.Next() {
		err = rows.Scan(&income,&sum)
		if err != nil {
			return err, 0
		}
	}

	sqlStr := "update account set income = ?, sum = ? where accountName = ? "
	_, err = DB.Exec(sqlStr, income+account.Income, sum+account.Income, account.AccountName)
	if err != nil {
		return err, 0
	}
	return nil,sum+account.Income
}
