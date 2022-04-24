package nulls

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// TestNewInt64 tests NewInt64.
func TestNewInt64(t *testing.T) {
	i := NewInt64(64)
	assert.True(t, i.Valid, "should be valid")
	assert.EqualValues(t, 64, i.Int64, "should contain correct value")
}

// Int64MarshalJSONSuite tests Int64.MarshalJSON.
type Int64MarshalJSONSuite struct {
	suite.Suite
}

func (suite *Int64MarshalJSONSuite) TestNotValid() {
	i := Int64{Int64: 64}
	raw, err := json.Marshal(i)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(jsonNull, raw, "should return correct value")
}

func (suite *Int64MarshalJSONSuite) TestOK() {
	i := NewInt64(64)
	raw, err := json.Marshal(i)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(marshalMust(64), raw, "should return correct value")
}

func TestInt64_MarshalJSON(t *testing.T) {
	suite.Run(t, new(Int64MarshalJSONSuite))
}

// Int64UnmarshalJSONSuite tests Int64.UnmarshalJSON.
type Int64UnmarshalJSONSuite struct {
	suite.Suite
}

func (suite *Int64UnmarshalJSONSuite) TestNull() {
	var i Int64
	err := json.Unmarshal(jsonNull, &i)
	suite.Require().NoError(err, "should not fail")
	suite.False(i.Valid, "should not be valid")
}

func (suite *Int64UnmarshalJSONSuite) TestOK() {
	var i Int64
	err := json.Unmarshal(marshalMust(64), &i)
	suite.Require().NoError(err, "should not fail")
	suite.True(i.Valid, "should be valid")
	suite.EqualValues(64, i.Int64, "should unmarshal correct value")
}

func TestInt64_UnmarshalJSON(t *testing.T) {
	suite.Run(t, new(Int64UnmarshalJSONSuite))
}

// Int64ScanSuite tests Int64.Scan.
type Int64ScanSuite struct {
	suite.Suite
}

func (suite *Int64ScanSuite) TestNull() {
	var i Int64
	err := i.Scan(nil)
	suite.Require().NoError(err, "should not fail")
	suite.False(i.Valid, "should not be valid")
}

func (suite *Int64ScanSuite) TestOK() {
	var i Int64
	err := i.Scan(64)
	suite.Require().NoError(err, "should not fail")
	suite.True(i.Valid, "should be valid")
	suite.EqualValues(64, i.Int64, "should scan correct value")
}

func TestInt64_Scan(t *testing.T) {
	suite.Run(t, new(Int64ScanSuite))
}

// Int64ValueSuite tests Int.Value.
type Int64ValueSuite struct {
	suite.Suite
}

func (suite *Int64ValueSuite) TestNull() {
	i := Int64{Int64: 64}
	raw, err := i.Value()
	suite.Require().NoError(err, "should not fail")
	suite.Nil(raw, "should return correct value")
}

func (suite *Int64ValueSuite) TestOK() {
	i := NewInt64(64)
	raw, err := i.Value()
	suite.Require().NoError(err, "should not fail")
	suite.EqualValues(64, raw, "should return correct value")
}

func TestInt64_Value(t *testing.T) {
	suite.Run(t, new(Int64ValueSuite))
}
