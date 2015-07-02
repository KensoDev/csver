package csver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	QueryFile string
	OutFile   string
}

func NewConfig(jsonBlob []byte) []Config {
	var c []Config
	err := json.Unmarshal(jsonBlob, &c)
	if err != nil {
		fmt.Println("error:", err)
	}
	return c
}

func (c *Config) getQuery() string {
	file, err := ioutil.ReadFile(c.QueryFile)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(file)
}
