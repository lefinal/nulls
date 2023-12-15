package nulls

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type myStruct struct {
	A string
}

func (m myStruct) ScanInto(src any, dst *myStruct) error {
	switch src := src.(type) {
	case string:
		dst.A = src
		return nil
	default:
		return fmt.Errorf("unsupported src type: %T", src)
	}
}

func (m myStruct) Value() (driver.Value, error) {
	return m.A, nil
}

func TestNullableInto2(t *testing.T) {
	m := myStruct{A: "Hello World!"}
	myNullable := NewNullableInto(m)
	assert.True(t, myNullable.Valid, "should be valid")
	assert.Equal(t, m, myNullable.V, "should have set correct value")
	err := myNullable.Scan("Ola!")
	require.NoError(t, err, "scan should not fail")
	assert.True(t, myNullable.Valid, "should still be valid")
	assert.Equal(t, "Ola!", myNullable.V.A, "should have scanned correctly")
}

// TestNewNullableInto tests NewNullableInto.
func TestNewNullableInto(t *testing.T) {
	n := NewNullableInto[myStruct](myStruct{A: "Hello World!"})
	assert.True(t, n.Valid, "should be valid")
	assert.Equal(t, "Hello World!", n.V.A, "should have set correct value")
}

// NullableIntoValueMock implements NullableIntoValue.
type NullableIntoValueMock struct {
	mock.Mock
}

func (n NullableIntoValueMock) MarshalJSON() ([]byte, error) {
	args := n.Called()
	var b []byte
	b, _ = args.Get(0).([]byte)
	return b, args.Error(1)
}

func (n *NullableIntoValueMock) UnmarshalJSON(data []byte) error {
	return n.Called(data).Error(0)
}

func (n NullableIntoValueMock) ScanInto(src any, dst *NullableIntoValueMock) error {
	return n.Called(src, dst).Error(0)
}

func (n NullableIntoValueMock) Value() (driver.Value, error) {
	args := n.Called()
	var v driver.Value
	v, _ = args.Get(0).(driver.Value)
	return v, args.Error(1)
}

// NullableIntoMarshalJSONSuite tests NullableInto.MarshalJSON.
type NullableIntoMarshalJSONSuite struct {
	suite.Suite
}

func (suite *NullableIntoMarshalJSONSuite) TestNotValid() {
	n := NullableInto[NullableIntoValueMock]{V: NullableIntoValueMock{}}
	raw, err := json.Marshal(n)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(jsonNull, raw, "should return correct value")
}

func (suite *NullableIntoMarshalJSONSuite) TestMarshalFail() {
	n := NewNullableInto(NullableIntoValueMock{})
	n.V.On("MarshalJSON").Return(nil, errors.New("sad life"))
	defer n.V.AssertExpectations(suite.T())
	_, err := json.Marshal(n)
	suite.Require().Error(err, "should fail")
}

func (suite *NullableIntoMarshalJSONSuite) TestOK() {
	n := NewNullableInto(NullableIntoValueMock{})
	expectRaw := marshalMust("meow")
	n.V.On("MarshalJSON").Return(expectRaw, nil)
	defer n.V.AssertExpectations(suite.T())
	raw, err := json.Marshal(n)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(expectRaw, raw, "should return correct value")
}

func TestNullableInto_MarshalJSON(t *testing.T) {
	suite.Run(t, new(NullableIntoMarshalJSONSuite))
}

// NullableIntoUnmarshalJSONSuite tests NullableInto.UnmarshalJSON.
type NullableIntoUnmarshalJSONSuite struct {
	suite.Suite
}

func (suite *NullableIntoUnmarshalJSONSuite) TestNull() {
	var n NullableInto[NullableIntoValueMock]
	err := json.Unmarshal(jsonNull, &n)
	suite.Require().NoError(err, "should not fail")
	suite.False(n.Valid, "should not be valid")
}

func (suite *NullableIntoUnmarshalJSONSuite) TestUnmarshalFail() {
	raw := marshalMust("meow")
	n := NullableInto[NullableIntoValueMock]{V: NullableIntoValueMock{}}
	n.V.On("UnmarshalJSON", raw).Return(errors.New("sad life"))
	defer n.V.AssertExpectations(suite.T())
	err := json.Unmarshal(raw, &n)
	suite.Require().Error(err, "should fail")
}

func (suite *NullableIntoUnmarshalJSONSuite) TestOK() {
	raw := marshalMust("meow")
	n := NullableInto[NullableIntoValueMock]{V: NullableIntoValueMock{}}
	n.V.On("UnmarshalJSON", raw).Return(nil)
	defer n.V.AssertExpectations(suite.T())
	err := json.Unmarshal(raw, &n)
	suite.Require().NoError(err, "should not fail")
	suite.True(n.Valid, "should be valid")
}

func TestNullableInto_UnmarshalJSON(t *testing.T) {
	suite.Run(t, new(NullableIntoUnmarshalJSONSuite))
}

// NullableIntoScanSuite tests NullableInto.Scan.
type NullableIntoScanSuite struct {
	suite.Suite
}

func (suite *NullableIntoScanSuite) TestNull() {
	n := NullableInto[NullableIntoValueMock]{V: NullableIntoValueMock{}}
	err := n.Scan(nil)
	suite.Require().NoError(err, "should not fail")
	suite.False(n.Valid, "should not be valid")
}

func (suite *NullableIntoScanSuite) TestScanFail() {
	src := "Hello World!"
	n := NullableInto[NullableIntoValueMock]{V: NullableIntoValueMock{}}
	n.V.On("ScanInto", src, &n.V).Return(errors.New("sad life"))
	defer n.V.AssertExpectations(suite.T())
	err := n.Scan(src)
	suite.Require().Error(err, "should fail")
}

func (suite *NullableIntoScanSuite) TestOK() {
	src := "Hello World!"
	n := NullableInto[NullableIntoValueMock]{V: NullableIntoValueMock{}}
	n.V.On("ScanInto", src, &n.V).Return(nil)
	defer n.V.AssertExpectations(suite.T())
	err := n.Scan(src)
	suite.Require().NoError(err, "should not fail")
	suite.True(n.Valid, "should be valid")
}

func TestNullableInto_ScanInto(t *testing.T) {
	suite.Run(t, new(NullableIntoScanSuite))
}

// NullableIntoValueSuite tests NullableInto.Value.
type NullableIntoValueSuite struct {
	suite.Suite
}

func (suite *NullableIntoValueSuite) TestNull() {
	n := NullableInto[NullableIntoValueMock]{V: NullableIntoValueMock{}}
	raw, err := n.Value()
	suite.Require().NoError(err, "should not fail")
	suite.Nil(raw, "should return correct value")
}

func (suite *NullableIntoValueSuite) TestValueFail() {
	n := NewNullableInto(NullableIntoValueMock{})
	n.V.On("Value").Return(nil, errors.New("sad life"))
	defer n.V.AssertExpectations(suite.T())
	_, err := n.Value()
	suite.Require().Error(err, "should fail")
}

func (suite *NullableIntoValueSuite) TestOK() {
	expectRaw := []byte("Hello World!")
	n := NewNullableInto(NullableIntoValueMock{})
	n.V.On("Value").Return(expectRaw, nil)
	defer n.V.AssertExpectations(suite.T())
	raw, err := n.Value()
	suite.Require().NoError(err, "should not fail")
	suite.Equal(expectRaw, raw, "should return correct value")
}

func TestNullableInto_Value(t *testing.T) {
	suite.Run(t, new(NullableIntoValueSuite))
}
