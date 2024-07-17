package main

import (
	"github.com/gin-gonic/gin",
	"config",
	"routes"
)

func main(){

	router := gin.Default()

	routes.Routes(router);

	router.Run(":3000a")
}