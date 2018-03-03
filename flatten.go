package arbitrary

import (
	"fmt"
	"strings"
)

// Flatten will walk the whole tree and create a new map[string]interface{}
// where all values in that map are basic datatypes. The path from the original
// structure will be converted into a string key. The default joiner will be '.'
func Flatten(data interface{}) (map[string]interface{}, error) {
	return FlattenWithJoiner(data, DefaultJoiner)
}

// FlattenWithJoiner will walk the whole tree and create a new map[string]interface{}
// where all values in that map are basic datatypes. The path from the original
// structure will be converted into a string key. The default joiner will be '.'
func FlattenWithJoiner(
	data interface{},
	joiner Joiner,
) (map[string]interface{}, error) {
	return joiner.flatten(data)
}

// Joiner is a typedef for functions that convert a string array to a string.
type Joiner func([]string) string

// DefaultJoiner will be used if no joiner is specified.
var DefaultJoiner = DotJoiner

// DotJoiner uses "." as the join arg for the array. "foo.bar.baz"
func DotJoiner(strs []string) string {
	return strings.Join(strs, ".")
}

func (j Joiner) cleanJoin(pre interface{}, suff string) string {
	preStr := fmt.Sprintf("%v", pre)
	if suff == "" {
		return preStr
	}
	return j([]string{preStr, suff})
}

func (j Joiner) flatten(data interface{}) (map[string]interface{}, error) {
	switch vv := data.(type) {
	case map[string]interface{}:
		return j.flattenMap(vv)
	case []interface{}:
		return j.flattenArray(vv)
	default:
		return map[string]interface{}{"": data}, nil
	}
}

func (j Joiner) flattenArray(data []interface{}) (map[string]interface{}, error) {
	retMap := map[string]interface{}{}
	for i, v := range data {
		flatVal, _ := j.flatten(v)
		for k, v := range flatVal {
			retMap[j.cleanJoin(i, k)] = v
		}
	}
	return retMap, nil
}

func (j Joiner) flattenMap(data map[string]interface{}) (map[string]interface{}, error) {
	retMap := map[string]interface{}{}
	for k, v := range data {
		flatVal, _ := j.flatten(v)
		for innerK, innerV := range flatVal {
			retMap[j.cleanJoin(k, innerK)] = innerV
		}
	}
	return retMap, nil
}
