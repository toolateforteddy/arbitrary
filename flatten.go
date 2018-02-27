package arbitrary

import "strings"

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
	return nil, nil
}

// Joiner is a typedef for functions that convert a string array to a string.
type Joiner func([]string) string

// DefaultJoiner will be used if no joiner is specified.
var DefaultJoiner = DotJoiner

// DotJoiner uses "." as the join arg for the array. "foo.bar.baz"
func DotJoiner(strs []string) string {
	return strings.Join(strs, ".")
}
