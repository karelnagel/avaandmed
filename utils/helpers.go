package utils

import (
	"time"
)

const dateFormat = "02.01.2006"

func Date(v interface{}) int64 {
	date, _ := time.Parse(dateFormat, v.(string))
	return date.Unix()
}

func DatePointer(v *string) *int64 {
	if v == nil {
		return nil
	}
	date, err := time.Parse(dateFormat, *v)
	if err != nil {
		return nil
	}
	unix := date.Unix()
	return &unix
}
