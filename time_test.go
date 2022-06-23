package nulls

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

// testTime for usage in tests.
var testTime = time.UnixMilli(1532765)

// TestNewTime tests NewTime.
func TestNewTime(t *testing.T) {
	tt := NewTime(testTime)
	assert.True(t, tt.Valid, "should be valid")
	assert.True(t, testTime.Equal(tt.Time), "should contain correct value")
}

// TimeMarshalJSONSuite tests Time.MarshalJSON.
type TimeMarshalJSONSuite struct {
	suite.Suite
}

func (suite *TimeMarshalJSONSuite) TestNotValid() {
	tt := Time{Time: testTime}
	raw, err := json.Marshal(tt)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(jsonNull, raw, "shoudl return correct value")
}

func (suite *TimeMarshalJSONSuite) TestOK() {
	tt := NewTime(testTime)
	raw, err := json.Marshal(tt)
	suite.Require().NoError(err, "should not fail")
	suite.Equal(marshalMust(testTime), raw, "should return correct value")
}

func TestTime_MarshalJSON(t *testing.T) {
	suite.Run(t, new(TimeMarshalJSONSuite))
}

// TimeUnmarshalJSONSuite tests Time.UnmarshalJSON.
type TimeUnmarshalJSONSuite struct {
	suite.Suite
}

func (suite *TimeUnmarshalJSONSuite) TestNull() {
	var tt Time
	err := tt.UnmarshalJSON(jsonNull)
	suite.Require().NoError(err, "should not fail")
	suite.False(tt.Valid, "should not be valid")
}

func (suite *TimeUnmarshalJSONSuite) TestOK() {
	var tt Time
	err := tt.UnmarshalJSON(marshalMust(testTime))
	suite.Require().NoError(err, "should not fail")
	suite.True(tt.Valid, "should be valid")
	suite.True(testTime.Equal(tt.Time), "should unmarshal correct value")
}

func TestTime_UnmarshalJSON(t *testing.T) {
	suite.Run(t, new(TimeUnmarshalJSONSuite))
}

// TimeScanSuite tests Time.Scan.
type TimeScanSuite struct {
	suite.Suite
}

func (suite *TimeScanSuite) TestNull() {
	var tt Time
	err := tt.Scan(nil)
	suite.Require().NoError(err, "should not fail")
	suite.False(tt.Valid, "should not be valid")
}

func (suite *TimeScanSuite) TestOK() {
	var tt Time
	err := tt.Scan(testTime)
	suite.Require().NoError(err, "should not fail")
	suite.True(tt.Valid, "should be valid")
	suite.True(testTime.Equal(tt.Time), "should scan correct value")
}

func TestTime_Scan(t *testing.T) {
	suite.Run(t, new(TimeScanSuite))
}

// TimeValueSuite tests Time.Value.
type TimeValueSuite struct {
	suite.Suite
}

func (suite *TimeValueSuite) TestNull() {
	tt := Time{Time: testTime}
	raw, err := tt.Value()
	suite.Require().NoError(err, "should not fail")
	suite.Nil(raw, "should return correct value")
}

func (suite *TimeValueSuite) TestOK() {
	tt := NewTime(testTime)
	raw, err := tt.Value()
	suite.Require().NoError(err, "should not fail")
	rawTime, ok := raw.(time.Time)
	suite.Require().True(ok, "returned valued should be time")
	suite.True(testTime.Equal(rawTime), "should return correct value")
}

func TestTime_Value(t *testing.T) {
	suite.Run(t, new(TimeValueSuite))
}

// TimeUTCSuite tests Time.UTC.
type TimeUTCSuite struct {
	suite.Suite
}

func (suite *TimeUTCSuite) TestNull() {
	tt := Time{Time: testTime}
	raw, err := tt.UTC().Value()
	suite.Require().NoError(err, "should not fail")
	suite.Nil(raw, "should return correct value")
}

func (suite *TimeUTCSuite) TestOK() {
	tt := NewTime(testTime)
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	tt.Time.In(loc)
	raw, err := tt.UTC().Value()
	suite.Require().NoError(err, "should not fail")
	rawTime, ok := raw.(time.Time)
	suite.Require().True(ok, "returned valued should be time")
	suite.NotEqual(rawTime, testTime, "times should differ")
	suite.Equal(rawTime, testTime.UTC(), "should return correct value")
}

func TestTime_UTC(t *testing.T) {
	suite.Run(t, new(TimeUTCSuite))
}
