package jsonconfig

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const FILE_NAME string = "test.json"
const NAME string = "Stuart"
const VALUE string = "Davies"
const PLAIN_JSON = "{\"Name\":\"Stuart\",\"Value\":\"Davies\"}"

type Validatable struct {
	Name  string
	Value string
}

func (p Validatable) Validate(filename string) error {

	if p.Name != NAME {
		return fmt.Errorf("Name is invalid [%s]", p.Name)
	}
	if p.Value != VALUE {
		return fmt.Errorf("Value is invalid [%s]", p.Value)
	}
	return nil
}

type NonValidatable struct {
	Name  string
	Value string
}

func TestFileLoad(t *testing.T) {
	v := Validatable{}
	err := LoadJson(FILE_NAME, &v)
	assertNotNull(err, t)
	assertContains(err.Error(), "no such file", t)

	ioutil.WriteFile(FILE_NAME, []byte("{Name\":\""+NAME+"\",\"Value\":\""+VALUE+"\"}"), 0777)
	v = Validatable{}
	err = LoadJson(FILE_NAME, &v)
	assertNotNull(err, t)
	assertContains(err.Error(), "invalid character", t)

}

func TestValidate(t *testing.T) {
	ioutil.WriteFile(FILE_NAME, []byte("{\"Name\":\""+NAME+"\",\"Value\":\""+VALUE+"\"}"), 0777)
	v := Validatable{}
	err := LoadJson(FILE_NAME, &v)
	assertNull(err, t)

	ioutil.WriteFile(FILE_NAME, []byte("{\"Name\":\""+NAME+"\",\"Value\":\"INVALID\"}"), 0777)
	v = Validatable{}
	err = LoadJson(FILE_NAME, &v)
	assertNotNull(err, t)
	assertContains(err.Error(), "Value is invalid", t)

	ioutil.WriteFile(FILE_NAME, []byte("{\"Name\":\"NASTY\",\"Value\":\""+VALUE+"\"}"), 0777)
	v = Validatable{}
	err = LoadJson(FILE_NAME, &v)
	assertNotNull(err, t)
	assertContains(err.Error(), "Name is invalid", t)

	ioutil.WriteFile(FILE_NAME, []byte("{\"Name\":\"NASTY\",\"Value\":\""+VALUE+"\"}"), 0777)
	nv := NonValidatable{}
	err = LoadJson(FILE_NAME, &nv)
	assertNull(err, t)
	assertEquals("NASTY", nv.Name, t)
	assertEquals("Davies", nv.Value, t)

	cleanup(t)
}

func TestStoreAndLoad(t *testing.T) {
	m1 := Validatable{Name: NAME, Value: VALUE}
	/*
		Store the json file
	*/
	err := StoreJson(FILE_NAME, m1)
	assertNull(err, t)
	/*
		Was the file created?
	*/
	_, err = os.Stat(FILE_NAME)
	assertNull(err, t)
	/*
		Load a new one from the file
	*/
	m2 := NonValidatable{}
	err = LoadJson(FILE_NAME, &m2)
	assertNull(err, t)
	/*
		Check the names
	*/
	assertEquals(m1.Name, m2.Name, t)
	assertEquals(m1.Value, m2.Value, t)
	/*
		Stringify m1
	*/
	string1, err := StringJson(m1)
	assertNull(err, t)
	assertEquals(string1, PLAIN_JSON, t)
	/*
		Stringify m2
	*/
	string2, err := StringJson(m2)
	assertNull(err, t)
	assertEquals(string2, PLAIN_JSON, t)
	cleanup(t)
}

func assertContains(container string, thingContained string, t *testing.T) {
	if strings.Contains(container, thingContained) {
		return
	}
	t.Errorf("ERROR: '%s' does not contain '%s'", container, thingContained)
}

func assertEquals(expected string, actual string, t *testing.T) {
	if expected == actual {
		return
	}
	t.Errorf("ERROR: Expected value '%s' not equal to actual '%s'", expected, actual)
}

func assertNull(expected interface{}, t *testing.T) {
	if expected == nil {
		return
	}
	t.Errorf("ERROR: value is NOT nil")
}

func assertNotNull(expected interface{}, t *testing.T) {
	if expected != nil {
		return
	}
	t.Errorf("ERROR: value is nil")
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
