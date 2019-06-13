package common

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//空结构体,用于给前端返回空
type null struct {

}

//返回给前端的格式，error:状态码 message：提示消息 data：返回的结构体（所有的数据都放在这里）
type rData struct {
	Error int64 `json:"error"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}


//json字符在转为结构体
func JSONTOSTRUCT(js string,t interface{}) {
	err := json.Unmarshal([]byte(string(js)), t)
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
	}
}

//如果用户没传参数或者参数为空格返回默认值 get方法
func GETINIT(r *http.Request,k string,v string) string {
	r.ParseForm()
	//判断元素是否存在或者是否传空
	if len(r.Form[k]) == 0 || strings.Replace(r.Form[k][0]," ","",-1) == ""{
		return v
	}else {
		return r.Form[k][0]
	}
}

//如果用户没传参数或者参数为空格返回默认值 post方法
func POSTINIT(r *http.Request,k string,v string) string {
	r.ParseForm()
	//判断元素是否存在或者是否传空
	if len(r.PostForm[k]) == 0 || strings.Replace(r.PostForm[k][0]," ","",-1) == ""{
		return v
	}else {
		return r.Form[k][0]
	}
}

//判断变量类型
func TYPYOF(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

//返回给前端的数据,error:状态码 message：提示消息 data：返回的结构体
func RETURNDATA(error int64,message string,data interface{}) string {

	if data == ""{
		data = null{}
	}

	var d = rData{
		error,
		message,
		data,
	}

	str, err := json.Marshal(d)

	if err != nil {
		fmt.Println(err)
	}

	return string(str)

}

//字符串转为int
func STRINGTOINT(s string) int {

	i, err := strconv.Atoi(s)

	if err != nil {
		return -1
	}else {
		return i
	}
}

//打乱字符串数组
func Random(strings []string, length int) ([]string) {

	var newString []string

	rand.Seed(time.Now().UnixNano())

	var j = 0
	for i := len(strings); i > 0; i-- {

		index := rand.Intn(i)

		newString = append(newString,strings[index])

		strings = append(strings[:index],strings[index+1:]...)

		j++

		if (j == length){
			break
		}
	}

	return newString

}

//float64转为int
func FLOAT64TOINT(f float64) (int) {

	int64 := int64(f)
	strInt64 := strconv.FormatInt(int64, 10)
	int16,_ := strconv.Atoi(strInt64)
	return int16
}

//int转为float64
func INTTOFLOAT64(i int) (float64) {

	return float64(int64(i))
}


