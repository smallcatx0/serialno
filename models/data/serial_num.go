package data

import (
	"fmt"
	"serialno/models/dao"
	"serialno/pkg/glog"
	"strconv"
	"time"

	"github.com/golang-module/carbon"
)

type SerNum struct {
	TrackID string
}

func (s *SerNum) BuildNo(prefix string, len int) (no string, err error) {
	t := carbon.Now()
	timeFactor := t.Format("ymd")
	rk := "ono:" + timeFactor
	hk := "pf_" + prefix
	res := dao.Rdb.HIncrBy(dao.Rdb.Context(), rk, hk, 1)
	if err = res.Err(); err != nil {
		glog.Error("[redis] 查询redis失败", s.TrackID, res.String(), err.Error())
		return
	}
	num := res.Val()
	if num == 1 {
		err = dao.Rdb.Expire(dao.Rdb.Context(), rk, 168*time.Hour).Err()
		glog.ErrorOnly("[redis] 添加过期时间失败", s.TrackID, err)
	}
	if len == 0 {
		len = 4
	}
	no = fmt.Sprintf("%s%s%0"+strconv.Itoa(len)+"d",
		prefix, timeFactor, num,
	)
	return
}

func (s *SerNum) PrexNum(day string, prefix string, num int) (max int, err error) {
	var t carbon.Carbon
	if day != "" {
		t = carbon.Parse(day)
	} else {
		t = carbon.Now()
	}
	timeFactor := t.Format("ymd")
	rk := "ono:" + timeFactor
	hk := "pf_" + prefix
	res := dao.Rdb.HSet(dao.Rdb.Context(), rk, hk, num)
	if err = res.Err(); err != nil {
		glog.Error("[redis] 查询redis失败", s.TrackID, res.String(), err.Error())
		return
	}
	max = num
	return
}
