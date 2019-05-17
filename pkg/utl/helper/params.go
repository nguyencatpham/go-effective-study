package helper

import (
	"errors"
	// "fmt"
	"github.com/thoas/go-funk"
	"strconv"
	"strings"
)

func Str2Ids(str string, splitChar string) ([]int, error) {
	temps := strings.Split(str, splitChar)
	var ids []int
	funk.ForEach(temps, func(item string) {
		id, err := strconv.Atoi(item)
		if err == nil {
			if !funk.Contains(ids, id) {
				ids = append(ids, id)
			}
		}

	})
	if len(ids) != len(temps) {
		return nil, errors.New("invalid params")
	} else {
		return ids, nil
	}

}
func Str2IdStrs(str string, splitChar string) ([]string, error) {
	temps := strings.Split(str, splitChar)
	var ids []string
	funk.ForEach(temps, func(item string) {
		if !funk.Contains(ids, item) {
			ids = append(ids, item)
		}

	})
	if len(ids) != len(temps) {
		return nil, errors.New("invalid params")
	} else {
		return ids, nil
	}

}
