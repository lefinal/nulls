package nulls

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// TestNewInt32 tests NewInt32.
func TestNewInt32(t *testing.T) {
	i := NewInt32(64)
	assert.True(t, i.Valid, "should be valid")
	assert.EqualValues(t, 64, i.Int32, "should contain correct value")
}

// Int32MarshalJSONSuite tests Int32.MarshalJSON.
type Int32MarshalJSONSuite struct {
	suite.Suite
}

func (suite *Int32MarshalJSONSuite) TestNotValid() {
	i := Int32{Int32: 64}
	raw, err := json.Marshal(i)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(jsonNull, raw, "should return correct value")
}

func (suite *Int32MarshalJSONSuite) TestOK() {
	i := NewInt32(64)
	raw, err := json.Marshal(i)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(marshalMust(64), raw, "should return correct value")
}

func TestInt32_MarshalJSON(t *testing.T) {
	suite.Run(t, new(Int32MarshalJSONSuite))
}

// Int32UnmarshalJSONSuite tests Int32.UnmarshalJSON.
type Int32UnmarshalJSONSuite struct {
	suite.Suite
}

func (suite *Int32UnmarshalJSONSuite) TestNull() {
	var i Int32
	err := json.Unmarshal(jsonNull, &i)
	suite.Require().NoError(err, "should not fail")
	suite.False(i.Valid, "should not be valid")
}

func (suite *Int32UnmarshalJSONSuite) TestOK() {
	var i Int32
	err := json.Unmarshal(marshalMust(64), &i)
	suite.Require().NoError(err, "should not fail")
	suite.True(i.Valid, "should be valid")
	suite.EqualValues(64, i.Int32, "should unmarshal correct value")
}

func TestInt32_UnmarshalJSON(t *testing.T) {
	suite.Run(t, new(Int32UnmarshalJSONSuite))
}

// Int32ScanSuite tests Int32.Scan.
type Int32ScanSuite struct {
	suite.Suite
}

func (suite *Int32ScanSuite) TestNull() {
	var i Int32
	err := i.Scan(nil)
	suite.Require().NoError(err, "should not fail")
	suite.False(i.Valid, "should not be valid")
}

func (suite *Int32ScanSuite) TestOK() {
	var i Int32
	err := i.Scan(64)
	suite.Require().NoError(err, "should not fail")
	suite.True(i.Valid, "should be valid")
	suite.EqualValues(64, i.Int32, "should scan correct value")
}

func TestInt32_Scan(t *testing.T) {
	suite.Run(t, new(Int32ScanSuite))
}

// Int32ValueSuite tests Int.Value.
type Int32ValueSuite struct {
	suite.Suite
}

func (suite *Int32ValueSuite) TestNull() {
	i := Int32{Int32: 64}
	raw, err := i.Value()
	suite.Require().NoError(err, "should not fail")
	suite.Nil(raw, "should return correct value")
}

func (suite *Int32ValueSuite) TestOK() {
	i := NewInt32(64)
	raw, err := i.Value()
	suite.Require().NoError(err, "should not fail")
	suite.EqualValues(64, raw, "should return correct value")
}

func TestInt32_Value(t *testing.T) {
	suite.Run(t, new(Int32ValueSuite))
}
