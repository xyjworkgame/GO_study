package utils

import (
	"fmt"
	"strconv"
)

func Float64ToInt(data interface{}) int {
	strTran := fmt.Sprintf( "%v",data)
	tran,_ := strconv.Atoi(strTran)
	return tran
}
