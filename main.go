package main

import (
	"fmt"
	"log"
	"net/http"
	"social-plus/src/config"
	"social-plus/src/router"
)

func main() {
	config.LoadSys()

	fmt.Println("Port", config.Port)
	fmt.Println("New services to main")

	r := router.Generate()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}