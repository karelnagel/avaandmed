package utils

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const COMPANIES = 346698
const DATE_FORMAT = "02.01.2006"

func Date(v interface{}) int64 {
	date, _ := time.Parse(DATE_FORMAT, v.(string))
	return date.Unix()
}

func DatePointer(v *string) *int64 {
	if v == nil {
		return nil
	}
	date, err := time.Parse(DATE_FORMAT, *v)
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

type ProgressBar struct {
	bar *progressbar.ProgressBar
}

func NewProgressBar(max int64, description string) *ProgressBar {
	disable := os.Getenv("DISABLE_PROGRESS")
	if disable == "true" {
		fmt.Println(description)
		return &ProgressBar{
			bar: nil,
		}
	}
	return &ProgressBar{
		bar: progressbar.Default(max, description),
	}
}

func (p *ProgressBar) Add(n int) {
	if p.bar != nil {
		p.bar.Add(n)
	}
}
