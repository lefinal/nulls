package nulls

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"testing"
)

// jsonNull is the marshalled nil-value.
var jsonNull []byte

func init() {
	var err error
	jsonNull, err = json.Marshal(nil)
	if err != nil {
		panic(err)
	}
}

// marshalMust marshals the given value and panics if marshalling fails.
func marshalMust(v any) []byte {
	raw, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return raw
}

// isNullSuite tests isNull.
type isNullSuite struct {
	suite.Suite
}

func (suite *isNullSuite) TestNil() {
	suite.True(isNull(nil), "should return correct value")
}

func (suite *isNullSuite) TestNull() {
	suite.True(isNull([]byte("null")), "should return correct value")
}

func (suite *isNullSuite) TestEmpty() {
	suite.False(isNull([]byte("")), "should return correct value")
}

func (suite *isNullSuite) TestOK() {
	suite.False(isNull([]byte("Hello World!")), "should return correct value")
}

func TestIsNull(t *testing.T) {
	suite.Run(t, new(isNullSuite))
}

// copyBytesSuite tests copyBytes.
type copyBytesSuite struct {
	suite.Suite
}

func (suite *copyBytesSuite) TestNil() {
	b := copyBytes(nil)
	suite.Nil(b, "should return correct value")
}

func (suite *copyBytesSuite) TestEmpty() {
	original := make([]byte, 0)
	b := copyBytes(original)
	suite.Equal(original, b, "should return correct value")
	suite.NotSame(original, b, "should return copy")
}

func (suite *copyBytesSuite) TestOK() {
	original := []byte("Hello World!")
	b := copyBytes(original)
	suite.Equal(original, b, "should return correct value")
	suite.NotSame(original, b, "should return copy")
}

func TestCopyBytes(t *testing.T) {
	suite.Run(t, new(copyBytesSuite))
}
