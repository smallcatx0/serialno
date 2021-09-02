package v1

import (
	"serialno/middleware/httpmd"
	"serialno/models/data"
	"serialno/valid"

	"github.com/gin-gonic/gin"
)

// 顺序流水号生成
func OrderNo(c *gin.Context) {
	param := struct {
		Prefix string `json:"prefix"`
		Len    int    `json:"len"`
	}{}
	err := valid.BindAndCheck(c, &param)
	if err != nil {
		r.Fail(c, err)
		return
	}
	serv := data.SerNum{TrackID: c.GetHeader(httpmd.RequestIDKey)}
	no, err := serv.BuildNo(param.Prefix, param.Len)
	if err != nil {
		r.Fail(c, err)
		return
	}
	r.Succ(c, no)
}

func PrefixNum(c *gin.Context) {
	param := struct {
		Prefix string `json:"prefix"`
		Day    string `json:"day"`
		Num    int    `json:"num" binding:"required"`
	}{}
	err := valid.BindAndCheck(c, &param)
	if err != nil {
		r.Fail(c, err)
		return
	}
	serv := data.SerNum{TrackID: c.GetHeader(httpmd.RequestIDKey)}
	num, err := serv.PrexNum(param.Day, param.Prefix, param.Num)
	if err != nil {
		r.Fail(c, err)
		return
	}
	r.Succ(c, num)
}
