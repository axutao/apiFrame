package model

import (
	"fmt"
	"io/ioutil"
	"mysql"
	"strings"
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

func (this *Login) AddData (name string,age int)(int64) {
	return -1
}
//
//	sql := fmt.Sprintf("insert into user(`name`,`age`) values('%s',%d)",name,age)
//	sql1 := fmt.Sprintf("insert into user(`namee`,`age`) values('%s',%d)",name,age)
//
//	db, err := mysql.NewDb()
//
//	tx, err := db.Begin()
//	if err != nil {
//		return -1
//	}
//
//	if err != nil {
//		fmt.Println("打开SQL时出错:", err.Error())
//		return -2
//	}
//
//	defer db.Close()
//
//	instid, err := tx.Insert(sql)
//
//	_, err1 := tx.Insert(sql1)
//
//	if err != nil || err1 != nil{
//		fmt.Println("insert SQL时出错:", err.Error(),sql)
//		tx.Rollback()
//
//		return -1
//	}
//
//	return instid
//
//}
//
func (this *Login) GetFileContentAsStringLines(filePath string) ([]string, error) {
	//logger.Infof("get file content as lines: %v", filePath)
	result := []string{}
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		//logger.Errorf("read file: %v error: %v", filePath, err)
		return result, err
	}
	s := string(b)
	for _, lineStr := range strings.Split(s, "\n") {
		lineStr = strings.TrimSpace(lineStr)
		if lineStr == "" {
			continue
		}
		result = append(result, lineStr)
	}
//	logger.Infof("get file content as lines: %v, size: %v", filePath, len(result))
	return result, nil
}

func (this *Login) IsPlay(uid int) (string) {
	sql := fmt.Sprintf("select count(1) as num from ndj_game_log where uid=%d",uid)
	db,err := mysql.NewDb()
	if err != nil {
		fmt.Println("打开SQL时出错:", err.Error())
		return ""
	}

	defer db.Close()

	JsonOut := db.QueryDataRowsToJson(sql)

	return JsonOut
}

func (this *Login) AllLuckList(limit int,page int) (string) {
	sql := fmt.Sprintf("select l.ctime,p.name,p.type,l.nickName,l.headImage from ndj_prize_log as l join ndj_prize as p on l.prizeId=p.id limit %d,%d",(page - 1)*limit,limit)

	db,err := mysql.NewDb()
	if err != nil {
		fmt.Println("打开SQL时出错:", err.Error())
		return ""
	}

	defer db.Close()

	JsonOut := db.QueryDataRowsToJson(sql)

	return JsonOut
}

func (this *Login) AllLuckListCount() (string) {
	sql := fmt.Sprintf("select count(1) as num from ndj_prize_log as l join ndj_prize as p on l.prizeId=p.id")

	db,err := mysql.NewDb()
	if err != nil {
		fmt.Println("打开SQL时出错:", err.Error())
		return ""
	}

	defer db.Close()

	JsonOut := db.QueryDataRowsToJson(sql)

	return JsonOut
}

func (this *Login) GetGameTimes(className string) (string) {
	sql := fmt.Sprintf("select %s as num from ndj_class",className)

	db,err := mysql.NewDb()
	if err != nil {
		fmt.Println("打开SQL时出错:", err.Error())
		return ""
	}

	defer db.Close()

	JsonOut := db.QueryDataRowsToJson(sql)

	return JsonOut
}

func (this *Login) GetGameTimesToDay(uid int) (string) {
	sql := fmt.Sprintf("select count(1) as num from ndj_game_log where uid=%d",uid)

	db,err := mysql.NewDb()
	if err != nil {
		fmt.Println("打开SQL时出错:", err.Error())
		return ""
	}

	defer db.Close()

	JsonOut := db.QueryDataRowsToJson(sql)

	return JsonOut
}
func (this *Login) GetPrizeList() (string) {
	sql := fmt.Sprintf("select id,rate from ndj_prize")

	db,err := mysql.NewDb()
	if err != nil {
		fmt.Println("打开SQL时出错:", err.Error())
		return ""
	}

	defer db.Close()

	JsonOut := db.QueryDataRowsToJson(sql)

	return JsonOut
}

func (this *Login) GetPrizeById(id string) (string) {
	sql := fmt.Sprintf("select id,name,num from ndj_prize where id=%s",id)

	db,err := mysql.NewDb()
	if err != nil {
		fmt.Println("打开SQL时出错:", err.Error())
		return ""
	}

	defer db.Close()

	JsonOut := db.QueryDataRowsToJson(sql)

	return JsonOut
}

func (this *Login) UpdateStatus(id int,status int) (int) {
	sql := fmt.Sprintf("update ndj_prize_log set status=%d where id=%d",status,id)

	db,err := mysql.NewDb()
	if err != nil {
		fmt.Println("打开SQL时出错:", err.Error())
		return -2
	}

	defer db.Close()

	affNum,_ := db.Update(sql)

	if affNum == 0{
		return -1
	}

	return 1
}
