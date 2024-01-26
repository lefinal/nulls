package nulls

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestOptional(t *testing.T) {
	m := myStruct{A: "Hello World!"}
	myOptional := NewOptional(m)
	assert.True(t, myOptional.Valid, "should be valid")
	assert.Equal(t, m, myOptional.V, "should have set correct value")
}

// TestNewOptional tests NewOptional.
func TestNewOptional(t *testing.T) {
	n := NewOptional[myStruct](myStruct{A: "Hello World!"})
	assert.True(t, n.Valid, "should be valid")
	assert.Equal(t, "Hello World!", n.V.A, "should have set correct value")
}

// OptionalValueMock implements OptionalValue.
type OptionalValueMock struct {
	mock.Mock
}

func (n OptionalValueMock) MarshalJSON() ([]byte, error) {
	args := n.Called()
	var b []byte
	b, _ = args.Get(0).([]byte)
	return b, args.Error(1)
}

func (n *OptionalValueMock) UnmarshalJSON(data []byte) error {
	return n.Called(data).Error(0)
}

func (n OptionalValueMock) Scan(src any) error {
	return n.Called(src).Error(0)
}

func (n OptionalValueMock) Value() (driver.Value, error) {
	args := n.Called()
	var v driver.Value
	v, _ = args.Get(0).(driver.Value)
	return v, args.Error(1)
}

// OptionalMarshalJSONSuite tests Optional.MarshalJSON.
type OptionalMarshalJSONSuite struct {
	suite.Suite
}

func (suite *OptionalMarshalJSONSuite) TestNotValid() {
	n := Optional[OptionalValueMock]{V: OptionalValueMock{}}
	raw, err := json.Marshal(n)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(jsonNull, raw, "should return correct value")
}

func (suite *OptionalMarshalJSONSuite) TestMarshalFail() {
	n := NewOptional(OptionalValueMock{})
	n.V.On("MarshalJSON").Return(nil, errors.New("sad life"))
	defer n.V.AssertExpectations(suite.T())
	_, err := json.Marshal(n)
	suite.Require().Error(err, "should fail")
}

func (suite *OptionalMarshalJSONSuite) TestOK() {
	n := NewOptional(OptionalValueMock{})
	expectRaw := marshalMust("meow")
	n.V.On("MarshalJSON").Return(expectRaw, nil)
	defer n.V.AssertExpectations(suite.T())
	raw, err := json.Marshal(n)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(expectRaw, raw, "should return correct value")
}

func TestOptional_MarshalJSON(t *testing.T) {
	suite.Run(t, new(OptionalMarshalJSONSuite))
}

// OptionalUnmarshalJSONSuite tests Optional.UnmarshalJSON.
type OptionalUnmarshalJSONSuite struct {
	suite.Suite
}

func (suite *OptionalUnmarshalJSONSuite) TestNull() {
	var n Optional[OptionalValueMock]
	err := json.Unmarshal(jsonNull, &n)
	suite.Require().NoError(err, "should not fail")
	suite.False(n.Valid, "should not be valid")
}

func (suite *OptionalUnmarshalJSONSuite) TestUnmarshalFail() {
	raw := marshalMust("meow")
	n := Optional[OptionalValueMock]{V: OptionalValueMock{}}
	n.V.On("UnmarshalJSON", raw).Return(errors.New("sad life"))
	defer n.V.AssertExpectations(suite.T())
	err := json.Unmarshal(raw, &n)
	suite.Require().Error(err, "should fail")
}

func (suite *OptionalUnmarshalJSONSuite) TestOK() {
	raw := marshalMust("meow")
	n := Optional[OptionalValueMock]{V: OptionalValueMock{}}
	n.V.On("UnmarshalJSON", raw).Return(nil)
	defer n.V.AssertExpectations(suite.T())
	err := json.Unmarshal(raw, &n)
	suite.Require().NoError(err, "should not fail")
	suite.True(n.Valid, "should be valid")
}

func TestOptional_UnmarshalJSON(t *testing.T) {
	suite.Run(t, new(OptionalUnmarshalJSONSuite))
}

// OptionalScanSuite tests Optional.Scan.
type OptionalScanSuite struct {
	suite.Suite
}

func (suite *OptionalScanSuite) TestNull() {
	n := Optional[OptionalValueMock]{V: OptionalValueMock{}}
	err := n.Scan(nil)
	suite.Error(err, "should fail")
}

func (suite *OptionalScanSuite) TestScan() {
	src := "Hello World!"
	n := Optional[OptionalValueMock]{V: OptionalValueMock{}}
	n.V.On("Scan", src).Return(nil).Maybe()
	defer n.V.AssertExpectations(suite.T())
	err := n.Scan(src)
	suite.Error(err, "should fail")
}

func TestOptional_Scan(t *testing.T) {
	suite.Run(t, new(OptionalScanSuite))
}

// OptionalValueSuite tests Optional.Value.
type OptionalValueSuite struct {
	suite.Suite
}

func (suite *OptionalValueSuite) TestNull() {
	n := Optional[OptionalValueMock]{V: OptionalValueMock{}}
	_, err := n.Value()
	suite.Error(err, "should fail")
}

func (suite *OptionalValueSuite) TestValue() {
	n := NewOptional(OptionalValueMock{})
	n.V.On("Value").Return("Hello", nil).Maybe()
	defer n.V.AssertExpectations(suite.T())
	_, err := n.Value()
	suite.Error(err, "should fail")
}

func TestOptional_Value(t *testing.T) {
	suite.Run(t, new(OptionalValueSuite))
}
