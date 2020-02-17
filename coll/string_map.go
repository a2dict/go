package coll

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/a2dict/go/str"
)

type KVMapper func(k, v string) (kk, vv string)
type BiStringMapper func(k, v string) string
type KVSplitter func(string) (k, v string)

var EqJoinedBiStringMapper BiStringMapper = func(k, v string) string { return fmt.Sprintf("%s=%s", k, v) }
var EqJoinedKVSplitter KVSplitter = func(s string) (k, v string) {
	sp := strings.SplitN(s, "=", 2)
	if len(sp) == 2 {
		return sp[0], sp[1]
	}
	return s, ""
}

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
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	prepend := ""
	var buffer bytes.Buffer
	for _, k := range keys {
		v := m[k]
		buffer.WriteString(prepend)
		buffer.WriteString(k)
		buffer.WriteString("=")
		buffer.WriteString(v)
		prepend = "&"
	}
	return buffer.String()
}

func TransMap2(m map[string]string, keyMapper str.StringMapper, valueMapper str.StringMapper) map[string]string {
	res := make(map[string]string)
	for k, v := range m {
		kk := keyMapper(k)
		vv := valueMapper(v)
		res[kk] = vv
	}
	return res
}

func Map2StringSlice(m map[string]string, mp BiStringMapper) []string {
	res := make([]string, 0, len(m))
	for k, v := range m {
		res = append(res, mp(k, v))
	}
	return res
}

func Map2UrlEncodedString(m map[string]string) string {
	return str.Join("&",
		Map2StringSlice(
			TransMap2(
				m, str.Urlencoder, str.Urlencoder),
			EqJoinedBiStringMapper)...)
}

func UrlEncodedString2Map(s string) map[string]string {
	res := make(map[string]string)
	kvs := strings.Split(s, "&")
	for _, kv := range kvs {
		k, v := EqJoinedKVSplitter(kv)
		kk := str.Urldecoder(k)
		vv := str.Urldecoder(v)
		res[kk] = vv
	}
	return res
}
