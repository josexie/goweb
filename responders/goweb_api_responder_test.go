package responders

import (
	context_test "github.com/stretchrcom/goweb/webcontext/test"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestAPI_Interface(t *testing.T) {

	assert.Implements(t, (*APIResponder)(nil), new(GowebAPIResponder))

}

func TestNewGowebAPIResponder(t *testing.T) {

	http := new(GowebHTTPResponder)
	api := NewGowebAPIResponder(http)

	assert.Equal(t, http, api.httpResponder)

}

func TestCodecService(t *testing.T) {

	// TODO: this

}

func TestRespond(t *testing.T) {

	http := new(GowebHTTPResponder)
	API := NewGowebAPIResponder(http)
	ctx := context_test.MakeTestContext()
	data := map[string]interface{}{"name": "Mat"}

	API.Respond(ctx, 200, data, nil)

	assert.Equal(t, context_test.TestResponseWriter.Output, "{\"d\":{\"name\":\"Mat\"},\"s\":200}")

}

func TestWriteResponseObject(t *testing.T) {

	http := new(GowebHTTPResponder)
	API := NewGowebAPIResponder(http)
	ctx := context_test.MakeTestContext()
	data := map[string]interface{}{"name": "Mat"}

	API.WriteResponseObject(ctx, 200, data)

	assert.Equal(t, context_test.TestResponseWriter.Output, "{\"name\":\"Mat\"}")

}