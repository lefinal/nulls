package nulls

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// TestNewInt16 tests NewInt16.
func TestNewInt16(t *testing.T) {
	i := NewInt16(16)
	assert.True(t, i.Valid, "should be valid")
	assert.EqualValues(t, 16, i.Int16, "should contain correct value")
}

// Int16MarshalJSONSuite tests Int16.MarshalJSON.
type Int16MarshalJSONSuite struct {
	suite.Suite
}

func (suite *Int16MarshalJSONSuite) TestNotValid() {
	i := Int16{Int16: 16}
	raw, err := json.Marshal(i)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(jsonNull, raw, "should return correct value")
}

func (suite *Int16MarshalJSONSuite) TestOK() {
	i := NewInt16(16)
	raw, err := json.Marshal(i)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(marshalMust(16), raw, "should return correct value")
}

func TestInt16_MarshalJSON(t *testing.T) {
	suite.Run(t, new(Int16MarshalJSONSuite))
}

// Int16UnmarshalJSONSuite tests Int16.UnmarshalJSON.
type Int16UnmarshalJSONSuite struct {
	suite.Suite
}

func (suite *Int16UnmarshalJSONSuite) TestNull() {
	var i Int16
	err := json.Unmarshal(jsonNull, &i)
	suite.Require().NoError(err, "should not fail")
	suite.False(i.Valid, "should not be valid")
}

func (suite *Int16UnmarshalJSONSuite) TestOK() {
	var i Int16
	err := json.Unmarshal(marshalMust(16), &i)
	suite.Require().NoError(err, "should not fail")
	suite.True(i.Valid, "should be valid")
	suite.EqualValues(16, i.Int16, "should unmarshal correct value")
}

func TestInt16_UnmarshalJSON(t *testing.T) {
	suite.Run(t, new(Int16UnmarshalJSONSuite))
}

// Int16ScanSuite tests Int16.Scan.
type Int16ScanSuite struct {
	suite.Suite
}

func (suite *Int16ScanSuite) TestNull() {
	var i Int16
	err := i.Scan(nil)
	suite.Require().NoError(err, "should not fail")
	suite.False(i.Valid, "should not be valid")
}

func (suite *Int16ScanSuite) TestOK() {
	var i Int16
	err := i.Scan(16)
	suite.Require().NoError(err, "should not fail")
	suite.True(i.Valid, "should be valid")
	suite.EqualValues(16, i.Int16, "should scan correct value")
}

func TestInt16_Scan(t *testing.T) {
	suite.Run(t, new(Int16ScanSuite))
}

// Int16ValueSuite tests Int.Value.
type Int16ValueSuite struct {
	suite.Suite
}

func (suite *Int16ValueSuite) TestNull() {
	i := Int16{Int16: 16}
	raw, err := i.Value()
	suite.Require().NoError(err, "should not fail")
	suite.Nil(raw, "should return correct value")
}

func (suite *Int16ValueSuite) TestOK() {
	i := NewInt16(16)
	raw, err := i.Value()
	suite.Require().NoError(err, "should not fail")
	suite.EqualValues(16, raw, "should return correct value")
}

func TestInt16_Value(t *testing.T) {
	suite.Run(t, new(Int16ValueSuite))
}
