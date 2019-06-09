package model

import (
	"fmt"
	"mysql"
)

type Login struct {

}

func LoginConstruct()(*Login){
	return &Login{}
}

func (this *Login) UserList(page int,limit int) (string) {

	sql := fmt.Sprintf("select * from user limit %d,%d",(page - 1)*limit,limit)

	db, err := mysql.NewDb()

	if err != nil {
		fmt.Println("打开SQL时出错:", err.Error())
		return ""
	}

	defer db.Close()

	JsonOut := db.QueryDataRowsToJson(sql)

	return JsonOut

}

func (this *Login) AddData (name string,age int)(int64){

	sql := fmt.Sprintf("insert into user(`name`,`age`) values('%s',%d)",name,age)

	db, err := mysql.NewDb()

	if err != nil {
		fmt.Println("打开SQL时出错:", err.Error())
		return -2
	}

	defer db.Close()

	instid, err := db.Insert(sql)
	if err != nil {
		fmt.Println("insert SQL时出错:", err.Error(),sql)
		return -1
	}

	return instid

}