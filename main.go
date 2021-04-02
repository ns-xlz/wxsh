// @Title  
// @Description  
// @Author  xianglizhao  2021/4/2 2:30 下午
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ns-xlz/wxsh/router"
)

func main() {
	engine := gin.New()
	router.RouteHttp(engine)

	engine.Run(":13344")
}

