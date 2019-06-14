package controller

import (
	"app/index/model"
	"common"
	"fmt"
	"math"
	"net/http"
	"strconv"

	//"strconv"
)

type DATA struct {
	List interface{} `json:"list"`
	Count int `json:"count"`
	Page int `json:"page"`
}

//获取用户信息返回给前端的数据字段
type uesrInfo struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Age string `json:"age"`
	CreateTime string `json:"create_time"`
}

type null struct {

}

type ONE struct {
	Id string `json:"id"`
	Rate string `json:"rate"`
}

//获取用户列表
func UserList(w http.ResponseWriter, r *http.Request)  {

	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	limit := common.STRINGTOINT(common.GETINIT(r,"limit","10"))
	page := common.STRINGTOINT(common.GETINIT(r,"page","10"))

	loginModel := model.LoginConstruct()
	JsonOut := loginModel.UserList(page,limit)

	var u []uesrInfo

	common.JSONTOSTRUCT(JsonOut,&u)

	dd := DATA{u,1,0}

	fmt.Fprintln(w, common.RETURNDATA(200,"成功",dd))

}

//新增数据
func AddData(w http.ResponseWriter, r *http.Request)  {

	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	name := common.GETINIT(r,"name","")
	age := common.STRINGTOINT(common.GETINIT(r,"age","10"))

	loginModel := model.LoginConstruct()

	id := loginModel.AddData(name,age)

	if (id < 0){
		fmt.Fprintln(w, common.RETURNDATA(201,"插入失败",""))
	}else {
		fmt.Fprintln(w, common.RETURNDATA(200,"插入成功",""))
	}

}

func GetWord(w http.ResponseWriter, r *http.Request)  {

	limit := common.STRINGTOINT(common.POSTINIT(r,"limit","10"))

	var path = "C:\\Users\\PC\\Desktop\\allFile\\go\\apiFrame\\upload/word.txt"
	loginModel := model.LoginConstruct()


	data,_:= loginModel.GetFileContentAsStringLines(path)
	s:= common.Random(data,limit)

	dd := DATA{s,1,0}

	fmt.Fprintln(w, common.RETURNDATA(200,"成功",dd))

}

func GetWordZ(w http.ResponseWriter, r *http.Request)  {

	limit := common.STRINGTOINT(common.POSTINIT(r,"limit","5"))

	var path = "C:\\Users\\PC\\Desktop\\allFile\\go\\apiFrame\\upload/word_z.txt"
	loginModel := model.LoginConstruct()


	data,_:= loginModel.GetFileContentAsStringLines(path)
	s:= common.Random(data,limit)

	dd := DATA{s,1,0}

	fmt.Fprintln(w, common.RETURNDATA(200,"成功",dd))

}

/**********************************/
func IsPlay(w http.ResponseWriter, r *http.Request)  {

	type NUM struct {
		Num string `json:"num"`
	}
	type COUNT struct {
		Count int `json:"count"`
	}

	r.ParseForm()//解析url传递的参数，对于POST则解析响应包的主体（request body）

	uid := common.STRINGTOINT(common.POSTINIT(r,"uid","0"))

	loginModel := model.LoginConstruct()

	JsonOut := loginModel.IsPlay(uid)

	var num []NUM

	flag := 0

	common.JSONTOSTRUCT(JsonOut,&num)

	if common.STRINGTOINT(num[0].Num) > 0{
		flag = 1
	}else {
		flag = 0
	}

	c := COUNT{flag}

	fmt.Fprintln(w, common.RETURNDATA(200,"成功",c))

}
//中奖用户
func AllLuckList(w http.ResponseWriter, r *http.Request)  {
	type LIST struct {
		Ctime string `json:"ctime"`
		Name string `json:"name"`
		Type string `json:"type"`
		NickName string `json:"nickName"`
		HeadImage string `json:"headImage"`
	}

	type NUM struct {
		Num string `json:"num"`
	}
	limit := common.STRINGTOINT(common.POSTINIT(r,"limit","10"))
	page := common.STRINGTOINT(common.POSTINIT(r,"page","1"))

	loginModel := model.LoginConstruct()

	//获取指定数据列表
	JsonOut := loginModel.AllLuckList(limit,page)
	var list []LIST

	common.JSONTOSTRUCT(JsonOut,&list)

	//获取数量
	JsonOutCount := loginModel.AllLuckListCount()
	var Count []NUM
	common.JSONTOSTRUCT(JsonOutCount,&Count)

	count := common.STRINGTOINT(Count[0].Num)

	//fmt.Fprintln(w, count/limit)


	dd := DATA{list,count,common.FLOAT64TOINT(math.Ceil(common.INTTOFLOAT64(count/limit)+0.5))}

	fmt.Fprintln(w, common.RETURNDATA(200,"成功",dd))


}

//获取当前用户游戏次数
func GetGameTimes(w http.ResponseWriter, r *http.Request)  {
	type NUM struct {
		Num string `json:"num"`
	}
	uid := common.STRINGTOINT(common.POSTINIT(r,"uid","10"))
	level := common.STRINGTOINT(common.POSTINIT(r,"level","1"))
	className := GetClassName(level)

	loginModel := model.LoginConstruct()

	//获取指定数据列表
	JsonOut := loginModel.GetGameTimes(className)

	var list []NUM

	common.JSONTOSTRUCT(JsonOut,&list)

	allNum,_ := strconv.Atoi(list[0].Num)

	//获取当前用户今天玩了多少次
	JsonOut1 := loginModel.GetGameTimesToDay(uid)

	var list1 []NUM

	common.JSONTOSTRUCT(JsonOut1,&list1)

	yNum,_ := strconv.Atoi(list1[0].Num)

	dd := NUM{strconv.Itoa(allNum-yNum)}
	fmt.Fprintln(w, common.RETURNDATA(200,"成功",dd))
}

//上传文件
func UploadFile(w http.ResponseWriter, r *http.Request)  {
	k := common.POSTINIT(r,"uploadfile","uploadfile")
	suffix := []string{"jpg","png"}
	filaName := "test.png"
	pathZ := "upload/"

	filePath := common.UPLOADFILEONE(r,k,pathZ,filaName,suffix)
	fmt.Fprintln(w, filePath)
}



func Luck(w http.ResponseWriter, r *http.Request)  {

	type prize struct {
		Id string `json:"id"`
		Name string `json:"name"`
		Num string `json:"num"`
	}

	//uid := common.STRINGTOINT(common.POSTINIT(r,"uid","10"))
	loginModel := model.LoginConstruct()
	//获取指定数据列表
	JsonOut := loginModel.GetPrizeList()

	var list []ONE

	common.JSONTOSTRUCT(JsonOut,&list)

	id := isLuck(list)
	JsonOut1 := loginModel.GetPrizeById(id)

	var prizes []prize

	common.JSONTOSTRUCT(JsonOut1,&prizes)

	fmt.Fprintln(w, common.RETURNDATA(200,"成功",prizes[0]))
}

//更改游戏状态
func UpdateStatus(w http.ResponseWriter, r *http.Request)  {
	goodId := common.STRINGTOINT(common.POSTINIT(r,"goodId","10"))
	status := common.STRINGTOINT(common.POSTINIT(r,"status","1"))
	loginModel := model.LoginConstruct()
	//获取指定数据列表
	loginModel.UpdateStatus(goodId,status)

	var dd null

	fmt.Fprintln(w, common.RETURNDATA(200,"成功",dd))
}

func GetClassName(level int) (string) {

	if level > 80{
		level = 70
	}

	classS := map[string]string{
		"0":"oneClass",
		"1":"twoClass",
		"2":"threeClass",
		"3":"fourClass",
		"4":"fiveClass",
		"5":"sixClass",
		"6":"sevenClass",
		"7":"eightClass",
	}

	return classS[strconv.Itoa(level/10)]
}

//判读是否中奖
func isLuck(s []ONE) (string) {
	var aa []string
	for i := 0; i < len(s); i++ {
		rate,_ := strconv.Atoi(s[i].Rate)
		for j := 0; j < rate; j++ {
			aa = append(aa,s[i].Id)
		}
	}

	aa = common.Random(aa,1)

	return aa[0]

}




