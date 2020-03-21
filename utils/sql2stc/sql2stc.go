package main

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"strings"
)

type Table struct {
	Name   string `db:"name"`
	Column []Column
}

type Column struct {
	Name     string `db:"name"`
	JsonName string
	Type     string `db:"type"`
}

func init() {
	var err error
	db, err = sqlx.Open("mysql", "root:123123@tcp(127.0.0.1:3306)/ojo?charset=utf8mb4")
	if err != nil {
		fmt.Printf("error:%v", err)
		os.Exit(-1)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("error:%v", err)
		os.Exit(-1)
	}
}
func HandleString(s string) string {
	res := strings.Split(s, "_")
	for j := 0; j < len(res); j++ {
		b := []byte(res[j])
		b[0] = b[0] - 32
		res[j] = string(b)
	}
	return strings.Join(res, "")
}
func LowerInit(s string) string {
	b := []byte(s)
	b[0] = b[0] + 32
	return string(b)
}

func GetData(db *sqlx.DB) ([]Table, error) {
	var tables []Table
	err := db.Select(&tables, "SELECT table_name 'name' FROM information_schema.tables WHERE table_schema='ojo'")
	if err != nil {
		fmt.Printf("error:%v", err)
		return nil, err
	}
	for i := 0; i < len(tables); i++ {
		//fmt.Println(tb.Name)
		var column []Column
		err := db.Select(&column, `SELECT
		COLUMN_NAME "name",
			COLUMN_TYPE "type"
		FROM
		information_schema. COLUMNS
		WHERE TABLE_SCHEMA = 'ojo' AND TABLE_NAME = ?;`, tables[i].Name)
		if err != nil {
			fmt.Printf("error:%v", err)
			return nil, err
		}
		tables[i].Column = column
	}
	//fmt.Println(tables)
	return tables, nil
}

func HandleData(tables []Table) []Table {
	for i := 0; i < len(tables); i++ {
		tables[i].Name = HandleString(tables[i].Name)
		for j := 0; j < len(tables[i].Column); j++ {
			tables[i].Column[j].JsonName = HandleString(tables[i].Column[j].Name)
			switch {
			case strings.Contains(tables[i].Column[j].Type, "tinyint"):
				tables[i].Column[j].Type = "bool"
			case strings.Contains(tables[i].Column[j].Type, "int"):
				tables[i].Column[j].Type = "int"
			case strings.Contains(tables[i].Column[j].Type, "varchar"):
				tables[i].Column[j].Type = "string"
			case strings.Contains(tables[i].Column[j].Type, "text"):
				tables[i].Column[j].Type = "string"
			case strings.Contains(tables[i].Column[j].Type, "timestamp"):
				tables[i].Column[j].Type = "string"
			default:
				tables[i].Column[j].Type = "NULL"
			}
		}
	}
	return tables
}

func WriteData(tables []Table) {
	fmt.Println(tables)
	file, err := os.OpenFile("./dto.go", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	_, err = file.WriteString("package entity\n\n")
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	file.WriteString("type Res struct {\n")
	file.WriteString("Error string `json:\"error\"`\n")
	file.WriteString("Data interface{} `json:\"data\"`\n")
	file.WriteString("}\n")
	for i := 0; i < len(tables); i++ {
		file.WriteString("type " + tables[i].Name + " struct {\n")
		for j := 0; j < len(tables[i].Column); j++ {
			file.WriteString("	" + tables[i].Column[j].JsonName + " " + tables[i].Column[j].Type +
				" `json:\"" + LowerInit(tables[i].Column[j].JsonName) + "\"" +
				" db:\"" + tables[i].Column[j].Name + "\"`" + "\n")
		}
		file.WriteString("}\n")
	}
}

var db *sqlx.DB

func main() {
	data, err := GetData(db)
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	res := HandleData(data)
	WriteData(res)
}
