package routers

import (
	"net/http"
	"social-plus/src/controllers"
)

var loginRoute = Router{
	URI: "/auth/login",
	Method: http.MethodPost,
	Function: controllers.Login,
	RequestAuthentication: false,
}