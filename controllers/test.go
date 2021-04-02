// @Title  
// @Description  
// @Author  xianglizhao  2021/4/2 1:50 下午
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ns-xlz/wxsh/service/attack"
	"github.com/ns-xlz/wxsh/util"
	"net/http"
)

type AttackCritCompareParam struct {
	X float64	`json:"x" form:"x" binding:"required"`
	Y float64	`json:"y" form:"y" binding:"required"`
	Z float64	`json:"z" form:"z" binding:"required"`
	A float64	`json:"a" form:"a" binding:"required"`
}

func CompareAttackCrit(ctx *gin.Context) {
	var param AttackCritCompareParam
	if err := ctx.ShouldBind(&param); err != nil {
		_ = ctx.Error(err)
		ctx.JSON(http.StatusOK, util.JsonResponse{
			ErrNo:  -1,
			ErrMsg: "参数不合法，请纠正参数",
			Data:   nil,
		})
		return
	}

	ori, now, b, sug := attack.GetAttackCompare(ctx, param.X, param.Y, param.Z, param.A)
	ctx.JSON(http.StatusOK, gin.H{
		"ori": ori,
		"now": now,
		"changeB": b,
		"suggest": sug,
	})
}