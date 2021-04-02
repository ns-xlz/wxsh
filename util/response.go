// @Title  
// @Description  
// @Author  xianglizhao  2021/4/2 2:03 下午
package util

import jsoniter "github.com/json-iterator/go"

type JsonResponse struct {
	ErrNo  int			`json:"errNo"`
	ErrMsg string		`json:"errMsg"`
	Data   interface{}	`json:"data"`
}

func (res *JsonResponse) String() string {
	s, _ := jsoniter.MarshalToString(res)
	return s
}

