// @Title  
// @Description  
// @Author  xianglizhao  2021/4/2 2:15 下午
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ns-xlz/wxsh/controllers"
)

func RouteHttp(engine *gin.Engine) {
	engine.Use(gin.Recovery())
	engine.POST("/compare/attack/crit", controllers.CompareAttackCrit)
}