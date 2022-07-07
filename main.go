package main

//https://mp.weixin.qq.com/s/jQlRLkSPyieaW0MLFd3W5Q

import (
	"github.com/gin-gonic/gin"
	"myGin/routes"
)

func main() {
	r := gin.Default()

	routes.Load(r)

	r.Run()
}
