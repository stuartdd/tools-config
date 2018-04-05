package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ConfigInterface interface {
	Validate(filename string) error
}

func LoadJsonRequired(filename string, configObject ConfigInterface, errorCode int) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File: [%s]. Error: %s\n", filename, err)
		os.Exit(errorCode)
	}
	err = json.Unmarshal(content, configObject)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Json content could not be parsed for file: [%s]. Error: %s\n", filename, err)
		os.Exit(errorCode)
	}
	path, _ := filepath.Abs(filename)
	return configObject.Validate(path)
}

func LoadJson(filename string, m interface{}) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(content, m)
}

func StoreJson(filename string, m interface{}) error {
	s, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, s, 0777)
}

func StringJson(m interface{}) (string, error) {
	s, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(s), nil
}
