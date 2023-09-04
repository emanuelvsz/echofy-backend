package strloader

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var strings = map[string]string{}

func Fetch(key string) string {
	v, ok := strings[key]
	if !ok {
		return key
	}

	return v
}

func Load(dir string) error {
	path := filepath.Join(dir)

	// check if the directory exists
	_, err := os.Stat(path)
	if err != nil {
		return errors.New(fmt.Sprint("directory not found:", dir))
	}

	// get a list of all JSON files in the directory
	files, err := filepath.Glob(filepath.Join(path, "*.json"))
	if err != nil {
		return errors.New(fmt.Sprint("it was not possible to read the JSON files in ", dir))
	}

	// load each JSON file into the strings map
	for _, file := range files {
		data, err := os.ReadFile(file)

		if err != nil {
			return errors.New(fmt.Sprintf("error reading file %s", file))
		}

		var m map[string]string
		err = json.Unmarshal(data, &m)
		if err != nil {
			return errors.New(fmt.Sprintf("error parsing JSON file %s", file))
		}

		for k, v := range m {
			strings[k] = v
		}
	}

	return nil
}
