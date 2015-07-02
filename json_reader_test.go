package csver

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestJsonReader(t *testing.T) { TestingT(t) }

type JsonReaderTestSuite struct{}

var _ = Suite(&JsonReaderTestSuite{})

func (s *JsonReaderTestSuite) TestReadJsonFileWhenFileNameWrong(c *C) {
	reader := JsonReader{FileName: "some-missing-filename"}
	content := reader.ReadFile()
	c.Assert(len(content), Equals, 0)
}
