package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Load loads settings from the given filename
func (s *Settings) Load(filename string) error {
	file, err := os.Open("file.go") // For read access.
	if err != nil {
		return err
	}

	rawJSON, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawJSON, s)
	return err
}
