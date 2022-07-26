package nulls

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type aStruct struct {
	An int `json:"an"`
}

// TestNewJSONNullable tests NewJSONNullable.
func TestNewJSONNullable(t *testing.T) {
	n := NewJSONNullable[aStruct](aStruct{An: 12})
	assert.True(t, n.Valid, "should be valid")
	assert.Equal(t, 12, n.V.An, "should have set correct value")
}

// jSONNullableValueMock implements JSONNullableValue.
type jSONNullableValueMock struct {
	mock.Mock
}

func (n *jSONNullableValueMock) MarshalJSON() ([]byte, error) {
	args := n.Called()
	var b []byte
	b, _ = args.Get(0).([]byte)
	return b, args.Error(1)
}

func (n *jSONNullableValueMock) UnmarshalJSON(data []byte) error {
	return n.Called(data).Error(0)
}

// JSONNullableMarshalJSONSuite tests JSONNullable.MarshalJSON.
type JSONNullableMarshalJSONSuite struct {
	suite.Suite
}

func (suite *JSONNullableMarshalJSONSuite) TestNotValid() {
	n := JSONNullable[*jSONNullableValueMock]{V: &jSONNullableValueMock{}}
	raw, err := json.Marshal(n)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(jsonNull, raw, "should return correct value")
}

func (suite *JSONNullableMarshalJSONSuite) TestMarshalFail() {
	n := NewJSONNullable(&jSONNullableValueMock{})
	n.V.On("MarshalJSON").Return(nil, errors.New("sad life"))
	defer n.V.AssertExpectations(suite.T())
	_, err := json.Marshal(n)
	suite.Require().Error(err, "should fail")
}

func (suite *JSONNullableMarshalJSONSuite) TestOK() {
	n := NewJSONNullable(&jSONNullableValueMock{})
	expectRaw := marshalMust("meow")
	n.V.On("MarshalJSON").Return(expectRaw, nil)
	defer n.V.AssertExpectations(suite.T())
	raw, err := json.Marshal(n)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(expectRaw, raw, "should return correct value")
}

func TestJSONNullable_MarshalJSON(t *testing.T) {
	suite.Run(t, new(JSONNullableMarshalJSONSuite))
}

// JSONNullableUnmarshalJSONSuite tests JSONNullable.UnmarshalJSON.
type JSONNullableUnmarshalJSONSuite struct {
	suite.Suite
}

func (suite *JSONNullableUnmarshalJSONSuite) TestNull() {
	var n JSONNullable[*jSONNullableValueMock]
	err := json.Unmarshal(jsonNull, &n)
	suite.Require().NoError(err, "should not fail")
	suite.False(n.Valid, "should not be valid")
}

func (suite *JSONNullableUnmarshalJSONSuite) TestUnmarshalFail() {
	raw := marshalMust("meow")
	n := JSONNullable[*jSONNullableValueMock]{V: &jSONNullableValueMock{}}
	n.V.On("UnmarshalJSON", raw).Return(errors.New("sad life"))
	defer n.V.AssertExpectations(suite.T())
	err := json.Unmarshal(raw, &n)
	suite.Require().Error(err, "should fail")
}

func (suite *JSONNullableUnmarshalJSONSuite) TestOK() {
	raw := marshalMust(`{"an": 12}`)
	n := JSONNullable[*jSONNullableValueMock]{V: &jSONNullableValueMock{}}
	n.V.On("UnmarshalJSON", raw).Return(nil)
	defer n.V.AssertExpectations(suite.T())
	err := json.Unmarshal(raw, &n)
	suite.Require().NoError(err, "should not fail")
	suite.True(n.Valid, "should be valid")
}

func TestJSONNullable_UnmarshalJSON(t *testing.T) {
	suite.Run(t, new(JSONNullableUnmarshalJSONSuite))
}

// JSONNullableScanSuite tests JSONNullable.Scan.
type JSONNullableScanSuite struct {
	suite.Suite
}

func (suite *JSONNullableScanSuite) TestNull() {
	n := JSONNullable[*jSONNullableValueMock]{V: &jSONNullableValueMock{}}
	err := n.Scan(nil)
	suite.Error(err, "should fail")
}

func (suite *JSONNullableScanSuite) TestOK() {
	src := "Hello World!"
	n := JSONNullable[*jSONNullableValueMock]{V: &jSONNullableValueMock{}}
	err := n.Scan(src)
	suite.Error(err, "should fail")
}

func TestJSONNullable_Scan(t *testing.T) {
	suite.Run(t, new(JSONNullableScanSuite))
}

// JSONNullableValueSuite tests JSONNullable.Value.
type JSONNullableValueSuite struct {
	suite.Suite
}

func (suite *JSONNullableValueSuite) TestNull() {
	n := JSONNullable[*jSONNullableValueMock]{V: &jSONNullableValueMock{}}
	_, err := n.Value()
	suite.Error(err, "should fail")
}

func (suite *JSONNullableValueSuite) TestOK() {
	n := NewJSONNullable(&jSONNullableValueMock{})
	_, err := n.Value()
	suite.Error(err, "should fail")
}

func TestJSONNullable_Value(t *testing.T) {
	suite.Run(t, new(JSONNullableValueSuite))
}
