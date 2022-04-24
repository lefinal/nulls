package nulls

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// TestNewFloat32 tests NewFloat32.
func TestNewFloat32(t *testing.T) {
	f := NewFloat32(32)
	assert.True(t, f.Valid, "should be valid")
	assert.EqualValues(t, 32, f.Float32, "should contain correct value")
}

// Float32MarshalJSONSuite tests Float32.MarshalJSON.
type Float32MarshalJSONSuite struct {
	suite.Suite
}

func (suite *Float32MarshalJSONSuite) TestNotValid() {
	f := Float32{Float32: 32}
	raw, err := json.Marshal(f)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(jsonNull, raw, "should return correct value")
}

func (suite *Float32MarshalJSONSuite) TestOK() {
	f := NewFloat32(32)
	raw, err := json.Marshal(f)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(marshalMust(32), raw, "should return correct value")
}

func TestFloat32_MarshalJSON(t *testing.T) {
	suite.Run(t, new(Float32MarshalJSONSuite))
}

// Float32UnmarshalJSONSuite tests Float32.UnmarshalJSON.
type Float32UnmarshalJSONSuite struct {
	suite.Suite
}

func (suite *Float32UnmarshalJSONSuite) TestNull() {
	var f Float32
	err := json.Unmarshal(jsonNull, &f)
	suite.Require().NoError(err, "should not fail")
	suite.False(f.Valid, "should not be valid")
}

func (suite *Float32UnmarshalJSONSuite) TestOK() {
	var f Float32
	err := json.Unmarshal(marshalMust(32), &f)
	suite.Require().NoError(err, "should not fail")
	suite.True(f.Valid, "should be valid")
	suite.EqualValues(32, f.Float32, "should unmarshal correct value")
}

func TestFloat32_UnmarshalJSON(t *testing.T) {
	suite.Run(t, new(Float32UnmarshalJSONSuite))
}

// Float32ScanSuite tests Float32.Scan.
type Float32ScanSuite struct {
	suite.Suite
}

func (suite *Float32ScanSuite) TestNull() {
	var f Float32
	err := f.Scan(nil)
	suite.Require().NoError(err, "should not fail")
	suite.False(f.Valid, "should not be valid")
}

func (suite *Float32ScanSuite) TestOK() {
	var f Float32
	err := f.Scan(float64(64))
	suite.Require().NoError(err, "should not fail")
	suite.True(f.Valid, "should be valid")
	suite.EqualValues(64, f.Float32, "should scan correct value")
}

func TestFloat32_Scan(t *testing.T) {
	suite.Run(t, new(Float32ScanSuite))
}

// Float32ValueSuite tests Float32.Value.
type Float32ValueSuite struct {
	suite.Suite
}

func (suite *Float32ValueSuite) TestNull() {
	f := Float32{Float32: 32}
	raw, err := f.Value()
	suite.Require().NoError(err, "should not fail")
	suite.Nil(raw, "should return correct value")
}

func (suite *Float32ValueSuite) TestOK() {
	f := NewFloat32(32)
	raw, err := f.Value()
	suite.Require().NoError(err, "should not fail")
	suite.Equal(float64(32), raw, "should return correct value")
}

func TestFloat32_Value(t *testing.T) {
	suite.Run(t, new(Float32ValueSuite))
}
