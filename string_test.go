package nulls

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// TestNewString tests NewString.
func TestNewString(t *testing.T) {
	s := NewString("Hello World!")
	assert.True(t, s.Valid, "should be valid")
	assert.Equal(t, "Hello World!", s.String, "should contain correct value")
}

// StringMarshalJSONSuite tests String.MarshalJSON.
type StringMarshalJSONSuite struct {
	suite.Suite
}

func (suite *StringMarshalJSONSuite) TestNotValid() {
	s := String{String: "Hello World!"}
	raw, err := json.Marshal(s)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(jsonNull, raw, "should return correct value")
}

func (suite *StringMarshalJSONSuite) TestOK() {
	s := NewString("Hello World!")
	raw, err := json.Marshal(s)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(marshalMust("Hello World!"), raw, "should return correct value")
}

func TestString_MarshalJSON(t *testing.T) {
	suite.Run(t, new(StringMarshalJSONSuite))
}

// StringUnmarshalJSONSuite tests String.UnmarshalJSON.
type StringUnmarshalJSONSuite struct {
	suite.Suite
}

func (suite *StringUnmarshalJSONSuite) TestNull() {
	var s String
	err := json.Unmarshal(jsonNull, &s)
	suite.Require().NoError(err, "should not fail")
	suite.False(s.Valid, "should not be valid")
}

func (suite *StringUnmarshalJSONSuite) TestOK() {
	var s String
	err := json.Unmarshal(marshalMust("Hello World!"), &s)
	suite.Require().NoError(err, "should not fail")
	suite.True(s.Valid, "should be valid")
	suite.Equal("Hello World!", s.String, "should unmarshal correct value")
}

func TestString_UnmarshalJSON(t *testing.T) {
	suite.Run(t, new(StringUnmarshalJSONSuite))
}

// StringScanSuite tests String.Scan.
type StringScanSuite struct {
	suite.Suite
}

func (suite *StringScanSuite) TestNull() {
	var s String
	err := s.Scan(nil)
	suite.Require().NoError(err, "should not fail")
	suite.False(s.Valid, "should not be valid")
}

func (suite *StringScanSuite) TestOK() {
	var s String
	err := s.Scan("Hello World!")
	suite.Require().NoError(err, "should not fail")
	suite.True(s.Valid, "should be valid")
	suite.Equal("Hello World!", s.String, "should scan correct value")
}

func TestString_Scan(t *testing.T) {
	suite.Run(t, new(StringScanSuite))
}

// StringValueSuite tests String.Value.
type StringValueSuite struct {
	suite.Suite
}

func (suite *StringValueSuite) TestNull() {
	s := String{String: "Hello World!"}
	raw, err := s.Value()
	suite.Require().NoError(err, "should not fail")
	suite.Nil(raw, "should return correct value")
}

func (suite *StringValueSuite) TestOK() {
	s := NewString("Hello World!")
	raw, err := s.Value()
	suite.Require().NoError(err, "should not fail")
	suite.EqualValues("Hello World!", raw, "should return correct value")
}

func TestString_Value(t *testing.T) {
	suite.Run(t, new(StringValueSuite))
}
