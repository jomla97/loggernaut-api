package parsing

import (
	"encoding/json"
	"fmt"
	"strings"
)

// JSONParser is a parser for JSON data
type JSONParser struct {
	data *[]byte
}

// Parse parses the data as JSON, returning an array of objects
func (p JSONParser) Parse(data *[]byte) (entries []map[string]interface{}, err error) {
	// Attempt to parse the data as an array of objects
	err = json.Unmarshal(*data, &entries)
	if err == nil {
		fmt.Println("Parsed as array")
		return entries, nil
	}

	// Attempt to parse the data as one object
	var obj map[string]interface{}
	err = json.Unmarshal(*data, &obj)
	if err == nil {
		fmt.Println("Parsed as object")
		entries = append(entries, obj)
		return entries, nil
	}

	// Attempt to parse each entry as an object
	for _, entry := range strings.Split(string(*data), "\n") {
		var obj map[string]interface{}
		err = json.Unmarshal([]byte(entry), &obj)
		if err != nil {
			break
		}
		fmt.Printf("Parsed line as object %v\n", obj)
		entries = append(entries, obj)
	}

	if err != nil {
		return []map[string]interface{}{}, err
	}

	return entries, err
}
