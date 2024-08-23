package utils

import (
	"time"
)

const dateFormat = "02.01.2006"

func String(v interface{}) string {
	return v.(string)
}
func Date(v interface{}) int64 {
	date, _ := time.Parse(dateFormat, v.(string))
	return date.Unix()
}
func Bool(v interface{}) bool {
	return v.(bool)
}
func Int(v interface{}) int64 {
	return int64(v.(float64))
}

func StringPointer(v interface{}) *string {
	if v == nil {
		return nil
	}
	s, ok := v.(string)
	if !ok {
		return nil
	}
	return &s
}

func BoolPointer(v interface{}) *bool {
	if v == nil {
		return nil
	}
	b, ok := v.(bool)
	if !ok {
		return nil
	}
	return &b
}

func IntPointer(v interface{}) *int64 {
	if v == nil {
		return nil
	}
	f, ok := v.(float64)
	if !ok {
		return nil
	}
	i := int64(f)
	return &i
}

func DatePointer(v interface{}) *int64 {
	if v == nil {
		return nil
	}
	date, err := time.Parse(dateFormat, v.(string))
	if err != nil {
		return nil
	}
	unix := date.Unix()
	return &unix
}
