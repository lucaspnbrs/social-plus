package routers

import (
	"net/http"
	"social-plus/src/controllers"
)

//In the Go language, the imports request the name of package later the function name

var routesUsers = []Router {
	 {
		URI: "/users",
		Method: http.MethodPost,
		Function: controllers.CreateUser,
		RequestAuthentication: false,
	 },
	 {
		URI: "/users/{userID}",
		Method: http.MethodPut,
		Function: controllers.UpUsers,
		RequestAuthentication: true,
	 },
	 {
		URI: "/users/{userID}",
		Method: http.MethodDelete,
		Function: controllers.DeleteUsers,
		RequestAuthentication: true,
	 },
	 {
		URI: "/users",
		Method: http.MethodGet,
		Function: controllers.FetchUsers,
		RequestAuthentication: true,
	 },
	 {
		URI: "/users/{userID}",
		Method: http.MethodGet,
		Function: controllers.FetchUsersByID,
		RequestAuthentication: true,
	 },
}