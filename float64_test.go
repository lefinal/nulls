package nulls

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// TestNewFloat64 tests NewFloat64.
func TestNewFloat64(t *testing.T) {
	f := NewFloat64(64)
	assert.True(t, f.Valid, "should be valid")
	assert.EqualValues(t, 64, f.Float64, "should contain correct value")
}

// Float64MarshalJSONSuite tests Float64.MarshalJSON.
type Float64MarshalJSONSuite struct {
	suite.Suite
}

func (suite *Float64MarshalJSONSuite) TestNotValid() {
	f := Float64{Float64: 64}
	raw, err := json.Marshal(f)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(jsonNull, raw, "should return correct value")
}

func (suite *Float64MarshalJSONSuite) TestOK() {
	f := NewFloat64(64)
	raw, err := json.Marshal(f)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(marshalMust(64), raw, "should return correct value")
}

func TestFloat64_MarshalJSON(t *testing.T) {
	suite.Run(t, new(Float64MarshalJSONSuite))
}

// Float64UnmarshalJSONSuite tests Float64.UnmarshalJSON.
type Float64UnmarshalJSONSuite struct {
	suite.Suite
}

func (suite *Float64UnmarshalJSONSuite) TestNull() {
	var f Float64
	err := json.Unmarshal(jsonNull, &f)
	suite.Require().NoError(err, "should not fail")
	suite.False(f.Valid, "should not be valid")
}

func (suite *Float64UnmarshalJSONSuite) TestOK() {
	var f Float64
	err := json.Unmarshal(marshalMust(64), &f)
	suite.Require().NoError(err, "should not fail")
	suite.True(f.Valid, "should be valid")
	suite.EqualValues(64, f.Float64, "should unmarshal correct value")
}

func TestFloat64_UnmarshalJSON(t *testing.T) {
	suite.Run(t, new(Float64UnmarshalJSONSuite))
}

// Float64ScanSuite tests Float64.Scan.
type Float64ScanSuite struct {
	suite.Suite
}

func (suite *Float64ScanSuite) TestNull() {
	var f Float64
	err := f.Scan(nil)
	suite.Require().NoError(err, "should not fail")
	suite.False(f.Valid, "should not be valid")
}

func (suite *Float64ScanSuite) TestOK() {
	var f Float64
	err := f.Scan(64)
	suite.Require().NoError(err, "should not fail")
	suite.True(f.Valid, "should be valid")
	suite.EqualValues(64, f.Float64, "should scan correct value")
}

func TestFloat64_Scan(t *testing.T) {
	suite.Run(t, new(Float64ScanSuite))
}

// Float64ValueSuite tests Float64.Value.
type Float64ValueSuite struct {
	suite.Suite
}

func (suite *Float64ValueSuite) TestNull() {
	f := Float64{Float64: 64}
	raw, err := f.Value()
	suite.Require().NoError(err, "should not fail")
	suite.Nil(raw, "should return correct value")
}

func (suite *Float64ValueSuite) TestOK() {
	f := NewFloat64(64)
	raw, err := f.Value()
	suite.Require().NoError(err, "should not fail")
	suite.Equal(float64(64), raw, "should return correct value")
}

func TestFloat64_Value(t *testing.T) {
	suite.Run(t, new(Float64ValueSuite))
}
