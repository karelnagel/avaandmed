package utils

import (
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io"
	"strconv"
	"strings"
	"time"
)

const COMPANIES = 346698
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

func stringReplace(v string) string {
	s := strings.ReplaceAll(v, " ", "")
	s = strings.ReplaceAll(s, ",", ".")
	if s == "" {
		return "0"
	}
	return s
}

func ParseFloat(v string) float64 {
	f, err := strconv.ParseFloat(stringReplace(v), 64)
	if err != nil {
		panic(err)
	}
	return f
}

func ParseInt(v string) int {
	i, err := strconv.Atoi(stringReplace(v))
	if err != nil {
		panic(err)
	}
	return i
}

func NewUTF8Reader(r io.Reader) io.Reader {
	return transform.NewReader(r, charmap.ISO8859_1.NewDecoder())
}
