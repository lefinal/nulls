package nulls

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"testing"
)

// NewBoolSuite tests NewBool.
type NewBoolSuite struct {
	suite.Suite
}

func (suite *NewBoolSuite) TestNewTrue() {
	b := NewBool(true)
	suite.True(b.Valid, "should be valid")
	suite.True(b.Bool, "should have correct value")
}

func (suite *NewBoolSuite) TestNewFalse() {
	b := NewBool(false)
	suite.True(b.Valid, "should be valid")
	suite.False(b.Bool, "should have correct value")
}

func TestNewBool(t *testing.T) {
	suite.Run(t, new(NewBoolSuite))
}

// BoolMarshalJSONSuite tests Bool.MarshalJSON.
type BoolMarshalJSONSuite struct {
	suite.Suite
}

func (suite *BoolMarshalJSONSuite) TestNotValid() {
	b := Bool{Bool: true, Valid: false}
	raw, err := json.Marshal(b)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(jsonNull, raw, "should return correct value")
}

func (suite *BoolMarshalJSONSuite) TestOK1() {
	b := NewBool(true)
	raw, err := json.Marshal(b)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(marshalMust(true), raw, "should return correct value")
}

func (suite *BoolMarshalJSONSuite) TestOK2() {
	b := NewBool(false)
	raw, err := json.Marshal(b)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(marshalMust(false), raw, "should return correct value")
}

func TestBool_MarshalJSON(t *testing.T) {
	suite.Run(t, new(BoolMarshalJSONSuite))
}

// BoolUnmarshalJSONSuite tests Bool.UnmarshalJSON.
type BoolUnmarshalJSONSuite struct {
	suite.Suite
}

func (suite *BoolUnmarshalJSONSuite) TestNull() {
	var b Bool
	err := json.Unmarshal(jsonNull, &b)
	suite.Require().NoError(err, "should not fail")
	suite.False(b.Valid, "should not be valid")
}

func (suite *BoolUnmarshalJSONSuite) TestOK1() {
	var b Bool
	err := json.Unmarshal([]byte("true"), &b)
	suite.Require().NoError(err, "should not fail")
	suite.True(b.Valid, "should be valid")
	suite.True(b.Bool, "should unmarshal correct value")
}

func (suite *BoolUnmarshalJSONSuite) TestOK2() {
	var b Bool
	err := json.Unmarshal([]byte("false"), &b)
	suite.Require().NoError(err, "should not fail")
	suite.True(b.Valid, "should be valid")
	suite.False(b.Bool, "should unmarshal correct value")
}

func TestBool_UnmarshalJSON(t *testing.T) {
	suite.Run(t, new(BoolUnmarshalJSONSuite))
}

// BoolScanSuite tests Bool.Scan.
type BoolScanSuite struct {
	suite.Suite
}

func (suite *BoolScanSuite) TestNull() {
	var b Bool
	err := b.Scan(nil)
	suite.Require().NoError(err, "should not fail")
	suite.False(b.Valid, "should not be valid")
}

func (suite *BoolScanSuite) TestOK1() {
	var b Bool
	err := b.Scan([]byte("true"))
	suite.Require().NoError(err, "should not fail")
	suite.True(b.Valid, "should be valid")
	suite.True(b.Bool, "should return correct value")
}

func (suite *BoolScanSuite) TestOK2() {
	var b Bool
	err := b.Scan([]byte("false"))
	suite.Require().NoError(err, "should not fail")
	suite.True(b.Valid, "should be valid")
	suite.False(b.Bool, "should return correct value")
}

func TestBool_Scan(t *testing.T) {
	suite.Run(t, new(BoolScanSuite))
}

// BoolValueSuite tests Bool.Value.
type BoolValueSuite struct {
	suite.Suite
}

func (suite *BoolValueSuite) TestNull() {
	b := Bool{Valid: false}
	v, err := b.Value()
	suite.Require().NoError(err, "should not fail")
	suite.Nil(v, "should return correct value")
}

func (suite *BoolValueSuite) TestOK1() {
	b := NewBool(true)
	v, err := b.Value()
	suite.Require().NoError(err, "should not fail")
	suite.EqualValues(true, v, "should return correct value")
}

func (suite *BoolValueSuite) TestOK2() {
	b := NewBool(false)
	v, err := b.Value()
	suite.Require().NoError(err, "should not fail")
	suite.EqualValues(false, v, "should not fail")
}

func TestBool_Value(t *testing.T) {
	suite.Run(t, new(BoolValueSuite))
}
