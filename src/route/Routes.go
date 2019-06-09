package route

import (
	"app/index/controller"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{Name:"UserList", Method:"GET",   Pattern:"/UserList",           HandlerFunc:controller.UserList},
	Route{Name:"AddData", Method:"POST",   Pattern:"/AddData",           HandlerFunc:controller.AddData},
}
