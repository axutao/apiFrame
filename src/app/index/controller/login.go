package controller

import (
	"app/index/model"
	"common"
	"fmt"
	"net/http"
)

type data struct {
	List interface{} `json:"list"`
	Count int `json:"count"`
}

//获取用户信息返回给前端的数据字段
type uesrInfo struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Age string `json:"age"`
	CreateTime string `json:"create_time"`
}

//获取用户列表
func UserList(w http.ResponseWriter, r *http.Request)  {

	r.ParseForm()//解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	limit := common.STRINGTOINT(common.GETINIT(r,"limit","10"))
	page := common.STRINGTOINT(common.GETINIT(r,"page","10"))

	loginModel := model.LoginConstruct()
	JsonOut := loginModel.UserList(page,limit)

	var u []uesrInfo

	common.JSONTOSTRUCT(JsonOut,&u)

	dd := data{u,1}

	fmt.Fprintln(w, common.RETURNDATA(200,"成功",dd))

}

//新增数据
func AddData(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()//解析url传递的参数，对于POST则解析响应包的主体（request body）
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



