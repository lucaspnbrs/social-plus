package routers

import (
	"net/http"
	"social-plus/src/middleware"

	"github.com/gorilla/mux"
)

//Router to be reference all the routes of the API
type Router struct {
	URI string
	Method string
	Function func( http.ResponseWriter, *http.Request)
	RequestAuthentication bool
}


//Apply all the routes inside the router
func Settings( r *mux.Router) *mux.Router {
	  routes := routesUsers
	  routes = append(routes, loginRoute)

	  for _, router := range  routes {	
		if router.RequestAuthentication {
			r.HandleFunc(router.URI, 
				middleware.Logger(middleware.Authenticate(router.Function))).Methods(router.Method)
		} else {
			r.HandleFunc(router.URI, middleware.Logger(router.Function)).Methods(router.Method)
		}
	  }

	  return r
}

