package csver

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestConfig(t *testing.T) { TestingT(t) }

type ConfigTestSuite struct{}

var _ = Suite(&ConfigTestSuite{})

func (s *ConfigTestSuite) TestCorrectConfiguration(c *C) {
	jsonReader := JsonReader{FileName: "fixtures/configuration.json"}
	config := NewConfig(jsonReader.ReadFile())
	c.Assert(len(config), Equals, 1)
	c.Assert(config[0].QueryFile, Equals, "fixtures/query.sql")
	c.Assert(config[0].OutFile, Equals, "fixtures/output/x.csv")
}

func (s *ConfigTestSuite) TestGetQuery(c *C) {
	jsonReader := JsonReader{FileName: "fixtures/configuration.json"}
	config := NewConfig(jsonReader.ReadFile())
	c.Assert(config[0].getQuery(), Matches, "*select name from users;\n")
}
