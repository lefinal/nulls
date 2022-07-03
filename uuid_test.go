package nulls

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

func newUUIDV4() uuid.UUID {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return id
}

// TestNewUUID tests NewUUID.
func TestNewUUID(t *testing.T) {
	id := newUUIDV4()
	s := NewUUID(id)
	assert.True(t, s.Valid, "should be valid")
	assert.Equal(t, id, s.UUID, "should contain correct value")
}

// UUIDMarshalJSONSuite tests UUID.MarshalJSON.
type UUIDMarshalJSONSuite struct {
	suite.Suite
}

func (suite *UUIDMarshalJSONSuite) TestNotValid() {
	s := uuid.NullUUID{UUID: newUUIDV4()}
	raw, err := json.Marshal(s)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(jsonNull, raw, "should return correct value")
}

func (suite *UUIDMarshalJSONSuite) TestOK() {
	id := newUUIDV4()
	s := NewUUID(id)
	raw, err := json.Marshal(s)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(marshalMust(id), raw, "should return correct value")
}

func TestUUID_MarshalJSON(t *testing.T) {
	suite.Run(t, new(UUIDMarshalJSONSuite))
}

// UUIDUnmarshalJSONSuite tests UUID.UnmarshalJSON.
type UUIDUnmarshalJSONSuite struct {
	suite.Suite
}

func (suite *UUIDUnmarshalJSONSuite) TestNull() {
	var s uuid.NullUUID
	err := json.Unmarshal(jsonNull, &s)
	suite.Require().NoError(err, "should not fail")
	suite.False(s.Valid, "should not be valid")
}

func (suite *UUIDUnmarshalJSONSuite) TestOK() {
	id := newUUIDV4()
	var s uuid.NullUUID
	err := json.Unmarshal(marshalMust(id), &s)
	suite.Require().NoError(err, "should not fail")
	suite.True(s.Valid, "should be valid")
	suite.Equal(id, s.UUID, "should unmarshal correct value")
}

func TestUUID_UnmarshalJSON(t *testing.T) {
	suite.Run(t, new(UUIDUnmarshalJSONSuite))
}

// UUIDScanSuite tests UUID.Scan.
type UUIDScanSuite struct {
	suite.Suite
}

func (suite *UUIDScanSuite) TestNull() {
	var s uuid.NullUUID
	err := s.Scan(nil)
	suite.Require().NoError(err, "should not fail")
	suite.False(s.Valid, "should not be valid")
}

func (suite *UUIDScanSuite) TestOK() {
	id := newUUIDV4()
	var s uuid.NullUUID
	err := s.Scan(id)
	suite.Require().NoError(err, "should not fail")
	suite.True(s.Valid, "should be valid")
	suite.Equal(id, s.UUID, "should scan correct value")
}

func TestUUID_Scan(t *testing.T) {
	suite.Run(t, new(UUIDScanSuite))
}

// UUIDValueSuite tests UUID.Value.
type UUIDValueSuite struct {
	suite.Suite
}

func (suite *UUIDValueSuite) TestNull() {
	s := uuid.NullUUID{UUID: newUUIDV4()}
	raw, err := s.Value()
	suite.Require().NoError(err, "should not fail")
	suite.Nil(raw, "should return correct value")
}

func (suite *UUIDValueSuite) TestOK() {
	id := newUUIDV4()
	s := NewUUID(id)
	raw, err := s.Value()
	suite.Require().NoError(err, "should not fail")
	suite.EqualValues(id.String(), raw, "should return correct value")
}

func TestUUID_Value(t *testing.T) {
	suite.Run(t, new(UUIDValueSuite))
}
