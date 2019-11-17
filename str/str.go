package str

import (
	"encoding/json"
	"fmt"
	"time"
)

// AllTrueFilter ...
var AllTrueFilter StringFilter = func(string) bool { return true }

// NotEmptyFilter ...
var NotEmptyFilter StringFilter = func(s string) bool { return s != "" }

// IdenticalMapper ...
var IdenticalMapper StringMapper = func(s string) string { return s }

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
