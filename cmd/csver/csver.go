package main

import (
	"fmt"
	"github.com/KensoDev/csver"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	configurationFile = kingpin.Flag("configuration-file", "JSON configuration file").String()
	userName          = kingpin.Flag("username", "DB username").String()
	password          = kingpin.Flag("pass", "DB password").String()
	dbHost            = kingpin.Flag("host", "DB host").String()
	dbName            = kingpin.Flag("db-name", "DB Name").String()
)

func main() {
	kingpin.Parse()

	dbConfig := &csver.DBConfig{
		User:   *userName,
		Pass:   *password,
		Host:   *dbHost,
		Dbname: *dbName,
	}

	fmt.Println(dbConfig)

	jsonReader := csver.JsonReader{FileName: *configurationFile}
	config := csver.NewConfig(jsonReader.ReadFile())

	c := csver.NewCsver(config, dbConfig)
	c.Execute()
}
