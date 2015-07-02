package csver

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"os"
)

type CSVer struct {
	config   []Config
	dbConfig *DBConfig
}

type DBConfig struct {
	User   string
	Pass   string
	Host   string
	Dbname string
}

func NewCsver(c []Config, dConfig *DBConfig) (csver *CSVer) {
	csver = &CSVer{config: c, dbConfig: dConfig}
	return
}

func (c *CSVer) Execute() {
	fmt.Printf("%v:%v@tcp(%v:3306)/%v?charset=utf8", c.dbConfig.User, c.dbConfig.Pass, c.dbConfig.Host, c.dbConfig.Dbname)
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:3306)/%v?charset=utf8", c.dbConfig.User, c.dbConfig.Pass, c.dbConfig.Host, c.dbConfig.Dbname))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for _, configuration := range c.config {
		rows, err := db.Query(configuration.getQuery())
		if err != nil {
			fmt.Println("Failed to run query", err)
			return
		}
		fo, _ := os.Create(configuration.OutFile)
		w := bufio.NewWriter(fo)
		dumpTable(rows, w)
	}
}

func dumpTable(rows *sql.Rows, out io.Writer) error {
	colNames, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(out)
	readCols := make([]interface{}, len(colNames))
	writeCols := make([]string, len(colNames))
	for i, _ := range writeCols {
		readCols[i] = &writeCols[i]
	}
	for rows.Next() {
		err := rows.Scan(readCols...)
		if err != nil {
			panic(err)
		}
		writer.Write(writeCols)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	writer.Flush()
	return nil
}
