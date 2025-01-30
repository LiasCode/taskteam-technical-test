package internals

import (
	"encoding/json"
	"fmt"
)

func LogJSON(v interface{}) error {
	b, err := json.MarshalIndent(v, "", "  ")

	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}
