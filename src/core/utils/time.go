package utils

import "time"

var Time TimeInterface = &timeInterface{}

type TimeInterface interface {
	FromUnix(int64) time.Time
	Now() time.Time
	NowUnix() int64
}

type timeInterface struct{}

func (t *timeInterface) FromUnix(i int64) time.Time {
	return time.Unix(i, 0)
}

func (t *timeInterface) Now() time.Time {
	//TODO check time to where to be
	return time.Now().Local()
}

func (t *timeInterface) NowUnix() int64 {
	return time.Now().Unix()
}
