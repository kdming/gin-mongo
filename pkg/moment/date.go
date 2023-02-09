package moment

import (
	"strconv"
	"time"
)

func Get_Ytd() string {
	return time.Now().Format("2006-01-02")
}

func GetYtd() string {
	return time.Now().Format("20060102")
}

func GetTimeStampStr() string {
	timestamp := time.Now().Unix()
	return strconv.FormatInt(timestamp, 10)
}
