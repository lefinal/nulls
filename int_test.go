package nulls

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// TestNewInt tests NewInt.
func TestNewInt(t *testing.T) {
	i := NewInt(64)
	assert.True(t, i.Valid, "should be valid")
	assert.EqualValues(t, 64, i.Int, "should contain correct value")
}

// IntMarshalJSONSuite tests Int.MarshalJSON.
type IntMarshalJSONSuite struct {
	suite.Suite
}

func (suite *IntMarshalJSONSuite) TestNotValid() {
	i := Int{Int: 64}
	raw, err := json.Marshal(i)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(jsonNull, raw, "should return correct value")
}

func (suite *IntMarshalJSONSuite) TestOK() {
	i := NewInt(64)
	raw, err := json.Marshal(i)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(marshalMust(64), raw, "should return correct value")
}

func TestInt_MarshalJSON(t *testing.T) {
	suite.Run(t, new(IntMarshalJSONSuite))
}

// IntUnmarshalJSONSuite tests Int.UnmarshalJSON.
type IntUnmarshalJSONSuite struct {
	suite.Suite
}

func (suite *IntUnmarshalJSONSuite) TestNull() {
	var i Int
	err := json.Unmarshal(jsonNull, &i)
	suite.Require().NoError(err, "should not fail")
	suite.False(i.Valid, "should not be valid")
}

func (suite *IntUnmarshalJSONSuite) TestOK() {
	var i Int
	err := json.Unmarshal(marshalMust(64), &i)
	suite.Require().NoError(err, "should not fail")
	suite.True(i.Valid, "should be valid")
	suite.EqualValues(64, i.Int, "should unmarshal correct value")
}

func TestInt_UnmarshalJSON(t *testing.T) {
	suite.Run(t, new(IntUnmarshalJSONSuite))
}

// IntScanSuite tests Int.Scan.
type IntScanSuite struct {
	suite.Suite
}

func (suite *IntScanSuite) TestNull() {
	var i Int
	err := i.Scan(nil)
	suite.Require().NoError(err, "should not fail")
	suite.False(i.Valid, "should not be valid")
}

func (suite *IntScanSuite) TestOK() {
	var i Int
	err := i.Scan(64)
	suite.Require().NoError(err, "should not fail")
	suite.True(i.Valid, "should be valid")
	suite.EqualValues(64, i.Int, "should scan correct value")
}

func TestInt_Scan(t *testing.T) {
	suite.Run(t, new(IntScanSuite))
}

// IntValueSuite tests Int.Value.
type IntValueSuite struct {
	suite.Suite
}

func (suite *IntValueSuite) TestNull() {
	i := Int{Int: 64}
	raw, err := i.Value()
	suite.Require().NoError(err, "should not fail")
	suite.Nil(raw, "should return correct value")
}

func (suite *IntValueSuite) TestOK() {
	i := NewInt(64)
	raw, err := i.Value()
	suite.Require().NoError(err, "should not fail")
	suite.EqualValues(64, raw, "should return correct value")
}

func TestInt_Value(t *testing.T) {
	suite.Run(t, new(IntValueSuite))
}
