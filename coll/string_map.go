package coll

import (
	"bytes"
	"sort"

	"github.com/a2dict/go/str"
)

// MapOf ...
func MapOf(kvs ...string) map[string]string {
	res := map[string]string{}
	n := len(kvs)
	for i := 0; i < n; i += 2 {
		k := kvs[i]
		v := ""
		if i+1 < n {
			v = kvs[i+1]
		}
		res[k] = v
	}
	return res
}

// MergeMap ...
func MergeMap(ms ...map[string]string) map[string]string {
	res := map[string]string{}
	for _, m := range ms {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}

// AssocMap ...
func AssocMap(m map[string]string, kvs ...string) map[string]string {
	return MergeMap(m, MapOf(kvs...))
}

// FilterMap ...
func FilterMap(m map[string]string, f str.StringFilter) map[string]string {
	return FilterMap2(m, f, str.AllTrueFilter)
}

// FilterMap2 ...
func FilterMap2(m map[string]string, keyFilter str.StringFilter, valFilter str.StringFilter) map[string]string {
	res := map[string]string{}
	for k, v := range m {
		if keyFilter(k) && valFilter(v) {
			res[k] = v
		}
	}
	return res
}

// SelectKeys ...
func SelectKeys(m map[string]string, ks ...string) map[string]string {
	return FilterMap(m, str.ContainStringFilter(ks...))
}

// ExcludeKeys ...
func ExcludeKeys(m map[string]string, ks ...string) map[string]string {
	return FilterMap(m, str.ExcludeStringFilter(ks...))
}

// PurifyMap ...
func PurifyMap(m map[string]interface{}) map[string]string {
	res := map[string]string{}
	for k, v := range m {
		res[k] = str.Purify(v)
	}
	return res
}

// PolluteMap ...
func PolluteMap(m map[string]string) map[string]interface{} {
	res := make(map[string]interface{})
	for k, v := range m {
		res[k] = v
	}
	return res
}

// Map2SortString ...
func Map2SortString(m map[string]string) string {
	return Map2SortString_(m, "=", "&", str.IdenticalMapper, str.IdenticalMapper, str.AllTrueFilter, str.AllTrueFilter)
}

// Map2SortString_ ...
func Map2SortString_(m map[string]string, join, split string, keyMapper, valueMapper str.StringMapper, keyFilter, valueFilter str.StringFilter) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	s := ""
	var buffer bytes.Buffer
	for _, k := range keys {
		v := m[k]
		if valueFilter(v) && keyFilter(k) {
			buffer.WriteString(s)
			buffer.WriteString(keyMapper(k))
			buffer.WriteString(join)
			buffer.WriteString(valueMapper(v))
			s = split
		}
	}
	ss := buffer.String()
	return ss
}
