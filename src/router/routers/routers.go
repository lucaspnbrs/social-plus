package routers

import (
	"net/http"

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
		r.HandleFunc(router.URI, router.Function).Methods(router.Method)
	  }

	  return r
}

