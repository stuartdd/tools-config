package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

const FILE_NAME string = "test.json"
const NAME string = "Stuart"
const VALUE string = "Davies"

type StructConfigValidate struct {
	Name  string
	Value string
}

func (p StructConfigValidate) Validate(filename string) error {
	if p.Name != NAME {
		return fmt.Errorf("Name is invalid [%s]", p.Name)
	}
	if p.Value != VALUE {
		return fmt.Errorf("Value is invalid [%s]", p.Value)
	}
	return nil
}

type StructConfig struct {
	Name  string
	Value string
}

func TestValidate(t *testing.T) {
	ioutil.WriteFile(FILE_NAME, []byte("{\"Name\":\""+NAME+"\",\"Value\":\""+VALUE+"\"}"), 0777)
	m2 := StructConfigValidate{}
	err := LoadJsonRequired(FILE_NAME, &m2, 1)
	if err != nil {
		t.Errorf("ERROR: %s", err.Error())
		t.FailNow()
	}

	ioutil.WriteFile(FILE_NAME, []byte("{\"Name\":\""+NAME+"\",\"Value\":\"INVALID\"}"), 0777)
	m2 = StructConfigValidate{}
	err = LoadJsonRequired(FILE_NAME, &m2, 1)
	if err == nil {
		t.Errorf("Should return an error")
		t.FailNow()
	}

	ioutil.WriteFile(FILE_NAME, []byte("{\"Name\":\"NASTY\",\"Value\":\""+VALUE+"\"}"), 0777)
	m2 = StructConfigValidate{}
	err = LoadJsonRequired(FILE_NAME, &m2, 1)
	if err == nil {
		t.Errorf("Should return an error")
		t.FailNow()
	}
	cleanup(t)
}

func TestStoreAndLoad(t *testing.T) {
	m1 := StructConfig{Name: NAME, Value: VALUE}
	/*
		Store the json file
	*/
	err := StoreJson(FILE_NAME, m1)
	if err != nil {
		t.Errorf("ERROR: %s", err.Error())
		t.FailNow()
	}
	/*
		Was the file created?
	*/
	_, err = os.Stat(FILE_NAME)
	if err != nil {
		t.Errorf("ERROR: %s", err.Error())
		t.FailNow()
	}
	/*
		Load a new one from the file
	*/
	m2 := StructConfig{}
	err = LoadJson(FILE_NAME, &m2)
	if err != nil {
		t.Errorf("ERROR: %s", err.Error())
		t.FailNow()
	}
	/*
		Check the names
	*/
	if m1.Name != m2.Name {
		t.Errorf("expected [%s], actual [%s]", m1.Name, m2.Name)
	}
	/*
		Check the values
	*/
	if m1.Value != m2.Value {
		t.Errorf("expected [%s], actual [%s]", m1.Value, m2.Value)
	}
	/*
		Stringify m1
	*/
	string1, err := StringJson(m1)
	if err != nil {
		t.Errorf("ERROR: %s", err.Error())
		t.FailNow()
	}
	/*
		Stringify m2
	*/
	string2, err := StringJson(m2)
	if err != nil {
		t.Errorf("ERROR: %s", err.Error())
		t.FailNow()
	}
	/*
		test the strings are the same
	*/
	if string1 != string2 {
		t.Errorf("expected [%s], actual [%s]", string1, string2)
	}
	cleanup(t)
}

func cleanup(t *testing.T) {
	/*
		Remove the file
	*/
	err := os.Remove(FILE_NAME)
	if err != nil {
		t.Errorf("ERROR: %s", err.Error())
		t.FailNow()
	}

}
