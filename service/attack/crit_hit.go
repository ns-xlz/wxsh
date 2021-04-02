// @Title  
// @Description  
// @Author  xianglizhao  2021/4/2 1:53 下午
package attack

import (
	"github.com/gin-gonic/gin"
	"math"
)

// 获取原始攻击数值
func originAttackHits(x, y, z float64) float64 {
	return x*y*z + x*(1-y)
}

// 获取更改后的攻击数值
func nowAttackHits(x, y, z, a, b float64) float64 {
	return (x-a)*(y+b)*z + (x-a)*(1-y-b)
}

// 获取应提升的暴击数值
func getB(x, y, z, a float64) (float64, float64) {
	aa := 1.0
	bb := -1 * (x + 1 - y - x*z + a*z)
	cc := -1 * a * y * z

	return (-bb + math.Sqrt(bb*bb-4*aa*cc)) / (2 * aa), (-bb - math.Sqrt(bb*bb-4*aa*cc)) / (2 * aa)
}

// 对比更改前后的数值
func GetAttackCompare(ctx *gin.Context, x, y, z, a float64) (oriAttackHit, nowAttackHit, changeCritPercent float64, suggest bool) {
	oriAttackHit = originAttackHits(x, y, z)
	changeCritPercent, _ = getB(x, y, z, a)
	nowAttackHit = nowAttackHits(x, y, z, a, changeCritPercent)
	return oriAttackHit, nowAttackHit, changeCritPercent, !(changeCritPercent + y > 1 || changeCritPercent + y < 0)
}