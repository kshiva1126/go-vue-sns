package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"strings"
)

var db *gorm.DB
var err error

func Init() {
	connStr := getConnectionString()
	db, err = gorm.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
}

func GetConnection() *gorm.DB {
	return db
}

func getConnectionString() string {
	host := getParamString("MYSQL_DB_HOST", "localhost")
	port := getParamString("MYSQL_PORT", "3306")
	user := getParamString("MYSQL_USER", "root")
	pass := getParamString("MYSQL_PASSWORD", "root")
	dbname := getParamString("MYSQL_DB", "go-vue-sns")
	protocol := getParamString("MYSQL_PROTOCOL", "tcp")
	dbargs := getParamString("MYSQL_DBARGS", "parseTime=true")

	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s", user, pass, protocol, host, port, dbname, dbargs)
}

func getParamString(param string, defaultValue string) string {
	env := os.Getenv(param)
	if env != "" {
		return env
	}
	return defaultValue
}
