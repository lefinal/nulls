package nulls

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// TestNewJSONRawMessage tests NewJSONRawMessage.
func TestNewJSONRawMessage(t *testing.T) {
	raw := json.RawMessage("Hello World!")
	rm := NewJSONRawMessage(raw)
	assert.True(t, rm.Valid, "should be valid")
	assert.Equal(t, raw, rm.RawMessage, "should contain correct value")
}

// JSONRawMessageMarshalJSONSuite tests JSONRawMessage.MarshalJSON.
type JSONRawMessageMarshalJSONSuite struct {
	suite.Suite
}

func (suite *JSONRawMessageMarshalJSONSuite) TestNull() {
	rm := NewJSONRawMessage(json.RawMessage("null"))
	raw, err := json.Marshal(rm)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(jsonNull, raw, "should return correct value")
}

func (suite *JSONRawMessageMarshalJSONSuite) TestNotValid() {
	rm := JSONRawMessage{RawMessage: json.RawMessage(`{"hello":"world"}`)}
	raw, err := json.Marshal(rm)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(jsonNull, raw, "should return correct value")
}

func (suite *JSONRawMessageMarshalJSONSuite) TestOK() {
	v := json.RawMessage(`{"hello":"world"}`)
	rm := NewJSONRawMessage(v)
	raw, err := json.Marshal(rm)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(marshalMust(v), raw, "should return correct value")
}

func TestJSONRawMessage_MarshalJSON(t *testing.T) {
	suite.Run(t, new(JSONRawMessageMarshalJSONSuite))
}

// JSONRawMessageUnmarshalJSONSuite tests JSONRawMessage.UnmarshalJSON.
type JSONRawMessageUnmarshalJSONSuite struct {
	suite.Suite
}

func (suite *JSONRawMessageUnmarshalJSONSuite) TestUndefinedAsField() {
	var s struct {
		RM JSONRawMessage `json:"rm"`
	}
	err := json.Unmarshal([]byte(`{}`), &s)
	suite.Require().NoError(err, "should not fail")
	suite.False(s.RM.Valid, "should not be valid")
}

func (suite *JSONRawMessageUnmarshalJSONSuite) TestDirectly() {
	var rm JSONRawMessage
	err := rm.UnmarshalJSON(nil)
	suite.Require().NoError(err, "should not fail")
	suite.False(rm.Valid, "should not be valid")
}

func (suite *JSONRawMessageUnmarshalJSONSuite) TestNull() {
	var s struct {
		RM JSONRawMessage `json:"rm"`
	}
	err := json.Unmarshal([]byte(`{"rm": null}`), &s)
	suite.Require().NoError(err, "should not fail")
	suite.False(s.RM.Valid, "should not be valid")
}

func (suite *JSONRawMessageUnmarshalJSONSuite) TestOK() {
	v := json.RawMessage(`{"hello":"world"}`)
	var RM JSONRawMessage
	err := json.Unmarshal(marshalMust(v), &RM)
	suite.Require().NoError(err, "should not fail")
	suite.True(RM.Valid, "should be valid")
	suite.Equal(v, RM.RawMessage, "should unmarshal correct value")
}

func TestJSONRawMessage_UnmarshalJSON(t *testing.T) {
	suite.Run(t, new(JSONRawMessageUnmarshalJSONSuite))
}

// JSONRawMessageScanSuite tests JSONRawMessage.Scan.
type JSONRawMessageScanSuite struct {
	suite.Suite
}

func (suite *JSONRawMessageScanSuite) TestNull() {
	var rm JSONRawMessage
	err := rm.Scan(nil)
	suite.Require().NoError(err, "should not fail")
	suite.False(rm.Valid, "should not be valid")
}

func (suite *JSONRawMessageScanSuite) TestJSONNull() {
	var rm JSONRawMessage
	err := rm.Scan(jsonNull)
	suite.Require().NoError(err, "should not fail")
	suite.True(rm.Valid, "should not be valid")
}

func (suite *JSONRawMessageScanSuite) TestUnexpectedValue() {
	var rm JSONRawMessage
	err := rm.Scan("I'm not a byte slice.")
	suite.Error(err, "should fail")
}

func (suite *JSONRawMessageScanSuite) TestOK() {
	v := json.RawMessage(`{"meow":"woof"}`)
	var rm JSONRawMessage
	err := rm.Scan([]byte(v))
	suite.Require().NoError(err, "should not fail")
	suite.True(rm.Valid, "should be valid")
	suite.Equal(v, rm.RawMessage, "should scan correct value")
}

func TestJSONRawMessage_Scan(t *testing.T) {
	suite.Run(t, new(JSONRawMessageScanSuite))
}

// JSONRawMessageValueSuite tests JSONRawMessage.Value.
type JSONRawMessageValueSuite struct {
	suite.Suite
}

func (suite *JSONRawMessageValueSuite) TestUndefined() {
	rm := JSONRawMessage{RawMessage: json.RawMessage(`{"hello": "world"}`)}
	raw, err := rm.Value()
	suite.Require().NoError(err, "should not fail")
	suite.Nil(raw, "should return correct value")
}

func (suite *JSONRawMessageValueSuite) TestNull() {
	rm := NewJSONRawMessage(json.RawMessage("null"))
	raw, err := rm.Value()
	suite.Require().NoError(err, "should not fail")
	suite.Equal([]byte("null"), raw, "should return correct value")
}

func (suite *JSONRawMessageValueSuite) TestOK() {
	v := json.RawMessage(`{"hello":"world"}`)
	rm := NewJSONRawMessage(v)
	raw, err := rm.Value()
	suite.Require().NoError(err, "should not fail")
	suite.Equal([]byte(v), raw, "should return correct value")
}

func TestJSONRawMessage_Value(t *testing.T) {
	suite.Run(t, new(JSONRawMessageValueSuite))
}
