package goweb

import (
	"testing"
	"url"
	"bytes"
	"http"
	"io/ioutil"
)

type personTestStruct struct {
	Name  string
	Age   int
	Atoms int64
}

func makeFormData() string {
	form := make(url.Values)
	form.Add("Name", "Alice")
	form.Add("Age", "25")
	form.Add("Atoms", "29357029322375092")
	return form.Encode()
}

func makeTestContextWithContentTypeAndBody(ct, body string) *Context {
	var request *http.Request = new(http.Request)
	request.URL, _ = url.Parse("http://www.example.com/test")
	// add content type
	request.Header = make(http.Header)
	request.Header.Add("Content-Type", ct)
	// add form data as ReadCloser
	data := ioutil.NopCloser(bytes.NewBufferString(body))
	request.Body = data
	request.ContentLength = int64(len(body))
	request.Method = "POST"
	// setup context
	var responseWriter http.ResponseWriter
	var pathParams ParameterValueMap = make(ParameterValueMap)
	return makeContext(request, responseWriter, pathParams)
}

func TestFormDecoding(t *testing.T) {
	cx := makeTestContextWithContentTypeAndBody(
		"application/x-www-form-urlencoded; charset=utf8", makeFormData())
	// fill struct 
	var person personTestStruct
	err := cx.Fill(&person)
	if err != nil {
		t.Errorf("form-decoder:", err)
	}
	// check it
	if person.Name != "Alice" {
		t.Errorf("form-decoder: expected 'alice' got %v", person.Name)
	}
	if person.Age != 25 {
		t.Errorf("form-decoder: expected '25' got %v", person.Age)
	}
	if person.Atoms != int64(29357029322375092) {
		t.Errorf("form-decoders: expected int64 '29357029322375092' got %v", person.Atoms)
	}
}

func TestFormDecodingPtrPtr(t *testing.T) {
	cx := makeTestContextWithContentTypeAndBody(
		"application/x-www-form-urlencoded; charset=utf8", makeFormData())
	// fill struct via ptr -> ptr
	person := &personTestStruct{}
	err := cx.Fill(&person)
	if err != nil {
		t.Errorf("form-decoder:", err)
	}
	// check it
	if person.Name != "Alice" {
		t.Errorf("form-decoder: expected 'alice' got %v", person.Name)
	}
	if person.Age != 25 {
		t.Errorf("form-decoder: expected '25' got %v", person.Age)
	}
	if person.Atoms != int64(29357029322375092) {
		t.Errorf("form-decoders: expected int64 '29357029322375092' got %v", person.Atoms)
	}
}

func makeJsonData() string {
	return `{"Name":"Alice", "Age":25, "Atoms":29357029322375092}`
}

func TestJsonDecoding(t *testing.T) {
	cx := makeTestContextWithContentTypeAndBody("application/json", makeJsonData())
	// fill struct 
	person := new(personTestStruct)
	err := cx.Fill(person)
	if err != nil {
		t.Errorf("form-decoder:", err)
	}
	// check it
	if person.Name != "Alice" {
		t.Errorf("form-decoder: expected 'alice' got %v", person.Name)
	}
	if person.Age != 25 {
		t.Errorf("form-decoder: expected '25' got %v", person.Age)
	}
	if person.Atoms != int64(29357029322375092) {
		t.Errorf("form-decoders: expected int64 '29357029322375092' got %v", person.Atoms)
	}
}

func TestUnknownDecoding(t *testing.T) {
	cx := makeTestContextWithContentTypeAndBody("application/junk", "<<junk>>")
	// fill struct 
	person := new(personTestStruct)
	err := cx.Fill(person)
	if err == nil {
		t.Errorf("form-decoder: should have raised error for unknown content-type: application/junk")
	}
}
