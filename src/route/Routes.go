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
	Route{Name:"GetWord", Method:"POST",   Pattern:"/GetWord",           HandlerFunc:controller.GetWord},
	Route{Name:"GetWordZ", Method:"POST",   Pattern:"/GetWordZ",           HandlerFunc:controller.GetWordZ},
	Route{Name:"IsPlay", Method:"POST",   Pattern:"/IsPlay",           HandlerFunc:controller.IsPlay},
	Route{Name:"AllLuckList", Method:"POST",   Pattern:"/AllLuckList",           HandlerFunc:controller.AllLuckList},
	Route{Name:"GetGameTimes", Method:"POST",   Pattern:"/GetGameTimes",           HandlerFunc:controller.GetGameTimes},
	Route{Name:"Luck", Method:"POST",   Pattern:"/Luck",           HandlerFunc:controller.Luck},
	Route{Name:"UpdateStatus", Method:"POST",   Pattern:"/UpdateStatus",           HandlerFunc:controller.UpdateStatus},
}
