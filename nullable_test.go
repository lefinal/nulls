package nulls

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

// TestNewNullable tests NewNullable.
func TestNewNullable(t *testing.T) {
	n := NewNullable[*sql.NullBool](&sql.NullBool{Bool: true})
	assert.True(t, n.Valid, "should be valid")
	assert.True(t, n.V.Bool, "should have set correct value")
}

// nullableValueMock implements NullableValue.
type nullableValueMock struct {
	mock.Mock
}

func (n *nullableValueMock) MarshalJSON() ([]byte, error) {
	args := n.Called()
	var b []byte
	b, _ = args.Get(0).([]byte)
	return b, args.Error(1)
}

func (n *nullableValueMock) UnmarshalJSON(data []byte) error {
	return n.Called(data).Error(0)
}

func (n *nullableValueMock) Scan(src any) error {
	return n.Called(src).Error(0)
}

func (n *nullableValueMock) Value() (driver.Value, error) {
	args := n.Called()
	var v driver.Value
	v, _ = args.Get(0).(driver.Value)
	return v, args.Error(1)
}

// NullableMarshalJSONSuite tests Nullable.MarshalJSON.
type NullableMarshalJSONSuite struct {
	suite.Suite
}

func (suite *NullableMarshalJSONSuite) TestNotValid() {
	n := Nullable[*nullableValueMock]{V: &nullableValueMock{}}
	raw, err := json.Marshal(n)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(jsonNull, raw, "should return correct value")
}

func (suite *NullableMarshalJSONSuite) TestMarshalFail() {
	n := NewNullable(&nullableValueMock{})
	n.V.On("MarshalJSON").Return(nil, errors.New("sad life"))
	defer n.V.AssertExpectations(suite.T())
	_, err := json.Marshal(n)
	suite.Require().Error(err, "should fail")
}

func (suite *NullableMarshalJSONSuite) TestOK() {
	n := NewNullable(&nullableValueMock{})
	expectRaw := marshalMust("meow")
	n.V.On("MarshalJSON").Return(expectRaw, nil)
	defer n.V.AssertExpectations(suite.T())
	raw, err := json.Marshal(n)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(expectRaw, raw, "should return correct value")
}

func TestNullable_MarshalJSON(t *testing.T) {
	suite.Run(t, new(NullableMarshalJSONSuite))
}

// NullableUnmarshalJSONSuite tests Nullable.UnmarshalJSON.
type NullableUnmarshalJSONSuite struct {
	suite.Suite
}

func (suite *NullableUnmarshalJSONSuite) TestNull() {
	var n Nullable[*nullableValueMock]
	err := json.Unmarshal(jsonNull, &n)
	suite.Require().NoError(err, "should not fail")
	suite.False(n.Valid, "should not be valid")
}

func (suite *NullableUnmarshalJSONSuite) TestUnmarshalFail() {
	raw := marshalMust("meow")
	n := Nullable[*nullableValueMock]{V: &nullableValueMock{}}
	n.V.On("UnmarshalJSON", raw).Return(errors.New("sad life"))
	defer n.V.AssertExpectations(suite.T())
	err := json.Unmarshal(raw, &n)
	suite.Require().Error(err, "should fail")
}

func (suite *NullableUnmarshalJSONSuite) TestOK() {
	raw := marshalMust("meow")
	n := Nullable[*nullableValueMock]{V: &nullableValueMock{}}
	n.V.On("UnmarshalJSON", raw).Return(nil)
	defer n.V.AssertExpectations(suite.T())
	err := json.Unmarshal(raw, &n)
	suite.Require().NoError(err, "should not fail")
	suite.True(n.Valid, "should be valid")
}

func TestNullable_UnmarshalJSON(t *testing.T) {
	suite.Run(t, new(NullableUnmarshalJSONSuite))
}

// NullableScanSuite tests Nullable.Scan.
type NullableScanSuite struct {
	suite.Suite
}

func (suite *NullableScanSuite) TestNull() {
	n := Nullable[*nullableValueMock]{V: &nullableValueMock{}}
	err := n.Scan(nil)
	suite.Require().NoError(err, "should not fail")
	suite.False(n.Valid, "should not be valid")
}

func (suite *NullableScanSuite) TestScanFail() {
	src := "Hello World!"
	n := Nullable[*nullableValueMock]{V: &nullableValueMock{}}
	n.V.On("Scan", src).Return(errors.New("sad life"))
	defer n.V.AssertExpectations(suite.T())
	err := n.Scan(src)
	suite.Require().Error(err, "should fail")
}

func (suite *NullableScanSuite) TestOK() {
	src := "Hello World!"
	n := Nullable[*nullableValueMock]{V: &nullableValueMock{}}
	n.V.On("Scan", src).Return(nil)
	defer n.V.AssertExpectations(suite.T())
	err := n.Scan(src)
	suite.Require().NoError(err, "should not fail")
	suite.True(n.Valid, "should be valid")
}

func TestNullable_Scan(t *testing.T) {
	suite.Run(t, new(NullableScanSuite))
}

// NullableValueSuite tests Nullable.Value.
type NullableValueSuite struct {
	suite.Suite
}

func (suite *NullableValueSuite) TestNull() {
	n := Nullable[*nullableValueMock]{V: &nullableValueMock{}}
	raw, err := n.Value()
	suite.Require().NoError(err, "should not fail")
	suite.Nil(raw, "should return correct value")
}

func (suite *NullableValueSuite) TestValueFail() {
	n := NewNullable(&nullableValueMock{})
	n.V.On("Value").Return(nil, errors.New("sad life"))
	defer n.V.AssertExpectations(suite.T())
	_, err := n.Value()
	suite.Require().Error(err, "should fail")
}

func (suite *NullableValueSuite) TestOK() {
	expectRaw := []byte("Hello World!")
	n := NewNullable(&nullableValueMock{})
	n.V.On("Value").Return(expectRaw, nil)
	defer n.V.AssertExpectations(suite.T())
	raw, err := n.Value()
	suite.Require().NoError(err, "should not fail")
	suite.Equal(expectRaw, raw, "should return correct value")
}

func TestNullable_Value(t *testing.T) {
	suite.Run(t, new(NullableValueSuite))
}
