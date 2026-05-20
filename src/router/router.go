package router

import (
	"github.com/gorilla/mux" 
	"social-plus/src/router/routers" 
)

func Generate() *mux.Router {
	r := mux.NewRouter()

	return routers.Settings(r)
}