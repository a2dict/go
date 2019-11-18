package str

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// AllTrueFilter ...
var AllTrueFilter StringFilter = func(string) bool { return true }

// NotEmptyFilter ...
var NotEmptyFilter StringFilter = func(s string) bool { return s != "" }

// IdenticalMapper ...
var IdenticalMapper StringMapper = func(s string) string { return s }

var Urlencoder StringMapper = url.QueryEscape
var Urldecoder StringMapper = func(s string) string { return MustReturn(url.QueryUnescape(s)) }

// ContainStringFilter ...
func ContainStringFilter(strs ...string) StringFilter {
	return func(s string) bool {
		for _, t := range strs {
			if t == s {
				return true
			}
		}
		return false
	}
}

// ExcludeStringFilter ...
func ExcludeStringFilter(strs ...string) StringFilter {
	cf := ContainStringFilter(strs...)
	return func(s string) bool {
		return !cf(s)
	}
}

// Or return first NotEmpty string
func Or(vs ...string) string {
	for _, v := range vs {
		if v != "" {
			return v
		}
	}
	return ""
}

// Purify ...
func Purify(v interface{}) string {
	switch vv := v.(type) {
	case string:
		return vv
	case time.Time:
		// rfc3339 time
		return vv.Format("2006-01-02T15:04:05Z07:00")
	case *time.Time:
		return vv.Format("2006-01-02T15:04:05Z07:00")
	default:
		return fmt.Sprintf("%v", v)
	}
}

// JsonStr ...
func JsonStr(v interface{}) string {
	bs, _ := json.Marshal(v)
	return string(bs)
}

func Join(delimiter string, elements ...string) string {
	var buffer bytes.Buffer
	prepend := ""
	for _, ele := range elements {
		buffer.WriteString(prepend)
		buffer.WriteString(ele)
		prepend = delimiter
	}
	return buffer.String()
}

// TransString map []string to []string
func TransString(mapper StringMapper, vs ...string) []string {
	res := make([]string, 0, len(vs))
	for _, v := range vs {
		res = append(res, mapper(v))
	}
	return res
}

func MustReturn(res string, _ error) string {
	return res
}
