package helper

import (
	"strconv"
)

// "fmt"
// "strconv"

func CheckErr(err error) {
	if err != nil {
		// fmt.Printf("panic program due to the error:%v\n", err)
		panic(err.Error())
	}
}

func ConvertStringToInt(text string) int {
	value, err := strconv.Atoi(text)
	if err != nil {
		return 0
	}
	return value
}
