package nulls

import (
	"encoding/base64"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// TestNewByteSlice tests NewByteSlice.
func TestNewByteSlice(t *testing.T) {
	raw := []byte("Hello World!")
	b := NewByteSlice(raw)
	assert.True(t, b.Valid, "should be valid")
	assert.Equal(t, raw, b.ByteSlice, "should contain correct value")
}

// ByteSliceMarshalJSONSuite tests ByteSlice.MarshalJSON.
type ByteSliceMarshalJSONSuite struct {
	suite.Suite
}

func (suite *ByteSliceMarshalJSONSuite) TestNotValid() {
	b := ByteSlice{ByteSlice: []byte("meow")}
	raw, err := json.Marshal(b)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(jsonNull, raw, "should return correct value")
}

func (suite *ByteSliceMarshalJSONSuite) TestOK() {
	v := []byte("Hello World!")
	b := NewByteSlice(v)
	raw, err := json.Marshal(b)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(marshalMust(v), raw, "should return correct value")
}

func TestByteSlice_MarshalJSON(t *testing.T) {
	suite.Run(t, new(ByteSliceMarshalJSONSuite))
}

// ByteSliceUnmarshalJSONSuite tests ByteSlice.UnmarshalJSON.
type ByteSliceUnmarshalJSONSuite struct {
	suite.Suite
}

func (suite *ByteSliceUnmarshalJSONSuite) TestNull() {
	var b ByteSlice
	err := json.Unmarshal(jsonNull, &b)
	suite.Require().NoError(err, "should not fail")
	suite.False(b.Valid, "should not be valid")
}

func (suite *ByteSliceUnmarshalJSONSuite) TestOK() {
	v := []byte("Hello World!")
	var b ByteSlice
	err := json.Unmarshal(marshalMust(v), &b)
	suite.Require().NoError(err, "should not fail")
	suite.True(b.Valid, "should be valid")
	suite.Equal(v, b.ByteSlice, "should unmarshal correct value")
}

func TestByteSlice_UnmarshalJSON(t *testing.T) {
	suite.Run(t, new(ByteSliceUnmarshalJSONSuite))
}

// ByteSliceScanSuite tests ByteSlice.Scan.
type ByteSliceScanSuite struct {
	suite.Suite
}

func (suite *ByteSliceScanSuite) TestNull() {
	var b ByteSlice
	err := b.Scan(nil)
	suite.Require().NoError(err, "should not fail")
	suite.False(b.Valid, "should not be valid")
}

func (suite *ByteSliceScanSuite) TestOK() {
	v := []byte("Hello World!")
	var b ByteSlice
	err := b.Scan(base64.StdEncoding.EncodeToString(v))
	suite.Require().NoError(err, "should not fail")
	suite.True(b.Valid, "should be valid")
	suite.Equal(v, b.ByteSlice, "should scan correct value")
}

func TestByteSlice_Scan(t *testing.T) {
	suite.Run(t, new(ByteSliceScanSuite))
}

// ByteSliceValueSuite tests ByteSlice.Value.
type ByteSliceValueSuite struct {
	suite.Suite
}

func (suite *ByteSliceValueSuite) TestNull() {
	b := ByteSlice{ByteSlice: []byte("Hello World!")}
	raw, err := b.Value()
	suite.Require().NoError(err, "should not fail")
	suite.Nil(raw, "should return correct value")
}

func (suite *ByteSliceValueSuite) TestOK() {
	v := []byte("Hello World")
	b := NewByteSlice(v)
	raw, err := b.Value()
	suite.Require().NoError(err, "should not fail")
	suite.Equal(base64.StdEncoding.EncodeToString(v), raw, "should return correct value")
}

func TestByteSlice_Value(t *testing.T) {
	suite.Run(t, new(ByteSliceValueSuite))
}
