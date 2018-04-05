package jsonconfig

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type ValidateInterface interface {
	Validate(filename string) error
}

/*
Load and object from a JSON configuration file

This process is very forgiving. If the JSON is valid the response err will be null.

If there are no matching properties then the config data will be unchanged!

If your config object implements 'Validate(filename string) error' then it will be called.

Usage:
	config := Config{
		Timeout:    TIMEOUT_DEFAULT,
		Port:       (PORT_MIN - 1)}
	err := jsonconfig.LoadJson(configFileName, &config)

Note dont forget the '&' on the config parameter.

*/
func LoadJson(filename string, configObject interface{}) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, configObject)
	if err != nil {
		return err
	}

	configObjectValidate, ok := configObject.(ValidateInterface)
	if ok {
		path, err := filepath.Abs(filename)
		if err != nil {
			return err
		}
		err = configObjectValidate.Validate(path)
		return err
	}
	return nil
}

func StoreJson(filename string, configObject interface{}) error {
	s, err := json.Marshal(configObject)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, s, 0777)
}

func StringJson(configObject interface{}) (string, error) {
	s, err := json.Marshal(configObject)
	if err != nil {
		return "", err
	}
	return string(s), nil
}
