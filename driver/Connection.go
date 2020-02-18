package driver

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type DB struct {
	SQL *gorm.DB
}

var dbConn = &DB{}

func ConnectSQL(host, port, user, pass, dbname string) (*DB, error) {

	dbSource := fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		pass,
		dbname,
	)

	log.Println(dbSource)

	db, err := gorm.Open("mysql", dbSource)
	if err != nil {
		log.Println("error connecting to the database: ", err)
		panic(err)
	}

	dbConn.SQL = db
	return dbConn, err
}
