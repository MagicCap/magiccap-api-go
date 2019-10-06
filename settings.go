package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Load loads settings from the given filename
func (s *Settings) Load(filename string) error {
	file, err := os.Open(filename) // For read access.
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

// Save saves the current config to the given filename
func (s *Settings) Save(filename string) error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, 0644)
	return err
}
