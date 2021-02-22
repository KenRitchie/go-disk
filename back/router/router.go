package router

import (
	"go-disk/back/controller"
	"net/http"
)

// Map 路由
type Map struct {
	Path string
	Fn   func(http.ResponseWriter, *http.Request)
}

// Maps 路由列表
var Maps = []*Map{

	{
		Path: "/",
		Fn:   controller.IndexHandler,
	},
	{
		Path: "/view",
		Fn:   controller.ViewHandler,
	},
	{
		Path: "/upload",
		Fn:   controller.UploadHandler,
	},

}

// Routes 操作
func Routes() {
	for i := 0; i < len(Maps); i++ {
		cRoute := Maps[i]
		http.HandleFunc(cRoute.Path, cRoute.Fn)
	}
}
