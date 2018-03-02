package arbitrary

import (
	"bytes"
	"encoding/json"
	"io"
)

// TODO: Benchmark various serialization encodings to unsure that this is a wise
//       implementation. Also make sure this is faster than hand coding the conversion.

// Hydrate will convert the source arbitrary object into the dest typed object.
func Hydrate(data interface{}, typedData interface{}) error {
	buf := bytes.Buffer{}
	if err := json.NewEncoder(io.Writer(&buf)).Encode(data); err != nil {
		return err
	}
	if err := json.NewDecoder(io.Reader(&buf)).Decode(typedData); err != nil {
		return err
	}
	return nil
}
